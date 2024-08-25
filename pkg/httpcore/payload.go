package httpcore

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

var ErrInvalidQuery = errors.New("invalid query parameter type")

var validate = validator.New(validator.WithRequiredStructEnabled())

// DecodeBody decodes the request body into a given type T,
// validates it, and returns the decoded payload and any error encountered during decoding.
// It renders the JSON response to the given http.ResponseWriter and http.Request.
func DecodeBody[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	var payload T
	errDecode := render.DecodeJSON(r.Body, &payload)
	errValidate := validate.Struct(payload)
	err := errors.Join(errDecode, errValidate)

	if err != nil {
		return payload, err
	}
	return payload, nil
}
