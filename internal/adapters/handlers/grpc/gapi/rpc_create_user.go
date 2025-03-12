package gapi

import (
	"context"
	"net/mail"

	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/handlers/grpc/generated"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/dtos"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateUser(ctx context.Context, req *generated.CreateUserRequest) (*generated.CreateUserResponse, error) {
	violations := vefifyCreateUserRequest(req)

	if len(violations) > 0 {
		return nil, getErrorResponseFromFieldViolations(violations)
	}

	user, err := server.user_service.CreateUser(dtos.CreateUserDto{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return nil, getErrorResponseFromUserServiceError(err)
	}

	return &generated.CreateUserResponse{
		Id:                    user.Id.String(),
		Email:                 user.Email,
		EmailVerified:         user.EmailVerified,
		EmailVerificationCode: user.EmailVerificationCode,
		CreatedAt:             timestamppb.New(user.CreatedAt),
	}, nil
}

func vefifyCreateUserRequest(req *generated.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if len(req.GetPassword()) < 8 || len(req.GetPassword()) > 20 {
		violations = append(violations, fieldViolation("password", "Password string must be between 8 and 20 chars"))
	}

	if _, err := mail.ParseAddress(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", "Invalid email provided"))
	}

	return violations
}
