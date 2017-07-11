package app

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/kenkoii/Analytics/api/routers"
	"github.com/rs/cors"
)

func init() {
	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"},
	})
	router := routers.InitRoutes()
	router.Handle("/", http.FileServer(http.Dir("./public")))
	n := negroni.Classic()
	handler := c.Handler(router)
	n.UseHandler(handler)
	http.Handle("/", n)
}
