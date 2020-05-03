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
	baseURI                     = "https://api.yelp.com/v3"
	businessesEndpoint          = "/businesses/search"
	businessDetailEndpoint      = "/businesses/"
	businessSearchPhoneEndpoint = "/businesses/search/phone"
)

type Client struct {
	APIKey     string
	HTTPClient *http.Client
	BaseURI    string
}

// Init creates a new Yelp client to interface with Yelp API
func Init(apiKey string, httpClient *http.Client) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("access token is required but not provided")
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{APIKey: apiKey, BaseURI: baseURI, HTTPClient: httpClient}, nil

}

// BusinessSearch dispatches a request to the Yelp Business Search API
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

	err = c.dispatchRequest(businessesEndpoint, filteredParams, &res)
	if err != nil {
		return &BusinessSearchResponse{}, err
	}
	return res, nil
}

// BusinessDetails dispatches a request to the Yelp Business Detail API
func (c *Client) BusinessDetails(id string, locale string) (res *BusinessDetailsResponse, err error) {
	if id == "" {
		return &BusinessDetailsResponse{}, errors.New("id is required")
	}

	params := make(map[string]interface{})

	if locale != "" {
		params["locale"] = locale
	}

	err = c.dispatchRequest(fmt.Sprintf("%s%s", businessDetailEndpoint, id), params, &res)
	if err != nil {
		return &BusinessDetailsResponse{}, err
	}
	return res, nil
}

// BusinessPhoneSearch dispatches a request to the Yelp Phone Search API
func (c *Client) BusinessPhoneSearch(phoneNumber string, locale string) (res *BusinessPhoneSearchResponse, err error) {
	if phoneNumber == "" {
		return &BusinessPhoneSearchResponse{}, errors.New("phone number is required")
	}

	params := make(map[string]interface{})

	if locale != "" {
		params["locale"] = locale
	}

	err = c.dispatchRequest(fmt.Sprintf("%s%s", businessSearchPhoneEndpoint, phoneNumber), params, &res)
	if err != nil {
		return &BusinessPhoneSearchResponse{}, err
	}
	return res, nil
}

// dispatchRequest formats request and dispatches it to Yelp API
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
