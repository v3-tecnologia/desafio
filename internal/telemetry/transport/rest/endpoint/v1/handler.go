package v1

import (
	"github.com/charmingruby/g3/internal/telemetry/domain/usecase"
	"github.com/gin-gonic/gin"
)

func NewHandler(router *gin.Engine, telemetryService usecase.TelemetryUseCase) *Handler {
	return &Handler{
		router:           router,
		telemetryService: telemetryService,
	}
}

type Handler struct {
	router           *gin.Engine
	telemetryService usecase.TelemetryUseCase
}

func (h *Handler) Register() {
	basePath := "/telemetry"
	v1 := h.router.Group(basePath)
	{
		v1.POST("/photo", h.createPhotoEndpoint)
		v1.POST("/gyroscope", h.createGyroscopeEndpoint)
		v1.POST("/gps", h.createGPSEndpoint)
	}
}
