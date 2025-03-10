package routes

import (
	controllers "Naruto/src/Controllers"
	"net/http"
)

func tailedBeastsRoute() {
	http.HandleFunc("/tailed-beasts", controllers.TailedBeastsController)
}
