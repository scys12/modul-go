package user

import (
	"errors"

	"github.com/scys12/modul-go/pkg/payload"
	"github.com/scys12/modul-go/pkg/tokenizer"
	"golang.org/x/crypto/bcrypt"
)

const role = "user"

func (us *UserService) LoginUser(loginReq payload.LoginRequest) (string, int, error) {
	userChecked, err := us.userRepo.GetUser(loginReq.Username)
	if err != nil {
		return "", 0, err
	}
	passwordCheck := bcrypt.CompareHashAndPassword([]byte(userChecked.Password), []byte(loginReq.Password))
	if passwordCheck != nil {
		err = errors.New("username/password is wrong")
		return "", 0, err
	}
	accessToken, err := tokenizer.CreateToken(role, userChecked.ID)
	if err != nil {
		return "", 0, err
	}
	return accessToken, userChecked.ID, nil
}
