package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/scys12/modul-go/pkg/payload"
	"github.com/scys12/modul-go/pkg/tokenizer"
	"golang.org/x/crypto/bcrypt"
)

const role = "user"

func (user *UserDelivery) LoginUser(w http.ResponseWriter, r *http.Request) {
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
	userChecked, err := user.userRepo.GetUser(loginReq.Username)
	if err != nil {
		payload.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	passwordCheck := bcrypt.CompareHashAndPassword([]byte(userChecked.Password), []byte(loginReq.Password))
	if passwordCheck != nil {
		err = errors.New("username/password is wrong")
		payload.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	accessToken, err := tokenizer.CreateToken(role, userChecked.ID)
	if err != nil {
		payload.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	resp := payload.AuthResponse{
		Token:  accessToken,
		UserID: userChecked.ID,
	}
	payload.ResponseOK(w, resp)
}
