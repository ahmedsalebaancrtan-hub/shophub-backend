package handlers

import (
	"errors"
	"net/http"

	"github.com/ahmed/shophub-backend/dto"
	service "github.com/ahmed/shophub-backend/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *service.UserService
}

func RegisterUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		Service: userService,
	}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {

	var request dto.RegisterRequest

	// Validate incoming JSON.
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid registration data",
			"error":   err.Error(),
		})
		return
	}

	// Call service.
	response, err := h.Service.CreateUser(
		c.Request.Context(),
		request,
	)

	if err != nil {

		if errors.Is(err, service.ErrEmailAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{
				"success": false,
				"message": "Email already registered",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to register user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User registered successfully",
		"data":    response,
	})
}
