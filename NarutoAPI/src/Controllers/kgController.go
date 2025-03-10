package controllers

import (
	temp "Naruto/assets/Templates"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func KekkeiGenkaiController(w http.ResponseWriter, r *http.Request) {
	var allKG []KekkeiGenkai

	for page := 1; page <= maxPages; page++ {
		data, err := getFromApi[KekkeiGenkaiData](kekkeiGenkaiURL, page, charactersPerPage)
		if err != nil {
			log.Printf("Error fetching kekkei genkai page %d: %v", page, err)
			break
		}

		if len(data.KekkeiGenkai) == 0 {
			break
		}

		allKG = append(allKG, data.KekkeiGenkai...)
	}

	if len(allKG) == 0 {
		http.Error(w, "No kekkei genkai found", http.StatusInternalServerError)
		return
	}

	query := r.URL.Query().Get("search")
	filteredKG := searchKekkeiGenkai(allKG, query)

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	itemsPerPage := 12
	totalItems := len(filteredKG)
	totalPages := (totalItems + itemsPerPage - 1) / itemsPerPage

	if page > totalPages {
		page = totalPages
	}

	start := (page - 1) * itemsPerPage
	end := start + itemsPerPage
	if end > totalItems {
		end = totalItems
	}

	pageData := struct {
		KekkeiGenkai []KekkeiGenkai
		Query        string
		CurrentPage  int
		TotalPages   int
		TotalKG      int
	}{
		KekkeiGenkai: filteredKG[start:end],
		Query:        query,
		CurrentPage:  page,
		TotalPages:   totalPages,
		TotalKG:      totalItems,
	}

	if err := temp.Temp.ExecuteTemplate(w, "kgPage.html", pageData); err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, fmt.Sprintf("Failed to render page: %v", err), http.StatusInternalServerError)
	}
}
