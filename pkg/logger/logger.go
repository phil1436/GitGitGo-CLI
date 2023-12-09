package logger

import (
	"fmt"
	"time"
)

// quiet mode
var quietly bool = false

// the buffer for error messages
var errorMessages []string = []string{}

// Log with newline
func Log(message string) {
	if quietly {
		return
	}
	fmt.Println(message)
}

// Log without newline (in same line)
func LogSL(message string) {
	if quietly {
		return
	}
	fmt.Print(message)
}

// add error message to the buffer
func AddError(message string) {
	message = GetTimestamp() + " " + message

	errorMessages = append(errorMessages, message)
}

// add error message to the buffer
func AddErrObj(message string, err error) {
	message += ": " + err.Error()
	message = GetTimestamp() + " " + message

	errorMessages = append(errorMessages, message)
}

// prints all error messages and clears the buffer
func PrintErrors() {
	for _, message := range errorMessages {
		fmt.Println(message)
	}
	errorMessages = []string{}
}

// returns true if there are any error messages
func IsError() bool {
	return len(errorMessages) > 0
}

// enable quiet mode
func Quiet() {
	quietly = true
}

// disable quiet mode
func Unquiet() {
	quietly = false
}

// returns the current timestamp
func GetTimestamp() string {
	return "[" + time.Now().Format("2006-01-02 15:04:05") + "]"
}
