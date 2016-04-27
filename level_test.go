// Copyright (c) 2016 Ventu.io, Oleg Sklyar, contributors
// The use of this source code is governed by a MIT style license found in the LICENSE file

package slf

import (
	"encoding/json"
	"testing"
)

type testLevel struct {
	Level Level `json:"level"`
}

func TestLevel_marshal_success(t *testing.T) {
	res, err := json.Marshal(testLevel{Level: LevelDebug})
	if err != nil {
		t.Errorf("unexpected error, %v", err)
	}
	if string(res) != `{"level":"DEBUG"}` {
		t.Errorf("unexpected result, %v", string(res))
	}
	res, _ = json.Marshal(testLevel{Level: LevelInfo})
	if string(res) != `{"level":"INFO"}` {
		t.Errorf("unexpected result, %v", string(res))
	}
	res, _ = json.Marshal(testLevel{Level: LevelWarn})
	if string(res) != `{"level":"WARN"}` {
		t.Errorf("unexpected result, %v", string(res))
	}
	res, _ = json.Marshal(testLevel{Level: LevelError})
	if string(res) != `{"level":"ERROR"}` {
		t.Errorf("unexpected result, %v", string(res))
	}
	res, _ = json.Marshal(testLevel{Level: LevelPanic})
	if string(res) != `{"level":"PANIC"}` {
		t.Errorf("unexpected result, %v", string(res))
	}
	res, _ = json.Marshal(testLevel{Level: LevelFatal})
	if string(res) != `{"level":"FATAL"}` {
		t.Errorf("unexpected result, %v", string(res))
	}
}

func TestLevel_marshal_withWrongLevel_error(t *testing.T) {
	_, err := json.Marshal(testLevel{Level: Level(33)})
	if err == nil || err.Error() != "json: error calling MarshalJSON for type slf.Level: slf: unknown level 33" {
		t.Error("expected an error")
	}
}

func TestLevel_unmarshal_success(t *testing.T) {
	exp := testLevel{Level: LevelDebug}
	act := testLevel{}
	if err := json.Unmarshal([]byte(`{"level":"debug"}`), &act); err != nil || exp != act {
		t.Errorf("expected match, %v, %v", err, act)
	}
	exp = testLevel{Level: LevelInfo}
	act = testLevel{}
	if err := json.Unmarshal([]byte(`{"level":"info"}`), &act); err != nil || exp != act {
		t.Errorf("expected match, %v, %v", err, act)
	}
	exp = testLevel{Level: LevelWarn}
	act = testLevel{}
	if err := json.Unmarshal([]byte(`{"level":"warn"}`), &act); err != nil || exp != act {
		t.Errorf("expected match, %v, %v", err, act)
	}
	exp = testLevel{Level: LevelError}
	act = testLevel{}
	if err := json.Unmarshal([]byte(`{"level":"error"}`), &act); err != nil || exp != act {
		t.Errorf("expected match, %v, %v", err, act)
	}
	exp = testLevel{Level: LevelPanic}
	act = testLevel{}
	if err := json.Unmarshal([]byte(`{"level":"panic"}`), &act); err != nil || exp != act {
		t.Errorf("expected match, %v, %v", err, act)
	}
	exp = testLevel{Level: LevelFatal}
	act = testLevel{}
	if err := json.Unmarshal([]byte(`{"level":"fatal"}`), &act); err != nil || exp != act {
		t.Errorf("expected match, %v, %v", err, act)
	}
}

func TestLevel_unmarshal_withWrongLevel_error(t *testing.T) {
	act := testLevel{}
	err := json.Unmarshal([]byte(`{"level":"foo"}`), &act)
	if err == nil || err.Error() != `slf: unknown level "foo"` {
		t.Errorf("expected an error, %v", err)
	}
}

func TestLevel_String_success(t *testing.T) {
	if LevelError.String() != "ERROR" {
		t.Errorf("expected ERROR, %v", LevelError.String())
	}
	if Level(30).String() != "<NA>" {
		t.Errorf("expected <NA>, %v", Level(30).String())
	}
}
