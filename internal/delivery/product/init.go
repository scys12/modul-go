package product

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/scys12/modul-go/internal/service/product"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type IProductDelivery interface {
	BuyProduct(http.ResponseWriter, *http.Request)
	DeleteProduct(http.ResponseWriter, *http.Request)
	GetAvailableProducts(http.ResponseWriter, *http.Request)
	GetProduct(http.ResponseWriter, *http.Request)
	GetSellerProducts(http.ResponseWriter, *http.Request)
	InsertProduct(http.ResponseWriter, *http.Request)
	UpdateProduct(http.ResponseWriter, *http.Request)
}

type ProductDelivery struct {
	productService product.IProductService
}

func NewInstance(service product.IProductService) IProductDelivery {
	delivery := &ProductDelivery{
		productService: service,
	}
	return delivery
}
