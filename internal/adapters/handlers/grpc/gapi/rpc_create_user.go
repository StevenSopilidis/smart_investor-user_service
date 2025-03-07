package gapi

import (
	"context"
	"net/mail"

	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/handlers/grpc/generated"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/dtos"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *Server) CreateUser(ctx context.Context, req *generated.CreateUserRequest) (*generated.CreateUserResponse, error) {
	violations := vefifyCreateUserRequest(req)

	if violations != nil {
		return nil, getErrorResponseFromFieldViolations(violations)
	}

	_, err := server.user_service.CreateUser(dtos.CreateUserDto{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return nil, getErrorResponseFromUserServiceError(err)
	}

	return &generated.CreateUserResponse{}, nil
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
