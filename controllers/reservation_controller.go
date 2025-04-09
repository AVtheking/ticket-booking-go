package controllers

import (
	"net/http"
	"strconv"

	"github.com/AVtheking/ticketo/models"
	"github.com/AVtheking/ticketo/services"
	"github.com/gin-gonic/gin"
)

type RegisteredControllers struct {
	ReservationService *services.ReservationService
}

func NewRegisteredControllers(reservationService *services.ReservationService) *RegisteredControllers {
	return &RegisteredControllers{ReservationService: reservationService}
}

func (c *RegisteredControllers) CreateReservation(ctx *gin.Context) {
	var reservation *models.Reservation
	if err := ctx.ShouldBindJSON(&reservation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	reservation.UserID = uint(userId.(float64))

	reservation, err := c.ReservationService.CreateReservation(reservation)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, reservation)
}

func (c *RegisteredControllers) GetAvailableSeats(ctx *gin.Context) {
	showTimeID := ctx.Param("showTimeID")
	showTimeIDUint, err := strconv.ParseUint(showTimeID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid show time ID"})
		return
	}

	availableSeats, err := c.ReservationService.GetAvailableSeats(uint(showTimeIDUint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, availableSeats)
}

func (c *RegisteredControllers) GetReservation(ctx *gin.Context) {
	reservationID := ctx.Param("reservationID")
	reservationIDUint, err := strconv.ParseUint(reservationID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reservation ID"})
		return
	}

	reservation, err := c.ReservationService.GetReservation(uint(reservationIDUint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reservation)
}

func (c *RegisteredControllers) GetUserReservations(ctx *gin.Context) {
	userID := ctx.Param("userID")
	userIDUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	reservations, err := c.ReservationService.GetUserReservations(uint(userIDUint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reservations)
}

func (c *RegisteredControllers) GetAllReservations(ctx *gin.Context) {
	reservations, err := c.ReservationService.GetAllReservations()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reservations)
}
