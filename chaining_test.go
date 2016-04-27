// Copyright (c) 2016 Ventu.io, Oleg Sklyar, contributors
// The use of this source code is governed by a MIT style license found in the LICENSE file

package slf

import (
	"testing"
)

func TestChaining_withContext_success(t *testing.T) {
	log := WithContext("context")
	log.Log(LevelInfo, "").Trace(nil)
	log.WithError(nil).Debugf("%v", "").Trace(nil)
	log.WithField("a", "b").Info("").Trace(nil)
	log.WithFields(Fields{"a": "b"}).Error("").Trace(nil)
	log.WithFields(Fields{"a": "b"}).WithError(nil).Infof("%v", "").Trace(nil)
	log.WithError(nil).Errorf("%v", "").Trace(nil)
	log.WithField("a", "b").WithField("c", "d").Info("").Trace(nil)
	log.Debug("")
	log.WithField("a", "b").Debug("").Trace(nil)
	log.WithError(nil).Warn("").Trace(nil)
	log.WithError(nil).Warnf("%v", "").Trace(nil)
	log.WithCaller(CallerShort).WithCaller(CallerLong).WithCaller(CallerNone).Error("test")
}

func TestChaning_panic_success(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("panic expecte")
		}
	}()
	WithContext("context").Panic("")
}

func TestChaning_panicf_success(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("panic expecte")
		}
	}()
	WithContext("context").WithField("a", "b").Panicf("%v", "")
}
