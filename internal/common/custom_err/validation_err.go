package custom_err

import "fmt"

func NewValidationErr(msg string) error {
	return &ErrValidation{
		Message: msg,
	}
}

type ErrValidation struct {
	Message string `json:"message"`
}

func (e *ErrValidation) Error() string {
	return e.Message
}

func NewRequiredErrMessage(field string) string {
	return fmt.Sprintf("%s cannot be blank", field)
}
