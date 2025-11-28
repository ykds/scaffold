package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	InitLogger(Config{})
	logger.Debugf("debug")
	logger.Infof("info")
	logger.Errorf("error")
}
