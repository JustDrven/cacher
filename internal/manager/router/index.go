package router

import (
	"cacher/internal/security"
	"net/http"
)

func RegisterRouter(requestType string, path string, handler func(http.ResponseWriter, *http.Request), useMiddleware bool) {
	if useMiddleware {
		http.HandleFunc(requestType+" "+path, security.AuthMiddleware(handler))
	} else {
		http.HandleFunc(requestType+" "+path, handler)
	}
}
