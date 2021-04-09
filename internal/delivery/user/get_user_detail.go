package user

import (
	"net/http"
	"strconv"

	"github.com/scys12/modul-go/pkg/payload"
)

func (ud *UserDelivery) GetUserDetail(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.Header.Get("X-Header-UserID"))
	if err != nil {
		payload.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	currentUser, err := ud.userService.GetUserByID(userID)
	if err != nil {
		payload.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	payload.ResponseOK(w, currentUser)
}
