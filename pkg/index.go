package pkg

import (
	"encoding/base64"
	"net/http"
)

const (
	SOURCE      string = "CACHER:"
	API_VERSION string = "/v1/"
)

func SetETag(data string, w http.ResponseWriter) {
	hash := base64.StdEncoding.EncodeToString([]byte(data))

	w.Header().Add("ETag", hash)
}
