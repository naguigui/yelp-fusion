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
	baseURI            = "https://api.yelp.com/v3"
	businessesEndpoint = "/businesses/search"
	businessDetailEndpoint = "/businesses/"
)

type Client struct {
	AccessToken string
}

// Init creates a new Yelp client to interface with Yelp API
func Init(accessToken string) (*Client, error) {
	if accessToken == "" {
		return nil, errors.New("access token is required but not provided")
	}

	return &Client{AccessToken: accessToken}, nil

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

// BusinessDetail dispatches a request to the Yelp Business Detail API
func (c *Client) BusinessDetails(id string, locale string) (res *BusinessDetailsResponse, err error) {
	params := make(map[string]interface{})

	if locale == "" {
		params["locale"] = locale
	}

	if id == "" {
		return &BusinessDetailsResponse{}, errors.New("id is required")
	}

	err = c.dispatchRequest(fmt.Sprintf("%s%s", businessDetailEndpoint, id), params, &res)

	if err != nil {
		return &BusinessDetailsResponse{}, err
	}

	return res, nil

}

// dispatchRequest formats request and dispatches it to Yelp API
func (c *Client) dispatchRequest(endpoint string, params map[string]interface{}, payload interface{}) error {

	httpClient := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", baseURI, endpoint), nil)

	q := req.URL.Query()

	for key, val := range params {
		q.Add(key, fmt.Sprint(val))
	}

	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))

	res, err := httpClient.Do(req)

	defer res.Body.Close()

	if err != nil {
		return err
	}

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
