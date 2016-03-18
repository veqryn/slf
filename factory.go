// Copyright (c) 2016 Ventu.io, Oleg Sklyar, contributors
// The use of this source code is governed by a MIT style license found in the LICENSE file

package slf

var root Logger = &Noop{}

// IsSet checks if the root logger has been defined. Default is false due to the Noop logger
// (No Operation).
func IsSet() bool {
	if _, ok := root.(*Noop); !ok {
		return true
	}
	return false
}

// Set defines the root logger.
func Set(log Logger) {
	root = log
}

// WithContext returns a logger with context set to a string.
func WithContext(context string) StructuredLogger {
	return root.WithContext(context)
}

// WithCaller returns a logger with context set to the caller.
func WithCaller() StructuredLogger {
	return root.WithCaller()
}
