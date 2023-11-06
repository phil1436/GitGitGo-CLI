package logger

import (
	"fmt"
	"time"
)

var quietly bool = false

var errorMessages []string = []string{}

// Log with newline
func Log(message string) {
	if quietly {
		return
	}
	fmt.Println(message)
}

// Log without newline (same line)
func LogSL(message string) {
	if quietly {
		return
	}
	fmt.Print(message)
}

// add error message to the buffer
func AddError(message string) {
	message = GetTimestamp() + message

	errorMessages = append(errorMessages, message)
}

func AddErrObj(message string, err error) {
	message += ": " + err.Error()
	message = GetTimestamp() + message

	errorMessages = append(errorMessages, message)
}

func PrintErrors() {
	for _, message := range errorMessages {
		fmt.Println(message)
	}
	errorMessages = []string{}
}

func Quiet() {
	quietly = true
}

func Unquiet() {
	quietly = false
}

func GetTimestamp() string {
	return "[" + time.Now().Format("2006-01-02 15:04:05") + "]"
}
