package api

import (
	"cacher/app"
	"cacher/security"
	"cacher/utility"
	"log"
	"net/http"
)

func StartAPI(app app.Application) {

	http.HandleFunc("GET "+utility.API_VERSION+"valid", security.AuthMiddleware(IsValid))
	http.HandleFunc("GET "+utility.API_VERSION+"get", security.AuthMiddleware(GetData))
	http.HandleFunc("POST "+utility.API_VERSION+"set", security.AuthMiddleware(SaveData))
	http.HandleFunc("DELETE "+utility.API_VERSION+"remove", security.AuthMiddleware(RemoveData))
	http.HandleFunc("POST "+utility.API_VERSION+"replace", security.AuthMiddleware(ReplaceData))

	http.HandleFunc("/", NotFound)

	log.Println("The server is starting at http://localhost" + app.Config.Addr + "..")
	log.Fatal(http.ListenAndServe(app.Config.Addr, nil))
}
