// Package yelp consists of wrapper functions to interface with the Yelp v3 Fusion API.
// The package supports business, event, and category endpoints.
package yelp

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/naguigui/yelp-fusion/yelp/utility"
	"io/ioutil"
	"net/http"
)

const (
	BASE_URI                             = "https://api.yelp.com/v3"
	BUSINESS_ENDPOINT                    = "/businesses"
	BUSINESS_SEARCH_ENDPOINT             = "/search"
	BUSINESS_SEARCH_PHONE_ENDPOINT       = "/search/phone"
	BUSINESS_REVIEWS_ENDPOINT            = "/reviews"
	BUSINESS_TRANSACTION_SEARCH_ENDPOINT = "/transactions/delivery/search"
	BUSINESS_AUTOCOMPLETE_ENDPOINT       = "/autocomplete"
)

// Client is responsible for dispatching requests to the Yelp Fusion API via its methods.
// An instance is created from Init()
type Client struct {
	APIKey     string
	HTTPClient *http.Client
	BaseURI    string
}

// ClientOptions is provided as an argument to create an instance of the Client.
// It provides the Yelp API Key needed to authenticate API requests
type ClientOptions struct {
	APIKey     string
	HTTPClient *http.Client
}

// Init creates a new Yelp Client to interface with Yelp API.
func Init(c *ClientOptions) (*Client, error) {
	if c.APIKey == "" {
		return nil, errors.New("access token is required but not provided")
	}

	if c.HTTPClient == nil {
		c.HTTPClient = http.DefaultClient
	}

	return &Client{APIKey: c.APIKey, BaseURI: BASE_URI, HTTPClient: c.HTTPClient}, nil

}

// BusinessSearch dispatches a request to the Yelp Business Search API.
func (c *Client) BusinessSearch(b BusinessSearchReq) (res BusinessSearchRes, err error) {
	params, err := utility.StructToMap(b)

	if err != nil {
		return BusinessSearchRes{}, fmt.Errorf("unable to process business params: %v", err)
	}

	filteredParams := make(map[string]interface{})

	for k, v := range params {
		switch v := v.(type) {
		case int:
			if v != 0 {
				filteredParams[k] = v
			}
		case float32:
			if v != 0 {
				filteredParams[k] = v
			}
		case string:
			if v != "" {
				filteredParams[k] = v
			}
		}
	}

	if err = c.dispatchRequest(fmt.Sprintf("%s%s", BUSINESS_ENDPOINT, BUSINESS_SEARCH_ENDPOINT), filteredParams, &res); err != nil {
		return BusinessSearchRes{}, err
	}

	return res, nil
}

// BusinessDetails dispatches a request to the Yelp Business Detail API.
func (c *Client) BusinessDetails(id string, locale string) (res BusinessDetailsRes, err error) {
	if id == "" {
		return BusinessDetailsRes{}, errors.New("id is required")
	}

	params := make(map[string]interface{})

	if locale != "" {
		params["locale"] = locale
	}

	if err = c.dispatchRequest(fmt.Sprintf("%s/%s", BUSINESS_ENDPOINT, id), params, &res); err != nil {
		return BusinessDetailsRes{}, err
	}
	return res, nil
}

// BusinessPhoneSearch dispatches a request to the Yelp Phone Search API.
func (c *Client) BusinessPhoneSearch(phoneNumber string, locale string) (res BusinessPhoneSearchRes, err error) {
	if phoneNumber == "" {
		return BusinessPhoneSearchRes{}, errors.New("phone number is required")
	}

	params := make(map[string]interface{})

	params["phone"] = phoneNumber

	if locale != "" {
		params["locale"] = locale
	}

	if err = c.dispatchRequest(fmt.Sprintf("%s%s", BUSINESS_ENDPOINT, BUSINESS_SEARCH_PHONE_ENDPOINT), params, &res); err != nil {
		return BusinessPhoneSearchRes{}, err
	}

	return res, nil
}

// BusinessReviews dispatches a request to the Yelp Business Reviews API.
func (c *Client) BusinessReviews(id string, locale string) (res BusinessReviewsRes, err error) {
	if id == "" {
		return BusinessReviewsRes{}, errors.New("business id is required")
	}

	params := make(map[string]interface{})

	if locale != "" {
		params["locale"] = locale
	}

	if err = c.dispatchRequest(fmt.Sprintf("%s/%s%s", BUSINESS_ENDPOINT, id, BUSINESS_REVIEWS_ENDPOINT), params, &res); err != nil {
		return BusinessReviewsRes{}, err
	}
	return res, nil
}

// TransactionSearch dispatches a request to the Yelp Business Transaction Search API.
// Default value for transaction type is delivery.
func (c *Client) TransactionSearch(b BusinessTransactionReq) (res BusinessTransactionSearchRes, err error) {
	params := make(map[string]interface{})

	// Use location if specified, otherwise use latitude/longitude
	if b.Location != "" {
		params["location"] = b.Location
	} else {
		if b.Latitude == 0 || b.Longitude == 0 {
			return BusinessTransactionSearchRes{}, errors.New("latitude and longitude is required if location is not specified")
		}
		params["latitude"] = b.Latitude
		params["longitude"] = b.Longitude
	}

	if err = c.dispatchRequest(BUSINESS_TRANSACTION_SEARCH_ENDPOINT, params, &res); err != nil {
		return BusinessTransactionSearchRes{}, err
	}

	return res, nil
}

// Autocomplete dispatches a request to the Yelp Autocomplete API.
func (c *Client) Autocomplete(b BusinessAutoCompleteReq) (res BusinessAutocompleteRes, err error) {
	params := make(map[string]interface{})

	if b.Text == "" {
		return BusinessAutocompleteRes{}, errors.New("text is required")
	}

	if b.Coordinates.Latitude == 0 {
		return BusinessAutocompleteRes{}, errors.New("latitude is required")
	}

	if b.Coordinates.Longitude == 0 {
		return BusinessAutocompleteRes{}, errors.New("longitude is required")
	}

	params["text"] = b.Text
	params["latitude"] = b.Coordinates.Latitude
	params["longitude"] = b.Coordinates.Longitude

	if b.Locale != "" {
		params["locale"] = b.Locale
	}

	if err = c.dispatchRequest(BUSINESS_AUTOCOMPLETE_ENDPOINT, params, &res); err != nil {
		return BusinessAutocompleteRes{}, err
	}

	return res, nil
}

// dispatchRequest formats request and dispatches it to Yelp API.
func (c *Client) dispatchRequest(endpoint string, params map[string]interface{}, payload interface{}) error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.BaseURI, endpoint), nil)
	q := req.URL.Query()

	for key, val := range params {
		q.Add(key, fmt.Sprint(val))
	}

	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &payload)
	if err != nil {
		return err
	}

	return err
}
