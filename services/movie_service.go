package services

import (
	"errors"

	"github.com/AVtheking/ticketo/models"
	"gorm.io/gorm"
)

type MovieService struct {
	db *gorm.DB
}

func NewMovieService(db *gorm.DB) *MovieService {
	return &MovieService{db: db}
}

func (s *MovieService) CheckifMovieExists(id string) (bool, error) {
	var existingMovie models.Movie
	if err := s.db.First(&existingMovie, "id = ?", id).Error; err != nil {
		return false, err
	}

	return existingMovie.ID != 0, nil
}

func (s *MovieService) GetMovies(page int, pageSize int) ([]models.Movie, error) {
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
	var existingMovie models.Movie
	result := s.db.Where("title = ? AND description = ? AND year = ? AND cast = ?",
		movie.Title, movie.Description, movie.Year, movie.Cast).First(&existingMovie)

	if result.Error == nil {
		return nil, errors.New("movie with same details already exists")
	}

	result = s.db.Create(movie)
	if result.Error != nil {
		return nil, result.Error
	}

	return movie, nil
}

func (s *MovieService) UpdateMovie(id string, movie *models.Movie) (*models.Movie, error) {
	var existingMovie models.Movie

	if err := s.db.First(&existingMovie, "id = ?", id).Error; err != nil {
		return nil, err
	}

	result := s.db.Model(&models.Movie{}).Where("id = ?", id).Updates(movie)
	if result.Error != nil {
		return nil, result.Error
	}

	return movie, nil
}

func (s *MovieService) DeleteMovie(id string) error {
	_, err := s.CheckifMovieExists(id)
	if err != nil {
		return err
	}

	result := s.db.Delete(&models.Movie{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
