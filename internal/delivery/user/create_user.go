package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/scys12/modul-go/internal/model"
	"github.com/scys12/modul-go/pkg/payload"
	"golang.org/x/crypto/bcrypt"
)

func (ud *UserDelivery) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var registerReq payload.RegisterRequest
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
	isExist, err := ud.userRepo.CheckUsernameAndEmail(registerReq.Username, registerReq.Email)
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
	user := model.User{
		Email:    registerReq.Email,
		Password: string(pwd),
		Username: registerReq.Username,
	}
	err = ud.userRepo.InsertUser(user)
	if err != nil {
		payload.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	payload.ResponseOK(w, "User inserted")
}
