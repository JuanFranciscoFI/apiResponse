package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func InternalServerError(msg string) Response {
	return err(msg, http.StatusInternalServerError)
}

func BadRequest(msg string) Response {
	return err(msg, http.StatusBadRequest)
}

func NotFound(msg string) Response {
	return err(msg, http.StatusNotFound)
}

func Unauthorized(msg string) Response {
	return err(msg, http.StatusUnauthorized)
}

func err(message string, status int) Response {
	return &ErrorResponse{
		Message: message,
		Status:  status,
	}
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func (e *ErrorResponse) StatusCode() int {
	return e.Status
}

func (e *ErrorResponse) GetBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e *ErrorResponse) GetData() interface{} {
	return nil
}
