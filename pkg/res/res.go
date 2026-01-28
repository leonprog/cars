package res

import (
	"encoding/json"
	"net/http"
)

func JsonRes(writer http.ResponseWriter, data any, statusCode int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	json.NewEncoder(writer).Encode(data)
}

func JsonError(writer http.ResponseWriter, data string, statusCode int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	errorJson := Error{
		Error: data,
	}

	json.NewEncoder(writer).Encode(errorJson)
}
