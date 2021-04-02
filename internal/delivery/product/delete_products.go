package product

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/scys12/modul-go/pkg/payload"
)

func (pd *ProductDelivery) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(r.Header.Get("X-Header-UserID"))
	vars := mux.Vars(r)
	productID, _ := strconv.Atoi(vars["id"])
	err := pd.productRepo.DeleteProduct(productID, userID)
	if err != nil {
		payload.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	payload.ResponseOK(w, "Product deleted")
}
