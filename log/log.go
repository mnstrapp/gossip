package log

import (
	"fmt"
	"os"
	"time"
)

func Info(msg string) {
	fmt.Fprintf(os.Stdout, "[INFO] %s - %s\n", time.Now(), msg)
}

func Infof(format string, args ...any) {
	Info(fmt.Sprintf(format, args...))
}

func Error(err error) {
	fmt.Fprintf(os.Stderr, "[ERROR] %s - %s\n", time.Now(), err)
}

func Errorf(format string, args ...any) {
	Error(fmt.Errorf(format, args...))
}
