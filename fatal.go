package smaug

import "log"

var (
	Fatal   = log.Fatal
	Fatalf  = log.Fatalf
	Fatalln = log.Fatalln
)

func FatalIf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func FatalIff(err error, format string, v ...interface{}) {
	if err != nil {
		log.Fatalf(format, v...)
	}
}
