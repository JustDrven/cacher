package main

import (
	"cacher/router"
	"cacher/security"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/v1/valid", security.AuthMiddleware(router.IsValid))
	http.HandleFunc("/v1/get", security.AuthMiddleware(router.GetData))
	http.HandleFunc("/v1/set", security.AuthMiddleware(router.SaveData))
	http.HandleFunc("/v1/remove", security.AuthMiddleware(router.RemoveData))

	fmt.Println("The server is starting..")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
