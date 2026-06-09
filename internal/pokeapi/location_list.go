package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	//if the function passes in a URL then use that... Otherwise it's the above logic.
	if pageURL != nil {
		url = *pageURL
	}

	//added new Cache code here...
	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationsResp, nil
	}

	//end of the cache update

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//return an empty struct...
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		//return an empty struct
		return RespShallowLocations{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		//return an empty struct
		return RespShallowLocations{}, err
	}

	//declare an empty locations resp to be hydrated by the Unmarshal command.
	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		//return an empty struct
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil

}
