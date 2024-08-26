package multipart

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/test/dummy"
	adapter "github.com/kevenmiano/v3/internal/test/mock/internal_/adapter/multipart"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultipartWithFomData(t *testing.T) {

	adapterMultipart := new(adapter.AdapterInterface)

	newFileDto := &domain.FileDto{
		Name:    gofakeit.Name(),
		Content: gofakeit.ImagePng(255, 255),
		Ext:     ".png",
	}

	file := domain.NewFile(newFileDto)

	request := dummy.NewMultipart(file)

	adapterMultipart.On("Read", request).Return(file, nil).Once()

	read, _ := adapterMultipart.Read(request)

	assert.Equal(t, file, read)
}

func TestMultipartFormData_Read(t *testing.T) {

	adapterMultipart := new(adapter.AdapterInterface)

	newFileDto := &domain.FileDto{
		Name:    gofakeit.Name(),
		Content: gofakeit.ImagePng(255, 255),
		Ext:     ".png",
	}

	file := domain.NewFile(newFileDto)

	request := dummy.NewMultipart(file)

	adapterMultipart.On("Read", request).Return(file, nil).Once()

	read, _ := adapterMultipart.Read(request)

	assert.Equal(t, file, read)

}
