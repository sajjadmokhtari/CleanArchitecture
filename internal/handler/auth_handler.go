package handler

import (
	"CleanArchitecture/internal/handler/dto"
	"CleanArchitecture/internal/usecase/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// لایه‌ی هندلر (ورودی HTTP)
type AuthHandler struct {
    authUsecase *auth.AuthUsecase // یوزکیس احراز هویت
}

// سازنده‌ی هندلر
func NewAuthHandler(authUsecase *auth.AuthUsecase) *AuthHandler {
    return &AuthHandler{
        authUsecase: authUsecase, // تزریق یوزکیس
    }
}



// SendOtpHandler godoc
// @Summary Send OTP
// @Description Send OTP to phone number
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.SendOtpRequest true "Phone number"
// @Success 200 {object} map[string]string
func (h *AuthHandler) SendOtpHandler(c *gin.Context) {
	var req dto.SendOtpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	otp, err := h.authUsecase.SendOTP(req.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send otp"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "otp sent successfully",
		"otp":     otp, // اینجا در Swagger هم نمایش داده می‌شود
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
func (h *AuthHandler) VerifyOtpHandler(c *gin.Context) {
	var req dto.VerifyOtpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// 1️⃣ فقط Usecase صدا زده می‌شود
	user, err := h.authUsecase.VerifyOTPAndCreateUser(req.Phone, req.OTP)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	// 🔥 بعداً می‌توانیم JWT بسازیم
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"user":   user,
	})
}
