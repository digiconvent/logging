package logging

func (l *Logger) Mute() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = LevelSilent
}

func (l *Logger) Unmute() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = LevelInfo
}

func (l *Logger) ShowFrom(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}
