package test

import (
	"cacher/repository/storage"
	"testing"
)

func TestSetDataIntoEnv(t *testing.T) {
	key, value := "functionNameA", "TestSetDataIntoEnv"

	storage.Set(key, value)

	if !storage.Exist(key) {
		t.Fatalf("(%s = %v) are not saved!", key, value)
	}
}

func TestGetDataFromEnv(t *testing.T) {
	key, value := "functionNameB", "TestGetDataIntoEnv"

	storage.Set(key, value)

	data, err := storage.Get(key)

	if err {
		t.Fatalf("There is problem with getting data from storage")
		return
	}

	if data != value {
		t.Fatalf("%s and %s are not equivalent", value, data)
	}
}
