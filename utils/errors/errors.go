package errors

import (
	"fmt"
	json "github.com/json-iterator/go"
	"net/http"
)

type CauseList []interface{}

type ApiError interface {
	Message() string
	Code() string
	Status() int
	Cause() CauseList
	Error() string
}

type apiErr struct {
	ErrorMessage string    `json:"message"`
	ErrorCode    string    `json:"error"`
	ErrorStatus  int       `json:"status"`
	ErrorCause   CauseList `json:"cause"`
}

func (c CauseList) ToString() string {
	return fmt.Sprint(c)
}

func (e apiErr) Code() string {
	return e.ErrorCode
}

func (e apiErr) Error() string {
	return fmt.Sprintf("Message: %s;Error Code: %s;Status: %d;Cause: %v", e.ErrorMessage, e.ErrorCode, e.ErrorStatus, e.ErrorCause)
}

func (e apiErr) Status() int {
	return e.ErrorStatus
}

func (e apiErr) Cause() CauseList {
	return e.ErrorCause
}

func (e apiErr) Message() string {
	return e.ErrorMessage
}

func NewApiError(message string, error string, status int, cause CauseList) ApiError {
	return apiErr{message, error, status, cause}
}

func NewNotFoundApiError(message string) ApiError {
	return apiErr{message, "not_found", http.StatusNotFound, CauseList{}}
}

func NewTooManyRequestsError(message string) ApiError {
	return apiErr{message, "too_many_requests", http.StatusTooManyRequests, CauseList{}}
}

func NewBadRequestApiError(message string) ApiError {
	return apiErr{message, "bad_request", http.StatusBadRequest, CauseList{}}
}

func NewValidationApiError(message string, error string, cause CauseList) ApiError {
	return apiErr{message, error, http.StatusBadRequest, cause}
}

func NewMethodNotAllowedApiError() ApiError {
	return apiErr{"Method not allowed", "method_not_allowed", http.StatusMethodNotAllowed, CauseList{}}
}

func NewInternalServerApiError(message string, err error) ApiError {
	cause := CauseList{}
	if err != nil {
		cause = append(cause, err.Error())
	}
	return apiErr{message, "internal_server_error", http.StatusInternalServerError, cause}
}

func NewForbiddenApiError(message string) ApiError {
	return apiErr{message, "forbidden", http.StatusForbidden, CauseList{}}
}

func NewUnauthorizedApiError(message string) ApiError {
	return apiErr{message, "unauthorized_scopes", http.StatusUnauthorized, CauseList{}}
}

func NewConflictApiError(id string) ApiError {
	return apiErr{"Can't update " + id + " due to a conflict error", "conflict_error", http.StatusConflict, CauseList{}}
}

func NewApiErrorFromBytes(data []byte) (ApiError, error) {
	err := apiErr{}
	e := json.Unmarshal(data, &err)
	return err, e
}