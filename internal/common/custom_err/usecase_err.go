package custom_err

func NewInternalErr() error {
	return &ErrValidation{
		Message: "internal error",
	}
}

type ErrInternal struct {
	Message string `json:"message"`
}

func (e *ErrInternal) Error() string {
	return e.Message
}
