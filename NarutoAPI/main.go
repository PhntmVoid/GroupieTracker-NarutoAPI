package main

import (
	temp "Naruto/assets/Templates"
	"Naruto/src/routes"
)

func main() {

	temp.InitTemplates()
	routes.InitServer()

}
