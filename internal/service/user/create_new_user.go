package user

import (
	"errors"

	"github.com/scys12/modul-go/internal/model"
	"github.com/scys12/modul-go/pkg/payload"
	"golang.org/x/crypto/bcrypt"
)

func (us *UserService) CreateNewUser(registerReq payload.RegisterRequest) (err error) {
	isExist, err := us.userRepo.CheckUsernameAndEmail(registerReq.Username, registerReq.Email)
	if err != nil {
		return
	}
	if isExist {
		err = errors.New("username or email exists")
		return
	}
	pwd, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	user := model.User{
		Email:    registerReq.Email,
		Password: string(pwd),
		Username: registerReq.Username,
	}
	err = us.userRepo.InsertUser(user)
	return
}
