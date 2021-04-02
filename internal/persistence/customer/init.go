package customer

import (
	"database/sql"
)

type CustomerRepo struct {
	dbContext *sql.DB
}

func NewInstance(db *sql.DB) CustomerRepo {
	return CustomerRepo{
		dbContext: db,
	}
}
