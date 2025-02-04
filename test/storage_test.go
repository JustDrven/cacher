package test

import (
	"cacher/manager"
	"testing"
)

func TestSetDataIntoEnv(t *testing.T) {
	key, value := "functionNameA", "TestSetDataIntoEnv"

	manager.Set(key, value)

	if !manager.Exist(key) {
		t.Fatalf("(%s = %v) are not saved!", key, value)
	}
}

func TestGetDataFromEnv(t *testing.T) {
	key, value := "functionNameB", "TestGetDataIntoEnv"

	manager.Set(key, value)

	data, err := manager.Get(key)

	if err {
		t.Fatalf("There is problem with getting data from storage")
		return
	}

	if data != value {
		t.Fatalf("%s and %s are not equivalent", value, data)
	}
}
