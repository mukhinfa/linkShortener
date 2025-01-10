package res

import (
	"encoding/json"
	"net/http"
)

// JSON sends data in JSON format with the specified status code
func JSON(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
