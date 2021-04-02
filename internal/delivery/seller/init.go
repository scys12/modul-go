package seller

import (
	"github.com/scys12/modul-go/internal/persistence/seller"
)

type SellerDelivery struct {
	sellerRepo seller.SellerRepo
}

func NewInstance(seller seller.SellerRepo) SellerDelivery {
	delivery := SellerDelivery{
		sellerRepo: seller,
	}
	return delivery
}
