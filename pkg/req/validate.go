package req

import (
	"github.com/go-playground/validator/v10"
)

// IsValid validates the given payload.
// It returns an error if the payload is not valid.
func IsValid[T any](payload T) error {
	validate := validator.New()
	err := validate.Struct(payload)
	return err
}
