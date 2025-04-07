package routers

import "github.com/gin-gonic/gin"

func MovieRouter(ctx *gin.RouterGroup) {
	movieRouter := ctx.Group("/movie")
	{
		movieRouter.GET("", movieController.GetMovies)
		movieRouter.POST("", movieController.PostMovie)
		movieRouter.GET("/:id", movieController.GetMovieById)
		movieRouter.PUT("/:id", movieController.UpdateMovie)
		movieRouter.DELETE("/:id", movieController.DeleteMovie)

	}
}
