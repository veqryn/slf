// Copyright (c) 2016 Ventu.io, Oleg Sklyar, contributors
// The use of this source code is governed by a MIT style license found in the LICENSE file

package slf_test

import (
	"fmt"
	"github.com/ventu-io/slf"
	"testing"
)

type factory struct {
	*slf.Noop
	context string
}

func (log *factory) WithField(key string, value interface{}) slf.StructuredLogger {
	if key == "context" {
		log.context = fmt.Sprint(value)
	}
	return log
}

func TestFactory_define_success(t *testing.T) {
	slf.Define(&slf.Noop{})
	if slf.IsDefined() {
		t.Error("expected undefined")
	}
	slf.Define(&factory{&slf.Noop{}, ""})
	if !slf.IsDefined() {
		t.Error("expected defined")
	}
	slf.Define(&slf.Noop{})
	if slf.IsDefined() {
		t.Error("expected undefined")
	}
}

func TestFactory_get_success(t *testing.T) {
	logger := slf.Get()
	if _, ok := logger.(*slf.Noop); !ok {
		t.Error("expected Noop")
	}
	slf.Define(&factory{&slf.Noop{}, ""})
	logger = slf.Get()
	if _, ok := logger.(*factory); !ok {
		t.Error("expected factory")
	}
}

func TestFactory_withContext_success(t *testing.T) {
	slf.Define(&factory{&slf.Noop{}, ""})
	logger := slf.WithContext("test")
	f, ok := logger.(*factory)
	if !ok {
		t.Error("expected factory")
	}
	if f.context != "test" {
		t.Errorf("expected test, %v", f.context)
	}
}
