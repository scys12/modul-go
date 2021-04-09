package user_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/scys12/modul-go/internal/delivery/user"
	"github.com/scys12/modul-go/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoginUser(t *testing.T) {
	testUser := []struct {
		name        string
		err         error
		resultCode  int
		requestBody string
		token       string
		id          int
	}{
		{
			name:        "Success Login",
			err:         nil,
			resultCode:  http.StatusOK,
			requestBody: `{"password": "customer123", "username": "customer123"}`,
			token:       "abcd",
			id:          1,
		},
		{
			name:        "Failed Bad Request Body",
			err:         errors.New("error"),
			resultCode:  http.StatusBadRequest,
			requestBody: `{"password": "customer123", "username": }`,
		},
		{
			name:        "Failed Struct Validation",
			err:         errors.New("error"),
			resultCode:  http.StatusBadRequest,
			requestBody: `{"password": "c", "username": "c"}`,
		},
		{
			name:        "Failed Login User",
			err:         errors.New("error"),
			resultCode:  http.StatusInternalServerError,
			requestBody: `{"password": "customer123", "username": "customer123"}`,
			token:       "",
			id:          0,
		},
	}
	for _, tu := range testUser {
		t.Run(tu.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/user/login", strings.NewReader(tu.requestBody))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			mockService := new(mocks.IUserService)
			mockService.On("LoginUser", mock.Anything).Return(tu.token, tu.id, tu.err)
			handler := user.NewInstance(mockService)
			handler.LoginUser(rec, req)
			assert.Equal(t, tu.resultCode, rec.Code)
		})
	}
}
