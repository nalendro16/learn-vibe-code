package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/nalendro16/learn-vibe-code/internal/dto"
	"github.com/nalendro16/learn-vibe-code/internal/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		dto.SendBadRequest(c, "Invalid request body: "+err.Error())
		return
	}

	if err := h.userService.Register(c.Request.Context(), req); err != nil {
		if errors.Is(err, service.ErrEmailAlreadyRegistered) {
			dto.SendConflict(c, "Email is already registered")
			return
		}

		dto.SendInternalError(c, "Failed to process registration")
		return
	}

	dto.SendCreated(c, "User registered successfully", nil)
}
