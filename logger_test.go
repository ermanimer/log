package log

import (
	"bytes"
	"fmt"
	"io"
	"testing"
	"time"
)

func TestDebug(t *testing.T) {
	output := bytes.NewBuffer([]byte{})
	l := NewLogger()
	l.SetOutput(output)

	l.hookFunction = func(prefix, message string) {
		t.Logf("hookFunction called (prefix: %s, message: %s)", prefix, message)
	}

	prefix := debugPrefix
	message := "test"
	expected := fmt.Sprintf("%s %s %s\n", time.Now().Format(l.timeFormat), prefix, message)

	l.Debug(message)

	if bytes.Compare(output.Bytes(), []byte(expected)) != 0 {
		t.Error("output doesn't match expected")
	}
}

func TestDebugf(t *testing.T) {
	output := bytes.NewBuffer([]byte{})
	l := NewLogger()
	l.SetOutput(output)

	l.hookFunction = func(prefix, message string) {
		t.Logf("hookFunction called (prefix: %s, message: %s)", prefix, message)
	}

	prefix := debugPrefix
	format := "%s"
	message := "test"
	formattedMessage := fmt.Sprintf(format, message)
	expected := fmt.Sprintf("%s %s %s\n", time.Now().Format(l.timeFormat), prefix, formattedMessage)

	l.Debugf(format, message)

	if bytes.Compare(output.Bytes(), []byte(expected)) != 0 {
		t.Error("output doesn't match expected")
	}
}

func TestInfo(t *testing.T) {
	output := bytes.NewBuffer([]byte{})
	l := NewLogger()
	l.SetOutput(output)

	l.hookFunction = func(prefix, message string) {
		t.Logf("hookFunction called (prefix: %s, message: %s)", prefix, message)
	}

	prefix := infoPrefix
	message := "test"
	expected := fmt.Sprintf("%s %s %s\n", time.Now().Format(l.timeFormat), prefix, message)

	l.Info(message)

	if bytes.Compare(output.Bytes(), []byte(expected)) != 0 {
		t.Error("output doesn't match expected")
	}
}

func TestInfof(t *testing.T) {
	output := bytes.NewBuffer([]byte{})
	l := NewLogger()
	l.SetOutput(output)

	l.hookFunction = func(prefix, message string) {
		t.Logf("hookFunction called (prefix: %s, message: %s)", prefix, message)
	}

	prefix := infoPrefix
	format := "%s"
	message := "test"
	formattedMessage := fmt.Sprintf(format, message)
	expected := fmt.Sprintf("%s %s %s\n", time.Now().Format(l.timeFormat), prefix, formattedMessage)

	l.Infof(format, message)

	if bytes.Compare(output.Bytes(), []byte(expected)) != 0 {
		t.Error("output doesn't match expected")
	}
}

func TestWarning(t *testing.T) {
	output := bytes.NewBuffer([]byte{})
	l := NewLogger()
	l.SetOutput(output)

	l.hookFunction = func(prefix, message string) {
		t.Logf("hookFunction called (prefix: %s, message: %s)", prefix, message)
	}

	prefix := warningPrefix
	message := "test"
	expected := fmt.Sprintf("%s %s %s\n", time.Now().Format(l.timeFormat), prefix, message)

	l.Warning(message)

	if bytes.Compare(output.Bytes(), []byte(expected)) != 0 {
		t.Error("output doesn't match expected")
	}
}

func TestWarningf(t *testing.T) {
	output := bytes.NewBuffer([]byte{})
	l := NewLogger()
	l.SetOutput(output)

	l.hookFunction = func(prefix, message string) {
		t.Logf("hookFunction called (prefix: %s, message: %s)", prefix, message)
	}

	prefix := warningPrefix
	format := "%s"
	message := "test"
	formattedMessage := fmt.Sprintf(format, message)
	expected := fmt.Sprintf("%s %s %s\n", time.Now().Format(l.timeFormat), prefix, formattedMessage)

	l.Warningf(format, message)

	if bytes.Compare(output.Bytes(), []byte(expected)) != 0 {
		t.Error("output doesn't match expected")
	}
}

