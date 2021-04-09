package user

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/scys12/modul-go/internal/service/user"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type IUserDelivery interface {
	CreateNewUser(http.ResponseWriter, *http.Request)
	GetUserDetail(http.ResponseWriter, *http.Request)
	LoginUser(http.ResponseWriter, *http.Request)
}

type UserDelivery struct {
	userService user.IUserService
}

func NewInstance(user user.IUserService) IUserDelivery {
	delivery := UserDelivery{
		userService: user,
	}
	return &delivery
}
