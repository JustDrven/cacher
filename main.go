package main

import (
	"cacher/api"
	"cacher/app/settings"
)

func main() {
	api.StartAPI(settings.App)
}
