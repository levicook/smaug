package smaug

import "log"

var (
	Panic   = log.Panic
	Panicf  = log.Panicf
	Panicln = log.Panicln
)

func PanicIf(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func PanicIff(err error, format string, v ...interface{}) {
	if err != nil {
		log.Panicf(format, v...)
	}
}
