package user

import (
	"database/sql"
)

type UserRepo struct {
	dbContext *sql.DB
}

func NewInstance(db *sql.DB) UserRepo {
	return UserRepo{
		dbContext: db,
	}
}
