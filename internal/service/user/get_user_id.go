package user

import (
	"github.com/scys12/modul-go/internal/model"
)

func (us *UserService) GetUserByID(userID int) (user model.User, err error) {
	user, err = us.userRepo.GetUserByID(userID)
	return
}
