package controllers

import (
	temp "Naruto/assets/Templates"
	"fmt"
	"log"
	"net/http"
)

const charactersPerPage = 50
const maxPages = 2000

func FilterCharacters(characters []Character, filters map[string]string) []Character {
	if len(filters) == 0 {
		return characters
	}

	var filteredChars []Character
	for _, char := range characters {
		if matchesAllFilters(char, filters) {
			filteredChars = append(filteredChars, char)
		}
	}

	return filteredChars
}

func matchesAllFilters(char Character, filters map[string]string) bool {
	for filterKey, filterValue := range filters {
		if filterValue == "" {
			continue
		}

		switch filterKey {
		case "sex":
			if char.Personal.Sex != filterValue {
				return false
			}
		case "bloodType":
			if char.Personal.BloodType != filterValue {
				return false
			}
		case "clan":
			clanStr, ok := char.Personal.Clan.(string)
			if !ok {
				clanSlice, ok := char.Personal.Clan.([]string)
				if !ok {
					return false
				}

				found := false
				for _, c := range clanSlice {
					if c == filterValue {
						found = true
						break
					}
				}
				if !found {
					return false
				}
			} else if clanStr != filterValue {
				return false
			}
		case "natureType":
			found := false
			for _, nature := range char.NatureType {
				if nature == filterValue {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		case "kekkeiGenkai":
			kgStr, ok := char.Personal.KekkeiGenkai.(string)
			if !ok {
				kgSlice, ok := char.Personal.KekkeiGenkai.([]string)
				if !ok {
					return false
				}

				found := false
				for _, kg := range kgSlice {
					if kg == filterValue {
						found = true
						break
					}
				}
				if !found {
					return false
				}
			} else if kgStr != filterValue {
				return false
			}
		case "ninjaRank":
			if char.Rank.NinjaRank.PartI != filterValue {
				return false
			}
		}
	}
	return true
}

func CharactersController(w http.ResponseWriter, r *http.Request) {
	var allCharacters []Character

	for page := 1; page <= maxPages; page++ {
		data, err := getFromApi[CharactersData](charactersURL, page, charactersPerPage)
		if err != nil {
			log.Printf("Error fetching characters page %d: %v", page, err)
			break
		}

		if len(data.Characters) == 0 {
			break
		}

		allCharacters = append(allCharacters, data.Characters...)
	}

	if len(allCharacters) == 0 {
		http.Error(w, "No characters found", http.StatusInternalServerError)
		return
	}

	query := r.URL.Query().Get("search")

	filters := map[string]string{
		"sex":          r.URL.Query().Get("sex"),
		"bloodType":    r.URL.Query().Get("bloodType"),
		"clan":         r.URL.Query().Get("clan"),
		"natureType":   r.URL.Query().Get("natureType"),
		"kekkeiGenkai": r.URL.Query().Get("kekkeiGenkai"),
		"ninjaRank":    r.URL.Query().Get("ninjaRank"),
	}

	searchedCharacters := searchCharacters(allCharacters, query)

	filteredCharacters := FilterCharacters(searchedCharacters, filters)

	pagData := GetPaginationData(r, len(filteredCharacters), 15)

	end := pagData.Start + pagData.End
	if end > len(filteredCharacters) {
		end = len(filteredCharacters)
	}

	var displayCharacters []Character
	if pagData.Start < len(filteredCharacters) {
		displayCharacters = filteredCharacters[pagData.Start:end]
	}

	pageData := struct {
		Characters      []Character
		CurrentPage     int
		TotalPages      int
		Query           string
		TotalCharacters int
		Filters         map[string]string
	}{
		Characters:      displayCharacters,
		CurrentPage:     pagData.CurrentPage,
		TotalPages:      pagData.TotalPages,
		Query:           pagData.Query,
		TotalCharacters: pagData.TotalItems,
		Filters:         filters,
	}

	if err := temp.Temp.ExecuteTemplate(w, "charactersPage.html", pageData); err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, fmt.Sprintf("Failed to render page: %v", err), http.StatusInternalServerError)
	}
}
