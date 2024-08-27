package util

const (
	GeneralParseError = "PARSE_ERROR"
)

type ErrorUtil struct {
	Field string `json:"field"`
	Error string `json:"error"`
}
