package v1

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/charmingruby/g3/internal/common/api/rest"
	"github.com/charmingruby/g3/internal/common/custom_err"
	"github.com/charmingruby/g3/internal/telemetry/domain/dto"
	"github.com/charmingruby/g3/internal/telemetry/transport/rest/endpoint/v1/presenter"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createPhotoEndpoint(c *gin.Context) {
	file, header, err := c.Request.FormFile("photo")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Erro ao receber o arquivo: %v", err))
		return
	}
	defer file.Close()

	fileName := header.Filename

	const maxFileSize = 10 << 20 // 10 MB
	if header.Size > maxFileSize {
		err := fmt.Errorf("file %s exceeds the maximum limit of 10MB", fileName)
		rest.NewPayloadError(c, err)
		return
	}

	ext := strings.ToLower(filepath.Ext(fileName))
	allowedExtensions := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
	}
	if !allowedExtensions[ext] {
		err := fmt.Errorf("file extension %s not permitted", ext)
		rest.NewPayloadError(c, err)
		return
	}

	input := dto.CreatePhotoInputDTO{
		File:     file,
		FileName: fileName,
	}

	output, err := h.telemetryService.CreatePhotoUseCase(input)
	if err != nil {
		validationErr, ok := err.(*custom_err.ErrValidation)
		if ok {
			rest.NewEntityError(c, validationErr)
			return
		}

		rest.NewInternalServerError(c, err)
		return
	}

	mappedPhoto := presenter.DomainPhotoToHTTP(output.Photo)

	rest.NewCreatedResponse(c, "photo", mappedPhoto)
}
