package product

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/scys12/modul-go/pkg/payload"
)

func (pd *ProductDelivery) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])
	product, err := pd.productRepo.GetProduct(userID)
	if err != nil {
		payload.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	payload.ResponseOK(w, product)
}
