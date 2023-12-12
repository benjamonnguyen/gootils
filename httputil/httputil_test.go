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
	wantStatusMessage := "msg"
	got := httputil.HttpErrorFromErr(err, wantStatusMessage)
	if got.StatusCode() != wantStatusCode {
		t.Fatalf("got StatusCode %d, want %d", got.StatusCode(), wantStatusCode)
	}
	if got.StatusMessage() != wantStatusMessage {
		t.Fatalf("got StatusMessage %s, want %s", got.StatusMessage(), wantStatusMessage)
	}

	// HttpError
	wantStatusCode = http.StatusBadRequest
	err = httputil.NewHttpError(wantStatusCode, "", "")
	got = httputil.HttpErrorFromErr(err, "")
	if got.StatusCode() != wantStatusCode {
		t.Fatalf("got StatusCode %d, want %d", got.StatusCode(), wantStatusCode)
	}
}
