package product

import (
	"net/http"

	"github.com/scys12/modul-go/pkg/payload"
)

func (pd *ProductDelivery) GetAvailableProducts(w http.ResponseWriter, r *http.Request) {
	products, err := pd.productRepo.GetAvailableProducts()
	if err != nil {
		payload.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	payload.ResponseOK(w, products)
}
