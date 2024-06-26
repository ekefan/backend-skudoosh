package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ekefan/backend-skudoosh/internal/amadeus"
)

func (server *Server) amadeusClient(searchParams string) (amadeus.CityAndAirPortSearchResponse, error) {
	//create new amadeus client
	client := amadeus.New(server.config)

	httpClient := &http.Client{Timeout: time.Duration(2) * time.Minute}

	url := fmt.Sprintf("%s%s%s", client.BaseURL, "/city-and-airport-search/:", searchParams)

	resp, err := httpClient.Get(url)
	if err != nil {
		return amadeus.CityAndAirPortSearchResponse{}, fmt.Errorf("could not make request to api server: %v", err)
	}
	defer resp.Body.Close()
	var res amadeus.CityAndAirPortSearchResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return amadeus.CityAndAirPortSearchResponse{}, fmt.Errorf("could not read json data %v", err)
	}
	err = json.Unmarshal(body, &res)
	// fmt.Println(string(body), res)
	if err != nil {
		return amadeus.CityAndAirPortSearchResponse{}, fmt.Errorf("couldn't unmarshall json data %v", err)
	}
	return res, nil
}
