package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(loc string) (Location, error) {
	url := baseURL + "/location-area/" + loc

    val, ok := c.cache.Get(url)
    if ok {
        location := Location{}
        err := json.Unmarshal(val, &location)
        return location, err
    }

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, nil
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, nil
	}

	location := Location{}
	err = json.Unmarshal(data, &location)
	if err != nil {
		return Location{}, nil
	}
    
    c.cache.Add(url, data)
	return location, nil
}
