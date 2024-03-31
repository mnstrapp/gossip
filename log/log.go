package log

import (
	"fmt"
	"os"
	"time"
)

func Info(msg string) {
	fmt.Fprintf(os.Stdout, "[INFO] %s - %s\n", time.Now(), msg)
}

func Error(err error) {
	fmt.Fprintf(os.Stderr, "[Error] %s - %s\n", time.Now(), err)
}
