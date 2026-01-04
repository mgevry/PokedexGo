package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)
//pokeapi client method
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	
	// getting the specific url for the locations data
	url := baseURL + "/location-area"
	//if we have a specific url argument for ListLocations then we use this instead of the url we got above
	if pageURL != nil {
		url = *pageURL
	}
	//if we have data in the cache, we use that instead 
	if val, ok := c.cache.Get(url); ok {
		cacheLocationResp := RespShallowLocations{}
		err := json.Unmarshal(val, &cacheLocationResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return cacheLocationResp, nil
	}
	//get request on the url we just built, a get request recieves a copy of information from the original source, this does not send the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}
	//sends the http request and returns an http response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()
	//reads data from the response and puts it into a variable, bytes -> JSON
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}
	//decodes json bytes into provided struct
	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}

func (c *Client) ExploreLocation(areaName string) (LocationArea, error){
	url := baseURL + "/location-area/" + areaName

	//cache
	if val, ok := c.cache.Get(url); ok {
		cacheLocationArea := LocationArea{}
		err := json.Unmarshal(val, &cacheLocationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return cacheLocationArea, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	return locationArea, nil
}