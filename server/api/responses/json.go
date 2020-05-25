package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSON - returns a well formated response with a status code
func JSON(respWriter http.ResponseWriter, statusCode int, data interface{}) {
	respWriter.WriteHeader(statusCode)
	err := json.NewEncoder(respWriter).Encode(data)
	if err != nil {
		fmt.Fprintf(respWriter, "%s", err.Error())
	}
}

// ERROR - returns a jsonified error response along with a status code
func ERROR(respWriter http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(respWriter, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(respWriter, http.StatusBadRequest, nil)
}
