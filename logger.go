package smaug

import (
	"log"

	"github.com/davecgh/go-spew/spew"
)

const (
	Debug logger = iota
	Info
	Warn
	Error
)

type logger uint8

func (l logger) Dump(v ...interface{}) {
	if level(l) >= Level {
		spew.Dump(v...)
	}
}

func (l logger) Print(v ...interface{}) {
	if level(l) >= Level {
		log.Print(v...)
	}
}

func (l logger) Printf(format string, v ...interface{}) {
	if level(l) >= Level {
		log.Printf(format, v...)
	}
}

func (l logger) Println(v ...interface{}) {
	if level(l) >= Level {
		log.Println(v...)
	}
}
