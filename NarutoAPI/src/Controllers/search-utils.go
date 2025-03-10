package controllers

import "strings"

type SearchResult struct {
	Item    interface{}
	Score   int
	Matches []string
}

func searchCharacters(characters []Character, query string) []Character {
	if query == "" {
		return characters
	}

	var results []SearchResult
	query = strings.ToLower(query)

	for _, char := range characters {
		score := 0
		matches := []string{}

		if strings.Contains(strings.ToLower(char.Name), query) {
			score += 1
			matches = append(matches, "Name")
		}

		if clan, ok := char.Personal.Clan.(string); ok {
			if strings.Contains(strings.ToLower(clan), query) {
				score += 2
				matches = append(matches, "clan")
			}
		}

		for _, nature := range char.NatureType {
			if strings.Contains(strings.ToLower(nature), query) {
				score += 1
				matches = append(matches, "nature type")
				break
			}
		}

		if kg, ok := char.Personal.KekkeiGenkai.(string); ok {
			if strings.Contains(strings.ToLower(kg), query) {
				score += 2
				matches = append(matches, "kekkei genkai")
			}
		}

		for _, j := range char.Jutsu {
			if strings.Contains(strings.ToLower(j), query) {
				score += 1
				matches = append(matches, "jutsu")
				break
			}
		}

		if score > 0 {
			results = append(results, SearchResult{Item: char, Score: score, Matches: matches})
		}
	}

	sortSearchResults(results)

	var filteredChars []Character
	for _, result := range results {
		filteredChars = append(filteredChars, result.Item.(Character))
	}

	return filteredChars
}

func searchClans(clans []Clan, query string) []Clan {
	if query == "" {
		return getUniqueClans(clans)
	}

	var results []SearchResult
	query = strings.ToLower(query)
	seenClans := make(map[string]bool)

	for _, clan := range clans {
		if seenClans[clan.Name] {
			continue
		}

		score := 0
		matches := []string{}

		if strings.Contains(strings.ToLower(clan.Name), query) {
			score += 2
			matches = append(matches, "name")
		}

		if score > 0 {
			results = append(results, SearchResult{Item: clan, Score: score, Matches: matches})
			seenClans[clan.Name] = true
		}
	}

	sortSearchResults(results)

	var filteredClans []Clan
	for _, result := range results {
		filteredClans = append(filteredClans, result.Item.(Clan))
	}

	return filteredClans
}

func getUniqueClans(clans []Clan) []Clan {
	seen := make(map[string]bool)
	var uniqueClans []Clan

	for _, clan := range clans {
		if !seen[clan.Name] {
			seen[clan.Name] = true
			uniqueClans = append(uniqueClans, clan)
		}
	}
	return uniqueClans
}

func searchKekkeiGenkai(kgs []KekkeiGenkai, query string) []KekkeiGenkai {
	if query == "" {
		return kgs
	}

	var results []SearchResult
	query = strings.ToLower(query)

	for _, kg := range kgs {
		score := 0
		matches := []string{}

		if strings.Contains(strings.ToLower(kg.Name), query) {
			score += 2
			matches = append(matches, "name")
		}

		if score > 0 {
			results = append(results, SearchResult{Item: kg, Score: score, Matches: matches})
		}
	}

	sortSearchResults(results)

	var filteredKGs []KekkeiGenkai
	for _, result := range results {
		filteredKGs = append(filteredKGs, result.Item.(KekkeiGenkai))
	}

	return filteredKGs
}

func searchTailedBeasts(tbs []TailedBeast, query string) []TailedBeast {
	if query == "" {
		return tbs
	}

	var results []SearchResult
	query = strings.ToLower(query)

	for _, tb := range tbs {
		score := 0
		matches := []string{}

		if strings.Contains(strings.ToLower(tb.Name), query) {
			score += 3
			matches = append(matches, "name")
		}

		for _, nature := range tb.NatureType {
			if strings.Contains(strings.ToLower(nature), query) {
				score += 1
				matches = append(matches, "nature type")
				break
			}
		}

		// Search in unique traits
		for _, trait := range tb.UniqueTraits {
			if strings.Contains(strings.ToLower(trait), query) {
				score += 1
				matches = append(matches, "unique trait")
				break
			}
		}

		for _, jin := range tb.Personal.Jinchuriki {
			if strings.Contains(strings.ToLower(jin), query) {
				score += 2
				matches = append(matches, "jinchuriki")
				break
			}
		}

		if score > 0 {
			results = append(results, SearchResult{Item: tb, Score: score, Matches: matches})
		}
	}

	sortSearchResults(results)

	var filteredTBs []TailedBeast
	for _, result := range results {
		filteredTBs = append(filteredTBs, result.Item.(TailedBeast))
	}

	return filteredTBs
}

func sortSearchResults(results []SearchResult) {
	for i := 0; i < len(results)-1; i++ {
		for j := i + 1; j < len(results); j++ {
			if results[i].Score < results[j].Score {
				results[i], results[j] = results[j], results[i]
			}
		}
	}
}
