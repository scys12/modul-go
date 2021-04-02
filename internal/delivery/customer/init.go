package customer

import (
	"github.com/go-playground/validator/v10"
	"github.com/scys12/modul-go/internal/persistence/customer"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type CustomerDelivery struct {
	customerRepo customer.CustomerRepo
}

func NewInstance(customer customer.CustomerRepo) CustomerDelivery {
	delivery := CustomerDelivery{
		customerRepo: customer,
	}
	return delivery
}
