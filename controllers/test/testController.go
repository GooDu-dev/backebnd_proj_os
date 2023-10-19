package test

import (
	"net/http"

	testService "backend_proj_os/services/test"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}

func TestModel(c *gin.Context) {
	if response, err := testService.GetAllData(); err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(200, response)
	}
}
