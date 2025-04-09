package controllers

import (
	"net/http"
	"strconv"

	"github.com/AVtheking/ticketo/models"
	"github.com/AVtheking/ticketo/services"
	"github.com/gin-gonic/gin"
)

type ShowtimeController struct {
	ShowtimeService *services.ShowTimeService
}

func NewShowtimeController(showtimeService *services.ShowTimeService) *ShowtimeController {
	return &ShowtimeController{ShowtimeService: showtimeService}
}

func (c *ShowtimeController) CreateShowtime(ctx *gin.Context) {
	var showtime *models.ShowTime
	if err := ctx.ShouldBindJSON(&showtime); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	showtime, err := c.ShowtimeService.ScheduleShowTime(*showtime)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, showtime)
}

func (c *ShowtimeController) GetShowtimesByMovieID(ctx *gin.Context) {
	movieID := ctx.Param("movieID")
	movieIDUint, err := strconv.ParseUint(movieID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	showtimes, err := c.ShowtimeService.GetShowTimesByMovieID(uint(movieIDUint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, showtimes)
}

func (c *ShowtimeController) UpdateShowtime(ctx *gin.Context) {
	showtimeID := ctx.Param("showtimeID")
	showtimeIDUint, err := strconv.ParseUint(showtimeID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid showtime ID"})
		return
	}

	var showtime *models.ShowTime
	if err := ctx.ShouldBindJSON(&showtime); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	showtime, err = c.ShowtimeService.UpdateShowTime(uint(showtimeIDUint), *showtime)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, showtime)
}

func (c *ShowtimeController) DeleteShowtime(ctx *gin.Context) {
	showtimeID := ctx.Param("showtimeID")
	showtimeIDUint, err := strconv.ParseUint(showtimeID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid showtime ID"})
		return
	}

	err = c.ShowtimeService.DeleteShowTime(uint(showtimeIDUint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Showtime deleted successfully"})
}
