package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createGPSEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
