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

// logging levels
const (
	DebugLevel = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
)

// Logger represents logger
type Logger struct {
	mutex               *sync.Mutex
	output              io.Writer
	internalErrorOutput io.Writer
	buffer              []byte
	loggingLevel        int
	timeFormat          string
	debugPrefix         string
	infoPrefix          string
	warningPrefix       string
	errorPrefix         string
	fatalPrefix         string
	exitFunction        func()
	hookFunction        func(prefix, message string)
}

// NewLogger creates and returns a new logger instance with given output and default parameters
func NewLogger() *Logger {
	l := &Logger{
		mutex:               &sync.Mutex{},
		output:              os.Stdout,
		internalErrorOutput: os.Stderr,
		loggingLevel:        DebugLevel,
		timeFormat:          time.RFC3339,
		debugPrefix:         debugPrefix,
		infoPrefix:          infoPrefix,
		warningPrefix:       warningPrefix,
		errorPrefix:         errorPrefix,
		fatalPrefix:         fatalPrefix,
		exitFunction:        func() { os.Exit(1) },
		hookFunction:        func(prefix, message string) {},
	}
	return l
}

// SetOutput sets output of logger
func (l *Logger) SetOutput(output io.Writer) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.output = output
}

// SetInternalErrorOutput sets internal error output of logger
func (l *Logger) SetInternalErrorOutput(internalErrorOutput io.Writer) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.internalErrorOutput = internalErrorOutput
}

// SetLoggingLevel sets logging level of logger
func (l *Logger) SetLoggingLevel(loggingLevel int) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.loggingLevel = loggingLevel
}

// SetTimeFormat sets time format of logger
func (l *Logger) SetTimeFormat(timeFormat string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.timeFormat = timeFormat
}

// SetDebugPrefix sets debug prefix of logger
func (l *Logger) SetDebugPrefix(debugPrefix string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.debugPrefix = debugPrefix
}

// SetInfoPrefix sets info prefix of logger
func (l *Logger) SetInfoPrefix(infoPrefix string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.infoPrefix = infoPrefix
}

// SetWarningPrefix sets warning prefix of logger
func (l *Logger) SetWarningPrefix(warningPrefix string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.warningPrefix = warningPrefix
}

// SetErrorPrefix sets error prefix of logger
func (l *Logger) SetErrorPrefix(errorPrefix string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.errorPrefix = errorPrefix
}

// SetFatalPrefix sets fatal prefix of logger
func (l *Logger) SetFatalPrefix(fatalPrefix string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.fatalPrefix = fatalPrefix
}

// SetHookFunction sets hook function
func (l *Logger) SetHookFunction(hookFunction func(prefix, message string)) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.hookFunction = hookFunction
}

// Debug logs a message with debug prefix
func (l *Logger) Debug(values ...interface{}) {
	if l.loggingLevel > DebugLevel {
		return
	}

	l.log(l.debugPrefix, fmt.Sprint(values...))
}

// Debugf logs a formatted message with debug prefix
func (l *Logger) Debugf(format string, values ...interface{}) {
	if l.loggingLevel > DebugLevel {
		return
	}

	l.log(l.debugPrefix, fmt.Sprintf(format, values...))
}

// Info logs a message with info prefix
func (l *Logger) Info(values ...interface{}) {
	if l.loggingLevel > InfoLevel {
		return
	}
	l.log(l.infoPrefix, fmt.Sprint(values...))
}

// Infof logs a formatted message with info prefix
func (l *Logger) Infof(format string, values ...interface{}) {
	if l.loggingLevel > InfoLevel {
		return
	}

	l.log(l.infoPrefix, fmt.Sprintf(format, values...))
}

// Warning logs a message with warning prefix
func (l *Logger) Warning(values ...interface{}) {
	if l.loggingLevel > WarningLevel {
		return
	}
	l.log(l.warningPrefix, fmt.Sprint(values...))
}

// Warningf logs a formatted message with warning prefix
func (l *Logger) Warningf(format string, values ...interface{}) {
	if l.loggingLevel > WarningLevel {
		return
	}

	l.log(l.warningPrefix, fmt.Sprintf(format, values...))
}

// Error logs a message with error prefix
func (l *Logger) Error(values ...interface{}) {
	if l.loggingLevel > ErrorLevel {
		return
	}
	l.log(l.errorPrefix, fmt.Sprint(values...))
}

// Errorf logs a formatted message with error prefix
func (l *Logger) Errorf(format string, values ...interface{}) {
	if l.loggingLevel > ErrorLevel {
		return
	}

	l.log(l.errorPrefix, fmt.Sprintf(format, values...))
}

// Fatal logs a message with fatal prefix and calls os.Exit(1)
func (l *Logger) Fatal(values ...interface{}) {
	if l.loggingLevel > FatalLevel {
		return
	}

	l.log(l.fatalPrefix, fmt.Sprint(values...))

	l.exitFunction()
}

// Fatalf logs a formatted message with fatal prefix and calls os.Exit(1)
func (l *Logger) Fatalf(format string, values ...interface{}) {
	if l.loggingLevel > FatalLevel {
		return
	}

	l.log(l.fatalPrefix, fmt.Sprintf(format, values...))

	l.exitFunction()
}

func (l *Logger) log(prefix, message string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.buffer = l.buffer[:0]

	l.buffer = time.Now().AppendFormat(l.buffer, l.timeFormat)
	l.buffer = append(l.buffer, " "...)

	l.buffer = append(l.buffer, prefix...)
	l.buffer = append(l.buffer, " "...)

	l.buffer = append(l.buffer, message...)

	l.buffer = append(l.buffer, "\n"...)

	_, err := l.output.Write(l.buffer)
	if err != nil {
		fmt.Fprintf(l.internalErrorOutput, "writing to log's output failed, %s", err.Error())
	}

	l.hookFunction(prefix, message)
}
