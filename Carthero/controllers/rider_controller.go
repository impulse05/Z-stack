package controllers

import (
	"Carthero/model"
	"Carthero/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RiderController struct {
	riderService services.RiderService
}

func NewRiderController(riderService services.RiderService) *RiderController {
	return &RiderController{riderService: riderService}
}

func (rc *RiderController) CreateRider(c *gin.Context) {
	var rider model.Rider
	if err := c.ShouldBindJSON(&rider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := rc.riderService.CreateRider(&rider); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rider)
}

func (rc *RiderController) DeleteRider(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rider ID"})
		return
	}
	if err := rc.riderService.DeleteRider(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Rider deleted successfully"})
}

func (rc *RiderController) UpdateRiderState(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rider ID"})
		return
	}
	// Assigned is also present in the params
	state := c.Param("state")
	assigned, err := strconv.ParseBool(state)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rider state"})
		return
	}
	if err := rc.riderService.UpdateRiderState(uint(id), assigned); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Rider state updated successfully"})
}

func (rc *RiderController) GetFreeRiders(c *gin.Context) {
	riders, err := rc.riderService.GetFreeRiders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, riders)
}
