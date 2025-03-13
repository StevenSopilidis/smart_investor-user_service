package gapi

import (
	"context"

	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/handlers/grpc/generated"
)

func (server *Server) DeleteUser(ctx context.Context, req *generated.DeleteUserRequest) (*generated.DeleteUserResponse, error) {
	err := server.user_service.DeleteUser(req.Email)
	if err != nil {
		return nil, getErrorResponseFromUserServiceError(err)
	}

	return &generated.DeleteUserResponse{}, nil
}
