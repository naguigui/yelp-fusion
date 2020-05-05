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
	baseURI                           = "https://api.yelp.com/v3"
	businessesEndpoint                = "/businesses"
	businessesSearchEndpoint          = "/search"
	businessSearchPhoneEndpoint       = "/search/phone"
	businessReviewsEndpoint           = "/reviews"
	businessTransactionSearchEndpoint = "/transactions/delivery/search"
)

// Client is responsible for dispatching requests to the Yelp Fusion API via its methods.
// An instance is created from Init()
type Client struct {
	APIKey     string
	HTTPClient *http.Client
	BaseURI    string
}

// Options is provided as an argument to create an instance of the Client.
// It provides the Yelp API Key needed to authenticate API requests
type Options struct {
	APIKey     string
	HTTPClient *http.Client
}

// Init creates a new Yelp Client to interface with Yelp API.
func Init(o *Options) (*Client, error) {
	if o.APIKey == "" {
		return nil, errors.New("access token is required but not provided")
	}

	if o.HTTPClient == nil {
		o.HTTPClient = http.DefaultClient
	}

	return &Client{APIKey: o.APIKey, BaseURI: baseURI, HTTPClient: o.HTTPClient}, nil

}

// BusinessSearch dispatches a request to the Yelp Business Search API.
func (c *Client) BusinessSearch(business BusinessSearch) (res *BusinessSearchResponse, err error) {
	params, err := utility.StructToMap(business)

	if err != nil {
		return &BusinessSearchResponse{}, fmt.Errorf("unable to process business params: %v", err)
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

	err = c.dispatchRequest(fmt.Sprintf("%s%s", businessesEndpoint, businessesSearchEndpoint), filteredParams, &res)
	if err != nil {
		return &BusinessSearchResponse{}, err
	}
	return res, nil
}

// BusinessDetails dispatches a request to the Yelp Business Detail API.
func (c *Client) BusinessDetails(id string, locale string) (res *BusinessDetailsResponse, err error) {
	if id == "" {
		return &BusinessDetailsResponse{}, errors.New("id is required")
	}

	params := make(map[string]interface{})

	if locale != "" {
		params["locale"] = locale
	}

	err = c.dispatchRequest(fmt.Sprintf("%s/%s", businessesEndpoint, id), params, &res)
	if err != nil {
		return &BusinessDetailsResponse{}, err
	}
	return res, nil
}

// BusinessPhoneSearch dispatches a request to the Yelp Phone Search API.
func (c *Client) BusinessPhoneSearch(phoneNumber string, locale string) (res *BusinessPhoneSearchResponse, err error) {
	if phoneNumber == "" {
		return &BusinessPhoneSearchResponse{}, errors.New("phone number is required")
	}

	params := make(map[string]interface{})

	params["phone"] = phoneNumber

	if locale != "" {
		params["locale"] = locale
	}

	err = c.dispatchRequest(fmt.Sprintf("%s%s", businessesEndpoint, businessSearchPhoneEndpoint), params, &res)
	if err != nil {
		return &BusinessPhoneSearchResponse{}, err
	}
	return res, nil
}

// BusinessReviews dispatches a request to the Yelp Business Reviews API.
func (c *Client) BusinessReviews(id string, locale string) (res *BusinessReviewsResponse, err error) {
	if id == "" {
		return &BusinessReviewsResponse{}, errors.New("business id is required")
	}

	params := make(map[string]interface{})

	if locale != "" {
		params["locale"] = locale
	}

	err = c.dispatchRequest(fmt.Sprintf("%s/%s%s", businessesEndpoint, id, businessReviewsEndpoint), params, &res)
	if err != nil {
		return &BusinessReviewsResponse{}, err
	}
	return res, nil
}

// TransactionSearch dispatches a request to the Yelp Business Transaction Search API.
// Default value for transaction type is delivery.
func (c *Client) TransactionSearch(b *BusinessTransactionRequest) (res *BusinessTransactionSearchResponse, err error) {
	params := make(map[string]interface{})

	// Use location if specified, otherwise use latitude/longitude
	if b.Location != "" {
		params["location"] = b.Location
	} else {
		if b.Latitude == 0 || b.Longitude == 0 {
			return &BusinessTransactionSearchResponse{}, errors.New("latitude and longitude is required if location is not specified")
		}
		params["latitude"] = b.Latitude
		params["longitude"] = b.Longitude
	}

	err = c.dispatchRequest(businessTransactionSearchEndpoint, params, &res)
	if err != nil {
		return &BusinessTransactionSearchResponse{}, err
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
