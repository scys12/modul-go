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

func TestRegisterUser(t *testing.T) {
	testUser := []struct {
		name        string
		err         error
		resultCode  int
		requestBody string
	}{
		{
			name:        "Success Register",
			err:         nil,
			resultCode:  http.StatusOK,
			requestBody: `{"email": "customer@mail.com", "password": "customer123", "username": "customer123"}`,
		},
		{
			name:        "Failed Bad Request Body",
			err:         errors.New("error"),
			resultCode:  http.StatusBadRequest,
			requestBody: `{"email": "customer@mail.com", "password": "customer123", "username": }`,
		},
		{
			name:        "Failed Struct Validation",
			err:         errors.New("error"),
			resultCode:  http.StatusBadRequest,
			requestBody: `{"email": "customer@mail.com", "password": "c", "username": "c"}`,
		},
		{
			name:        "Failed Create User",
			err:         errors.New("error"),
			resultCode:  http.StatusInternalServerError,
			requestBody: `{"email": "customer@mail.com", "password": "customer123", "username": "customer123"}`,
		},
	}
	for _, tu := range testUser {
		t.Run(tu.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/user/register", strings.NewReader(tu.requestBody))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			mockService := new(mocks.IUserService)
			mockService.On("CreateNewUser", mock.Anything).Return(tu.err)
			handler := user.NewInstance(mockService)
			handler.CreateNewUser(rec, req)
			assert.Equal(t, tu.resultCode, rec.Code)
		})
	}
}
