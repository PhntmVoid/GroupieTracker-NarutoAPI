package routes

import (
	controllers "Naruto/src/Controllers"
	"net/http"
)

func aboutRoute() {
	http.HandleFunc("/about", controllers.AboutController)
}
