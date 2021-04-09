package user

import (
	"database/sql"

	"github.com/scys12/modul-go/internal/model"
)

type UserRepo struct {
	dbContext *sql.DB
}

type IUserRepo interface {
	CheckUsernameAndEmail(string, string) (bool, error)
	InsertUser(model.User) error
	GetUser(string) (model.User, error)
	GetUserByID(int) (model.User, error)
}

func NewInstance(db *sql.DB) IUserRepo {
	return &UserRepo{
		dbContext: db,
	}
}
