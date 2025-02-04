package controllers

import (
	"cacher/factory"
	"cacher/manager"
	"cacher/utility"
	"encoding/json"
	"net/http"
)

func SaveData(w http.ResponseWriter, r *http.Request) {
	writer := json.NewEncoder(w)

	var key string = r.Header.Get("key")
	if key == "" {
		w.WriteHeader(http.StatusNotFound)

		utility.SetETag("false", w)

		writer.Encode(factory.NewErrorResponse(404, "The key is missing!"))

		return
	}

	var value string = r.Header.Get("value")
	if value == "" {
		w.WriteHeader(http.StatusNotFound)

		utility.SetETag("false", w)

		writer.Encode(factory.NewErrorResponse(404, "The value is missing!"))

		return
	}

	if !manager.Exist(key) {

		utility.SetETag("true", w)

		manager.Set(key, value)

		writer.Encode(factory.NewDataResponse(key, value))
	} else {
		w.WriteHeader(http.StatusBadRequest)

		utility.SetETag("false", w)

		writer.Encode(factory.NewErrorResponse(400, "The value already exist!"))

	}

}
