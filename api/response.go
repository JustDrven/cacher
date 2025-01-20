package api

import (
	"cacher/data"
	"encoding/json"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(data.ErrorResponse{
		Error:   404,
		Message: "This page not found!",
	})

}
