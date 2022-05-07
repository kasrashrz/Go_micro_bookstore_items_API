package http_utils

import (
	"encoding/json"
	"github.com/kasrashrz/Go_micro_bookstore_OAth-go/oath/errors"
	"net/http"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, err *errors.RestErr) {
	RespondJson(w, err.Status, err)
}