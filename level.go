package smaug

const (
	DEBUG level = iota
	INFO
	WARN
	ERROR
)

var Level level = DEBUG

type level uint8

func (l level) String() (s string) {
	switch l {
	case DEBUG:
		s = "DEBUG"
	case INFO:
		s = "INFO"
	case WARN:
		s = "WARN"
	case ERROR:
		s = "ERROR"
	}
	return
}
