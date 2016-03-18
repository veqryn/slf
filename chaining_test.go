// Copyright (c) 2016 Ventu.io, Oleg Sklyar, contributors
// The use of this source code is governed by a MIT style license found in the LICENSE file

package slf_test

import (
	"github.com/ventu-io/slf"
	"testing"
)

func TestChaining_uotoerror_success(t *testing.T) {
	log := slf.Get()
	log.Log(slf.LevelInfo, "").Trace(nil)
	log.Logf(slf.LevelInfo, "%v", "").Trace(nil)
	log.WithCaller(true).Debug("").Trace(nil)
	log.WithError(nil).Debugf("%v", "").Trace(nil)
	log.WithField("a", "b").Info("").Trace(nil)
	log.WithCaller(true).WithField("a", "b").Debug("").Trace(nil)
	log.WithFields(slf.Fields{"a": "b"}).Error("").Trace(nil)
	log.WithFields(slf.Fields{"a": "b"}).WithError(nil).Infof("%v", "").Trace(nil)
	log.WithError(nil).Errorf("%v", "").Trace(nil)
	log.WithField("a", "b").WithField("c", "d").Info("").Trace(nil)
	log.WithCaller(true).WithError(nil).Warn("").Trace(nil)
	log.WithCaller(true).WithError(nil).Warnf("%v", "").Trace(nil)
}

func TestChaning_panic_success(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("panic expecte")
		}
	}()
	slf.Get().Panic("")
}

func TestChaning_panicf_success(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("panic expecte")
		}
	}()
	slf.Get().WithField("a", "b").Panicf("%v", "")
}
