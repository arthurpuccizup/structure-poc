package errors

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Error interface {
	Error() customError
	SensitiveError() sensitiveError
	AddMeta(key, value string) *customError
	AddOperation(operation string) *customError
	Marshal() ([]byte, error)
}

type customError struct {
	sensitiveError
	Detail     string   `json:"detail"`
	Operations []string `json:"operations"`
}

type sensitiveError struct {
	ID    uuid.UUID         `json:"id"`
	Title string            `json:"title"`
	Meta  map[string]string `json:"meta"`
}

func (e *customError) Error() customError {
	return *e
}

func (e *customError) SensitiveError() sensitiveError {
	return e.sensitiveError
}

func (e *customError) AddMeta(key, value string) *customError {
	e.Meta[key] = value
	return e
}

func (e *customError) AddOperation(operation string) *customError {
	e.Operations = append(e.Operations, operation)
	return e
}

func (e *customError) Marshal() ([]byte, error) {
	return json.Marshal(&e)
}

func New(title, detail string) Error {
	return &customError{
		sensitiveError: sensitiveError{
			ID:    uuid.New(),
			Title: title,
			Meta: map[string]string{
				"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
			},
		},
		Detail:     detail,
		Operations: []string{},
	}
}
