package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
)

var (
	reset                    = "\033[0m"
	red                      = "\033[31m"
	green                    = "\033[32m"
	yellow                   = "\033[33m"
	white                    = "\033[97m"
	Out       io.Writer      = os.Stdout
	Exit      func(code int) = os.Exit
	IsWindows                = runtime.GOOS == "windows"
	buffer                   = ""
)

func adjustEnv() {
	if IsWindows {
		reset = ""
		red = ""
		green = ""
		yellow = ""
		white = ""
	}
}

func GetBuffer() string {
	return buffer
}

func print(message string, color string) {
	buffer = color + message + reset
	fmt.Fprint(Out, buffer)
}

func printf(message string, color string, obj ...interface{}) {
	print(fmt.Sprintf(message, obj...), color)
}

func Print(message string) {
	adjustEnv()
	print(message, white)
}

func Printf(message string, obj ...interface{}) {
	adjustEnv()
	printf(message, white, obj...)
}

func Error(message string) {
	adjustEnv()
	print(message, red)
}

func ErrorF(message string, obj ...interface{}) {
	adjustEnv()
	printf(message, red, obj...)
}

func FatalF(message string, obj ...interface{}) {
	ErrorF(message, obj...)
	Exit(1)
}

func Fatal(message string) {
	Error(message)
	Exit(1)
}

func Success(message string) {
	adjustEnv()
	print(message, green)
}

func SuccessF(message string, obj ...interface{}) {
	adjustEnv()
	printf(message, green, obj...)
}

func Warning(message string) {
	adjustEnv()
	print(message, yellow)
}

func Warningf(message string, obj ...interface{}) {
	adjustEnv()
	printf(message, yellow, obj...)
}
