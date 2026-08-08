package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetLocationAreas(pageURL *string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	data, err := c.get(url)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	var locations LocationAreaResponse
	if err := json.Unmarshal(data, &locations); err != nil {
		return LocationAreaResponse{}, fmt.Errorf("could not decode response body: %w", err)
	}

	return locations, nil
}

func (c *Client) GetLocationAreaPokemons(areaName string) (LocationAreaPokemons, error) {
	url := baseURL + "/location-area/" + areaName

	data, err := c.get(url)
	if err != nil {
		return LocationAreaPokemons{}, err
	}

	var areaPokemons LocationAreaPokemons
	if err := json.Unmarshal(data, &areaPokemons); err != nil {
		return LocationAreaPokemons{}, fmt.Errorf("could not decode response body: %w", err)
	}

	return areaPokemons, nil
}
