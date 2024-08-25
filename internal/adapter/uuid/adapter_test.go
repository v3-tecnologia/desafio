package uuid_test

import (
	"github.com/brianvoe/gofakeit/v7"
	adapter "github.com/kevenmiano/v3/internal/test/mock/internal_/adapter/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUUIDAdapter(t *testing.T) {

	uuidAdapter := new(adapter.AdapterInterface)

	uuidAdapter.On("Value").Return(gofakeit.UUID()).Once()

	uuid := uuidAdapter.Value()

	assert.NotEmpty(t, uuid)

}
