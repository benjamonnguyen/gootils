package httputil

import (
	"errors"
	"fmt"
	"net/http"
)

// HttpError is a custom error type that provides status code
// and message for client-side error handling.
type HttpError interface {
	error
	StatusCode() int
	Status() string
}

type httpErr struct {
	code   int
	status string
	err    error
}

func (e httpErr) Error() string {
	return fmt.Sprintf("%s (status code: %d)", e.err, e.code)
}

func (e httpErr) StatusCode() int {
	return e.code
}

func (e httpErr) Status() string {
	return e.status
}

func NewHttpError(statusCode int, status string, errorMessage string) HttpError {
	return httpErr{
		code:   statusCode,
		status: status,
		err:    errors.New(errorMessage),
	}
}

func HttpErrorFromErr(err error) HttpError {
	if err == nil {
		return nil
	}

	code := http.StatusInternalServerError
	status := ""
	httperr, _ := err.(HttpError)
	if httperr != nil {
		code = httperr.StatusCode()
		status = httperr.Status()
	}

	return httpErr{
		code:   code,
		status: status,
		err:    err,
	}
}
