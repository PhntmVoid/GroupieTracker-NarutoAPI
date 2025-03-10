package routes

import (
	controllers "Naruto/src/Controllers"
	"net/http"
)

func homeRoute() {
	http.HandleFunc("/", controllers.HomeController)
}
