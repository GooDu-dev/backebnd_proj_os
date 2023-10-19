package test

import (
	testController "backend_proj_os/controllers/test"

	"github.com/gin-gonic/gin"
)

func testRoute(router *gin.Engine) {
	path := "/test"
	router.GET(path, testController.Ping)
	router.GET(path+"/data", testController.TestModel)
}
