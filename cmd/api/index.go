package api

import (
	"cacher/cmd/api/controllers"
	"cacher/cmd/app"
	"cacher/internal/manager/files"
	"cacher/internal/manager/router"
	"log"
	"net/http"
	"strconv"
)

func StartAPI(app app.Application) {
	files.CheckFiles()

	router.RegisterRouter("GET", "/query/cache/validation", controllers.IsValid, true)
	router.RegisterRouter("GET", "/query/cache", controllers.GetData, true)
	router.RegisterRouter("POST", "/mutation/cache", controllers.SaveData, true)
	router.RegisterRouter("DELETE", "/mutation/cache", controllers.RemoveData, true)
	router.RegisterRouter("PUT", "/mutation/cache", controllers.ReplaceData, true)

	router.RegisterRouter("GET", "/service/ping", controllers.Ping, false)

	log.Println("The server is starting at http://127.0.0.1:" + app.Config.Addr + "..")
	log.Println("Debug is " + strconv.FormatBool(app.Debug))

	log.Fatal(http.ListenAndServe(app.Config.Addr, nil))
}
