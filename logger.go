package log

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// default prefixes
const (
	debugPrefix   = "debug"
	infoPrefix    = "info"
	warningPrefix = "warning"
	errorPrefix   = "error"
	fatalPrefix   = "fatal"
)

// default time format
const (
	timeFormat = time.RFC3339
)

// logging levels
const (
	DebugLevel = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
)

// logger represents logger
type logger struct {
	mutex         *sync.Mutex
	output        io.Writer
	buffer        []byte
	loggingLevel  int
	timeFormat    string
	debugPrefix   string
	infoPrefix    string
	warningPrefix string
	errorPrefix   string
	fatalPrefix   string
	exitFunction  func()
	hookFunction  func(prefix, message string)
}

// NewLogger creates and returns a new logger instance with given output and default parameters
func NewLogger(output io.Writer) *logger {
	l := &logger{
		mutex:         &sync.Mutex{},
		output:        output,
		loggingLevel:  DebugLevel,
		timeFormat:    timeFormat,
		debugPrefix:   debugPrefix,
		infoPrefix:    infoPrefix,
		warningPrefix: warningPrefix,
		errorPrefix:   errorPrefix,
		fatalPrefix:   fatalPrefix,
		exitFunction:  func() { os.Exit(1) },
		hookFunction:  func(prefix, message string) {},
	}
	return l
}

// LoggingLevel returns logging level of logger
func (l *logger) LoggingLevel() int {
	return l.loggingLevel
}

// SetLoggingLevel sets logging level of logger
func (l *logger) SetLoggingLevel(loggingLevel int) {
	l.loggingLevel = loggingLevel
}

// TimeFormat returns time format of logger
func (l *logger) TimeFormat() string {
	return l.timeFormat
}

// SetTimeFormat sets time format of loggger
func (l *logger) SetTimeFormat(timeFormat string) {
	l.timeFormat = timeFormat
}

// DebugPrefix returns debug prefix of logger
func (l *logger) DebugPrefix() string {
	return l.debugPrefix
}

// SetDebugPrefix sets debug prefix of logger
func (l *logger) SetDebugPrefix(debugPrefix string) {
	l.debugPrefix = debugPrefix
}

// InfoPrefix returns info prefix of logger
func (l *logger) InfoPrefix() string {
	return l.infoPrefix
}

// SetInfoPrefix sets info prefix of logger
func (l *logger) SetInfoPrefix(infoPrefix string) {
	l.infoPrefix = infoPrefix
}

// WarningPrefix returns warning prefix of logger
func (l *logger) WarningPrefix() string {
	return l.warningPrefix
}

// SetWarningPrefix sets warning prefix of logger
func (l *logger) SetWarningPrefix(warningPrefix string) {
	l.warningPrefix = warningPrefix
}

// ErrorPrefix returns error prefix of logger
func (l *logger) ErrorPrefix() string {
	return l.errorPrefix
}

// SetErrorPrefix sets error prefix of logger
func (l *logger) SetErrorPrefix(errorPrefix string) {
	l.errorPrefix = errorPrefix
}

// FatalPrefix returns fatal prefix of logger
func (l *logger) FatalPrefix() string {
	return l.fatalPrefix
}

// SetFatalPrefix sets fatal prefix of logger
func (l *logger) SetFatalPrefix(fatalPrefix string) {
	l.fatalPrefix = fatalPrefix
}

// SetHookFunction sets hook function
func (l *logger) SetHookFunction(hookFunction func(prefix, message string)) {
	l.hookFunction = hookFunction
}

// Debug logs a message with debug prefix
func (l *logger) Debug(values ...interface{}) {
	if l.loggingLevel > DebugLevel {
		return
	}
	l.log(l.debugPrefix, fmt.Sprint(values...))
}

// Debugf logs a formatted message with debug prefix
func (l *logger) Debugf(format string, values ...interface{}) {
	if l.loggingLevel > DebugLevel {
		return
	}
	l.log(l.debugPrefix, fmt.Sprintf(format, values...))
}

// Info logs a message with info prefix
func (l *logger) Info(values ...interface{}) {
	if l.loggingLevel > InfoLevel {
		return
	}
	l.log(l.infoPrefix, fmt.Sprint(values...))
}

// Infof logs a formatted message with info prefix
func (l *logger) Infof(format string, values ...interface{}) {
	if l.loggingLevel > InfoLevel {
		return
	}
	l.log(l.infoPrefix, fmt.Sprintf(format, values...))
}

// Warning logs a message with warning prefix
func (l *logger) Warning(values ...interface{}) {
	if l.loggingLevel > WarningLevel {
		return
	}
	l.log(l.warningPrefix, fmt.Sprint(values...))
}

// Warningf logs a formatted message with warning prefix
func (l *logger) Warningf(format string, values ...interface{}) {
	if l.loggingLevel > WarningLevel {
		return
	}
	l.log(l.warningPrefix, fmt.Sprintf(format, values...))
}

// Error logs a message with error prefix
func (l *logger) Error(values ...interface{}) {
	if l.loggingLevel > ErrorLevel {
		return
	}
	l.log(l.errorPrefix, fmt.Sprint(values...))
}

// Errorf logs a formatted message with error prefix
func (l *logger) Errorf(format string, values ...interface{}) {
	if l.loggingLevel > ErrorLevel {
		return
	}
	l.log(l.errorPrefix, fmt.Sprintf(format, values...))
}

// Fatal logs a message with fatal prefix and calls os.Exit(1)
func (l *logger) Fatal(values ...interface{}) {
	if l.loggingLevel > FatalLevel {
		return
	}
	l.log(l.fatalPrefix, fmt.Sprint(values...))
	l.exitFunction()
}

// Fatalf logs a formatted message with fatal prefix and calls os.Exit(1)
func (l *logger) Fatalf(format string, values ...interface{}) {
	if l.loggingLevel > FatalLevel {
		return
	}
	l.log(l.fatalPrefix, fmt.Sprintf(format, values...))
	l.exitFunction()
}

func (l *logger) log(prefix, message string) {
	// synchronization
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// clear buffer
	l.buffer = l.buffer[:0]

	// append timestamp to buffer
	l.buffer = time.Now().AppendFormat(l.buffer, l.timeFormat)
	l.buffer = append(l.buffer, " "...)

	//append prefix to buffer
	l.buffer = append(l.buffer, prefix...)
	l.buffer = append(l.buffer, " "...)

	// append message to buffer
	l.buffer = append(l.buffer, message...)

	// append new line character
	l.buffer = append(l.buffer, "\n"...)

	// write buffer to output
	l.output.Write(l.buffer)

	// call hook function
	l.hookFunction(prefix, message)
}
