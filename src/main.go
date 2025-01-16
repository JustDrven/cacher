package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var SOURCE = "CACHER:"
var data = make(map[string]string)

func main() {
	http.HandleFunc("/v1/list", getData)
	http.HandleFunc("/v1/get", getData)
	http.HandleFunc("/v1/set", saveData)

	fmt.Println("Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func saveData(w http.ResponseWriter, r *http.Request) {
	if checkAuthorization(r.Header) {
		var key string = r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, `{"code": 404, "message": "The key is missing!"}`, http.StatusNotFound)
			return
		}

		var value string = r.URL.Query().Get("value")
		if value == "" {
			http.Error(w, `{"code": 404, "message": "The value is missing!"}`, http.StatusNotFound)
			return
		}

		if os.Getenv(SOURCE+key) != "" {
			os.Setenv(SOURCE+key, value)
			json.NewEncoder(w).Encode(`{"ok": true}`)
			return
		} else {
			json.NewEncoder(w).Encode(`{"ok": false}`)
			return
		}

	} else {
		http.Error(w, `{"code": 401, "message": "Unauthorized!"}`, http.StatusNotFound)
		return
	}
}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if checkAuthorization(r.Header) {
		var key string = r.URL.Query().Get("key")

		if key != "" && os.Getenv(SOURCE+key) != "" {
			var value string = os.Getenv(SOURCE + key)

			json.NewEncoder(w).Encode("{\"key\":\"" + key + "\",\"value\":\"" + value + "\"}")
			return
		} else {
			http.Error(w, `{"code": 404, "message": "The key is missing!"}`, http.StatusNotFound)
			return
		}
	} else {
		http.Error(w, `{"code": 401, "message": "Unauthorized!"}`, http.StatusNotFound)
		return
	}
}

func checkAuthorization(header http.Header) bool {
	if header.Get("X-API-Key") == os.Getenv("API_KEY") {
		return false
	}
	return true
}
