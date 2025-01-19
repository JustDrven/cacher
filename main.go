package main

import "cacher/api"

func main() {
	app := api.App{
		Addr: ":8080",
	}

	api.StartAPI(app)
}
