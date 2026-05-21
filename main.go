package main

import (
	"cacher/cmd/api"
	"cacher/cmd/app/settings"
)

func main() {
	api.StartAPI(settings.App)
}
