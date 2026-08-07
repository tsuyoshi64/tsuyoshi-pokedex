package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetLocationAreas(pageURL *string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if data, ok := c.cache.Get(url); ok {
		var locations LocationAreaResponse
		if err := json.Unmarshal(data, &locations); err != nil {
			return LocationAreaResponse{}, fmt.Errorf("could not decode cached body: %w", err)
		}
		return locations, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("could not send request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return LocationAreaResponse{}, fmt.Errorf("request failed: %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("could not read response's body: %w", err)
	}

	c.cache.Add(url, data)

	var locations LocationAreaResponse
	if err := json.Unmarshal(data, &locations); err != nil {
		return LocationAreaResponse{}, fmt.Errorf("could not decode response body: %w", err)
	}

	return locations, nil
}
