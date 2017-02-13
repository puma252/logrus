package logrus

import (
	"io"
)

var (
	// std is the name of the standard logger in stdlib `log`
	std = New()
)

func StandardLogger() *Logger {
	return std
}

// SetOutput sets the standard logger output.
func SetOutput(out io.Writer) {
	std.mu.Lock()
	defer std.mu.Unlock()
	std.Out = out
}

// SetFormatter sets the standard logger formatter.
func SetFormatter(formatter Formatter) {
	std.mu.Lock()
	defer std.mu.Unlock()
	std.Formatter = formatter
}

// SetLevel sets the standard logger level.
func SetLevel(level Level) {
	std.mu.Lock()
	defer std.mu.Unlock()
	std.Level = level
}

// GetLevel returns the standard logger level.
func GetLevel() Level {
	std.mu.Lock()
	defer std.mu.Unlock()
	return std.Level
}

// AddHook adds a hook to the standard logger hooks.
func AddHook(hook Hook) {
	std.mu.Lock()
	defer std.mu.Unlock()
	std.Hooks.Add(hook)
}

// WithError creates an entry from the standard logger and adds an error to it, using the value defined in ErrorKey as key.
func WithError(err error) *Entry {
	return NewEntry(std).WithField(ErrorKey, err)
}

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithField(key string, value interface{}) *Entry {
	return NewEntry(std).WithField(key, value)
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithFields(fields Fields) *Entry {
	return NewEntry(std).WithFields(fields)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	if std.Level >= DebugLevel {
		newEntry(std).Debug(args...)
	}
}

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
	if std.Level >= DebugLevel {
		newEntry(std).Info(args...)
	}
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	if std.Level >= InfoLevel {
		newEntry(std).Info(args...)
	}
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	if std.Level >= WarnLevel {
		newEntry(std).Warn(args...)
	}
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...interface{}) {
	if std.Level >= WarnLevel {
		newEntry(std).Warn(args...)
	}
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	if std.Level >= ErrorLevel {
		newEntry(std).Error(args...)
	}
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	if std.Level >= PanicLevel {
		newEntry(std).Panic(args...)
	}
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(args ...interface{}) {
	if std.Level >= FatalLevel {
		newEntry(std).Fatal(args...)
	}
	//os.Exit(1)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	if std.Level >= DebugLevel {
		newEntry(std).Debugf(format, args...)
	}
}

// Printf logs a message at level Info on the standard logger.
func Printf(format string, args ...interface{}) {
	newEntry(std).Printf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	if std.Level >= InfoLevel {
		newEntry(std).Infof(format, args...)
	}
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	if std.Level >= WarnLevel {
		newEntry(std).Warnf(format, args...)
	}
}

// Warningf logs a message at level Warn on the standard logger.
func Warningf(format string, args ...interface{}) {
	if std.Level >= WarnLevel {
		newEntry(std).Warnf(format, args...)
	}
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	if std.Level >= ErrorLevel {
		newEntry(std).Errorf(format, args...)
	}
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(format string, args ...interface{}) {
	if std.Level >= PanicLevel {
		newEntry(std).Panicf(fomat, args...)
	}
}

// Fatalf logs a message at level Fatal on the standard logger.
func Fatalf(format string, args ...interface{}) {
	if std.Level >= FatalLevel {
		newEntry(std).Fatalf(format, args...)
	}
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(args ...interface{}) {
	if std.Level >= DebugLevel {
		newEntry(std).Debugln(args...)
	}
}

// Println logs a message at level Info on the standard logger.
func Println(args ...interface{}) {
	newEntry(std).Println(args...)
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(args ...interface{}) {
	if std.Level >= InfoLevel {
		newEntry(std).Infoln(args...)
	}
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(args ...interface{}) {
	if std.Level >= WarnLevel {
		newEntry(std).Warnln(args...)
	}
}

// Warningln logs a message at level Warn on the standard logger.
func Warningln(args ...interface{}) {
	if std.Level >= WarnLevel {
		newEntry(std).Warningln(args...)
	}
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(args ...interface{}) {
	if std.Level >= ErrorLevel {
		newEntry(std).Errorln(args...)
	}
}

// Panicln logs a message at level Panic on the standard logger.
func Panicln(args ...interface{}) {
	if std.Level >= PanicLevel {
		newEntry(std).Panicln(args...)
	}
}

// Fatalln logs a message at level Fatal on the standard logger.
func Fatalln(args ...interface{}) {
	if std.Level >= FatalLevel {
		newEntry(std).Fatalln(args...)
	}
}
