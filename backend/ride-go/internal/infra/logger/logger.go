package logger

type Logger struct{}

func (l *Logger) Log(message string) {
	println(message)
}
