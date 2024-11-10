package logger

import (
	"log"
)

func Error(err error) {
	if err != nil {
		log.Printf("=====Error=====: %v", err)
	}
}

func ErrorWithMessage(message string, err error) {
	if err != nil {
		log.Printf("=====Error=====: %s: %v", message, err)
	}
}

func Fatal(err error) {
	if err != nil {
		log.Fatalf("=====Fatal=====: %v", err)
	}
}

func FatalWithMessage(message string, err error) {
	if err != nil {
		log.Fatalf("=====Fatal=====: %s: %v", message, err)
	}
}