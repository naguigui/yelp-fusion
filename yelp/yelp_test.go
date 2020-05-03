package yelp_test

import (
	"fmt"
	"github.com/naguigui/yelp-fusion/yelp"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	httpClient *http.Client
	server     *httptest.Server
)

var businessSearchSuccessResponse = `
{
  "total": 1,
  "businesses": [
    {
      "rating": 4.5,
      "price": "$$$$",
      "phone": "+14157492060",
      "id": "WavvLdfdP6g8aZTtbBQHTw",
      "alias": "gary-danko-san-francisco",
      "categories": [
        {
          "alias": "newamerican",
          "title": "American (New)"
        }
      ],
      "review_count": 4525,
      "name": "Gary Danko",
      "url": "https://www.yelp.com/biz/gary-danko-san-francisco",
      "coordinates": {
        "latitude": 37.80587,
        "longitude": -122.42058
      },
      "image_url": "https://s3-media3.fl.yelpcdn.com/bphoto/--8oiPVp0AsjoWHqaY1rDQ/o.jpg",
      "is_closed": false,
      "location": {
        "city": "San Francisco",
        "country": "US",
        "address2": "",
        "address3": "",
        "state": "CA",
        "address1": "800 N Point St",
        "zip_code": "94109"
      },
      "transactions": ["restaurant_reservation"]
    }
  ]
}
`

func setup() *yelp.Client {
	client, _ := yelp.Init("yelp-key", &http.Client{})

	return client

}

func TestBusinessSearch(t *testing.T) {
	client := setup()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintln(w, businessSearchSuccessResponse)
	}))

	defer ts.Close()

	client.BaseURI = ts.URL

	params := yelp.BusinessSearch{
		Term:     "restaurant",
		Location: "222 Yonge St. Toronto, ON",
	}

	res, err := client.BusinessSearch(params)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.Total, 1)
	assert.Equal(t, res.Businesses[0].ID, "WavvLdfdP6g8aZTtbBQHTw")
	assert.Equal(t, res.Businesses[0].Phone, "+14157492060")
	assert.Equal(t, res.Businesses[0].Alias, "gary-danko-san-francisco")
	assert.Equal(t, res.Businesses[0].Location.Country, "US")
	assert.Equal(t, res.Businesses[0].Location.City, "San Francisco")
	assert.Equal(t, res.Businesses[0].Location.Address1, "800 N Point St")
	assert.Equal(t, res.Businesses[0].Categories[0].Alias, "newamerican")
	assert.Equal(t, res.Businesses[0].Categories[0].Title, "American (New)")
	assert.Equal(t, res.Businesses[0].Coordinates.Latitude, float32(37.80587))
	assert.Equal(t, res.Businesses[0].Coordinates.Longitude, float32(-122.42058))
}

func TestBusinessSearchError(t *testing.T) {
	client := setup()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
	}))

	defer ts.Close()

	client.BaseURI = ts.URL

	params := yelp.BusinessSearch{
		Term:     "restaurant",
		Location: "222 Yonge St. Toronto, ON",
	}

	_, err := client.BusinessSearch(params)

	assert.EqualError(t, err, "400 Bad Request")
}
