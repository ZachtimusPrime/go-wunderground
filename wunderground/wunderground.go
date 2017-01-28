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
type currentObservation struct {
	Weather weatherResponse `json:"current_observation"`
}

// weatherResponse breaks all the weather stats into logical groupings (i.e. location data, temperature data, wind data etc.)
type weatherResponse struct {
	Location location `json:"display_location"`
	*wind
	*temperature
}

// location holds all the locational data
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

// wind holds all the data relating to wind speeds and direction
type wind struct {
	Description string  `json:"wind_string"`
	Direction   string  `json:"wind_dir"`
	Degrees     float32 `json:"wind_degrees"`
	MPH         float32 `json:"wind_mph"`
	GustMPH     string  `json:"wind_gust_mph"`
	KPH         float32 `json:"wind_kph"`
	GustKPH     string  `json:"wind_gust_kph"`
}

// temperature holds all the temperature data
type temperature struct {
	TempString string  `json:"temperature_string"`
	Fahrenheit float32 `json:"temp_f"`
	Celsius    float32 `json:"temp_c"`
}

// client manages communication with the Wunderground weather API.
// New client objects should be created using the NewClient function.
//
// The URL field is populated by the parameters supplied to NewClient.
type client struct {
	HTTPClient *http.Client // HTTP client used to communicate with the API
	URL        string
}

// NewClient instantiates a new client to the Wunderground weather API. Instantiating a new client requires the
// state abbreviation and city name for weather data from a city, as well as a personal API key which can be obtained
// for free by registering at https://www.wunderground.com/signup?mode=api_signup
func NewClient(httpClient *http.Client, State string, City string, APIKey string) *client {
	// Create a new client
	if httpClient == nil {
		tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}} // turn off certificate checking
		httpClient = &http.Client{Timeout: time.Second * 20, Transport: tr}
	}

	url := "http://api.wunderground.com/api/" + APIKey + "/conditions/q/" + State + "/" + City + ".json"
	fmt.Print(url)

	c := &client{HTTPClient: httpClient, URL: url}

	return c
}

// GetWeather is used to obtain weather data from the city specified during client instantiation.
func (c *client) GetWeather() (currentObservation, error) {

	// make new request
	req, err := http.NewRequest("GET", c.URL, nil)
	req.Header.Add("Accept", "application/json")

	// receive response
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return currentObservation{}, err
	}

	// read the body to bytes for un-marshalling
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return currentObservation{}, err
	}

	// If statusCode is not good, return error string
	switch res.StatusCode {
	case 200:
		response := currentObservation{}
		json.Unmarshal(body, &response)
		return response, nil
	default:
		// Turn response into string and return it
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		responseBody := buf.String()
		err = errors.New(responseBody)
		//log.Print(responseBody)	// print error to screen for checking/debugging

		return currentObservation{}, err
	}
}
