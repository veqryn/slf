// Copyright (c) 2016 Ventu.io, Oleg Sklyar, contributors
// The use of this source code is governed by a MIT style license found in the LICENSE file

package slf

import (
	"errors"
	"fmt"
)

// Noop implements the Logger interface with no operation (except for panic in Panic and Panicf).
type Noop struct{}

// WithContext implements the Logger interface.
func (log *Noop) WithContext(string) StructuredLogger {
	return log
}

// WithField implements the Logger interface.
func (log *Noop) WithField(string, interface{}) StructuredLogger {
	return log
}

// WithFields implements the Logger interface.
func (log *Noop) WithFields(Fields) StructuredLogger {
	return log
}

// WithError implements the Logger interface.
func (log *Noop) WithError(error) BasicLogger {
	return log
}

// Log implements the Logger interface.
func (log *Noop) Log(Level, string) Tracer {
	return log
}

// Logf implements the Logger interface.
func (log *Noop) Logf(Level, string, ...interface{}) Tracer {
	return log
}

// Debug implements the Logger interface.
func (log *Noop) Debug(string) Tracer {
	return log
}

// Debugf implements the Logger interface.
func (log *Noop) Debugf(string, ...interface{}) Tracer {
	return log
}

// Info implements the Logger interface.
func (log *Noop) Info(string) Tracer {
	return log
}

// Infof implements the Logger interface.
func (log *Noop) Infof(string, ...interface{}) Tracer {
	return log
}

// Warn implements the Logger interface.
func (log *Noop) Warn(string) Tracer {
	return log
}

// Warnf implements the Logger interface.
func (log *Noop) Warnf(string, ...interface{}) Tracer {
	return log
}

// Error implements the Logger interface.
func (log *Noop) Error(string) Tracer {
	return log
}

// Errorf implements the Logger interface.
func (log *Noop) Errorf(string, ...interface{}) Tracer {
	return log
}

// Panic implements the Logger interface.
func (*Noop) Panic(message string) {
	panic(errors.New(message))
}

// Panicf implements the Logger interface.
func (*Noop) Panicf(message string, args ...interface{}) {
	panic(fmt.Errorf(message, args...))
}

// Trace implements the Logger interface.
func (*Noop) Trace(*error) {
}
