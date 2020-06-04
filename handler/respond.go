package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ReplyError would response errors to api request
func ReplyError(w *http.ResponseWriter, code int, args ...interface{}) {
	response := make(map[string]string)
	response["error"] = fmt.Sprint(args...)
	responseJSON, _ := json.Marshal(response)

	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(code)
	(*w).Write(responseJSON)
}

// ReplySuccess would response success to api request
func ReplySuccess(w *http.ResponseWriter, responseData interface{}) {
	responseJSON, _ := json.Marshal(responseData)
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Write(responseJSON)
}
