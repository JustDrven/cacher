package network

import "net/http"

func Status(statusCode int, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
}

func OkStatus(w http.ResponseWriter) {
	Status(http.StatusOK, w)
}

func NotFoundStatus(w http.ResponseWriter) {
	Status(http.StatusNotFound, w)
}

func BadRequestStatus(w http.ResponseWriter) {
	Status(http.StatusBadRequest, w)
}

func UnauthorizedStatus(w http.ResponseWriter) {
	Status(http.StatusUnauthorized, w)
}
