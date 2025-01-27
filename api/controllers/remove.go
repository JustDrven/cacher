package controllers

import (
	"cacher/data"
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

		writer.Encode(data.ErrorResponse{
			Error:   404,
			Message: "The key is missing!",
		})

		return
	} else {
		var value string = os.Getenv(utility.SOURCE + requestKey)

		if value != "" {
			w.WriteHeader(http.StatusOK)

			utility.SetETag("true", w)

			os.Unsetenv(utility.SOURCE + requestKey)

			writer.Encode(data.Data{
				Key:   requestKey,
				Value: value,
			})

		} else {
			w.WriteHeader(http.StatusNotFound)

			utility.SetETag("false", w)

			writer.Encode(data.ErrorResponse{
				Error:   404,
				Message: "The data doesn't exist!",
			})
		}
	}
}
