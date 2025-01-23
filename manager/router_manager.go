package manager

import (
	"cacher/security"
	"cacher/utility"
	"net/http"
)

func RegisterRouter(requestType string, path string, handler func(http.ResponseWriter, *http.Request), useMiddleware bool) {
	if useMiddleware {
		http.HandleFunc(requestType+" "+utility.API_VERSION+path, security.AuthMiddleware(handler))
	} else {
		http.HandleFunc(requestType+" "+utility.API_VERSION+path, handler)
	}
}
