package logger

// Logger is an interface of log application
type Logger interface {
	// InitLogger created logger instance
	InitLogger()
	// Debug represents a debug log
	Debug(args ...interface{})
	// Info represents a debug log
	Info(args ...interface{})
	// Error represents a debug log
	Error(args ...interface{})
	// Fatal represents a debug log
	Fatal(args ...interface{})
}
