package main

import (
	"cacher/api"
	"cacher/security"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/v1/valid", security.AuthMiddleware(api.IsValid))
	http.HandleFunc("/v1/get", security.AuthMiddleware(api.GetData))
	http.HandleFunc("/v1/set", security.AuthMiddleware(api.SaveData))
	http.HandleFunc("/v1/remove", security.AuthMiddleware(api.RemoveData))

	fmt.Println("The server is starting..")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
