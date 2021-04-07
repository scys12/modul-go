package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/scys12/modul-go/internal/persistence/user"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type UserDelivery struct {
	userRepo user.UserRepo
}

func NewInstance(user user.UserRepo) UserDelivery {
	delivery := UserDelivery{
		userRepo: user,
	}
	return delivery
}
