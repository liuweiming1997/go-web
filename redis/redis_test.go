package redis

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"
)

var notMatchErr = errors.New("do not match")

type testStruct struct {
	Name string
	Sex  string
}

func TestBase(t *testing.T) {
	s := &testStruct{"test_name", "test_sex"}

	key := []byte("test_key")
	value, _ := json.Marshal(s)

	defer Delete(key)

	err := Set(key, value)
	if err != nil {
		t.Fatal(err)
	}

	res := Get(key)

	if bytes.Equal(res, value) == false {
		t.Fatal(notMatchErr)
	}

	res1 := Exist(key)
	if res1 == false {
		t.Fatal(notMatchErr)
	}
}
