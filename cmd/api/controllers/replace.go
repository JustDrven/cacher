package controllers

import (
	"cacher/internal/factory"
	"cacher/internal/repository/storage"
	"cacher/pkg"
	"cacher/pkg/network"
	"encoding/json"
	"net/http"
)

func ReplaceData(w http.ResponseWriter, r *http.Request) {
	writer := json.NewEncoder(w)

	var key string = r.Header.Get("key")
	if key == "" {
		network.NotFoundStatus(w)

		pkg.SetETag("false", w)

		writer.Encode(factory.NewErrorResponse(404, "The key is missing!"))

		return
	}

	var value string = r.Header.Get("value")
	if value == "" {
		network.NotFoundStatus(w)

		pkg.SetETag("false", w)

		writer.Encode(factory.NewErrorResponse(404, "The value is missing!"))

		return
	}

	if storage.Exist(key) {
		network.OkStatus(w)

		storage.Remove(key)
		storage.Set(key, value)

		pkg.SetETag("true", w)

		writer.Encode(factory.NewValidResponse(true))
	} else {
		network.BadRequestStatus(w)

		pkg.SetETag("false", w)

		writer.Encode(factory.NewValidResponse(false))
	}

}
