package controllers

import (
	"cacher/factory"
	"cacher/repository/storage"
	"cacher/utility"
	"cacher/utility/network"
	"encoding/json"
	"net/http"
)

func RemoveData(w http.ResponseWriter, r *http.Request) {
	writer := json.NewEncoder(w)

	var requestKey string = r.Header.Get("key")
	if requestKey == "" {
		network.NotFoundStatus(w)

		utility.SetETag("false", w)

		writer.Encode(factory.NewErrorResponse(404, "The key is missing!"))

		return
	} else {
		var value, err = storage.Get(requestKey)

		if !err {
			network.OkStatus(w)

			utility.SetETag("true", w)

			storage.Remove(requestKey)

			writer.Encode(factory.NewDataResponse(requestKey, value))

		} else {
			network.NotFoundStatus(w)

			utility.SetETag("false", w)

			writer.Encode(factory.NewErrorResponse(404, "The data doesn't exist!"))
		}
	}
}
