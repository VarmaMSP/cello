package model

import "net/http"

type AppError struct {
	id            string
	comment       string
	detailedError string
	statusCode    int
	params        map[string]interface{}
	retry         bool
}

func (e *AppError) Error() string {
	return e.id + ": " + e.detailedError
}

func (e *AppError) Id(id string) *AppError {
	e.id = id
	return e
}

func (e *AppError) GetId() string {
	return e.id
}

func (e *AppError) Comment(comment string) *AppError {
	e.comment = comment
	return e
}

func (e *AppError) GetComment() string {
	if e.comment != "" {
		return e.comment
	}
	return e.detailedError
}

func (e *AppError) DetailedError(detailedError string) *AppError {
	e.detailedError = detailedError
	return e
}

func (e *AppError) StatusCode(statusCode int) *AppError {
	e.statusCode = statusCode
	return e
}

func (e *AppError) GetStatusCode() int {
	return e.statusCode
}

func (e *AppError) Params(params map[string]interface{}) *AppError {
	e.params = params
	return e
}

func (e *AppError) Retry() *AppError {
	e.retry = true
	return e
}

func (e *AppError) CanRetry() bool {
	return e.retry
}

func NewAppError(where string, detailedError string, statusCode int, params map[string]interface{}) *AppError {
	return &AppError{
		id:            where,
		detailedError: detailedError,
		statusCode:    statusCode,
		params:        params,
	}
}

func New500Error(where string, detailedError string, params map[string]interface{}) *AppError {
	return &AppError{
		id:            where,
		detailedError: detailedError,
		statusCode:    http.StatusInternalServerError,
		params:        params,
	}
}

func New400Error(where string, detailedError string, params map[string]interface{}) *AppError {
	return &AppError{
		id:            where,
		detailedError: detailedError,
		statusCode:    http.StatusBadRequest,
		params:        params,
	}
}

func NewAppErrorC(where string, statusCode int, params map[string]interface{}) func(detailedError string) *AppError {
	return func(detailedError string) *AppError {
		return &AppError{
			id:            where,
			detailedError: detailedError,
			statusCode:    statusCode,
			params:        params,
		}
	}
}
