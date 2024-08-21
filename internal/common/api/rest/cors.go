package rest

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupCORS(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Origin", "Accept", "Content-Type", "Authorization", "User-Agent"},
		ExposeHeaders:   []string{"Content-Length"},
	}))
}
