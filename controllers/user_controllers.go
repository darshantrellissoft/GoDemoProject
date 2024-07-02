package controllers

import (
	"MyTransactAPP/config"
	"MyTransactAPP/models"
	"MyTransactAPP/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// UpdateUser godoc
// @Summary Update user details
// @Description Update the user's first name and last name
// @Tags user
// @Accept json
// @Produce json
// @Param user body utils.UpdateUserInput true "User details"
// @Success 200 {string} string "User updated successfully"
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Server error"
// @Security BearerAuth
// @Router /user/update [put]
func UpdateUser(c *gin.Context) {
	var input utils.UpdateUserInput
	updateLogger := config.Log.WithFields(logrus.Fields{
		"API handler": "UpdateUser",
	})

	// Bind JSON to input model
	if err := c.ShouldBindJSON(&input); err != nil {
		updateLogger.Error("Invalid payload")
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid payload", gin.H{"error": err.Error()})
		return
	}

	// Fetch the user by email
	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		updateLogger.Error("User not found")
		utils.JSONResponse(c, http.StatusNotFound, input.Email, gin.H{"error": "User not found"})
		return
	}

	// Update user fields
	user.FirstName = input.FirstName
	user.LastName = input.LastName

	// Save updated user in database
	if err := config.DB.Save(&user).Error; err != nil {
		updateLogger.Error("Failed to update user")
		utils.JSONResponse(c, http.StatusInternalServerError, "Failed to update user", gin.H{"error": err.Error()})
		return
	}

	updateLogger.Info("User updated successfully")
	utils.JSONResponse(c, http.StatusOK, "User updated successfully", gin.H{"message": "User updated successfully"})
}
