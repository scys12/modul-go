package product

import (
	"database/sql"
)

type ProductRepo struct {
	dbContext *sql.DB
}

func NewInstance(db *sql.DB) ProductRepo {
	return ProductRepo{
		dbContext: db,
	}
}
