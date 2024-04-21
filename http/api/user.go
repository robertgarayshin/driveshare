package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robertgarayshin/driveshare/proto/users"
	"net/http"
)

func (h *Handler) GetAllUsers(c *gin.Context) {
	usersResponse, err := h.GetAll(c, &users.Request{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"users": usersResponse.Users,
	})
}

func (h *Handler) GetUserById(c *gin.Context) {
	var req *users.User

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})

		return
	}

	userResponse, err := h.Get(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user": userResponse.User,
	})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	var req *users.User
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})

		return
	}

	response, err := h.DeleteProfile(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("user %d successfully deleted", response.User.Id),
	})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var req *users.User
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
		return
	}

	response, err := h.EditProfile(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("user %d successfully updated", response.User.Id),
	})
}
