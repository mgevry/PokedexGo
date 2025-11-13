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
	//reads data from the response and puts it into a variable
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