package gapi

import (
	"context"

	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/handlers/grpc/generated"
)

func (server *Server) VerifyEmail(ctx context.Context, req *generated.VerifyEmailRequest) (*generated.VerifyEmailResponse, error) {
	err := server.user_service.ValidateEmail(req.Email, req.EmailVerificiationCode)
	if err != nil {
		return nil, getErrorResponseFromUserServiceError(err)
	}

	return &generated.VerifyEmailResponse{}, nil
}
