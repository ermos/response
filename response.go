package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Success[T any](w http.ResponseWriter, status int, content T) {
	w.WriteHeader(status)
	if content == nil {
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := json.NewEncoder(w).Encode(content); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func Error(w http.ResponseWriter, status int, code string, message interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	messageString := ""

	switch msg := message.(type) {
	case string:
		messageString = msg
	case error:
		messageString = msg.Error()
	default:
		messageString = fmt.Sprintf("%v", message)
	}

	err := json.NewEncoder(w).Encode(ErrorResponse{
		Code:    code,
		Message: messageString,
	})
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func InternalServerError(w http.ResponseWriter, code string, err error) {
	Error(w, http.StatusInternalServerError, code, err)
}

func BadRequest(w http.ResponseWriter, code string, err error) {
	Error(w, http.StatusBadRequest, code, err)
}

func Unauthorized(w http.ResponseWriter, code string, err error) {
	Error(w, http.StatusUnauthorized, code, err)
}

func Forbidden(w http.ResponseWriter, code string, err error) {
	Error(w, http.StatusForbidden, code, err)
}

func NotFound(w http.ResponseWriter, code string, err error) {
	Error(w, http.StatusNotFound, code, err)
}

func Conflict(w http.ResponseWriter, code string, err error) {
	Error(w, http.StatusConflict, code, err)
}

func TooManyRequests(w http.ResponseWriter, code string, err error) {
	Error(w, http.StatusTooManyRequests, code, err)
}

func ServiceUnavailable(w http.ResponseWriter, code string, err error) {
	Error(w, http.StatusServiceUnavailable, code, err)
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
