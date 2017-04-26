package routers

import (
	"github.com/gorilla/mux"
	"github.com/kenkoii/Analytics/api/routers/EnglishStory"
)

// InitRoutes is for initializing all routes/endpoints
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	//Initiate User Routes
	router = routers.SetUserRoutes(router)
	return router
}
