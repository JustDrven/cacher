package factory

import "cacher/data"

func NewErrorResponse(error int, msg string) data.ErrorResponse {
	return data.ErrorResponse{
		Error:   error,
		Message: msg,
	}
}

func NewDataResponse(key string, value string) data.Data {
	return data.Data{
		Key:   key,
		Value: value,
	}
}

func NewValidResponse(ok bool) data.Valid {
	return data.Valid{
		Ok: ok,
	}
}
