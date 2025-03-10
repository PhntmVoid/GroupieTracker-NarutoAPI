package controllers

import (
	temp "Naruto/assets/Templates"
	"fmt"
	"log"
	"net/http"
)

func AboutController(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		log.Printf("404 Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	err := temp.Temp.ExecuteTemplate(w, "aboutPage.html", nil)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, fmt.Sprintf("Failed to render page: %v", err), http.StatusInternalServerError)
		return
	}
}
