package services

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/app_errors"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/domain"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/dtos"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/ports/mocks"
)

func TestCreateUser(t *testing.T) {
	testCases := []struct {
		testName      string
		user          dtos.CreateUserDto
		buildStubs    func(mockRepo *mocks.MockIUserRepo)
		checkResponse func(t *testing.T, err error)
	}{
		{
			testName: "UserCreatedSuccessfully",
			user: dtos.CreateUserDto{
				Email:    "test@test.com",
				Password: "test1235",
			},
			buildStubs: func(mockRepo *mocks.MockIUserRepo) {
				mockRepo.EXPECT().FindUserByEmail("test@test.com").
					Times(1).
					Return(domain.User{}, &app_errors.UserNotFound{})
				mockRepo.EXPECT().CreateUser(gomock.Any()).Times(1)
			},
			checkResponse: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		{
			testName: "EmailOrPasswordAlreadyTaken",
			user: dtos.CreateUserDto{
				Email:    "test@test.com",
				Password: "test1235",
			},
			buildStubs: func(mockRepo *mocks.MockIUserRepo) {
				mockRepo.EXPECT().FindUserByEmail("test@test.com").
					Times(1).
					Return(domain.User{}, nil)
				mockRepo.EXPECT().CreateUser(gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, err error) {
				require.ErrorIs(t, err, &app_errors.EmailOrPasswordAlreadyExist{})
			},
		},
		{
			testName: "InternalServerErrorFromFindUserByEmail",
			user: dtos.CreateUserDto{
				Email:    "test@test.com",
				Password: "test1235",
			},
			buildStubs: func(mockRepo *mocks.MockIUserRepo) {
				mockRepo.EXPECT().FindUserByEmail("test@test.com").
					Times(1).
					Return(domain.User{}, &app_errors.InternalServerError{})
				mockRepo.EXPECT().CreateUser(gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, err error) {
				require.ErrorIs(t, err, &app_errors.InternalServerError{})
			},
		},
		{
			testName: "InternalServerErrorFromCreateUser",
			user: dtos.CreateUserDto{
				Email:    "test@test.com",
				Password: "test1235",
			},
			buildStubs: func(mockRepo *mocks.MockIUserRepo) {
				mockRepo.EXPECT().FindUserByEmail("test@test.com").
					Times(1).
					Return(domain.User{}, &app_errors.UserNotFound{})
				mockRepo.EXPECT().CreateUser(gomock.Any()).
					Times(1).
					Return(&app_errors.InternalServerError{})
			},
			checkResponse: func(t *testing.T, err error) {
				require.ErrorIs(t, err, &app_errors.InternalServerError{})
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			storeCtrl := gomock.NewController(t)
			mockRepo := mocks.NewMockIUserRepo(storeCtrl)
			defer storeCtrl.Finish()

			tc.buildStubs(mockRepo)
			userService, err := NewUserService(mockRepo)
			require.NoError(t, err)
			err = userService.CreateUser(tc.user)
			tc.checkResponse(t, err)
		})
	}
}

func TestFindUserByEmail(t *testing.T) {
	testCases := []struct {
		testName      string
		email         string
		buildStubs    func(mockRepo *mocks.MockIUserRepo)
		checkResponse func(t *testing.T, err error)
	}{
		{
			testName: "UserFoundSuccessfully",
			email:    "test@test.com",
			buildStubs: func(mockRepo *mocks.MockIUserRepo) {
				mockRepo.EXPECT().FindUserByEmail("test@test.com").
					Times(1).
					Return(domain.User{}, nil)
			},
			checkResponse: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		{
			testName: "UserNotFound",
			email:    "test@test.com",
			buildStubs: func(mockRepo *mocks.MockIUserRepo) {
				mockRepo.EXPECT().FindUserByEmail("test@test.com").
					Times(1).
					Return(domain.User{}, &app_errors.UserNotFound{})
			},
			checkResponse: func(t *testing.T, err error) {
				require.ErrorIs(t, err, &app_errors.UserNotFound{})
			},
		},
		{
			testName: "InternalServerError",
			email:    "test@test.com",
			buildStubs: func(mockRepo *mocks.MockIUserRepo) {
				mockRepo.EXPECT().FindUserByEmail("test@test.com").
					Times(1).
					Return(domain.User{}, &app_errors.InternalServerError{})
			},
			checkResponse: func(t *testing.T, err error) {
				require.ErrorIs(t, err, &app_errors.InternalServerError{})
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			storeCtrl := gomock.NewController(t)
			mockRepo := mocks.NewMockIUserRepo(storeCtrl)
			defer storeCtrl.Finish()

			tc.buildStubs(mockRepo)
			userService, err := NewUserService(mockRepo)
			require.NoError(t, err)
			_, err = userService.FindUserByEmail(tc.email)
			tc.checkResponse(t, err)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	testCases := []struct {
		testName      string
		email         string
		buildStubs    func(mockRepo *mocks.MockIUserRepo)
		checkResponse func(t *testing.T, err error)
	}{
		{
			testName: "UserDeletedSuccessfully",
			email:    "test@test.com",
			buildStubs: func(mockRepo *mocks.MockIUserRepo) {
				mockRepo.EXPECT().FindUserByEmail("test@test.com").
					Times(1).
					Return(domain.User{}, nil)

				mockRepo.EXPECT().DeleteUser(gomock.Any()).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		{
			testName: "UserNotFound",
			email:    "test@test.com",
			buildStubs: func(mockRepo *mocks.MockIUserRepo) {
				mockRepo.EXPECT().FindUserByEmail("test@test.com").
					Times(1).
					Return(domain.User{}, &app_errors.UserNotFound{})

				mockRepo.EXPECT().DeleteUser(gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, err error) {
				require.ErrorIs(t, err, &app_errors.UserNotFound{})
			},
		},
		{
			testName: "InternalServerErrorFromFindUserByEmail",
			email:    "test@test.com",
			buildStubs: func(mockRepo *mocks.MockIUserRepo) {
				mockRepo.EXPECT().FindUserByEmail("test@test.com").
					Times(1).
					Return(domain.User{}, &app_errors.InternalServerError{})
				mockRepo.EXPECT().DeleteUser(gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, err error) {
				require.ErrorIs(t, err, &app_errors.InternalServerError{})
			},
		},
		{
			testName: "InternalServerErrorFromDeleteUser",
			email:    "test@test.com",
			buildStubs: func(mockRepo *mocks.MockIUserRepo) {
				mockRepo.EXPECT().FindUserByEmail("test@test.com").
					Times(1).
					Return(domain.User{}, nil)
				mockRepo.EXPECT().DeleteUser(gomock.Any()).
					Times(1).
					Return(&app_errors.InternalServerError{})
			},
			checkResponse: func(t *testing.T, err error) {
				require.ErrorIs(t, err, &app_errors.InternalServerError{})
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			storeCtrl := gomock.NewController(t)
			mockRepo := mocks.NewMockIUserRepo(storeCtrl)
			defer storeCtrl.Finish()

			tc.buildStubs(mockRepo)
			userService, err := NewUserService(mockRepo)
			require.NoError(t, err)
			err = userService.DeleteUser(tc.email)
			tc.checkResponse(t, err)
		})
	}
}

func TestVerifyEmail(t *testing.T) {
	testCases := []struct {
		testName      string
		email         string
		buildStubs    func(mockRepo *mocks.MockIUserRepo)
		checkResponse func(t *testing.T, err error)
	}{
		{
			testName: "EmailVerifiedSuccessfully",
			email:    "test@test.com",
			buildStubs: func(mockRepo *mocks.MockIUserRepo) {
				mockRepo.EXPECT().FindUserByEmail("test@test.com").
					Times(1).
					Return(domain.User{}, nil)

				mockRepo.EXPECT().ValidateEmail(gomock.Any()).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		{
			testName: "UserNotFound",
			email:    "test@test.com",
			buildStubs: func(mockRepo *mocks.MockIUserRepo) {
				mockRepo.EXPECT().FindUserByEmail("test@test.com").
					Times(1).
					Return(domain.User{}, &app_errors.UserNotFound{})

				mockRepo.EXPECT().ValidateEmail(gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, err error) {
				require.ErrorIs(t, err, &app_errors.UserNotFound{})
			},
		},
		{
			testName: "InternalServerErrorFromFindUserByEmail",
			email:    "test@test.com",
			buildStubs: func(mockRepo *mocks.MockIUserRepo) {
				mockRepo.EXPECT().FindUserByEmail("test@test.com").
					Times(1).
					Return(domain.User{}, &app_errors.InternalServerError{})
				mockRepo.EXPECT().ValidateEmail(gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, err error) {
				require.ErrorIs(t, err, &app_errors.InternalServerError{})
			},
		},
		{
			testName: "InternalServerErrorFromValidateEmail",
			email:    "test@test.com",
			buildStubs: func(mockRepo *mocks.MockIUserRepo) {
				mockRepo.EXPECT().FindUserByEmail("test@test.com").
					Times(1).
					Return(domain.User{}, nil)
				mockRepo.EXPECT().ValidateEmail(gomock.Any()).
					Times(1).
					Return(&app_errors.InternalServerError{})
			},
			checkResponse: func(t *testing.T, err error) {
				require.ErrorIs(t, err, &app_errors.InternalServerError{})
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			storeCtrl := gomock.NewController(t)
			mockRepo := mocks.NewMockIUserRepo(storeCtrl)
			defer storeCtrl.Finish()

			tc.buildStubs(mockRepo)
			userService, err := NewUserService(mockRepo)
			require.NoError(t, err)
			err = userService.ValidateEmail(tc.email)
			tc.checkResponse(t, err)
		})
	}
}
