package api

import (
	"cacher/app"
	"cacher/manager"
	"log"
	"net/http"
	"strconv"
)

func StartAPI(app app.Application) {
	manager.CheckFiles()

	manager.RegisterRouter("GET", "valid", IsValid, true)
	manager.RegisterRouter("GET", "get", GetData, true)
	manager.RegisterRouter("POST", "set", SaveData, true)
	manager.RegisterRouter("DELETE", "remove", RemoveData, true)
	manager.RegisterRouter("POST", "replace", ReplaceData, true)

	log.Println("The server is starting at http://localhost" + app.Config.Addr + "..")
	log.Println("Debug is " + strconv.FormatBool(app.Debug))

	log.Fatal(http.ListenAndServe(app.Config.Addr, nil))
}
