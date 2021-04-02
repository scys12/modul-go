package product

import (
	"github.com/go-playground/validator/v10"
	"github.com/scys12/modul-go/internal/persistence/product"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type ProductDelivery struct {
	productRepo product.ProductRepo
}

func NewInstance(product product.ProductRepo) ProductDelivery {
	delivery := ProductDelivery{
		productRepo: product,
	}
	return delivery
}
