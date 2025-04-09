package services

import (
	"errors"
	"time"

	"github.com/AVtheking/ticketo/models"
	"gorm.io/gorm"
)

type ShowTimeService struct {
	db *gorm.DB
}

func NewShowTimeService(db *gorm.DB) *ShowTimeService {
	return &ShowTimeService{db: db}
}

func (s *ShowTimeService) CheckIfShowTimeExists(id interface{}) (*models.ShowTime, error) {
	var showTime models.ShowTime
	if err := s.db.First(&showTime, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &showTime, nil
}

func (s *ShowTimeService) CheckConflictingShowTimes(showTime models.ShowTime) (int64, error) {
	var conflictCount int64
	if err := s.db.Model(&models.ShowTime{}).Where("? < end_time AND ? > start_time", showTime.EndTime, showTime.StartTime).Count(&conflictCount).Error; err != nil {
		return 0, err
	}
	return conflictCount, nil
}

func (s *ShowTimeService) ScheduleShowTime(showTime models.ShowTime) (*models.ShowTime, error) {
	var existingShowTime models.ShowTime

	if err := s.db.Where("movie_id = ? AND start_time = ?", showTime.MovieID, showTime.StartTime).First(&existingShowTime).Error; err != nil {
		return nil, err
	}

	endTime := showTime.StartTime.Add(time.Duration(showTime.Movie.Duration) * time.Minute)
	showTime.EndTime = endTime

	conflictCount, err := s.CheckConflictingShowTimes(showTime)
	if err != nil {
		return nil, err
	}

	if conflictCount > 0 {
		return nil, errors.New("show time conflicts with existing show times")
	}

	if err := s.db.Create(showTime).Error; err != nil {
		return nil, err
	}

	return &showTime, nil
}

func (s *ShowTimeService) GetShowTimesByMovieID(movieID uint) ([]models.ShowTime, error) {
	var showTimes []models.ShowTime
	if err := s.db.Where("movie_id = ?", movieID).Find(&showTimes).Error; err != nil {
		return nil, err
	}
	return showTimes, nil
}

func (s *ShowTimeService) UpdateShowTime(showTimeID uint, showTime models.ShowTime) (*models.ShowTime, error) {
	_, err := s.CheckIfShowTimeExists(showTimeID)
	if err != nil {
		return nil, err
	}

	conflictCount, err := s.CheckConflictingShowTimes(showTime)
	if err != nil {
		return nil, err
	}

	if conflictCount > 0 {
		return nil, errors.New("show time conflicts with existing show times")
	}
	result := s.db.Model(&models.ShowTime{}).Where("id = ?", showTimeID).Updates(showTime)
	if result.Error != nil {
		return nil, result.Error
	}
	return &showTime, nil
}

func (s *ShowTimeService) DeleteShowTime(id uint) error {
	_, err := s.CheckIfShowTimeExists(id)
	if err != nil {
		return err
	}

	if err := s.db.Delete(&models.ShowTime{}, id).Error; err != nil {
		return err
	}
	return nil
}
