package pokeapi

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
)

type LocationInfo struct {
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []areaInfo `json:"results"`
}

type areaInfo struct {
	Name string `json:"name"`
}


func GetPokeData(url string) (results LocationInfo, err error) {

	var pokeData LocationInfo

	resp, err := http.Get(url)
	if err != nil {
		return pokeData, fmt.Errorf("Error: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return pokeData, fmt.Errorf("Error reading body: %w", err)
	}

	err = json.Unmarshal(body, &pokeData)
	if err != nil {
		return pokeData, fmt.Errorf("Error unmarshaling data: %w", err)
	}

	return pokeData, nil
}