func TestError(t *testing.T) {
	output := bytes.NewBuffer([]byte{})
	l := NewLogger()
	l.SetOutput(output)

	l.hookFunction = func(prefix, message string) {
		t.Logf("hookFunction called (prefix: %s, message: %s)", prefix, message)
	}

	prefix := errorPrefix
	message := "test"
	expected := fmt.Sprintf("%s %s %s\n", time.Now().Format(l.timeFormat), prefix, message)

	l.Error(message)

	if bytes.Compare(output.Bytes(), []byte(expected)) != 0 {
		t.Error("output doesn't match expected")
	}
}

func TestErrorf(t *testing.T) {
	output := bytes.NewBuffer([]byte{})
	l := NewLogger()
	l.SetOutput(output)

	l.hookFunction = func(prefix, message string) {
		t.Logf("hookFunction called (prefix: %s, message: %s)", prefix, message)
	}

	prefix := errorPrefix
	format := "%s"
	message := "test"
	formattedMessage := fmt.Sprintf(format, message)
	expected := fmt.Sprintf("%s %s %s\n", time.Now().Format(l.timeFormat), prefix, formattedMessage)

	l.Errorf(format, message)

	if bytes.Compare(output.Bytes(), []byte(expected)) != 0 {
		t.Error("output doesn't match expected")
	}
}

func TestFatal(t *testing.T) {
	output := bytes.NewBuffer([]byte{})
	l := NewLogger()
	l.SetOutput(output)

	l.hookFunction = func(prefix, message string) {
		t.Logf("hookFunction called (prefix: %s, message: %s)", prefix, message)
	}

	l.exitFunction = func() { t.Log("exitFunction called") }

	prefix := fatalPrefix
	message := "test"
	expected := fmt.Sprintf("%s %s %s\n", time.Now().Format(l.timeFormat), prefix, message)

	l.Fatal(message)

	if bytes.Compare(output.Bytes(), []byte(expected)) != 0 {
		t.Error("output doesn't match expected")
	}
}

func TestFatalf(t *testing.T) {
	output := bytes.NewBuffer([]byte{})
	l := NewLogger()
	l.SetOutput(output)

	l.hookFunction = func(prefix, message string) {
		t.Logf("hookFunction called (prefix: %s, message: %s)", prefix, message)
	}

	l.exitFunction = func() { t.Log("exitFunction called") }

	prefix := fatalPrefix
	format := "%s"
	message := "test"
	formattedMessage := fmt.Sprintf(format, message)
	expected := fmt.Sprintf("%s %s %s\n", time.Now().Format(l.timeFormat), prefix, formattedMessage)

	l.Fatalf(format, message)

	if bytes.Compare(output.Bytes(), []byte(expected)) != 0 {
		t.Error("output doesn't match expected")
	}
}

func BenchmarkDebug(b *testing.B) {
	l := NewLogger()
	l.SetOutput(io.Discard)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Debug("test")
		}
	})
}

func BenchmarkDebugf(b *testing.B) {
	l := NewLogger()
	l.SetOutput(io.Discard)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Debugf("%s", "test")
		}
	})
}

func BenchmarkInfo(b *testing.B) {
	l := NewLogger()
	l.SetOutput(io.Discard)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Info("test")
		}
	})
}

func BenchmarkInfof(b *testing.B) {
	l := NewLogger()
	l.SetOutput(io.Discard)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Infof("%s", "test")
		}
	})
}

func BenchmarkWarning(b *testing.B) {
	l := NewLogger()
	l.SetOutput(io.Discard)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Warning("test")
		}
	})
}

func BenchmarkWarningf(b *testing.B) {
	l := NewLogger()
	l.SetOutput(io.Discard)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Warningf("%s", "test")
		}
	})
}

func BenchmarkError(b *testing.B) {
	l := NewLogger()
	l.SetOutput(io.Discard)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Error("test")
		}
	})
}

func BenchmarkErrorf(b *testing.B) {
	l := NewLogger()
	l.SetOutput(io.Discard)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Errorf("%s", "test")
		}
	})
}

func BenchmarkFatal(b *testing.B) {
	l := NewLogger()
	l.SetOutput(io.Discard)

	l.exitFunction = func() {}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Fatal("test")
		}
	})
}

func BenchmarkFatalf(b *testing.B) {
	l := NewLogger()
	l.SetOutput(io.Discard)

	l.exitFunction = func() {}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Fatalf("%s", "test")
		}
	})
}
