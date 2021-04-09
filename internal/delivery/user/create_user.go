package user

import (
	"encoding/json"
	"net/http"

	"github.com/scys12/modul-go/pkg/payload"
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
	err = ud.userService.CreateNewUser(registerReq)
	if err != nil {
		payload.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	payload.ResponseOK(w, "User inserted")
}
