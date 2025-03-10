package controllers

import (
	temp "Naruto/assets/Templates"
	"fmt"
	"log"
	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Printf("404 Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	err := temp.Temp.ExecuteTemplate(w, "homePage.html", nil)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, fmt.Sprintf("Failed to render page: %v", err), http.StatusInternalServerError)
		return
	}
}
