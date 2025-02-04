package controllers

import (
	"cacher/factory"
	"cacher/manager"
	"cacher/utility"
	"encoding/json"
	"net/http"
)

func IsValid(w http.ResponseWriter, r *http.Request) {
	writer := json.NewEncoder(w)
	var key string = r.Header.Get("key")

	if key == "" {
		w.WriteHeader(http.StatusNotFound)

		utility.SetETag("false", w)

		writer.Encode(factory.NewErrorResponse(404, "The key is missing!"))

		return
	}

	if manager.Exist(key) {
		w.WriteHeader(http.StatusOK)

		utility.SetETag("true", w)

		writer.Encode(factory.NewValidResponse(true))
	} else {
		w.WriteHeader(http.StatusNotFound)

		utility.SetETag("false", w)

		writer.Encode(factory.NewValidResponse(false))
	}

}
