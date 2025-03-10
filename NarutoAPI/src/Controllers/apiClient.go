package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func getFromApi[T any](baseURL string, page, limit int) (*T, error) {
	url := baseURL
	if !strings.Contains(url, "?") {
		url += "?"
	} else {
		url += "&"
	}
	url += fmt.Sprintf("page=%d&limit=%d", page, limit)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make API request to %s: %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, resp.Status)
	}

	var data T
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode API response: %v", err)
	}

	return &data, nil
}
