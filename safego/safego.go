package safego

import (
	"runtime/debug"
	"scaffold/pkg/logger"
)

func Go(f func()) {
	defer func() {
		if e := recover(); e != nil {
			logger.Errorf("panic recover: %+v", string(debug.Stack()))
		}
	}()
	go f()
}
