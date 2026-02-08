package handler

import (
	"CleanArchitecture/internal/handler/dto"
	"CleanArchitecture/internal/infrastructure/redis"
	"CleanArchitecture/internal/usecase/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

var req dto.SendOtpRequest

// SendOtpHandler godoc
// @Summary Send OTP
// @Description Send OTP to phone number
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.SendOtpRequest true "Phone number"
// @Success 200 {object} map[string]string
// @Router /auth/send-otp [post]
func SendOtpHandler(c *gin.Context) {
	var req dto.SendOtpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	redisClient := redis.NewRedisClient()
	otpRepo := redis.NewOtpRedisRepository(redisClient)
	usecase := auth.NewAuthUsecase(otpRepo)

	otpCode, err := usecase.SendOTP(req.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send otp"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"phone": req.Phone,
		"otp":   otpCode, // بعداً حذف میشه
	})
}

// VerifyOtpHandler godoc
// @Summary Verify OTP
// @Description Verify OTP code
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.VerifyOtpRequest true "Verify OTP"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/verify-otp [post]
func VerifyOtpHandler(c *gin.Context) {
	var req dto.VerifyOtpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	redisClient := redis.NewRedisClient()
	otpRepo := redis.NewOtpRedisRepository(redisClient)
	usecase := auth.NewAuthUsecase(otpRepo)

	err := usecase.VerifyOTP(req.Phone, req.OTP)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
