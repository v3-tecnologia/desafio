package err_rest

import "net/http"

type ErrRest struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Err  string `json:"err"`
}

func (e *ErrRest) Error() string {
	return e.Msg
}

func NewBadRequestErr(msg string) *ErrRest {
	return &ErrRest{
		Code: http.StatusBadRequest,
		Msg:  msg,
		Err:  "Err_status_bad_request",
	}
}

func NewInternalServerError(msg string) *ErrRest {
	return &ErrRest{
		Code: http.StatusInternalServerError,
		Msg:  msg,
		Err:  "Err_internal_server_error",
	}
}

func NewUnprocessableEntityError(msg string) *ErrRest {
	return &ErrRest{
		Code: http.StatusUnprocessableEntity,
		Msg:  msg,
		Err:  "Err_status_unprocessable_entity",
	}
}
