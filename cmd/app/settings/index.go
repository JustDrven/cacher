package settings

import (
	"cacher/cmd/app"
	"cacher/internal/file"
)

var App = app.Application{
	Debug: file.GetEnvToBool("DEBUG"),
	Config: app.ApplicationConfig{
		Addr: ":" + file.GetEnv("API_PORT"),
	},
}
