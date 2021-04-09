package product

import (
	"github.com/scys12/modul-go/internal/model"
	"github.com/scys12/modul-go/internal/persistence/product"
	"github.com/scys12/modul-go/pkg/payload"
)

type IProductService interface {
	BuyProduct(productID int) error
	DeleteProduct(productID, userID int) error
	GetAvailableProducts() ([]model.Product, error)
	GetProduct(productID int) (model.Product, error)
	GetSellerProducts(sellerID int) ([]model.Product, error)
	InsertProduct(payload.ProductRequest, int) error
	UpdateProduct(payload.ProductRequest, int, int) error
}

type ProductService struct {
	productRepo product.IProductRepo
}

func NewInstance(repo product.IProductRepo) IProductService {
	svc := &ProductService{
		productRepo: repo,
	}
	return svc
}
