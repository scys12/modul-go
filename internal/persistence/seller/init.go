package seller

import (
	"database/sql"
)

type SellerRepo struct {
	dbContext *sql.DB
}

func NewInstance(db *sql.DB) SellerRepo {
	return SellerRepo{
		dbContext: db,
	}
}
