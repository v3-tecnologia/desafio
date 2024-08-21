package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewResponse(c *gin.Context, code int, data any, message string) {
	res := Response{
		Message: message,
		Data:    data,
	}
	c.JSON(code, res)
}

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewCreatedResponse(c *gin.Context, entity string, data any) {
	msg := entity + " created successfully"
	NewResponse(c, http.StatusCreated, data, msg)
}

func NewEntityError(c *gin.Context, err error) {
	NewResponse(c, http.StatusUnprocessableEntity, nil, err.Error())
}

func NewPayloadError(c *gin.Context, err error) {
	NewResponse(c, http.StatusBadRequest, nil, "Payload error: "+err.Error())
}

func NewInternalServerError(c *gin.Context, err error) {
	NewResponse(c, http.StatusInternalServerError, nil, err.Error())
}
