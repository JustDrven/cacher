package controllers

import (
	"cacher/factory"
	"cacher/repository/storage"
	"cacher/utility"
	"cacher/utility/network"
	"encoding/json"
	"net/http"
)

func SaveData(w http.ResponseWriter, r *http.Request) {
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

	if !storage.Exist(key) {
		network.OkStatus(w)

		utility.SetETag("true", w)
		storage.Set(key, value)

		writer.Encode(factory.NewDataResponse(key, value))
	} else {
		network.BadRequestStatus(w)

		utility.SetETag("false", w)

		writer.Encode(factory.NewErrorResponse(400, "The value already exist!"))

	}

}
