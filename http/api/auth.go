package api

import (
	"github.com/gin-gonic/gin"
	proto "github.com/robertgarayshin/driveshare/proto/users"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var user *proto.User

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}

	userResponse, err := h.Create(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": userResponse.User.Id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var user *proto.User

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}

	tokenResponse, err := h.Auth(c, user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": err,
		})
	}

	if len(tokenResponse.Errors) > 0 {
		c.JSON(int(tokenResponse.Errors[0].Code), tokenResponse.Errors[0].Description)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token":    tokenResponse.Token,
		"is_valid": tokenResponse.Valid,
	})
}
