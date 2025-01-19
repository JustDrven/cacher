package security

import (
	"cacher/data"
	"cacher/file"
	"encoding/json"
	"net/http"
	"os"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var API_KEY string = ""
		if file.GetEnv("DEV_API_KEY") != "" {
			API_KEY = file.GetEnv("DEV_API_KEY")
		} else {
			API_KEY = os.Getenv("API_KEY")
		}

		if r.Header.Get("X-API-Key") != API_KEY {
			w.WriteHeader(http.StatusUnauthorized)

			json.NewEncoder(w).Encode(data.ErrorResponse{
				Error:   401,
				Message: "Unauthorized",
			})

			return
		}
		next(w, r)
	}
}
