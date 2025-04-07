package routers

import "github.com/gin-gonic/gin"

func CinemaRouter(ctx *gin.RouterGroup) {
	cinemaRouter := ctx.Group("/cinema")
	{
		cinemaRouter.GET("", cinemaController.GetCinemas)
		cinemaRouter.POST("", cinemaController.AddCinema)
		cinemaRouter.GET("/:id", cinemaController.GetCinemaById)
		cinemaRouter.PUT("/:id", cinemaController.UpdateCinema)
		cinemaRouter.DELETE("/:id", cinemaController.DeleteCinema)
		cinemaRouter.GET("/:id/movies", cinemaController.GetMoviesByCinemaId)
		
	}
}
