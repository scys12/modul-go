package middleware

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/scys12/modul-go/pkg/payload"
	"github.com/scys12/modul-go/pkg/tokenizer"
	"github.com/sirupsen/logrus"
)

const (
	headerRole   = "X-Header-Role"
	headerUserID = "X-Header-UserID"
)

func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		accessToken := tokenizer.ExtractToken(tokenStr)
		if accessToken == "" {
			err := errors.New("wrong token")
			payload.ResponseError(w, http.StatusUnauthorized, err)
			return
		}
		accessDetail, err := tokenizer.DecodeToken(accessToken)
		if err != nil {
			logrus.Error(err)
			payload.ResponseError(w, http.StatusUnauthorized, err)
			return
		}
		userID := int(accessDetail["user_id"].(float64))
		strUserID := strconv.Itoa(userID)
		r.Header.Set(headerRole, accessDetail["role"].(string))
		r.Header.Set(headerUserID, strUserID)
		next.ServeHTTP(w, r)
	})
}
