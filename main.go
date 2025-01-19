package main

import (
	"cacher/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/v1/valid", router.IsValid)
	http.HandleFunc("/v1/get", router.GetData)
	http.HandleFunc("/v1/set", router.SaveData)
	http.HandleFunc("/v1/remove", router.RemoveData)

	fmt.Println("The server is starting..")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
