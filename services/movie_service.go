package services

import (
	"github.com/AVtheking/ticketo/models"
	"gorm.io/gorm"
)

type MovieService struct {
	db *gorm.DB
}

func NewMovieService(db *gorm.DB) *MovieService {
	return &MovieService{db: db}
}

func (s *MovieService) GetMovies(page, pageSize int) ([]models.Movie, error) {
	var movies []models.Movie
	var totalCount int64

	if err := s.db.Model(&models.Movie{}).Count(&totalCount).Error; err != nil {
		return nil, err
	}

	offset := (page - 1) * pageSize

	if err := s.db.Limit(pageSize).Offset(offset).Find(&movies).Error; err != nil {
		return nil, err
	}

	return movies, nil
}

func (s *MovieService) GetMovieById(id string) (*models.Movie, error) {
	var movie models.Movie

	result := s.db.First(&movie, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &movie, nil
}

func (s *MovieService) CreateMovie(movie *models.Movie) (*models.Movie, error) {
	result := s.db.Create(movie)
	if result.Error != nil {
		return nil, result.Error
	}

	return movie, nil
}

func (s *MovieService) UpdateMovie(id string, movie *models.Movie) (*models.Movie, error) {
	result := s.db.Model(&models.Movie{}).Where("id = ?", id).Updates(movie)
	if result.Error != nil {
		return nil, result.Error
	}

	return movie, nil
}

func (s *MovieService) DeleteMovie(id string) error {
	result := s.db.Delete(&models.Movie{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
