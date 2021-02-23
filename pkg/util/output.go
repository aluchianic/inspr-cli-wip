package util

import (
	"fmt"
	"io"
	"os"
)

var (
	DEBUG = os.Getenv("DEBUG")
	// Stdout points to the output buffer to send screen output
	Stdout io.Writer = os.Stdout
	// Stderr points to the output buffer to send errors to the screen
	Stderr io.Writer = os.Stderr
)

func Errorf(format string, args ...interface{}) {
	f := format + "\n"
	fmt.Fprintf(Stderr, f, args...)
	os.Exit(0)
}

func Infof(format string, args ...interface{}) {
	f := format + "\n"
	fmt.Fprintf(Stdout, f, args...)
}

func Debugf(format string, args ...interface{}) {
	if DEBUG != "" {
		time := "0/0/0:0-0-1"
		f := "[" + time + "]" + format + "\n"
		fmt.Fprintf(Stdout, f, args...)
	}
}
