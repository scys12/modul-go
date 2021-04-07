package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/scys12/modul-go/internal/model"
	"github.com/scys12/modul-go/pkg/payload"
)

func (pd *ProductDelivery) InsertProduct(w http.ResponseWriter, r *http.Request) {
	var productReq payload.ProductRequest
	userID, _ := strconv.Atoi(r.Header.Get("X-Header-UserID"))

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
	product := model.Product{
		Name:        productReq.Name,
		Description: productReq.Description,
		Price:       productReq.Price,
		SellerID:    userID,
	}
	err = pd.productRepo.InsertProduct(product)
	if err != nil {
		payload.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	payload.ResponseOK(w, "Product inserted")
}
