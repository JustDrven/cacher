package controllers

import (
	"cacher/factory"
	"cacher/manager"
	"cacher/utility"
	"cacher/utility/network"
	"encoding/json"
	"net/http"
)

func IsValid(w http.ResponseWriter, r *http.Request) {
	writer := json.NewEncoder(w)
	var key string = r.Header.Get("key")

	if key == "" {
		network.NotFoundStatus(w)

		utility.SetETag("false", w)

		writer.Encode(factory.NewErrorResponse(404, "The key is missing!"))

		return
	}

	if manager.Exist(key) {
		network.OkStatus(w)

		utility.SetETag("true", w)

		writer.Encode(factory.NewValidResponse(true))
	} else {
		network.NotFoundStatus(w)

		utility.SetETag("false", w)

		writer.Encode(factory.NewValidResponse(false))
	}

}
