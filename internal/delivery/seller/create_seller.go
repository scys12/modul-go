package seller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/scys12/modul-go/internal/model"
	"github.com/scys12/modul-go/pkg/payload"
	"golang.org/x/crypto/bcrypt"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (cd SellerDelivery) CreateNewSeller(w http.ResponseWriter, r *http.Request) {
	var registerReq payload.SellerRegisterRequest
	err := json.NewDecoder(r.Body).Decode(&registerReq)
	if err != nil {
		payload.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	err = validate.Struct(registerReq)
	if err != nil {
		payload.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	isExist, err := cd.sellerRepo.CheckUsernameAndEmail(registerReq.Username, registerReq.Email)
	if err != nil {
		payload.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	if isExist {
		err = errors.New("username or email exists")
		payload.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	pwd, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), bcrypt.DefaultCost)
	if err != nil {
		payload.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	Seller := model.Seller{
		Email:    registerReq.Email,
		Password: string(pwd),
		Username: registerReq.Username,
	}
	err = cd.sellerRepo.InsertSeller(Seller)
	if err != nil {
		payload.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	payload.ResponseOK(w, "Seller inserted")
}
