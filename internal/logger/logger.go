package logger

import (
	"log"
)

var debug bool

func Init(verbose bool) {
	debug = verbose
}

func Infof(format string, v ...interface{}) {
	log.Printf("[I] "+format, v...)
}

func Warnf(format string, v ...interface{}) {
	log.Printf("[W] "+format, v...)
}

func Errorf(format string, v ...interface{}) {
	log.Printf("[E] "+format, v...)
}

func Debugf(format string, v ...interface{}) {
	if debug {
		log.Printf("[D] "+format, v...)
	}
}
