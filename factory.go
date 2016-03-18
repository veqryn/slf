// Copyright (c) 2016 Ventu.io, Oleg Sklyar, contributors
// The use of this source code is governed by a MIT style license found in the LICENSE file

package slf

var root Logger

func init() {
	root = &Noop{}
}

// IsDefined checks if the root logger has been defined (default is a Noop (No Operation) one).
func IsDefined() bool {
	if _, ok := root.(*Noop); !ok {
		return true
	}
	return false
}

// Define defines the root logger.
func Define(log Logger) {
	root = log
}

// Get returns the currently defined root logger.
func Get() Logger {
	return root
}

// WithContext returns a structured logger with "context" set.
func WithContext(context string) StructuredLogger {
	return root.WithField("context", context)
}
