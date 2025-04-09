package services

import (
	"errors"
	"strconv"
	"strings"

	"github.com/AVtheking/ticketo/models"
	"gorm.io/gorm"
)

type ReservationService struct {
	db *gorm.DB
}

func NewReservationService(db *gorm.DB) *ReservationService {
	return &ReservationService{db: db}
}

func (s *ReservationService) CreateReservation(reservation *models.Reservation) (*models.Reservation, error) {
	var existingReservation []models.Reservation
	if err := s.db.Where("show_time_id = ? AND user_id = ?", reservation.ShowTimeID, reservation.UserID).First(&existingReservation).Error; err != nil {
		return nil, err
	}

	bookedSeats := map[int]bool{}
	for _, seat := range existingReservation {
		seatNumbers := strings.Split(seat.SeatNumbers, ",")
		for _, seatNumber := range seatNumbers {
			seat, err := strconv.Atoi(seatNumber)
			if err != nil {
				return nil, err
			}

			bookedSeats[seat] = true
		}
	}

	seatNumbers := strings.Split(reservation.SeatNumbers, ",")
	for _, seatNumber := range seatNumbers {
		seat, err := strconv.Atoi(seatNumber)
		if err != nil {
			return nil, err
		}
		if bookedSeats[seat] {
			return nil, errors.New("seat already booked")
		}
	}

	if err := s.db.Create(reservation).Error; err != nil {
		return nil, err
	}
	return reservation, nil
}

func (s *ReservationService) GetAvailableSeats(showTimeID uint) ([]int, error) {
	var showTime models.ShowTime
	if err := s.db.First(&showTime, "id = ?", showTimeID).Error; err != nil {
		return nil, err
	}

	var reservations []models.Reservation
	if err := s.db.Where("show_time = ?", showTimeID).Find(&reservations).Error; err != nil {
		return nil, err
	}

	bookedSeats := map[int]bool{}
	for _, reservation := range reservations {
		seatNumbers := strings.Split(reservation.SeatNumbers, ",")
		for _, seatNumber := range seatNumbers {
			seat, err := strconv.Atoi(seatNumber)
			if err != nil {
				return nil, err
			}
			bookedSeats[seat] = true
		}
	}

	availableSeats := []int{}
	for i := 1; i <= showTime.AvailableSeats; i++ {
		if !bookedSeats[i] {
			availableSeats = append(availableSeats, i)
		}
	}
	return availableSeats, nil
}

func (s *ReservationService) GetReservation(reservationID uint) (*models.Reservation, error) {
	var reservation models.Reservation
	if err := s.db.First(&reservation, "id = ?", reservationID).Error; err != nil {
		return nil, err
	}
	return &reservation, nil
}

func (s *ReservationService) GetUserReservations(userID uint) ([]models.Reservation, error) {
	var reservations []models.Reservation
	if err := s.db.Where("user_id = ?", userID).Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}

func (s *ReservationService) GetAllReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation
	if err := s.db.Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}
