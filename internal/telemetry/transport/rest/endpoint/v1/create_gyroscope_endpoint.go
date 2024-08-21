package v1

import (
	"github.com/charmingruby/g3/internal/common/api/api_rest"
	"github.com/charmingruby/g3/internal/common/custom_err"
	"github.com/charmingruby/g3/internal/telemetry/domain/dto"
	"github.com/charmingruby/g3/internal/telemetry/transport/rest/endpoint/v1/v1_presenter"
	"github.com/gin-gonic/gin"
)

type CreateGyroscopeRequest struct {
	XPosition float64 `json:"x" binding:"required"`
	YPosition float64 `json:"y" binding:"required"`
	ZPosition float64 `json:"z" binding:"required"`
}

func (h *Handler) createGyroscopeEndpoint(c *gin.Context) {
	var req CreateGyroscopeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api_rest.NewPayloadError(c, err)
		return
	}

	dto := dto.CreateGyroscopeInputDTO{
		XPosition: req.XPosition,
		YPosition: req.YPosition,
		ZPosition: req.ZPosition,
	}

	output, err := h.telemetryService.CreateGyroscopeUseCase(dto)
	if err != nil {
		validationErr, ok := err.(*custom_err.ErrValidation)
		if ok {
			api_rest.NewEntityError(c, validationErr)
			return
		}

		api_rest.NewInternalServerError(c, err)
		return
	}

	mappedGyroscope := v1_presenter.DomainGyroscopeToHTTP(output.Gyroscope)

	api_rest.NewCreatedResponse(c, "gyroscope", mappedGyroscope)
}
