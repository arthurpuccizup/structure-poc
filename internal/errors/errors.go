package errors

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Error interface {
	Error() sensitiveError
	SensitiveError() customError
	AddMeta(key, value string) *sensitiveError
	AddOperation(operation string) *sensitiveError
	Marshal() ([]byte, error)
}

type sensitiveError struct {
	customError
	Detail     string   `json:"detail"`
	Operations []string `json:"operations"`
}

type customError struct {
	ID    uuid.UUID         `json:"id"`
	Title string            `json:"title"`
	Meta  map[string]string `json:"meta"`
}

func (e *sensitiveError) Error() sensitiveError {
	return *e
}

func (e *sensitiveError) SensitiveError() customError {
	return e.customError
}

func (e *sensitiveError) AddMeta(key, value string) *sensitiveError {
	e.Meta[key] = value
	return e
}

func (e *sensitiveError) AddOperation(operation string) *sensitiveError {
	e.Operations = append(e.Operations, operation)
	return e
}

func (e *sensitiveError) Marshal() ([]byte, error) {
	return json.Marshal(&e)
}

func New(title, detail string) *sensitiveError {
	return &sensitiveError{
		customError: customError{
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
