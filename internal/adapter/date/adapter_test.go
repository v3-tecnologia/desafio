package date_test

import (
	"github.com/brianvoe/gofakeit/v7"
	adapter "github.com/kevenmiano/v3/internal/test/mock/internal_/adapter/date"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateDateAdapter(t *testing.T) {

	adapterDate := new(adapter.AdapterInterface)

	date := gofakeit.Date()

	expectedUnix := date.Unix()

	adapterDate.On("Value").Return(expectedUnix).Once()

	unix := adapterDate.Value()

	assert.Equal(t, expectedUnix, unix)

}
