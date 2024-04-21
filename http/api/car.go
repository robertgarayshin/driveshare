package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robertgarayshin/driveshare/proto/cars"
	"net/http"
)

func (h *Handler) NewCar(c *gin.Context) {
	var req *cars.Car
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
		return
	}

	resp, err := h.CreateCar(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"car": resp.Car,
	})
}

func (h *Handler) GetCar(c *gin.Context) {
	var req *cars.Car
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
		return
	}

	car, err := h.GetCarById(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"car": car.Car,
	})
}

func (h *Handler) ListCars(c *gin.Context) {
	cars, err := h.GetAllCars(c, &cars.CarRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"cars": cars.Cars,
	})
}

func (h *Handler) EditCarRequest(c *gin.Context) {
	var req *cars.Car
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
		return
	}

	car, err := h.EditCar(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"car": car.Car,
	})
}

func (h *Handler) DeleteCarRequest(c *gin.Context) {
	var req *cars.Car
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
		return
	}

	car, err := h.DeleteCar(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("car %d successfully deleted", car.Car.Id),
	})
}
