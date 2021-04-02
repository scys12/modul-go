package product

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/scys12/modul-go/pkg/payload"
)

func (pd *ProductDelivery) BuyProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, _ := strconv.Atoi(vars["id"])
	err := pd.productRepo.BuyProduct(productID)
	if err != nil {
		payload.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	payload.ResponseOK(w, "Product deleted")
}
