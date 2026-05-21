package controllers

import (
	"cacher/internal/factory"
	"cacher/internal/repository/storage"
	"cacher/pkg"
	"cacher/pkg/network"
	"encoding/json"
	"net/http"
)

func RemoveData(w http.ResponseWriter, r *http.Request) {
	writer := json.NewEncoder(w)

	var requestKey string = r.Header.Get("key")
	if requestKey == "" {
		network.NotFoundStatus(w)

		pkg.SetETag("false", w)

		writer.Encode(factory.NewErrorResponse(404, "The key is missing!"))

		return
	} else {
		var value, err = storage.Get(requestKey)

		if !err {
			network.OkStatus(w)

			pkg.SetETag("true", w)

			storage.Remove(requestKey)

			writer.Encode(factory.NewDataResponse(requestKey, value))

		} else {
			network.NotFoundStatus(w)

			pkg.SetETag("false", w)

			writer.Encode(factory.NewErrorResponse(404, "The data doesn't exist!"))
		}
	}
}
