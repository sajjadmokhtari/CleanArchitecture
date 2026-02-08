package router

import (
	"CleanArchitecture/internal/handler"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "CleanArchitecture/docs" 

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// مسیر OTP
	r.POST("/auth/send-otp", handler.SendOtpHandler)
	r.POST("/auth/verify-otp", handler.VerifyOtpHandler)


	// مسیر Swagger (اینجا اضافه می‌کنیم فقط یکبار)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
