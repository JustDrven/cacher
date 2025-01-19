package security

import (
	"net/http"
	"os"
)

func CheckAuthorization(header http.Header) bool {
	if header.Get("X-API-Key") == os.Getenv("API_KEY") {
		return false
	}
	return true
}
