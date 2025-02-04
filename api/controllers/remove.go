package controllers

import (
	"cacher/factory"
	"cacher/manager"
	"cacher/utility"
	"encoding/json"
	"net/http"
)

func RemoveData(w http.ResponseWriter, r *http.Request) {
	writer := json.NewEncoder(w)

	var requestKey string = r.Header.Get("key")
	if requestKey == "" {
		w.WriteHeader(http.StatusNotFound)

		utility.SetETag("false", w)

		writer.Encode(factory.NewErrorResponse(404, "The key is missing!"))

		return
	} else {
		var value, err = manager.Get(requestKey)

		if !err {
			w.WriteHeader(http.StatusOK)

			utility.SetETag("true", w)

			manager.Remove(requestKey)

			writer.Encode(factory.NewDataResponse(requestKey, value))

		} else {
			w.WriteHeader(http.StatusNotFound)

			utility.SetETag("false", w)

			writer.Encode(factory.NewErrorResponse(404, "The data doesn't exist!"))
		}
	}
}
