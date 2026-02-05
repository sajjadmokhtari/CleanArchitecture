package router

import (
	"CleanArchitecture/internal/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	auth.POST("/send-otp", handler.SendOtpHandler)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
