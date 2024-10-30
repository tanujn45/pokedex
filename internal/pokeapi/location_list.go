package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreas{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, nil
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreas{}, nil
	}

	locationAreas := LocationAreas{}
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return LocationAreas{}, nil
	}

	return locationAreas, nil
}
