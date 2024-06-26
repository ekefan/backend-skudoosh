package amadeus

import "github.com/ekefan/backend-skudoosh/internal/utils"

// tokenResponse contains either a valid access token
// or an error that occurred while fetching the token
type tokenResponse struct {
	Token string
	Err   error
}

// Client is a client for the Amadeus API.
// It takes care of refreshing the access token regularly
// in the background while serving the currently valid
// token to clients
type Client struct {
	BaseURL     string
	AccessToken chan tokenResponse
}

// Create a new client and start the token refreshing goroutine.
func New(config utils.Config) *Client {
	c := &Client{
		BaseURL:     "https://test.api.amadeus.com/v1",
		AccessToken: make(chan tokenResponse),
	}
	go c.refreshToken(config)
	return c
}


// AuthResponse contains the unmarshaled response from the Amadeus
// authorization API.
// It is a blend of the success and error response, so that we can
// unmarshal the response first and check for success or error later.
type AuthResponse struct {
	AuthSuccessResponse
	AuthErrorResponse
}

type AuthSuccessResponse struct {
	Type            string `json:"type"`
	Username        string `json:"username"`
	ApplicationName string `json:"application_name"`
	ClientID        string `json:"client_id"`
	TokenType       string `json:"token_type"`
	AccessToken     string `json:"access_token"`
	ExpiresIn       int    `json:"expires_in"`
	State           string `json:"state"`
	Scope           string `json:"scope"`
}

type AuthErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	Code             int    `json:"code"`
	Title            string `json:"title"`
}




// type CityAndAirPortSearchResponse struct {
//     Meta MetaData `json:"meta"`
//     Data []Data   `json:"data"`
// }

// type MetaData struct {
//     Count int    `json:"count"`
//     Links Links  `json:"links"`
// }

// type Links struct {
//     Self string `json:"self"`
//     // Next string `json:"next"`
//     // Last string `json:"last"`
// }

// type Data struct {
//     Type          string     `json:"type"`
//     SubType       string     `json:"subType"`
//     Name          string     `json:"name"`
//     DetailedName  string     `json:"detailedName"`
//     ID            string     `json:"id"`
//     Self          Self       `json:"self"`
//     TimeZoneOffset string    `json:"timeZoneOffset"`
//     IataCode      string     `json:"iataCode"`
//     GeoCode       GeoCode    `json:"geoCode"`
//     Address       Address    `json:"address"`
// }

// type Self struct {
//     Href    string    `json:"href"`
//     Methods []Methods `json:"methods"`
// }

// type Methods struct {}

// type GeoCode struct {
//     Latitude  float64 `json:"latitude"`
//     Longitude float64 `json:"longitude"`
// }

// type Address struct {
//     CityName     string `json:"cityName"`
//     CityCode     string `json:"cityCode"`
//     CountryName  string `json:"countryName"`
//     CountryCode  string `json:"countryCode"`
//     RegionCode   string `json:"regionCode"`
// }