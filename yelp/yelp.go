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
)

type Client struct {
	AccessToken string
}

// Init creates a new Yelp client to interface with Yelp API
func Init(accessToken string) (*Client, error) {
	if accessToken == "" {
		return nil, fmt.Errorf("access token not provided")
	}

	return &Client{AccessToken: accessToken}, nil

}

// BusinessSearch dispatches a request to the Yelp Business API
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

	_, err = c.dispatchRequest(businessesEndpoint, "", filteredParams, &res)

	if err != nil {
		return &BusinessSearchResponse{}, err
	}

	return res, nil
}

// dispatchRequest formats request and dispatches it to Yelp API
func (c *Client) dispatchRequest(endpoint string, id string, params map[string]interface{}, v interface{}) (statusCode int, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", baseURI, endpoint), nil)

	q := req.URL.Query()

	for key, val := range params {
		q.Add(key, fmt.Sprint(val))
	}

	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))

	res, err := httpClient.Do(req)

	defer res.Body.Close()

	if err != nil {
		return http.StatusInternalServerError, err
	}

	// ensure the request returned a 200
	if res.StatusCode != 200 {
		return res.StatusCode, errors.New(res.Status)
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = json.Unmarshal(data, &v)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return res.StatusCode, err
}
