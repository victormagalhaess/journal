package log

import (
	"fmt"
	"os"
	"runtime"
)

var (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	white  = "\033[97m"
)

func adjustEnv() {
	if runtime.GOOS == "windows" {
		reset = ""
		red = ""
		green = ""
		yellow = ""
		white = ""
	}
}

func print(message string, color string) {
	adjustEnv()
	fmt.Print(color + message + reset)
}

func printf(message string, color string, obj ...interface{}) {
	adjustEnv()
	fmt.Printf(color+message+reset, obj...)
}

func Print(message string) {
	print(message, white)
}

func Printf(message string, obj ...interface{}) {
	printf(message, white, obj...)
}

func Error(message string) {
	print(message, red)
}

func ErrorF(message string, obj ...interface{}) {
	printf(message, red, obj...)
}

func FatalF(message string, obj ...interface{}) {
	ErrorF(message, obj...)
	os.Exit(1)
}

func Fatal(message string) {
	Error(message)
	os.Exit(1)
}

func Success(message string) {
	print(message, green)
}

func SuccessF(message string, obj ...interface{}) {
	printf(message, green, obj...)
}

func Warning(message string) {
	print(message, yellow)
}

func Warningf(message string, obj ...interface{}) {
	printf(message, yellow, obj...)
}
