package routers

import (
	"github.com/gorilla/mux"
	"github.com/kenkoii/Analytics/api/handlers/EnglishStory"
)

// SetUserRoutes initializes routes pertaining to users
func SetUserRoutes(router *mux.Router) *mux.Router {
	r := mux.NewRouter()

	// GET requests
	r.HandleFunc("/analytics/englishstory/userproperties", handlers.GetUserPropertiesEndpoint).Methods("GET")
	r.HandleFunc("/analytics/englishstory/userpurchases", handlers.GetUserPurchasesEndpoint).Methods("GET")
	r.HandleFunc("/analytics/englishstory/userdailyproperties", handlers.GetUserDailyPropertiesEndpoint).Methods("GET")

	// POST requests
	r.HandleFunc("/analytics/englishstory/userproperties", handlers.PostUserPropertyEndpoint).Methods("POST")
	r.HandleFunc("/analytics/englishstory/userpurchases", handlers.PostUserPurchaseEndpoint).Methods("POST")
	r.HandleFunc("/analytics/englishstory/userdailyproperties", handlers.PostUserDailyPropertyEndpoint).Methods("POST")

	r.PathPrefix("/api/v1/words").Handler(r)
	return router
}
