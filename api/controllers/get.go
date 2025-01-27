package controllers

import (
	"cacher/data"
	"cacher/utility"
	"encoding/json"
	"net/http"
	"os"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	writer := json.NewEncoder(w)

	var key string = r.Header.Get("key")
	var solidKey string = utility.SOURCE + key

	if key != "" && solidKey != "" {
		w.WriteHeader(http.StatusOK)
		var value string = os.Getenv(solidKey)

		if value != "" {
			utility.SetETag("true", w)

			writer.Encode(data.Data{
				Key:   key,
				Value: value,
			})
		} else {
			w.WriteHeader(http.StatusNotFound)

			utility.SetETag("false", w)

			writer.Encode(data.ErrorResponse{
				Error:   404,
				Message: "The value doesn't exist!",
			})
		}

		return
	} else {
		w.WriteHeader(http.StatusNotFound)

		utility.SetETag("false", w)

		writer.Encode(data.ErrorResponse{
			Error:   404,
			Message: "The key is missing!",
		})
	}
}
