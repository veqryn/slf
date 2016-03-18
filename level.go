// Copyright (c) 2016 Ventu.io, Oleg Sklyar, contributors
// The use of this source code is governed by a MIT style license found in the LICENSE file

package slf

import (
	"fmt"
	"strings"
)

// Level represents log level of the structured logger.
type Level int

// Log level constants.
const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelPanic
)

// MarshalJSON provides a JSON representation of the log level.
func (l Level) MarshalJSON() ([]byte, error) {
	var s string
	switch l {
	case LevelDebug:
		s = "debug"
	case LevelInfo:
		s = "info"
	case LevelWarn:
		s = "warn"
	case LevelError:
		s = "error"
	case LevelPanic:
		s = "panic"
	default:
		return nil, fmt.Errorf("slf: unknown level %v", l)
	}
	return []byte(`"` + s + `"`), nil
}

// UnmarshalJSON parses the JSON representation of the log level into a Level object.
func (l *Level) UnmarshalJSON(data []byte) error {
	s := strings.ToLower(string(data))
	switch s {
	case `"debug"`:
		fallthrough
	case "debug":
		*l = LevelDebug
	case `"info"`:
		fallthrough
	case "info":
		*l = LevelInfo
	case `"warn"`:
		fallthrough
	case "warn":
		*l = LevelWarn
	case `"error"`:
		fallthrough
	case "error":
		*l = LevelError
	case `"panic"`:
		fallthrough
	case "panic":
		*l = LevelPanic
	default:
		return fmt.Errorf("slf: unknown level %v", s)
	}
	return nil
}
