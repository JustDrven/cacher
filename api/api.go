package api

import (
	"cacher/app"
	"cacher/security"
	"log"
	"net/http"
)

func StartAPI(app app.Application) {

	http.HandleFunc("GET /v1/valid", security.AuthMiddleware(IsValid))
	http.HandleFunc("GET /v1/get", security.AuthMiddleware(GetData))
	http.HandleFunc("POST /v1/set", security.AuthMiddleware(SaveData))
	http.HandleFunc("DELETE /v1/remove", security.AuthMiddleware(RemoveData))

	http.HandleFunc("/", NotFound)

	log.Println("The server is starting at http://localhost" + app.Config.Addr + "..")
	log.Fatal(http.ListenAndServe(app.Config.Addr, nil))
}
