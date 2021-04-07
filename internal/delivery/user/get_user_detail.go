package user

import (
	"net/http"
	"strconv"

	"github.com/scys12/modul-go/pkg/payload"
)

func (ud *UserDelivery) GetUserDetail(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(r.Header.Get("X-Header-UserID"))
	currentUser, err := ud.userRepo.GetUserByID(userID)
	if err != nil {
		payload.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	payload.ResponseOK(w, currentUser)
}
