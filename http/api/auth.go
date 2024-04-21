package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/robertgarayshin/driveshare/proto"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var user *proto.User
	c.BindJSON(&user)
	test, err := h.Create(context.Background(), user)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": test.User.Id,
	})
}
