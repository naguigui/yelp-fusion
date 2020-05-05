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

const businessSearchResponse = `
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

const businessReviewsResponse = `
{
	"reviews": [
	  {
		"id": "review12345",
		"rating": 5,
		"user": {
		  "id": "user12345",
		  "profile_url": "https://profileurl",
		  "image_url": "https://myjpg.jpg",
		  "name": "Andrew Nguyen"
		},
		"text": "Omg I love Katsuya, I can eat it everyday",
		"time_created": "2020-05-04 00:41:13",
		"url": "https://mockurl.com"
	  }
	],
	"total": 1,
	"possible_languages": ["en"]
}
`

const businessTransactionResponse = `
{
	"total": 144,
	"businesses": [
	  {
		"id": "gR9DTbKCvezQlqvD7_FzPw",
		"alias": "north-india-restaurant-san-francisco",
		"price": "$$",
		"url": "https://www.yelp.com/biz/north-india-restaurant-san-francisco",
		"rating": 4,
		"location": {
		  "zip_code": "94105",
		  "state": "CA",
		  "country": "US",
		  "city": "San Francisco",
		  "address2": "",
		  "address3": "",
		  "address1": "123 Second St"
		},
		"categories": [
		  {
			"alias": "indpak",
			"title": "Indian"
		  }
		],
		"phone": "+14153481234",
		"coordinates": {
		  "longitude": -122.399305736113,
		  "latitude": 37.787789124691
		},
		"image_url": "http://s3-media4.fl.yelpcdn.com/bphoto/howYvOKNPXU9A5KUahEXLA/o.jpg",
		"is_closed": false,
		"name": "North India Restaurant",
		"review_count": 615,
		"transactions": ["pickup", "restaurant_reservation"]
	  }
	]
  }
`

const autocompleteResponse = `
{
	"terms": [
	  {
		"text": "Delivery"
	  }
	],
	"businesses": [
	  {
		"name": "Delfina",
		"id": "YqvoyaNvtoC8N5dA8pD2JA"
	  }
	],
	"categories": [
	  {
		"alias": "delis",
		"title": "Delis"
	  }
	]
}
`

func setup() *yelp.Client {
	client, _ := yelp.Init(&yelp.Options{APIKey: "yelp-key"})

	return client

}

func TestBusinessSearch(t *testing.T) {
	client := setup()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintln(w, businessSearchResponse)
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

func TestBusinessReviewsSuccess(t *testing.T) {
	client := setup()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintln(w, businessReviewsResponse)
	}))

	defer ts.Close()

	client.BaseURI = ts.URL

	res, err := client.BusinessReviews("review12345", "")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.Total, 1)
	assert.Equal(t, res.Reviews[0].ID, "review12345")
	assert.Equal(t, res.Reviews[0].Rating, 5)
	assert.Equal(t, res.Reviews[0].User.ID, "user12345")
	assert.Equal(t, res.Reviews[0].User.ProfileURL, "https://profileurl")
	assert.Equal(t, res.Reviews[0].User.ImageURL, "https://myjpg.jpg")
	assert.Equal(t, res.Reviews[0].Text, "Omg I love Katsuya, I can eat it everyday")
	assert.Equal(t, res.Reviews[0].TimeCreated, "2020-05-04 00:41:13")
	assert.Equal(t, res.Reviews[0].URL, "https://mockurl.com")
}

func TestBusinessReviewsError(t *testing.T) {
	client := setup()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
	}))

	defer ts.Close()

	client.BaseURI = ts.URL

	_, err := client.BusinessReviews("review12345", "")

	assert.EqualError(t, err, "500 Internal Server Error")
}

func TestBusinessTransactionSearchSuccess(t *testing.T) {
	client := setup()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintln(w, businessTransactionResponse)
	}))

	defer ts.Close()

	client.BaseURI = ts.URL

	params := &yelp.BusinessTransactionRequest{
		Latitude:  37.787789124691,
		Longitude: -122.399305736113,
	}

	res, err := client.TransactionSearch(params)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.Total, 144)
	assert.Equal(t, res.Businesses[0].ID, "gR9DTbKCvezQlqvD7_FzPw")
	assert.Equal(t, res.Businesses[0].Phone, "+14153481234")
	assert.Equal(t, res.Businesses[0].Alias, "north-india-restaurant-san-francisco")
	assert.Equal(t, res.Businesses[0].Location.Country, "US")
	assert.Equal(t, res.Businesses[0].Location.City, "San Francisco")
	assert.Equal(t, res.Businesses[0].Location.Address1, "123 Second St")
	assert.Equal(t, res.Businesses[0].Categories[0].Alias, "indpak")
	assert.Equal(t, res.Businesses[0].Categories[0].Title, "Indian")
	assert.Equal(t, res.Businesses[0].Coordinates.Latitude, float32(37.787789124691))
	assert.Equal(t, res.Businesses[0].Coordinates.Longitude, float32(-122.399305736113))
}

func TestBusinessTransactionSearchError(t *testing.T) {
	client := setup()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
	}))

	defer ts.Close()

	client.BaseURI = ts.URL

	params := &yelp.BusinessTransactionRequest{
		Latitude:  37.787789124691,
		Longitude: -122.399305736113,
	}

	_, err := client.TransactionSearch(params)

	assert.Error(t, err, "500 Internal Server Error")
}

func TestAutocompleteSuccess(t *testing.T) {
	// Arrange
	client := setup()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintln(w, autocompleteResponse)
	}))

	defer ts.Close()

	mockRes := &yelp.BusinessAutocompleteRes{
		Terms:      []yelp.Term{{Text: "Delivery"}},
		Categories: []yelp.Category{{Alias: "delis", Title: "Delis"}},
		Businesses: []yelp.AutocompleteBusiness{{Name: "Delfina", ID: "YqvoyaNvtoC8N5dA8pD2JA"}},
	}

	client.BaseURI = ts.URL

	params := &yelp.BusinessAutoCompleteReq{
		Coordinates: yelp.Coordinates{
			Latitude:  43.64784,
			Longitude: -79.38872,
		},
		Text: "test",
	}

	// Act
	res, err := client.Autocomplete(params)
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	assert.Equal(t, res, mockRes)
}

func TestAutocompleteError(t *testing.T) {
	// Arrange
	client := setup()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
	}))

	defer ts.Close()

	client.BaseURI = ts.URL

	params := &yelp.BusinessAutoCompleteReq{
		Coordinates: yelp.Coordinates{
			Latitude:  43.64784,
			Longitude: -79.38872,
		},
		Text: "test",
	}

	// Act
	_, err := client.Autocomplete(params)

	// Assert
	assert.Error(t, err, "500 Internal Server Error")
}
