package devlog

import (
	"log"
)

var (
	isInit    bool
	isEnabled bool
	logger    Logger
)

type Logger interface {
	Println(v ...any)
	Printf(format string, v ...any)
}

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

// TODO control which packages are logged
