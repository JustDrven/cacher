package controllers

import (
	"cacher/factory"
	"cacher/utility"
	"encoding/json"
	"net/http"
	"os"
)

func ReplaceData(w http.ResponseWriter, r *http.Request) {
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

	if os.Getenv(utility.SOURCE+key) != "" {
		os.Setenv(utility.SOURCE+key, value)

		utility.SetETag("true", w)

		writer.Encode(factory.NewValidResponse(true))
	} else {
		w.WriteHeader(http.StatusBadRequest)

		utility.SetETag("false", w)

		writer.Encode(factory.NewValidResponse(false))
	}

}
