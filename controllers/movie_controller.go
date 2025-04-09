package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/AVtheking/ticketo/models"
	"github.com/AVtheking/ticketo/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MovieController struct {
	movieService *services.MovieService
}

func NewMovieController(movieService *services.MovieService) *MovieController {
	return &MovieController{movieService: movieService}
}

func (c *MovieController) GetMovies(ctx *gin.Context) {
	page := ctx.Query("page")
	pageSize := ctx.Query("pageSize")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}

	movies, err := c.movieService.GetMovies(pageInt, pageSizeInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"movies": movies})
}

func (c *MovieController) GetMovieById(ctx *gin.Context) {
	id := ctx.Param("id")

	movie, err := c.movieService.GetMovieById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"movie": movie})
}

func (c *MovieController) CreateMovie(ctx *gin.Context) {
	var movie models.Movie
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdMovie, err := c.movieService.CreateMovie(&movie)
	if err != nil {
		if err.Error() == "movie with same details already exists" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"movie": createdMovie})
}

func (c *MovieController) UpdateMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	var movie models.Movie
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedMovie, err := c.movieService.UpdateMovie(id, &movie)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"movie": updatedMovie})
}

func (c *MovieController) DeleteMovie(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.movieService.DeleteMovie(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}
