package customer

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/scys12/modul-go/pkg/payload"
	"github.com/scys12/modul-go/pkg/tokenizer"
	"golang.org/x/crypto/bcrypt"
)

const role = "customer"

func (cd *CustomerDelivery) LoginCustomer(w http.ResponseWriter, r *http.Request) {
	var loginReq payload.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		payload.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	err = validate.Struct(loginReq)
	if err != nil {
		payload.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	customer, err := cd.customerRepo.GetCustomer(loginReq.Username)
	if err != nil {
		payload.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	passwordCheck := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(loginReq.Password))
	if passwordCheck != nil {
		err = errors.New("username/password is wrong")
		payload.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	accessToken, err := tokenizer.CreateToken(role, customer.ID)
	if err != nil {
		payload.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	resp := payload.AuthResponse{
		Token:  accessToken,
		UserID: customer.ID,
	}
	payload.ResponseOK(w, resp)
}
