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

	router.PathPrefix("/api/v1/analytics").Handler(r)
	return router
}
