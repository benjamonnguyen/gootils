package httpgootils

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
	StatusMessage() string
}

type httpErr struct {
	code int
	msg  string
	err  error
}

func (e httpErr) Error() string {
	return fmt.Sprintf("%s (status code: %d)", e.err, e.code)
}

func (e httpErr) StatusCode() int {
	return e.code
}

func (e httpErr) StatusMessage() string {
	return e.msg
}

func NewHttpError(statusCode int, statusMessage string, errorMessage string) HttpError {
	return httpErr{
		code: statusCode,
		msg:  statusMessage,
		err:  errors.New(errorMessage),
	}
}

func HttpErrorFromErr(err error, statusMessage string) HttpError {
	if err == nil {
		return nil
	}

	code := http.StatusInternalServerError
	httperr, _ := err.(HttpError)
	if httperr != nil {
		code = httperr.StatusCode()
	}

	return httpErr{
		code: code,
		msg:  statusMessage,
		err:  err,
	}
}
