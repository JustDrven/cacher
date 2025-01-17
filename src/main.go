package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Valid struct {
	Ok bool `json:"ok"`
}

type ErrorResponse struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}

var SOURCE = "CACHER:"

func main() {
	http.HandleFunc("/v1/valid", isValid)
	http.HandleFunc("/v1/get", getData)
	http.HandleFunc("/v1/set", saveData)

	fmt.Println("The server is starting..")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func saveData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if checkAuthorization(r.Header) {

		var key string = r.Header.Get("key")
		if key == "" {
			w.WriteHeader(http.StatusNotFound)

			var errorResponse = ErrorResponse{
				Error:   404,
				Message: "The key is missing!",
			}
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

		var value string = r.Header.Get("value")
		if value == "" {
			w.WriteHeader(http.StatusNotFound)

			var errorResponse = ErrorResponse{
				Error:   404,
				Message: "The value is missing!",
			}
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

		var data = Data{
			Key:   key,
			Value: value,
		}

		if os.Getenv(SOURCE+key) == "" {
			os.Setenv(SOURCE+key, value)
			json.NewEncoder(w).Encode(data)
			return
		} else {
			w.WriteHeader(http.StatusBadRequest)

			var error = ErrorResponse{
				Error:   400,
				Message: "The value already exist!",
			}

			json.NewEncoder(w).Encode(error)
			return
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)

		var error = ErrorResponse{
			Error:   401,
			Message: "Unauthorized!",
		}

		json.NewEncoder(w).Encode(error)
		return
	}
}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if checkAuthorization(r.Header) {
		var key string = r.Header.Get("key")

		if key != "" && os.Getenv(SOURCE+key) != "" {
			w.WriteHeader(http.StatusOK)
			var value string = os.Getenv(SOURCE + key)

			var data = Data{
				Key:   key,
				Value: value,
			}

			json.NewEncoder(w).Encode(data)
			return
		} else {
			w.WriteHeader(http.StatusNotFound)

			var errorResponse = ErrorResponse{
				Error:   404,
				Message: "The key is missing!",
			}
			json.NewEncoder(w).Encode(errorResponse)
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)

		var error = ErrorResponse{
			Error:   401,
			Message: "Unauthorized!",
		}

		json.NewEncoder(w).Encode(error)
		return
	}
}

func isValid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if checkAuthorization(r.Header) {
		var key string = r.Header.Get("key")

		if key == "" {
			w.WriteHeader(http.StatusNotFound)

			var error = ErrorResponse{
				Error:   404,
				Message: "The key is missing!",
			}

			json.NewEncoder(w).Encode(error)
			return
		}

		if os.Getenv(SOURCE+key) != "" {
			w.WriteHeader(http.StatusOK)

			var valid = Valid{
				Ok: true,
			}

			json.NewEncoder(w).Encode(valid)
			return
		} else {
			w.WriteHeader(http.StatusOK)

			var valid = Valid{
				Ok: false,
			}

			json.NewEncoder(w).Encode(valid)
			return
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)

		var error = ErrorResponse{
			Error:   401,
			Message: "Unauthorized!",
		}

		json.NewEncoder(w).Encode(error)
		return
	}
}

func checkAuthorization(header http.Header) bool {
	if header.Get("X-API-Key") == os.Getenv("API_KEY") {
		return false
	}
	return true
}
