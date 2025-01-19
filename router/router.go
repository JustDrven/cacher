package router

import (
	"cacher/data"
	"cacher/security"
	"encoding/json"
	"net/http"
	"os"
)

var SOURCE = "CACHER:"

func SaveData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	writer := json.NewEncoder(w)

	if security.CheckAuthorization(r.Header) {

		var key string = r.Header.Get("key")
		if key == "" {
			w.WriteHeader(http.StatusNotFound)

			writer.Encode(data.ErrorResponse{
				Error:   404,
				Message: "The key is missing!",
			})

			return
		}

		var value string = r.Header.Get("value")
		if value == "" {
			w.WriteHeader(http.StatusNotFound)

			writer.Encode(data.ErrorResponse{
				Error:   404,
				Message: "The value is missing!",
			})

			return
		}

		if os.Getenv(SOURCE+key) == "" {
			os.Setenv(SOURCE+key, value)

			writer.Encode(data.Data{
				Key:   key,
				Value: value,
			})

			return
		} else {
			w.WriteHeader(http.StatusBadRequest)

			writer.Encode(data.ErrorResponse{
				Error:   400,
				Message: "The value already exist!",
			})

			return
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)

		writer.Encode(data.ErrorResponse{
			Error:   401,
			Message: "Unauthorized!",
		})

		return
	}
}

func GetData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	writer := json.NewEncoder(w)

	if security.CheckAuthorization(r.Header) {
		var key string = r.Header.Get("key")
		var solidKey string = SOURCE + key

		if key != "" && solidKey != "" {
			w.WriteHeader(http.StatusOK)
			var value string = os.Getenv(solidKey)

			if value != "" {
				writer.Encode(data.Data{
					Key:   key,
					Value: value,
				})

				return
			} else {
				w.WriteHeader(http.StatusNotFound)

				writer.Encode(data.ErrorResponse{
					Error:   404,
					Message: "The value doesn't exist!",
				})
				return
			}

		} else {
			w.WriteHeader(http.StatusNotFound)

			writer.Encode(data.ErrorResponse{
				Error:   404,
				Message: "The key is missing!",
			})

			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)

		writer.Encode(data.ErrorResponse{
			Error:   401,
			Message: "Unauthorized!",
		})

		return
	}
}

func IsValid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	writer := json.NewEncoder(w)

	if security.CheckAuthorization(r.Header) {
		var key string = r.Header.Get("key")

		if key == "" {
			w.WriteHeader(http.StatusNotFound)

			writer.Encode(data.ErrorResponse{
				Error:   404,
				Message: "The key is missing!",
			})

			return
		}

		if os.Getenv(SOURCE+key) != "" {
			w.WriteHeader(http.StatusOK)

			writer.Encode(data.Valid{
				Ok: true,
			})

			return
		} else {
			w.WriteHeader(http.StatusNotFound)

			writer.Encode(data.Valid{
				Ok: false,
			})

			return
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)

		writer.Encode(data.ErrorResponse{
			Error:   401,
			Message: "Unauthorized!",
		})

		return
	}
}

func RemoveData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	writer := json.NewEncoder(w)

	if security.CheckAuthorization(r.Header) {
		var requestKey string = r.Header.Get("key")
		if requestKey == "" {
			w.WriteHeader(http.StatusNotFound)

			writer.Encode(data.ErrorResponse{
				Error:   404,
				Message: "The key is missing!",
			})

			return
		} else {
			var value string = os.Getenv(SOURCE + requestKey)

			if value != "" {
				w.WriteHeader(http.StatusOK)

				os.Unsetenv(SOURCE + requestKey)

				writer.Encode(data.Data{
					Key:   requestKey,
					Value: value,
				})

				return
			} else {
				w.WriteHeader(http.StatusNotFound)

				writer.Encode(data.ErrorResponse{
					Error:   404,
					Message: "The data doesn't exist!",
				})

				return
			}
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)

		writer.Encode(data.ErrorResponse{
			Error:   401,
			Message: "Unauthorized!",
		})

		return
	}
}
