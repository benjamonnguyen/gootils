package httputil_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/benjamonnguyen/gootils/httputil"
)

func TestHttpErrorFromErr(t *testing.T) {
	// error
	err := errors.New("")
	wantStatusCode := http.StatusInternalServerError
	wantStatus := ""
	got := httputil.HttpErrorFromErr(err)
	if got.StatusCode() != wantStatusCode {
		t.Fatalf("got StatusCode %d, want %d", got.StatusCode(), wantStatusCode)
	}
	if got.Status() != wantStatus {
		t.Fatalf("got StatusMessage %s, want %s", got.Status(), wantStatus)
	}

	// HttpError
	wantStatusCode = http.StatusBadRequest
	wantStatus = "msg"
	err = httputil.NewHttpError(wantStatusCode, wantStatus, "")
	got = httputil.HttpErrorFromErr(err)
	if got.StatusCode() != wantStatusCode {
		t.Fatalf("got StatusCode %d, want %d", got.StatusCode(), wantStatusCode)
	}
	if got.Status() != wantStatus {
		t.Fatalf("got StatusMessage %s, want %s", got.Status(), wantStatus)
	}
}
