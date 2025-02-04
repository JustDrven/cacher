package controllers

import (
	"cacher/factory"
	"cacher/manager"
	"cacher/utility"
	"cacher/utility/network"
	"encoding/json"
	"net/http"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	writer := json.NewEncoder(w)

	var key string = r.Header.Get("key")

	var value, err = manager.Get(key)

	if !err {
		if value != "" {
			network.OkStatus(w)

			utility.SetETag("true", w)

			writer.Encode(factory.NewDataResponse(key, value))
		} else {
			network.NotFoundStatus(w)

			utility.SetETag("false", w)

			writer.Encode(factory.NewErrorResponse(404, "The value doesn't exist!"))
		}

		return
	} else {
		network.NotFoundStatus(w)

		utility.SetETag("false", w)

		writer.Encode(factory.NewErrorResponse(404, "The key is missing!"))
	}
}
