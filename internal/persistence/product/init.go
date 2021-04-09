package product

import (
	"database/sql"

	"github.com/scys12/modul-go/internal/model"
)

type ProductRepo struct {
	dbContext *sql.DB
}

type IProductRepo interface {
	CheckUsernameAndEmail(string, string) (bool, error)
	InsertProduct(product model.Product) error
	GetCurrentUserProduct(int, int) (model.Product, error)
	GetProduct(int) (model.Product, error)
	UpdateProduct(model.Product) error
	DeleteProduct(int, int) error
	GetSellerProducts(int) ([]model.Product, error)
	BuyProduct(int) error
	GetAvailableProducts() ([]model.Product, error)
}

func NewInstance(db *sql.DB) IProductRepo {
	return &ProductRepo{
		dbContext: db,
	}
}
