package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	Info("info")
	Infof("infof %s", "infof")
	Warn("warn")
	Warnf("warnf %s", "warnf")
	// Fatal("fatal")
	Fatalf("fatalf %s", "fatalf")
	Error("error")
}
