package date

import "time"

type Adapter struct {
}

type AdapterInterface interface {
	Value() int64
}

func (d *Adapter) Value() int64 {

	now := time.Now()

	unix := now.Unix()

	return unix
}

func NewDateAdapter() *Adapter {
	return &Adapter{}
}
