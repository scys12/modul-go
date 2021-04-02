package customer

import (
	"net/http"
	"strconv"

	"github.com/scys12/modul-go/pkg/payload"
)

func (cd *CustomerDelivery) GetUserDetail(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(r.Header.Get("X-Header-UserID"))
	currentUser, err := cd.customerRepo.GetCustomerByID(userID)
	if err != nil {
		payload.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	payload.ResponseOK(w, currentUser)
}
