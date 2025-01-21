package api

import (
	"cacher/data"
	"cacher/utility"
	"encoding/json"
	"net/http"
	"os"
)

func SaveData(w http.ResponseWriter, r *http.Request) {
	writer := json.NewEncoder(w)

	var key string = r.Header.Get("key")
	if key == "" {
		w.WriteHeader(http.StatusNotFound)

		utility.SetETag("false", w)

		writer.Encode(data.ErrorResponse{
			Error:   404,
			Message: "The key is missing!",
		})

		return
	}

	var value string = r.Header.Get("value")
	if value == "" {
		w.WriteHeader(http.StatusNotFound)

		utility.SetETag("false", w)

		writer.Encode(data.ErrorResponse{
			Error:   404,
			Message: "The value is missing!",
		})

		return
	}

	if os.Getenv(utility.SOURCE+key) == "" {

		utility.SetETag("true", w)

		os.Setenv(utility.SOURCE+key, value)

		writer.Encode(data.Data{
			Key:   key,
			Value: value,
		})

		return
	} else {
		w.WriteHeader(http.StatusBadRequest)

		utility.SetETag("false", w)

		writer.Encode(data.ErrorResponse{
			Error:   400,
			Message: "The value already exist!",
		})

		return
	}

}

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

			return
		} else {
			w.WriteHeader(http.StatusNotFound)

			utility.SetETag("false", w)

			writer.Encode(data.ErrorResponse{
				Error:   404,
				Message: "The value doesn't exist!",
			})
			return
		}

	} else {
		w.WriteHeader(http.StatusNotFound)

		utility.SetETag("false", w)

		writer.Encode(data.ErrorResponse{
			Error:   404,
			Message: "The key is missing!",
		})

		return
	}
}

func IsValid(w http.ResponseWriter, r *http.Request) {
	writer := json.NewEncoder(w)
	var key string = r.Header.Get("key")

	if key == "" {
		w.WriteHeader(http.StatusNotFound)

		utility.SetETag("false", w)

		writer.Encode(data.ErrorResponse{
			Error:   404,
			Message: "The key is missing!",
		})

		return
	}

	if os.Getenv(utility.SOURCE+key) != "" {
		w.WriteHeader(http.StatusOK)

		utility.SetETag("true", w)

		writer.Encode(data.Valid{
			Ok: true,
		})

		return
	} else {
		w.WriteHeader(http.StatusNotFound)

		utility.SetETag("false", w)

		writer.Encode(data.Valid{
			Ok: false,
		})

		return
	}

}

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

			return
		} else {
			w.WriteHeader(http.StatusNotFound)

			utility.SetETag("false", w)

			writer.Encode(data.ErrorResponse{
				Error:   404,
				Message: "The data doesn't exist!",
			})

			return
		}
	}
}

func ReplaceData(w http.ResponseWriter, r *http.Request) {
	writer := json.NewEncoder(w)

	var key string = r.Header.Get("key")
	if key == "" {
		w.WriteHeader(http.StatusNotFound)

		utility.SetETag("false", w)

		writer.Encode(data.ErrorResponse{
			Error:   404,
			Message: "The key is missing!",
		})

		return
	}

	var value string = r.Header.Get("value")
	if value == "" {
		w.WriteHeader(http.StatusNotFound)

		utility.SetETag("false", w)

		writer.Encode(data.ErrorResponse{
			Error:   404,
			Message: "The value is missing!",
		})

		return
	}

	if os.Getenv(utility.SOURCE+key) != "" {
		os.Setenv(utility.SOURCE+key, value)

		utility.SetETag("true", w)

		writer.Encode(data.Valid{
			Ok: true,
		})

		return
	} else {
		w.WriteHeader(http.StatusBadRequest)

		utility.SetETag("false", w)

		writer.Encode(data.Valid{
			Ok: false,
		})

		return
	}

}
