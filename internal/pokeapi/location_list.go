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

	res, err := c.httpClient.Get(url)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("could not send request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return LocationAreaResponse{}, fmt.Errorf("request failed: %s", res.Status)
	}

	var locations LocationAreaResponse
	if err := json.NewDecoder(res.Body).Decode(&locations); err != nil {
		return LocationAreaResponse{}, fmt.Errorf("could not decode response body: %w", err)
	}

	return locations, nil
}
