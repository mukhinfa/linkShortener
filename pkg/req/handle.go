package req

import (
	"net/http"

	"github.com/muhinfa/linkShortener/pkg/res"
)

// HandleBody handles the request body and returns the decoded type and an error if the decoding fails.
func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		res.JSON(*w, err.Error(), 402)
		return nil, err
	}
	err = IsValid(body)
	if err != nil {
		res.JSON(*w, err.Error(), 402)
		return nil, err
	}
	return &body, nil
}
