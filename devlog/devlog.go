// Package devlog provides a singleton logger to print logs during development.
// When disabled in production environments, print methods are no-op.
package devlog

import (
	"fmt"
	"log"
)

var (
	isInit    bool
	isEnabled bool
	logger    Logger
)

type Logger interface {
	Printf(format string, v ...any)
}

// Init initializes the singleton logger and can only be called once.
// If Logger arg is nil, the default one from the standard log package is used.
func Init(enable bool, l Logger) {
	if isInit {
		panic("devlog: Init already called")
	}
	isInit = true
	isEnabled = enable
	if l != nil {
		logger = l
	} else {
		logger = log.Default()
	}
}

func Print(v ...any) {
	if isEnabled {
		logger.Printf("devlog: %s", fmt.Sprintln(v...))
	}
}

func Printf(format string, v ...any) {
	if isEnabled {
		logger.Printf("devlog: %s", fmt.Sprintf(format, v...))
	}
}

// TODO control which packages are logged
