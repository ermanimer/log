# log

Simple, customizable, leveled and efficient logging in Go

[![Go](https://github.com/ermanimer/log/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/ermanimer/log/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/ermanimer/log)](https://goreportcard.com/report/github.com/ermanimer/log)

# Installation

```bash
go get -u github.com/ermanimer/log
```

# Features

**log** is a simple logging package for Go. 

- You can set output
- You can set internal error output (slog writes Log method's errors to internal error output instead of returning them.)
- You can set time format and prefixes
- You can set logging level to filter out log messages
- You can set a hook function to be called after each log

**log** isn't the fastest logging solution, but **log** is efficient while being customizable. Please see [benchmark tests](https://github.com/ermanimer/log#benchmark-tests).

# Usage

```go
package main

import (
	"os"

	"github.com/ermanimer/log/v2"
)

func main() {
	// create a new logger instance with output and default parameters
	l := log.NewLogger()

	// log a message
	l.Debug("this is a debug message")

	// log a formatted message
	l.Debugf("this is a %s debug message", "formatted")
}
```

**Output:**

```bash
2021-06-07T16:46:26+03:00 debug this is a debug message
2021-06-07T16:46:26+03:00 debug this is a formatted debug message
```

# Logging Levels

 - Debug
 - Info
 - Warning
 - Error
 - Fatal

# Default Parameters:

| Parameter | Value |
|:----------|:-----:|
|Output|os.Stdout|
|Internal Error Output|os.Stderr|
|Time Format|RFC3339|
|Debug Prefix|debug|
|Info Prefix|info|
|Warning Prefix|warning|
|Error Prefix|error|
|Fatal Prefix|fatal|
|Logging Level|Debug|

# Set Output

```go
l.SetOutput(os.Stderr)
```

# Set Internal Error Output

```go
l.SetOutput(io.Discard)
```

# Set Logging Level

```go
l.SetLoggingLevel(InfoLevel)
```

# Set Time Format

```go
l.SetTimeFormat(time.RFC3339Nano)
```

# Set Prefixes

```go
l.SetDebugPrefix("DEB")
l.SetInfoPrefix("INF")
l.SetWarningPrefix("WAR")
l.SetErrorPrefix("ERR")
l.SetFatalPrefix("FAT")
```

# Set Hook Function:

```go
l.SetHookFunction(func(prefix, message string) {
  //you can filter messages with prefix and capture for Sentry.io
})
```

# Benchmark Tests

**Test Codes:**

```go
func BenchmarkDebug(b *testing.B) {
	// create logger
	l := NewLogger(ioutil.Discard)

	// start benchmark test
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Debug("test")
		}
	})
}

func BenchmarkDebugf(b *testing.B) {
	// create logger
	l := NewLogger(ioutil.Discard)

	// start benchmark test
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Debugf("%s", "test")
		}
	})
}
```
**Results:**

| Function | Time | Bytes Allocated | Objects Allocated |
|:---------|:----:|:---------------:|:-----------------:|
|Debug|410.7 ns/op|4 B/op|1 allocs/op|
|Debugf|408.7 ns/op|4 B/op|1 allocs/op|
|Info|404.0 ns/op|4 B/op|1 allocs/op|
|Infof|403.9 ns/op|4 B/op|1 allocs/op|
|Warning|407.0 ns/op|4 B/op|1 allocs/op|
|Warningf|409.4 ns/op|4 B/op|1 allocs/op|
|Error|404.6 ns/op|4 B/op|1 allocs/op|
|Errorf|406.2 ns/op|4 B/op|1 allocs/op|
|Fatal|402.1 ns/op |4 B/op|1 allocs/op|
|Fatalf|406.0 ns/op|4 B/op|1 allocs/op|
