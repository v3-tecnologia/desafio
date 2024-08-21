package log

import (
	"fmt"
	"log/slog"
)

func InternalErrLog(method string, location string, err error) {
	slog.Error(fmt.Sprintf("[%s - %s] %s", method, location, err.Error()))
}
