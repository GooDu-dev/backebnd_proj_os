package movies

import (
	moviesController "backend_proj_os/controllers/movies"
	// testController "backend_proj_os/controllers/test"

	"github.com/gin-gonic/gin"
)

func MoviesRouter(router *gin.Engine) {
	path := "/movies"
	router.GET(path+"/top/:count", moviesController.GetTopMovies)
	router.GET(path+"/category/:categoryName", moviesController.GetMoviesByCategory)
	router.GET(path+"/director/:directorName", moviesController.GetMoviesByDirectorName)
}
