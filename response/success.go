package response

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

func OK(msg string, data interface{}) Response {
	return success(msg, http.StatusOK, data)
}

func Created(msg string, data interface{}) Response {
	return success(msg, http.StatusCreated, data)
}

func Accepted(msg string, data interface{}) Response {
	return success(msg, http.StatusAccepted, data)
}

func NonAuthoritativeInformation(msg string, data interface{}) Response {
	return success(msg, http.StatusNonAuthoritativeInfo, data)
}

func NoContent(msg string, data interface{}) Response {
	return success(msg, http.StatusNoContent, data)
}

func ResetContent(msg string, data interface{}) Response {
	return success(msg, http.StatusResetContent, data)
}

func success(message string, status int, data interface{}) Response {
	return &SuccessResponse{
		Message: message,
		Status:  status,
		Data:    data,
	}
}

func (s *SuccessResponse) Error() string {
	return ""
}

func (s *SuccessResponse) StatusCode() int {
	return s.Status
}

func (s *SuccessResponse) GetBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s *SuccessResponse) GetData() interface{} {
	return s.Data
}
