package httpgootils_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/benjamonnguyen/gootils/httpgootils"
)

func TestHttpErrorFromErr(t *testing.T) {
	// error
	err := errors.New("")
	wantStatusCode := http.StatusInternalServerError
	wantStatusMessage := "msg"
	got := httpgootils.HttpErrorFromErr(err, wantStatusMessage)
	if got.StatusCode() != wantStatusCode {
		t.Fatalf("got StatusCode %d, want %d", got.StatusCode(), wantStatusCode)
	}
	if got.StatusMessage() != wantStatusMessage {
		t.Fatalf("got StatusMessage %s, want %s", got.StatusMessage(), wantStatusMessage)
	}

	// HttpError
	wantStatusCode = http.StatusBadRequest
	err = httpgootils.NewHttpError(wantStatusCode, "", "")
	got = httpgootils.HttpErrorFromErr(err, "")
	if got.StatusCode() != wantStatusCode {
		t.Fatalf("got StatusCode %d, want %d", got.StatusCode(), wantStatusCode)
	}
}
