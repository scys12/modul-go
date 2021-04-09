package user_test

import (
	"errors"
	"testing"

	"github.com/scys12/modul-go/internal/model"
	"github.com/scys12/modul-go/internal/service/user"
	"github.com/scys12/modul-go/mocks"
	"github.com/scys12/modul-go/pkg/payload"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestLoginUser(t *testing.T) {
	password := "customer123"
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	testUser := []struct {
		name          string
		err           error
		user          model.User
		loginReq      payload.LoginRequest
		expectedError error
	}{
		{
			name:          "Success login user",
			err:           nil,
			user:          model.User{ID: 1, Email: "customer@mail.com", Password: string(hashedPass)},
			loginReq:      payload.LoginRequest{Username: "customer123", Password: "customer123"},
			expectedError: nil,
		},
		{
			name:          "Failed get user",
			err:           errors.New("err"),
			user:          model.User{ID: 1, Email: "customer@mail.com", Password: string(hashedPass)},
			loginReq:      payload.LoginRequest{Username: "customer123", Password: "customer123"},
			expectedError: errors.New("err"),
		},
		{
			name:          "Failed Password Not Same",
			err:           nil,
			user:          model.User{ID: 1, Email: "customer@mail.com", Password: string(hashedPass)},
			loginReq:      payload.LoginRequest{Username: "customer123", Password: "custom"},
			expectedError: errors.New("username/password is wrong"),
		},
	}
	for _, tu := range testUser {
		t.Run(tu.name, func(t *testing.T) {
			mockRepo := new(mocks.IUserRepo)
			mockRepo.On("GetUser", mock.Anything).Return(tu.user, tu.err)
			service := user.NewInstance(mockRepo)
			_, _, err := service.LoginUser(tu.loginReq)
			t.Log(err)
			assert.Equal(t, tu.expectedError, err)
		})
	}
}
