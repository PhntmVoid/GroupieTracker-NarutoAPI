package routes

import (
	controllers "Naruto/src/Controllers"
	"net/http"
)

func clansRoute() {
	http.HandleFunc("/clans", controllers.ClansController)
}
