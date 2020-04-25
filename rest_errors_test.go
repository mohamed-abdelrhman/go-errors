package go_errors

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "Internal Server Error", err.Error())
}

func TestNewUnauthorizedError(t *testing.T) {
	err := NewUnauthorizedError("this is message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "this is message", err.Message())
	assert.EqualValues(t, "UnAuthorized", err.Error())
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("this is message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status())
	assert.EqualValues(t, "this is message", err.Message())
	assert.EqualValues(t, "Not Found", err.Error())
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("this is message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "this is message", err.Message())
	assert.EqualValues(t, "Bad Request", err.Error())
}

func TestNewRestError(t *testing.T) {
	err := NewRestError("This Is Message", 503, "Service Unavailable")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusServiceUnavailable, err.Status())
	assert.EqualValues(t, "This Is Message", err.Message())
	assert.EqualValues(t, "Service Unavailable", err.Error())
}

func TestNewRestErrorFromBytes(t *testing.T) {
	toMarshallErr := NewRestError("This Is Message", 503, "Service Unavailable")
	bytesErr, _ := json.Marshal(toMarshallErr)

	restErr, _ := NewRestErrorFromBytes(bytesErr)
	assert.NotNil(t, restErr)
	assert.EqualValues(t, 503, restErr.Status())
	assert.EqualValues(t, "This Is Message", restErr.Message())
	assert.EqualValues(t, "Service Unavailable", restErr.Error())
}

func TestNewUnprocessableEntity(t *testing.T) {
	var errorMessages []interface{}
	errorMessages = append(errorMessages, errors.New("ERROR ONE"))
	errorMessages = append(errorMessages, errors.New("ANOTHER ONE"))
	causes := []string{"ERROR ONE", "ANOTHER ONE"}
	err := NewUnprocessableEntity("Validation error", causes)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnprocessableEntity, err.Status())
	assert.EqualValues(t, "Validation error", err.Message())
	assert.EqualValues(t, "Unprocessable Entity", err.Error())
	assert.Contains(t, causes, "ERROR ONE")
	assert.Contains(t, causes, "ANOTHER ONE")
}
