package controllers

import (
	"cacher/internal/factory"
	"cacher/internal/repository/storage"
	"cacher/pkg"
	"cacher/pkg/network"

	"encoding/json"
	"net/http"
)

func IsValid(w http.ResponseWriter, r *http.Request) {
	writer := json.NewEncoder(w)
	var key string = r.Header.Get("key")

	if key == "" {
		network.NotFoundStatus(w)

		pkg.SetETag("false", w)

		writer.Encode(factory.NewErrorResponse(404, "The key is missing!"))

		return
	}

	if storage.Exist(key) {
		network.OkStatus(w)

		pkg.SetETag("true", w)

		writer.Encode(factory.NewValidResponse(true))
	} else {
		network.NotFoundStatus(w)

		pkg.SetETag("false", w)

		writer.Encode(factory.NewValidResponse(false))
	}

}
