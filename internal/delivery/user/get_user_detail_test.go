package user_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/scys12/modul-go/internal/delivery/user"
	"github.com/scys12/modul-go/internal/model"
	"github.com/scys12/modul-go/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserDetail(t *testing.T) {
	testUser := []struct {
		name       string
		err        error
		resultCode int
		user       model.User
		id         string
	}{
		{
			name:       "Success Get User Detail",
			err:        nil,
			resultCode: http.StatusOK,
			user:       model.User{ID: 1, Email: "customer@mail.com"},
			id:         "1",
		},
		{
			name:       "Failed Bad Header ID",
			err:        errors.New("error"),
			resultCode: http.StatusBadRequest,
			user:       model.User{ID: 1, Email: "customer@mail.com"},
			id:         "ab",
		},
		{
			name:       "Failed Get User Detail",
			err:        errors.New("error"),
			resultCode: http.StatusInternalServerError,
			user:       model.User{ID: 1, Email: "customer@mail.com"},
			id:         "1",
		},
	}
	for _, tu := range testUser {
		t.Run(tu.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/user/me", nil)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-Header-UserID", tu.id)
			rec := httptest.NewRecorder()

			mockService := new(mocks.IUserService)
			mockService.On("GetUserByID", mock.Anything).Return(tu.user, tu.err)
			handler := user.NewInstance(mockService)
			handler.GetUserDetail(rec, req)
			assert.Equal(t, tu.resultCode, rec.Code)
		})
	}
}
