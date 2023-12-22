package devlog_test

import (
	"bytes"
	"io"
	"log"
	"testing"

	"github.com/benjamonnguyen/gootils/devlog"
)

func init() {
	devlog.Init(true, nil)
}

func TestPrintln(t *testing.T) {
	var buf bytes.Buffer
	log.SetFlags(0)
	log.SetOutput(&buf)
	devlog.Print("hello", "world!")

	// log package always adds newline
	const want = "devlog: hello world!\n"
	data, _ := io.ReadAll(&buf)
	got := string(data)
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestPrintf(t *testing.T) {
	var buf bytes.Buffer
	log.SetFlags(0)
	log.SetOutput(&buf)
	devlog.Printf("%s %s", "hello", "world!")

	// log package always adds newline
	const want = "devlog: hello world!\n"
	data, _ := io.ReadAll(&buf)
	got := string(data)
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}
