package api

import (
	"cacher/security"
	"log"
	"net/http"
)

type App struct {
	Addr string
}

func StartAPI(app App) {
	http.HandleFunc("GET /v1/valid", security.AuthMiddleware(IsValid))
	http.HandleFunc("GET /v1/get", security.AuthMiddleware(GetData))
	http.HandleFunc("POST /v1/set", security.AuthMiddleware(SaveData))
	http.HandleFunc("DELETE /v1/remove", security.AuthMiddleware(RemoveData))

	log.Println("The server is starting at " + app.Addr + "..")
	log.Fatal(http.ListenAndServe(app.Addr, nil))
}
