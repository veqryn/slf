// Copyright (c) 2016 Ventu.io, Oleg Sklyar, contributors
// The use of this source code is governed by a MIT style license found in the LICENSE file

package slf_test

import (
	"encoding/json"
	"github.com/ventu-io/slf"
	"testing"
)

type test struct {
	Level slf.Level `json:"level"`
}

func TestLevel_marshal_success(t *testing.T) {
	res, err := json.Marshal(test{Level: slf.LevelDebug})
	if err != nil {
		t.Errorf("unexpected error, %v", err)
	}
	if string(res) != `{"level":"DEBUG"}` {
		t.Errorf("unexpected result, %v", string(res))
	}
	res, _ = json.Marshal(test{Level: slf.LevelInfo})
	if string(res) != `{"level":"INFO"}` {
		t.Errorf("unexpected result, %v", string(res))
	}
	res, _ = json.Marshal(test{Level: slf.LevelWarn})
	if string(res) != `{"level":"WARN"}` {
		t.Errorf("unexpected result, %v", string(res))
	}
	res, _ = json.Marshal(test{Level: slf.LevelError})
	if string(res) != `{"level":"ERROR"}` {
		t.Errorf("unexpected result, %v", string(res))
	}
	res, _ = json.Marshal(test{Level: slf.LevelPanic})
	if string(res) != `{"level":"PANIC"}` {
		t.Errorf("unexpected result, %v", string(res))
	}
}

func TestLevel_marshal_withWrongLevel_error(t *testing.T) {
	_, err := json.Marshal(test{Level: slf.Level(33)})
	if err == nil || err.Error() != "json: error calling MarshalJSON for type slf.Level: slf: unknown level 33" {
		t.Error("expected an error")
	}
}

func TestLevel_unmarshal_success(t *testing.T) {
	exp := test{Level: slf.LevelDebug}
	act := test{}
	if err := json.Unmarshal([]byte(`{"level":"debug"}`), &act); err != nil || exp != act {
		t.Errorf("expected match, %v, %v", err, act)
	}
	exp = test{Level: slf.LevelInfo}
	act = test{}
	if err := json.Unmarshal([]byte(`{"level":"info"}`), &act); err != nil || exp != act {
		t.Errorf("expected match, %v, %v", err, act)
	}
	exp = test{Level: slf.LevelWarn}
	act = test{}
	if err := json.Unmarshal([]byte(`{"level":"warn"}`), &act); err != nil || exp != act {
		t.Errorf("expected match, %v, %v", err, act)
	}
	exp = test{Level: slf.LevelError}
	act = test{}
	if err := json.Unmarshal([]byte(`{"level":"error"}`), &act); err != nil || exp != act {
		t.Errorf("expected match, %v, %v", err, act)
	}
	exp = test{Level: slf.LevelPanic}
	act = test{}
	if err := json.Unmarshal([]byte(`{"level":"panic"}`), &act); err != nil || exp != act {
		t.Errorf("expected match, %v, %v", err, act)
	}
}

func TestLevel_unmarshal_withWrongLevel_error(t *testing.T) {
	act := test{}
	err := json.Unmarshal([]byte(`{"level":"foo"}`), &act)
	if err == nil || err.Error() != `slf: unknown level "foo"` {
		t.Errorf("expected an error, %v", err)
	}
}

func TestLevel_String_success(t *testing.T) {
	if slf.LevelError.String() != "ERROR" {
		t.Errorf("expected ERROR, %v", slf.LevelError.String())
	}
	if slf.Level(30).String() != "<NA>" {
		t.Errorf("expected <NA>, %v", slf.Level(30).String())
	}
}
