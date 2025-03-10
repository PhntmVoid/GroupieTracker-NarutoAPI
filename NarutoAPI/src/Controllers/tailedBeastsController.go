package controllers

import (
	temp "Naruto/assets/Templates"
	"fmt"
	"log"
	"net/http"
)

func TailedBeastsController(w http.ResponseWriter, r *http.Request) {
	var allTB []TailedBeast

	for page := 1; page <= maxPages; page++ {
		data, err := getFromApi[TailedBeastsData](tailedBeastsURL, page, charactersPerPage)
		if err != nil {
			log.Printf("Error fetching tailed beasts page %d: %v", page, err)
			break
		}

		if len(data.TailedBeasts) == 0 {
			break
		}

		allTB = append(allTB, data.TailedBeasts...)
	}

	if len(allTB) == 0 {
		http.Error(w, "No tailed beasts found", http.StatusInternalServerError)
		return
	}

	query := r.URL.Query().Get("search")
	filteredTB := searchTailedBeasts(allTB, query)

	pagData := GetPaginationData(r, len(filteredTB), ItemsPerPage)

	pageData := struct {
		TailedBeasts []TailedBeast
		Query        string
		CurrentPage  int
		TotalPages   int
		TotalTB      int
	}{
		TailedBeasts: filteredTB[pagData.Start:pagData.End],
		Query:        pagData.Query,
		CurrentPage:  pagData.CurrentPage,
		TotalPages:   pagData.TotalPages,
		TotalTB:      pagData.TotalItems,
	}

	if err := temp.Temp.ExecuteTemplate(w, "tailedBeastsPage.html", pageData); err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, fmt.Sprintf("Failed to render page: %v", err), http.StatusInternalServerError)
	}
}
