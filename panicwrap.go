package panicwrap

import (
	"fmt"
)

type LoggerFunc func(error)

var (
	loggerFn = func(error) {}
)

func SetLoggerFunc(fn LoggerFunc) {
	loggerFn = fn
}

func Go(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				loggerFn(fmt.Errorf("panic: %v", r))
			}
		}()

		fn()
	}()
}

func GoWithPanic(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				loggerFn(fmt.Errorf("panic: %v", r))
				panic(r)
			}
		}()

		fn()
	}()
}
