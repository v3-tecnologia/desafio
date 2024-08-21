package port

import (
	"io"
)

type StoragePort interface {
	SaveFile(file io.Reader, fileName string) error
}
