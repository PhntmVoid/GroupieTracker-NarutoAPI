package controllers

import (
	temp "Naruto/assets/Templates"
	"log"
	"net/http"
	"strconv"
)

func ClansController(w http.ResponseWriter, r *http.Request) {
	var allClans []Clan

	for page := 1; page <= maxPages; page++ {
		data, err := getFromApi[ClansData](clansURl, page, charactersPerPage)
		if err != nil {
			log.Printf("Error fetching clans page %d: %v", page, err)
			break
		}

		if len(data.Clans) == 0 {
			break
		}

		allClans = append(allClans, data.Clans...)
	}

	if len(allClans) == 0 {
		http.Error(w, "No clans found", http.StatusInternalServerError)
		return
	}

	allClans = getUniqueClans(allClans)

	query := r.URL.Query().Get("search")
	filteredClans := searchClans(allClans, query)

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	itemsPerPage := 12
	totalItems := len(filteredClans)
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
		Clans       []Clan
		Query       string
		CurrentPage int
		TotalPages  int
		TotalClans  int
	}{
		Clans:       filteredClans[start:end],
		Query:       query,
		CurrentPage: page,
		TotalPages:  totalPages,
		TotalClans:  totalItems,
	}

	err = temp.Temp.ExecuteTemplate(w, "clansPage.html", pageData)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
