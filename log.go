package logging

import (
	"fmt"

	"github.com/digiconvent/logging/internal/color"
	"github.com/digiconvent/logging/internal/format"
)

func (l *Logger) write(level Level, severityColor string, msg any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if level < l.level {
		return
	}

	fmt.Fprintln(l.writer(), format.Line(severityColor, msg, l.NoColor))
}

func (l *Logger) Error(msg any) {
	l.write(LevelError, color.Red, msg)
}

func (l *Logger) Warning(msg any) {
	l.write(LevelWarning, color.Yellow, msg)
}

func (l *Logger) Info(msg any) {
	l.write(LevelInfo, color.Cyan, msg)
}

func (l *Logger) Success(msg any) {
	l.write(LevelInfo, color.Green, msg)
}
