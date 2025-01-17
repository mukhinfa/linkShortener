package req

import (
	"encoding/json"
	"io"
)

// Decode decodes the request body into the given type.
// It returns the decoded type and an error if the decoding fails.
func Decode[T any](body io.ReadCloser) (T, error) {
	var payload T
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}
