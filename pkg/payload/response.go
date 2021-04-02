package payload

import (
	"encoding/json"
	"net/http"
)

type AuthResponse struct {
	Token  string `json:"token"`
	UserID int    `json:"user_id"`
}

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func ResponseOK(w http.ResponseWriter, body interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := Response{
		Status: "success",
		Data:   body,
	}
	return json.NewEncoder(w).Encode(resp)
}

func ResponseError(w http.ResponseWriter, code int, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	resp := Response{
		Status: "failed",
		Data:   err.Error(),
	}
	return json.NewEncoder(w).Encode(resp)
}
