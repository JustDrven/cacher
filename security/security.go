package security

import (
	"cacher/app/settings"
	"cacher/factory"
	"cacher/file"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

var API_KEY string = ""

func init() {
	if file.GetEnv("DEV_API_KEY") != "" {
		API_KEY = file.GetEnv("DEV_API_KEY")
	} else {
		API_KEY = os.Getenv("API_KEY")
	}

	if settings.App.Debug {
		log.Println("Security has been successfully initialized")
	}
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if !IsAPIKeyValid(r) {
			w.WriteHeader(http.StatusUnauthorized)

			json.NewEncoder(w).Encode(factory.NewErrorResponse(401, "Unauthorized"))

			return
		}
		next(w, r)
	}
}

func IsAPIKeyValid(r *http.Request) bool {
	return r.Header.Get("X-API-Key") == API_KEY
}
