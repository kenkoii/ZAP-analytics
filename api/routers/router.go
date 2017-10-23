package routers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kenkoii/Analytics/api/handlers"

	"github.com/gorilla/mux"
)

// InitRoutes is for initializing all routes/endpoints
func InitRoutes() *mux.Router {
	log.Println("Initialized!")
	router := mux.NewRouter()
	//Initiate other routes
	router = SetAnalyticsRoutes(router)
	return router
}

func InitGinRoutes(router *gin.Engine) *gin.Engine {
	v1 := router.Group("/api/v1/analytics")
	v1.POST("/", handlers.PostAnalytics)
	v1.GET("/userproperties", handlers.GetUserProperties)
	return router
}
