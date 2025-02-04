package controllers

import (
	"cacher/factory"
	"cacher/utility"
	"encoding/json"
	"net/http"
	"os"
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
		var value string = os.Getenv(utility.SOURCE + requestKey)

		if value != "" {
			w.WriteHeader(http.StatusOK)

			utility.SetETag("true", w)

			os.Unsetenv(utility.SOURCE + requestKey)

			writer.Encode(factory.NewDataResponse(requestKey, value))

		} else {
			w.WriteHeader(http.StatusNotFound)

			utility.SetETag("false", w)

			writer.Encode(factory.NewErrorResponse(404, "The data doesn't exist!"))
		}
	}
}
