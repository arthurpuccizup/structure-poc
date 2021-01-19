package errors

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Error interface {
	AddMeta(key, value string) *customError
	AddOperation(operation string) *customError
	Marshal() ([]byte, error)
}

type customError struct {
	ID         uuid.UUID         `json:"id"`
	Title      string            `json:"title"`
	Detail     string            `json:"detail"`
	Meta       map[string]string `json:"meta"`
	Operations []string          `json:"operations"`
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
		ID:     uuid.New(),
		Title:  title,
		Detail: detail,
		Meta: map[string]string{
			"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
		},
		Operations: []string{},
	}
}
