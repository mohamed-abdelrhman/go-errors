package go_errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []string
}

type restErr struct {
	ErrMessage string   `json:"message"`
	ErrStatus  int      `json:"status"`
	ErrError   string   `json:"error"`
	ErrCauses  []string `json:"causes"`
}

//Assign error methods to the struct
func (e restErr) Error() string {
	return e.ErrError
}

func (e restErr) Message() string {
	return e.ErrMessage
}

func (e restErr) Status() int {
	return e.ErrStatus
}

func (e restErr) Causes() []string {
	return e.ErrCauses
}

//errors

func NewRestError(message string, status int, err string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   err,
	}
}

func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr restErr
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

func NewBadRequestError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "Bad Request",
	}
}

func NewNotFoundError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "Not Found",
	}
}

func NewUnauthorizedError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "UnAuthorized",
	}
}

func NewInternalServerError(message string) RestErr {
	result := restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "Internal Server Error",
	}
	return result
}

func NewUnprocessableEntity(message string, causes []string) restErr {

	result := restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnprocessableEntity,
		ErrError:   "Unprocessable Entity",
		ErrCauses:  causes,
	}

	return result
}
