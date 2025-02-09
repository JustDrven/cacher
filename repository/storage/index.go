package storage

import (
	"cacher/utility"
	"os"
)

const PREFIX string = utility.SOURCE

func Exist(key string) bool {
	return os.Getenv(PREFIX+key) != ""
}

func Set(key string, value string) {
	os.Setenv(PREFIX+key, value)
}

func Remove(key string) {
	if !Exist(key) {
		return
	}

	os.Unsetenv(PREFIX + key)
}

func Get(key string) (string, bool) {
	if !Exist(key) {
		return "", true
	}

	return os.Getenv(PREFIX + key), false
}
