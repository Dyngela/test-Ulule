package main

import (
	"Ulule/src/controller"
	"Ulule/src/exception"
	"Ulule/src/utils"
	"github.com/gin-gonic/gin"
)

const port = ":8080"

func init() {
	gin.SetMode(gin.ReleaseMode)
	gin.ForceConsoleColor()
	utils.ConnectToPostgres()
	router := gin.Default()
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	exception.CheckErrors(err, "error while setting up trusted proxy")
	controller.ContributionController(router)
	err = router.Run(port)
	exception.CheckErrors(err, "error while running")

}

func main() {

}
