package movies

import (
	templateError "backend_proj_os/errors"
	movieService "backend_proj_os/services/movies"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTopMovies(c *gin.Context) {
	countStr := c.Param("count")
	count, e := strconv.Atoi(countStr)
	if e != nil {
		c.JSON(500, e)
	}

	if response, err := movieService.GetTopMovies(count); err != nil {
		statusCode, errResponse := templateError.GetErrorResponse(err)
		c.AbortWithStatusJSON(statusCode, errResponse)

	} else {
		// c.JSON(http.StatusOK, response)
		c.JSON(http.StatusOK, response)
	}
}

func GetMoviesByCategory(c *gin.Context) {
	catName := c.Param("categoryName")

	if response, err := movieService.GetMoviesByCategory(catName); err != nil {
		statusCode, errResponse := templateError.GetErrorResponse(err)
		c.AbortWithStatusJSON(statusCode, errResponse)

	} else {
		// c.JSON(http.StatusOK, response)
		c.JSON(http.StatusOK, response)
	}
}

func GetMoviesByDirectorName(c *gin.Context) {
	directorName := c.Param("directorName")

	if response, err := movieService.GetMoviesByDirector(directorName); err != nil {
		statusCode, errResponse := templateError.GetErrorResponse(err)
		c.AbortWithStatusJSON(statusCode, errResponse)

	} else {
		// c.JSON(http.StatusOK, response)
		c.JSON(http.StatusOK, response)
	}
}
