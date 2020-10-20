package web

import (
	"encoding/json"
	"net/http"
)

func returnJson(writer http.ResponseWriter, statusCode int, o interface{}) error {
	writer.WriteHeader(statusCode)
	return json.NewEncoder(writer).Encode(o)
}
