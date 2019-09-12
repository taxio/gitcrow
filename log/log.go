package log

import (
	"fmt"
	"io"
	"os"
)

type Logger struct {
	Verbose bool
	out     io.Writer
}

var logger Logger

func L() *Logger {
	return &logger
}

func init() {
	logger = Logger{
		Verbose: false,
		out:     os.Stdout,
	}
}

func (l *Logger) SetVerbose(verbose bool) {
	l.Verbose = verbose
}

func (l *Logger) Printf(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(l.out, format, a...)
}

func (l *Logger) Printfln(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(l.out, format+"\n", a...)
}

func (l *Logger) Println(msg string) {
	_, _ = fmt.Fprintln(l.out, msg)
}

func (l *Logger) Error(err error) {
	if l.Verbose {
		_, _ = fmt.Fprintf(l.out, "%+v\n", err)
	} else {
		_, _ = fmt.Fprintf(l.out, "%v\n", err)
	}
}
