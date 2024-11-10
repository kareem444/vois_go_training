package logger

import (
	"log"
)

func Print(message string) {
	log.Println("=====Log=====:", message)
}

func Info(message string) {
	log.Println("=====Info=====:" + message)
}
