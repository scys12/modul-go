package seller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/scys12/modul-go/pkg/payload"
	"github.com/scys12/modul-go/pkg/tokenizer"
	"golang.org/x/crypto/bcrypt"
)

const role = "seller"

func (cd *SellerDelivery) LoginSeller(w http.ResponseWriter, r *http.Request) {
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
	seller, err := cd.sellerRepo.GetSeller(loginReq.Username)
	if err != nil {
		payload.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	passwordCheck := bcrypt.CompareHashAndPassword([]byte(seller.Password), []byte(loginReq.Password))
	if passwordCheck != nil {
		err = errors.New("username/password is wrong")
		payload.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	accessToken, err := tokenizer.CreateToken(role, seller.ID)
	if err != nil {
		payload.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	resp := payload.AuthResponse{
		Token:  accessToken,
		UserID: seller.ID,
	}
	payload.ResponseOK(w, resp)
}
