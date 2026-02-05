package handler

import (
	"CleanArchitecture/internal/handler/dto"
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
// @Param request body map[string]string true "Phone number"
// @Success 200 {object} map[string]string
// @Router /auth/send-otp [post]
func SendOtpHandler(c *gin.Context) {
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	usecase := auth.NewAuthUsecase()
	otp := usecase.SendOTP(req.Phone)

	c.JSON(http.StatusOK, gin.H{
		"phone": req.Phone,
		"otp":   otp,
	})

}
