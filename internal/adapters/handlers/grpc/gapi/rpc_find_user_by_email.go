package gapi

import (
	"context"

	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/handlers/grpc/generated"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) FindUserByEmail(ctx context.Context, req *generated.FindUserByEmailRequest) (*generated.FindUserByEmailResponse, error) {
	user, err := server.user_service.FindUserByEmail(req.Email)
	if err != nil {
		return nil, getErrorResponseFromUserServiceError(err)
	}
	return &generated.FindUserByEmailResponse{
		Id:                    user.Id.String(),
		Email:                 user.Email,
		EmailVerified:         user.EmailVerified,
		EmailVerificationCode: user.EmailVerificationCode,
		Password:              user.Password,
		CreatedAt:             timestamppb.New(user.CreatedAt),
	}, nil
}
