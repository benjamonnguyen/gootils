package devlog

import "log"

var (
	isEnabled bool
	logger    Logger = log.Default()
)

type Logger interface {
	Println(v ...any)
	Printf(format string, v ...any)
}

func Enable(b bool) {
	isEnabled = b
}

func SetLogger(l Logger) {
	logger = l
}

func Println(v ...any) {
	if isEnabled {
		logger.Println(v...)
	}
}

func Printf(format string, v ...any) {
	if isEnabled {
		logger.Printf(format, v...)
	}
}
