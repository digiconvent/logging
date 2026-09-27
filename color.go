package logging

import "github.com/digiconvent/logging/internal/color"

func (l *Logger) Red(s string) string {
	if l.NoColor {
		return s
	}
	return color.Red + s + color.Reset
}

func (l *Logger) Green(s string) string {
	if l.NoColor {
		return s
	}
	return color.Green + s + color.Reset
}

func (l *Logger) Disabled(s string) string {
	if l.NoColor {
		return s
	}
	return color.Gray + s + color.Reset
}
