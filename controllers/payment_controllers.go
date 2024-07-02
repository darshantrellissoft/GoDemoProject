package controllers

import (
	"MyTransactAPP/config"
	"MyTransactAPP/models"
	"MyTransactAPP/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreatePayment godoc
// @Summary Create a new payment
// @Description Create a new payment
// @Tags payment
// @Accept json
// @Produce json
// @Param payment body utils.PaymentInput true "Payment details"
// @Success 200 {string} string "Payment created successfully, confirmation email sent"
// @Failure 400 {string} string "Invalid payload"
// @Failure 500 {string} string "Failed to create payment"
// @Security BearerAuth
// @Router /payments [post]
func CreatePayment(c *gin.Context) {
	var input utils.PaymentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Please provide valid payload", gin.H{"error": err.Error()})
		return
	}

	// Extract user ID from token
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user models.User
	if err := config.DB.First(&user, "id = ?", userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
		return
	}

	// Validate card details (simplified)
	if len(input.CardNumber) != 16 || len(input.CVV) != 3 {
		utils.JSONResponse(c, http.StatusBadRequest, "Please provide valid payload", gin.H{"error": "Invalid card details"})
		return
	}

	// Create payment record
	payment := models.Transaction{
		ID:         uuid.New(),
		CreatedAt:  time.Now(),
		CardNumber: input.CardNumber,
		ExpiryDate: input.ExpiryDate,
		CVV:        input.CVV,
		Amount:     input.Amount,
		Status:     "Pending",
		UserID:     user.ID,
	}

	if err := config.DB.Create(&payment).Error; err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Database error", gin.H{"error": "Failed to create payment"})
		return
	}

	// Send confirmation email with payment link
	paymentLink := "http://localhost:8080/payments/confirm/" + payment.ID.String()
	if err := utils.SendConfirmationEmail(user.Email, paymentLink); err != nil {
		config.Log.Errorf("Failed to send confirmation email: %v", err)
	}

	utils.JSONResponse(c, http.StatusOK, "Confirmation email sent successfully", gin.H{"message": "Payment created successfully, confirmation email sent"})
}

// ConfirmPayment godoc
// @Summary Confirm a payment
// @Description Confirm a payment by transaction ID
// @Tags payment
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {string} string "Payment confirmed successfully"
// @Failure 404 {string} string "Payment not found"
// @Failure 500 {string} string "Failed to update payment status"
// @Router /payments/confirm/{id} [get]
func ConfirmPayment(c *gin.Context) {
	paymentID := c.Param("id")

	var payment models.Transaction
	if err := config.DB.First(&payment, "id = ?", paymentID).Error; err != nil {
		utils.JSONResponse(c, http.StatusNotFound, "failure", gin.H{"error": "Payment not found"})
		return
	}

	// Update payment status
	payment.Status = "Completed"
	if err := config.DB.Save(&payment).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "failure at creating entry", gin.H{"error": "Failed to update payment status"})
		return
	}

	utils.JSONResponse(c, http.StatusOK, "transaction successful", gin.H{
		"message":        "Payment confirmed successfully",
		"transaction_id": payment.ID,
		"payment_status": payment.Status,
	})
}

// GetTransactionDetails godoc
// @Summary Get transaction details
// @Description Get transaction details by transaction ID
// @Tags payment
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} models.Transaction
// @Failure 404 {string} string "Transaction not found"
// @Security BearerAuth
// @Router /payments/{id} [get]
func GetTransactionDetails(c *gin.Context) {
	transactionID := c.Param("id")

	var transaction models.Transaction
	if err := config.DB.First(&transaction, "id = ?", transactionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	utils.JSONResponse(c, http.StatusOK, "transaction data fetched successfully", gin.H{"data": transaction})
}
