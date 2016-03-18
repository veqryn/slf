// Copyright (c) 2016 Ventu.io, Oleg Sklyar, contributors
// The use of this source code is governed by a MIT style license found in the LICENSE file

package slf_test

import (
	"github.com/ventu-io/slf"
	"testing"
)

type factory struct {
	*slf.Noop
	context string
}

func (log *factory) WithContext(context string) slf.StructuredLogger {
	log.context = context
	return log
}

func TestFactory_define_success(t *testing.T) {
	slf.Set(&slf.Noop{})
	if slf.IsSet() {
		t.Error("expected undefined")
	}
	slf.Set(&factory{&slf.Noop{}, ""})
	if !slf.IsSet() {
		t.Error("expected defined")
	}
	slf.Set(&slf.Noop{})
	if slf.IsSet() {
		t.Error("expected undefined")
	}
}

func TestFactory_withContext_success(t *testing.T) {
	slf.Set(&factory{&slf.Noop{}, ""})
	logger := slf.WithContext("test")
	f, ok := logger.(*factory)
	if !ok {
		t.Error("expected factory")
	}
	if f.context != "test" {
		t.Errorf("expected test, %v", f.context)
	}
}
