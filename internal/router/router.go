package router

import (
	"CleanArchitecture/internal/handler"

	_ "CleanArchitecture/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(authHandler *handler.AuthHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/auth/send-otp", authHandler.SendOtpHandler)
	r.POST("/auth/verify-otp", authHandler.VerifyOtpHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
