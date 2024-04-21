package api

import (
	"github.com/gin-gonic/gin"
	proto "github.com/robertgarayshin/driveshare/proto/users"
	"net/http"
)

func (h *Handler) GetAllUsers(c *gin.Context) {
	usersResponse, err := h.GetAll(c, &proto.Request{})
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
	var req *proto.User

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

}