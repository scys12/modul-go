package user_test

import (
	"errors"
	"testing"

	"github.com/scys12/modul-go/internal/model"
	"github.com/scys12/modul-go/internal/service/user"
	"github.com/scys12/modul-go/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUser(t *testing.T) {
	testUser := []struct {
		name   string
		userID int
		err    error
		user   model.User
	}{
		{
			name:   "Success get user",
			userID: 1,
			err:    nil,
			user:   model.User{ID: 1, Email: "customer@mail.com"},
		},
		{
			name:   "Failed get user",
			userID: 0,
			err:    errors.New("err"),
			user:   model.User{ID: 1, Email: "customer@mail.com"},
		},
	}
	for _, tu := range testUser {
		t.Run(tu.name, func(t *testing.T) {
			mockRepo := new(mocks.IUserRepo)
			mockRepo.On("GetUserByID", mock.Anything).Return(tu.user, tu.err)
			service := user.NewInstance(mockRepo)
			user, err := service.GetUserByID(tu.userID)
			assert.Equal(t, tu.err, err)
			assert.Equal(t, tu.user.ID, user.ID)
		})
	}
}
