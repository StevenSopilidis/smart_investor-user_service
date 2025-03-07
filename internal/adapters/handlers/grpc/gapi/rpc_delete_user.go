package gapi

import "gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/handlers/grpc/generated"

func (server *Server) DeleteUser(req *generated.DeleteUserRequest) (*generated.DeleteUserResponse, error) {
	err := server.user_service.DeleteUser(req.Email)
	if err != nil {
		return nil, getErrorResponseFromUserServiceError(err)
	}

	return &generated.DeleteUserResponse{}, nil
}
