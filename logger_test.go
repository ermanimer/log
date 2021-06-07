package log

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

func TestDebug(t *testing.T) {
	// create logger
	output := &bytes.Buffer{}
	l := NewLogger(output)

	// set hook function
	l.hookFunction = func(prefix, message string) {
		t.Logf("hookFunction called (prefix: %s, message: %s)", prefix, message)
	}

	//define prefix, message and expected
	prefix := debugPrefix
	message := "test"
	expected := fmt.Sprintf("%s %s %s\n", time.Now().Format(l.timeFormat), prefix, message)

	// log debug message
	l.Debug(message)

	// compare output and expected
	if bytes.Compare(output.Bytes(), []byte(expected)) != 0 {
		t.Error("output doesn't match expected")
	}
}

func TestDebugf(t *testing.T) {
	// create logger
	output := &bytes.Buffer{}
	l := NewLogger(output)

	// set hook function
	l.hookFunction = func(prefix, message string) {
		t.Logf("hookFunction called (prefix: %s, message: %s)", prefix, message)
	}

	//define prefix, message and expected
	prefix := debugPrefix
	format := "%s"
	message := "test"
	formattedMessage := fmt.Sprintf(format, message)
	expected := fmt.Sprintf("%s %s %s\n", time.Now().Format(l.timeFormat), prefix, formattedMessage)

	// log debug message
	l.Debugf(format, message)

	// compare output and expected
	if bytes.Compare(output.Bytes(), []byte(expected)) != 0 {
		t.Error("output doesn't match expected")
	}
}

func TestInfo(t *testing.T) {
	// create logger
	output := &bytes.Buffer{}
	l := NewLogger(output)

	// set hook function
	l.hookFunction = func(prefix, message string) {
		t.Logf("hookFunction called (prefix: %s, message: %s)", prefix, message)
	}

	//define prefix, message and expected
	prefix := infoPrefix
	message := "test"
	expected := fmt.Sprintf("%s %s %s\n", time.Now().Format(l.timeFormat), prefix, message)

	// log debug message
	l.Info(message)

	// compare output and expected
	if bytes.Compare(output.Bytes(), []byte(expected)) != 0 {
		t.Error("output doesn't match expected")
	}
}

func TestInfof(t *testing.T) {
	// create logger
	output := &bytes.Buffer{}
	l := NewLogger(output)

	// set hook function
	l.hookFunction = func(prefix, message string) {
		t.Logf("hookFunction called (prefix: %s, message: %s)", prefix, message)
	}

	//define prefix, message and expected
	prefix := infoPrefix
	format := "%s"
	message := "test"
	formattedMessage := fmt.Sprintf(format, message)
	expected := fmt.Sprintf("%s %s %s\n", time.Now().Format(l.timeFormat), prefix, formattedMessage)

	// log debug message
	l.Infof(format, message)

	// compare output and expected
	if bytes.Compare(output.Bytes(), []byte(expected)) != 0 {
		t.Error("output doesn't match expected")
	}
}

func TestWarning(t *testing.T) {
	// create logger
	output := &bytes.Buffer{}
	l := NewLogger(output)

	// set hook function
	l.hookFunction = func(prefix, message string) {
		t.Logf("hookFunction called (prefix: %s, message: %s)", prefix, message)
	}

	//define prefix, message and expected
	prefix := warningPrefix
	message := "test"
	expected := fmt.Sprintf("%s %s %s\n", time.Now().Format(l.timeFormat), prefix, message)

	// log debug message
	l.Warning(message)

	// compare output and expected
	if bytes.Compare(output.Bytes(), []byte(expected)) != 0 {
		t.Error("output doesn't match expected")
	}
}

func TestWarningf(t *testing.T) {
	// create logger
	output := &bytes.Buffer{}
	l := NewLogger(output)

	// set hook function
	l.hookFunction = func(prefix, message string) {
		t.Logf("hookFunction called (prefix: %s, message: %s)", prefix, message)
	}

	//define prefix, message and expected
	prefix := warningPrefix
	format := "%s"
	message := "test"
	formattedMessage := fmt.Sprintf(format, message)
	expected := fmt.Sprintf("%s %s %s\n", time.Now().Format(l.timeFormat), prefix, formattedMessage)

	// log debug message
	l.Warningf(format, message)

	// compare output and expected
	if bytes.Compare(output.Bytes(), []byte(expected)) != 0 {
		t.Error("output doesn't match expected")
	}
}

