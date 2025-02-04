package controllers

import (
	"cacher/utility/network"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	network.OkStatus(w)

	w.Write([]byte("Pong!"))
}
