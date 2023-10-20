package main

import (
	testRouter "backend_proj_os/routes/test"

	"github.com/gin-gonic/gin"
)

func main() {
	routes := gin.New()

	testRouter.TestRoute(routes)

	routes.Run() // default is port : 8080 , can set port by routes.Run(":3000")
}
