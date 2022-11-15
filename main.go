package main

import (
	_ "Ulule/docs"
	"Ulule/src/controller"
	"Ulule/src/exception"
	"Ulule/src/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

const port = ":8080"

func init() {
	gin.SetMode(gin.ReleaseMode)
	gin.ForceConsoleColor()
	utils.ConnectToPostgres()
}

// @title Ulule TP API documentation
// @version 1.0.0
// @host localhost:8080
// @BasePath /api/v1
func main() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	exception.CheckErrors(err, "error while setting up trusted proxy")
	controller.ContributionController(router)
	controller.VisitController(router)
	controller.ProjectController(router)
	//err = router.Run(port)
	err = router.Run(os.Getenv("APP_PORT"))
	exception.CheckErrors(err, "error while running")
}
