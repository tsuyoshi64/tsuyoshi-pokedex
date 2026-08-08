package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/tsuyoshi64/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      *pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache:      pokecache.NewCache(cacheInterval),
		httpClient: http.Client{Timeout: timeout},
	}
}

func (c *Client) get(url string) ([]byte, error) {
	if data, ok := c.cache.Get(url); ok {
		return data, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("could not send request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, fmt.Errorf("request failed: %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read the response's body: %w", err)
	}

	c.cache.Add(url, data)
	return data, nil
}
