package internal

// Logger log, warn, error with colors
type Logger struct{}

func (l *Logger) Log(msg string) {
	println(msg)
}

func (l *Logger) Warn(msg string) {
	println(msg)
}

func (l *Logger) Error(msg string) {
	println(msg)
}

func NewLogger() *Logger {
	return &Logger{}
}
