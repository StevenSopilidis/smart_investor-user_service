package gapi

import (
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/app_errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func getErrorResponseFromUserServiceError(err error) error {
	switch e := err.(type) {
	case *app_errors.EmailOrPasswordAlreadyExist:
		return status.Errorf(codes.AlreadyExists, e.Error())
	case *app_errors.InvalidVerificationCode:
		return status.Errorf(codes.InvalidArgument, e.Error())
	case *app_errors.UserNotFound:
		return status.Errorf(codes.NotFound, e.Error())
	case *app_errors.InvalidPassword:
		return status.Errorf(codes.Unauthenticated, e.Error())
	default:
		return status.Errorf(codes.Internal, e.Error())
	}
}
