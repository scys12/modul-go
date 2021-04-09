package user

import (
	"github.com/scys12/modul-go/internal/model"
	"github.com/scys12/modul-go/internal/persistence/user"
	"github.com/scys12/modul-go/pkg/payload"
)

type IUserService interface {
	CreateNewUser(payload.RegisterRequest) error
	GetUserByID(userID int) (model.User, error)
	LoginUser(payload.LoginRequest) (string, int, error)
}

type UserService struct {
	userRepo user.IUserRepo
}

func NewInstance(repo user.IUserRepo) IUserService {
	svc := &UserService{
		userRepo: repo,
	}
	return svc
}
