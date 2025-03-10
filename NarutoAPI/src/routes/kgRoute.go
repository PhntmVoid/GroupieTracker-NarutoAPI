package routes

import (
	controllers "Naruto/src/Controllers"
	"net/http"
)

func kekkeiGenkaiRoute() {
	http.HandleFunc("/kekkei-genkai", controllers.KekkeiGenkaiController)
}
