// Copyright (c) 2016 Ventu.io, Oleg Sklyar, contributors
// The use of this source code is governed by a MIT style license found in the LICENSE file

package slf

import (
	"testing"
)

type testFactory struct {
	*Noop
	context string
}

func (log *testFactory) WithContext(context string) StructuredLogger {
	log.context = context
	return log
}

func TestFactory_define_success(t *testing.T) {
	Set(&Noop{})
	if IsSet() {
		t.Error("expected undefined")
	}
	Set(&testFactory{&Noop{}, ""})
	if !IsSet() {
		t.Error("expected defined")
	}
	Set(&Noop{})
	if IsSet() {
		t.Error("expected undefined")
	}
}

func TestFactory_withContext_success(t *testing.T) {
	Set(&testFactory{&Noop{}, ""})
	logger := WithContext("test")
	f, ok := logger.(*testFactory)
	if !ok {
		t.Error("expected factory")
	}
	if f.context != "test" {
		t.Errorf("expected test, %v", f.context)
	}
}
