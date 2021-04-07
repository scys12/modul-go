package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/scys12/modul-go/pkg/payload"
)

func (pd *ProductDelivery) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var productReq payload.ProductRequest
	userID, _ := strconv.Atoi(r.Header.Get("X-Header-UserID"))
	vars := mux.Vars(r)
	productID, _ := strconv.Atoi(vars["id"])
	err := json.NewDecoder(r.Body).Decode(&productReq)
	if err != nil {
		payload.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	err = validate.Struct(productReq)
	if err != nil {
		payload.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	updatedProduct, err := pd.productRepo.GetCurrentUserProduct(productID, userID)
	if err != nil {
		payload.ResponseError(w, http.StatusNotFound, err)
		return
	}
	updatedProduct.Name = productReq.Name
	updatedProduct.Description = productReq.Description
	updatedProduct.Price = productReq.Price

	err = pd.productRepo.UpdateProduct(updatedProduct)
	if err != nil {
		payload.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	payload.ResponseOK(w, "Product updated")
}
