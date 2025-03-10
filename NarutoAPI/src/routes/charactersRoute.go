package routes

import (
	controllers "Naruto/src/Controllers"
	"net/http"
)

func charactersRoute() {
	http.HandleFunc("/characters", controllers.CharactersController)
}
