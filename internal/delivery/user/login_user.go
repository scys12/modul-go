package user

import (
	"encoding/json"
	"net/http"

	"github.com/scys12/modul-go/pkg/payload"
)

func (ud *UserDelivery) LoginUser(w http.ResponseWriter, r *http.Request) {
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

	accessToken, id, err := ud.userService.LoginUser(loginReq)
	if err != nil {
		payload.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	resp := payload.AuthResponse{
		Token:  accessToken,
		UserID: id,
	}
	payload.ResponseOK(w, resp)
}
