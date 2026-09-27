package logging

import (
	"io"
	"os"
	"sync"
)

type Level int

const (
	LevelInfo Level = iota
	LevelWarning
	LevelError
	LevelSilent
)

type Logger struct {
	Output  io.Writer
	NoColor bool

	mu    sync.Mutex
	level Level
}

func New(output io.Writer) *Logger {
	return &Logger{Output: output}
}

func (l *Logger) writer() io.Writer {
	if l.Output == nil {
		return os.Stdout
	}
	return l.Output
}
