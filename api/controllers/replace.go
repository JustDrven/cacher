package controllers

import (
	"cacher/factory"
	"cacher/repository/storage"
	"cacher/utility"
	"cacher/utility/network"
	"encoding/json"
	"net/http"
)

func ReplaceData(w http.ResponseWriter, r *http.Request) {
	writer := json.NewEncoder(w)

	var key string = r.Header.Get("key")
	if key == "" {
		network.NotFoundStatus(w)

		utility.SetETag("false", w)

		writer.Encode(factory.NewErrorResponse(404, "The key is missing!"))

		return
	}

	var value string = r.Header.Get("value")
	if value == "" {
		network.NotFoundStatus(w)

		utility.SetETag("false", w)

		writer.Encode(factory.NewErrorResponse(404, "The value is missing!"))

		return
	}

	if storage.Exist(key) {
		network.OkStatus(w)

		storage.Remove(key)
		storage.Set(key, value)

		utility.SetETag("true", w)

		writer.Encode(factory.NewValidResponse(true))
	} else {
		network.BadRequestStatus(w)

		utility.SetETag("false", w)

		writer.Encode(factory.NewValidResponse(false))
	}

}
