package user_test

import (
	"errors"
	"testing"

	"github.com/scys12/modul-go/internal/service/user"
	"github.com/scys12/modul-go/mocks"
	"github.com/scys12/modul-go/pkg/payload"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUser(t *testing.T) {
	testUser := []struct {
		name          string
		err           error
		registerReq   payload.RegisterRequest
		isExist       bool
		expectedError error
	}{
		{
			name:    "Success Register",
			err:     nil,
			isExist: false,
			registerReq: payload.RegisterRequest{
				Email:    "customer@gmail.com",
				Username: "customer123",
				Password: "customer123",
			},
			expectedError: nil,
		},
		{
			name:    "Failed User Exist",
			err:     nil,
			isExist: true,
			registerReq: payload.RegisterRequest{
				Email:    "customer@gmail.com",
				Username: "customer123",
				Password: "customer123",
			},
			expectedError: errors.New("username or email exists"),
		},
		{
			name:    "Failed Error Database",
			err:     errors.New("error"),
			isExist: true,
			registerReq: payload.RegisterRequest{
				Email:    "customer@gmail.com",
				Username: "customer123",
				Password: "customer123",
			},
			expectedError: errors.New("error"),
		},
		{
			name:    "Failed Error Password",
			err:     nil,
			isExist: false,
			registerReq: payload.RegisterRequest{
				Email:    "customer@gmail.com",
				Username: "customer123",
				Password: "",
			},
			expectedError: nil,
		},
	}
	for _, tu := range testUser {
		t.Run(tu.name, func(t *testing.T) {
			mockRepo := new(mocks.IUserRepo)
			mockRepo.On("CheckUsernameAndEmail", mock.Anything, mock.Anything).Return(tu.isExist, tu.err)
			mockRepo.On("InsertUser", mock.Anything).Return(tu.err)
			service := user.NewInstance(mockRepo)
			err := service.CreateNewUser(tu.registerReq)
			assert.Equal(t, tu.expectedError, err)
		})
	}
}
