package apibuildr

import "fmt"

type ApiError struct {
	Message    string            `json:"message"`
	ApiCode    string            `json:"api_code"`
	Cause      error             `json:"-"`
	StatusCode int               `json:"status_code"`
	Errors     map[string]string `json:"errors"`
}

func (e ApiError) Error() string {
	if e.ApiCode == "" {
		return e.Message
	}
	return fmt.Sprintf("%v : %s", e.ApiCode, e.Message)
}

func (e *ApiError) Unwrap() error {
	return e.Cause
}

func NewApiError(status int) *ApiError {
	return &ApiError{StatusCode: status}
}

func (e *ApiError) WithMessage(msg string) *ApiError {
	e.Message = msg
	return e
}

func (e *ApiError) WithApiCode(code string) *ApiError {
	e.ApiCode = code
	return e
}

func (e *ApiError) WithCause(cause *ApiError) *ApiError {
	e.Cause = cause
	return e
}
