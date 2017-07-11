package routers

import (
	"github.com/gorilla/mux"
	"github.com/kenkoii/Analytics/api/handlers"
)

// SetAnalyticsRoutes initializes routes pertaining to users
func SetAnalyticsRoutes(router *mux.Router) *mux.Router {
	r := mux.NewRouter()

	// POST requests
	r.HandleFunc("/api/v1/analytics", handlers.PostEntriesEndpoint).Methods("POST")
	r.HandleFunc("/api/v1/analytics/userpurchases", handlers.GetUserPurchaseEndpoint).Methods("GET")
	r.HandleFunc("/api/v1/analytics/userproperties", handlers.GetUserPropertiesEndpoint).Methods("GET")
	r.HandleFunc("/api/v1/analytics/userdailyproperties", handlers.GetUserDailyPropertiesEndpoint).Methods("GET")
	r.HandleFunc("/api/v1/analytics/tutorials", handlers.GetTutorialsEndpoint).Methods("GET")
	r.HandleFunc("/api/v1/analytics/stages", handlers.GetStagesEndpoint).Methods("GET")

	router.PathPrefix("/api/v1/analytics").Handler(r)
	return router
}