func TestError(t *testing.T) {
	// create logger
	output := &bytes.Buffer{}
	l := NewLogger(output)

	// set hook function
	l.hookFunction = func(prefix, message string) {
		t.Logf("hookFunction called (prefix: %s, message: %s)", prefix, message)
	}

	//define prefix, message and expected
	prefix := errorPrefix
	message := "test"
	expected := fmt.Sprintf("%s %s %s\n", time.Now().Format(l.timeFormat), prefix, message)

	// log debug message
	l.Error(message)

	// compare output and expected
	if bytes.Compare(output.Bytes(), []byte(expected)) != 0 {
		t.Error("output doesn't match expected")
	}
}

func TestErrorf(t *testing.T) {
	// create logger
	output := &bytes.Buffer{}
	l := NewLogger(output)

	// set hook function
	l.hookFunction = func(prefix, message string) {
		t.Logf("hookFunction called (prefix: %s, message: %s)", prefix, message)
	}

	//define prefix, message and expected
	prefix := errorPrefix
	format := "%s"
	message := "test"
	formattedMessage := fmt.Sprintf(format, message)
	expected := fmt.Sprintf("%s %s %s\n", time.Now().Format(l.timeFormat), prefix, formattedMessage)

	// log debug message
	l.Errorf(format, message)

	// compare output and expected
	if bytes.Compare(output.Bytes(), []byte(expected)) != 0 {
		t.Error("output doesn't match expected")
	}
}

func TestFatal(t *testing.T) {
	// create logger
	output := &bytes.Buffer{}
	l := NewLogger(output)

	// set hook function
	l.hookFunction = func(prefix, message string) {
		t.Logf("hookFunction called (prefix: %s, message: %s)", prefix, message)
	}

	// change default exit function
	l.exitFunction = func() { t.Log("exitFunction called") }

	//define prefix, message and expected
	prefix := fatalPrefix
	message := "test"
	expected := fmt.Sprintf("%s %s %s\n", time.Now().Format(l.timeFormat), prefix, message)

	// log debug message
	l.Fatal(message)

	// compare output and expected
	if bytes.Compare(output.Bytes(), []byte(expected)) != 0 {
		t.Error("output doesn't match expected")
	}
}

func TestFatalf(t *testing.T) {
	// create logger
	output := &bytes.Buffer{}
	l := NewLogger(output)

	// set hook function
	l.hookFunction = func(prefix, message string) {
		t.Logf("hookFunction called (prefix: %s, message: %s)", prefix, message)
	}

	// change default exit function
	l.exitFunction = func() { t.Log("exitFunction called") }

	//define prefix, message and expected
	prefix := fatalPrefix
	format := "%s"
	message := "test"
	formattedMessage := fmt.Sprintf(format, message)
	expected := fmt.Sprintf("%s %s %s\n", time.Now().Format(l.timeFormat), prefix, formattedMessage)

	// log debug message
	l.Fatalf(format, message)

	// compare output and expected
	if bytes.Compare(output.Bytes(), []byte(expected)) != 0 {
		t.Error("output doesn't match expected")
	}
}

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

func BenchmarkInfo(b *testing.B) {
	// create logger
	l := NewLogger(ioutil.Discard)

	// start benchmark test
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Info("test")
		}
	})
}

func BenchmarkInfof(b *testing.B) {
	// create logger
	l := NewLogger(ioutil.Discard)

	// start benchmark test
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Infof("%s", "test")
		}
	})
}

func BenchmarkWarning(b *testing.B) {
	// create logger
	l := NewLogger(ioutil.Discard)

	// start benchmark test
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Warning("test")
		}
	})
}

func BenchmarkWarningf(b *testing.B) {
	// create logger
	l := NewLogger(ioutil.Discard)

	// start benchmark test
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Warningf("%s", "test")
		}
	})
}

func BenchmarkError(b *testing.B) {
	// create logger
	l := NewLogger(ioutil.Discard)

	// start benchmark test
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Error("test")
		}
	})
}

func BenchmarkErrorf(b *testing.B) {
	// create logger
	l := NewLogger(ioutil.Discard)

	// start benchmark test
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Errorf("%s", "test")
		}
	})
}

func BenchmarkFatal(b *testing.B) {
	// create logger
	l := NewLogger(ioutil.Discard)

	// disable exit function
	l.exitFunction = func() {}

	// start benchmark test
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Fatal("test")
		}
	})
}

func BenchmarkFatalf(b *testing.B) {
	// create logger
	l := NewLogger(ioutil.Discard)

	// disable exit function
	l.exitFunction = func() {}

	// start benchmark test
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Fatalf("%s", "test")
		}
	})
}
