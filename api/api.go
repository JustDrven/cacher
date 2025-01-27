package api

import (
	"cacher/api/controllers"
	"cacher/app"
	"cacher/manager"
	"log"
	"net/http"
	"strconv"
)

func StartAPI(app app.Application) {
	manager.CheckFiles()

	manager.RegisterRouter("GET", "valid", controllers.IsValid, true)
	manager.RegisterRouter("GET", "get", controllers.GetData, true)
	manager.RegisterRouter("POST", "set", controllers.SaveData, true)
	manager.RegisterRouter("DELETE", "remove", controllers.RemoveData, true)
	manager.RegisterRouter("PUT", "replace", controllers.ReplaceData, true)

	manager.RegisterRouter("GET", "ping", controllers.Ping, false)

	log.Println("The server is starting at http://localhost" + app.Config.Addr + "..")
	log.Println("Debug is " + strconv.FormatBool(app.Debug))

	log.Fatal(http.ListenAndServe(app.Config.Addr, nil))
}
