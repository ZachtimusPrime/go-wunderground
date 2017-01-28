package wunderground

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// current_observation holds all the weather information, as last updated, when a request is made for the weather
type current_observation struct {
	Weather weather_response `json:"current_observation"`
}

type weather_response struct {
	Location location `json:"display_location"`
	*wind
	*temperature
}

type location struct {
	Full           string `json:"full"`
	City           string `json:"city"`
	State          string `json:"state"`
	StateName      string `json:"state_name"`
	Country        string `json:"country"`
	CountryISO3166 string `json:"country_iso3166"`
	Zipcode        string `json:"zip"`
	Latitude       string `json:"latitude"`
	Longitude      string `json:"longitude"`
	Elevation      string `json:"elevation"`
}

type wind struct {
	Description string  `json:"wind_string"`
	Direction   string  `json:"wind_dir"`
	Degrees     float32 `json:"wind_degrees"`
	MPH         float32 `json:"wind_mph"`
	GustMPH     string  `json:"wind_gust_mph"`
	KPH         float32 `json:"wind_kph"`
	GustKPH     string  `json:"wind_gust_kph"`
}

type temperature struct {
	TempString string  `json:"temperature_string"`
	Fahrenheit float32 `json:"temp_f"`
	Celcius    float32 `json:"temp_c"`
}

type WundergroundClient struct {
	HTTPClient *http.Client // HTTP client used to communicate with the API
	URL        string
}

func NewClient(httpClient *http.Client, State string, City string, APIKey string) *WundergroundClient {
	// Create a new client
	if httpClient == nil {
		tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}} // turn off certificate checking
		httpClient = &http.Client{Timeout: time.Second * 20, Transport: tr}
	}

	url := "http://api.wunderground.com/api/" + APIKey + "/conditions/q/" + State + "/" + City + ".json"
	fmt.Print(url)

	c := &WundergroundClient{HTTPClient: httpClient, URL: url}

	return c
}

// Client.Log is used to construct a new log event and POST it to the Splunk server.
//
// All that must be provided for a log event are the desired map[string]string key/val pairs. These can be anything
// that provide context or information for the situation you are trying to log (i.e. err messages, status codes, etc).
// The function auto-generates the event timestamp and hostname for you.
func (c *WundergroundClient) GetWeather() (current_observation, error) {

	// make new request
	req, err := http.NewRequest("GET", c.URL, nil)
	req.Header.Add("Accept", "application/json")

	// receive response
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return current_observation{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return current_observation{}, err
	}

	// If statusCode is not good, return error string

	switch res.StatusCode {
	case 200:
		response := current_observation{}
		json.Unmarshal(body, &response)
		return response, nil
	default:
		// Turn response into string and return it
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		responseBody := buf.String()
		err = errors.New(responseBody)
		//log.Print(responseBody)	// print error to screen for checking/debugging

		return current_observation{}, err
	}
}
