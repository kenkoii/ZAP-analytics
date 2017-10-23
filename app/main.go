package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenkoii/Analytics/api/controllers"
	"github.com/kenkoii/Analytics/api/middlewares"
	"github.com/kenkoii/Analytics/api/routers"
)

func init() {
	http.Handle("/", GetMainEngine())
}

func GetMainEngine() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())
	// router.GET("/", handlers.Greetings)
	router.LoadHTMLGlob("templates/*")
	router.Static("assets", "assets")
	router.GET("/", controllers.HomePageController)
	router.POST("/", controllers.ResultPageController)
	router = routers.InitGinRoutes(router)
	return router
}
