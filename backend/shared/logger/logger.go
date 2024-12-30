package logger

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "SERVICE: ", log.Ldate|log.Ltime|log.Lshortfile)

func Info(message string) {
	Logger.Println("INFO: " + message)
}

func Error(message string) {
	Logger.Println("ERROR: " + message)
}
