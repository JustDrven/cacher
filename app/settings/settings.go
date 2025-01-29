package settings

import (
	v1 "cacher/app"
	"cacher/file"
)

var App = v1.Application{
	Debug: file.GetEnvToBool("DEBUG"),
	Config: v1.ApplicationConfig{
		Addr: ":" + file.GetEnv("API_PORT"),
	},
}
