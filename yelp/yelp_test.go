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

const BUSINESS_SEARCH_RESPONSE = `
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

const BUSINESS_REVIEWS_RESPONSE = `
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

const BUSINESS_TRANSACTION_RESPONSE = `
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

const AUTOCOMPLETE_RESPONSE = `
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
	client, _ := yelp.Init(&yelp.ClientOptions{APIKey: "yelp-key"})

	return client

}

func TestBusinessSearch(t *testing.T) {
	client := setup()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprint(w, BUSINESS_SEARCH_RESPONSE)
	}))

	defer ts.Close()

	expected := yelp.BusinessSearchRes{
		Total: 1,
		Businesses: []yelp.Business{{
			Rating: 4.5,
			Price:  "$$$$",
			Phone:  "+14157492060",
			ID:     "WavvLdfdP6g8aZTtbBQHTw",
			Alias:  "gary-danko-san-francisco",
			Categories: []yelp.Category{{
				Alias: "newamerican",
				Title: "American (New)",
			}},
			ReviewCount: 4525,
			Name:        "Gary Danko",
			URL:         "https://www.yelp.com/biz/gary-danko-san-francisco",
			Coordinates: yelp.Coordinates{
				Latitude:  37.80587,
				Longitude: -122.42058,
			},
			ImageURL: "https://s3-media3.fl.yelpcdn.com/bphoto/--8oiPVp0AsjoWHqaY1rDQ/o.jpg",
			IsClosed: false,
			Location: yelp.Location{
				City:     "San Francisco",
				Country:  "US",
				State:    "CA",
				Address1: "800 N Point St",
				ZipCode:  "94109",
				Address2: "",
				Address3: "",
			},
			Transactions: []string{"restaurant_reservation"},
		}},
	}

	client.BaseURI = ts.URL

	params := yelp.BusinessSearchReq{
		Term:     "restaurant",
		Location: "222 Yonge St. Toronto, ON",
	}

	res, err := client.BusinessSearch(params)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res, expected)
}

func TestBusinessSearchError(t *testing.T) {
	// Arrange
	client := setup()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
	}))

	defer ts.Close()

	client.BaseURI = ts.URL

	params := yelp.BusinessSearchReq{
		Term:     "restaurant",
		Location: "222 Yonge St. Toronto, ON",
	}

	// Act
	_, err := client.BusinessSearch(params)

	// Assert
	assert.EqualError(t, err, "400 Bad Request")
}

func TestBusinessReviewsSuccess(t *testing.T) {
	// Arrange
	client := setup()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprint(w, BUSINESS_REVIEWS_RESPONSE)
	}))

	defer ts.Close()

	expected := yelp.BusinessReviewsRes{
		Total:             1,
		PossibleLanguages: []string{"en"},
		Reviews: []yelp.Review{
			{
				ID:     "review12345",
				Rating: 5,
				User: yelp.User{
					ID:         "user12345",
					ProfileURL: "https://profileurl",
					ImageURL:   "https://myjpg.jpg",
					Name:       "Andrew Nguyen",
				},
				Text:        "Omg I love Katsuya, I can eat it everyday",
				TimeCreated: "2020-05-04 00:41:13",
				URL:         "https://mockurl.com",
			},
		},
	}

	client.BaseURI = ts.URL

	// Act
	res, err := client.BusinessReviews("review12345", "")
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	assert.Equal(t, res, expected)
}

func TestBusinessReviewsError(t *testing.T) {
	// Arrange
	client := setup()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
	}))

	defer ts.Close()

	client.BaseURI = ts.URL

	// Act
	_, err := client.BusinessReviews("review12345", "")

	// Assert
	assert.EqualError(t, err, "500 Internal Server Error")
}

func TestBusinessTransactionSearchSuccess(t *testing.T) {
	// Arrange
	client := setup()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprint(w, BUSINESS_TRANSACTION_RESPONSE)
	}))

	defer ts.Close()

	expected := yelp.BusinessTransactionSearchRes{
		Total: 144,
		Businesses: []yelp.Business{
			{
				ID:     "gR9DTbKCvezQlqvD7_FzPw",
				Name:   "North India Restaurant",
				Phone:  "+14153481234",
				Alias:  "north-india-restaurant-san-francisco",
				Price:  "$$",
				Rating: 4,
				Location: yelp.Location{
					City:     "San Francisco",
					Country:  "US",
					State:    "CA",
					Address1: "123 Second St",
					Address2: "",
					Address3: "",
					ZipCode:  "94105",
				},
				Categories: []yelp.Category{{
					Alias: "indpak",
					Title: "Indian",
				}},
				Coordinates: yelp.Coordinates{
					Latitude:  37.787789124691,
					Longitude: -122.399305736113,
				},
				ImageURL:     "http://s3-media4.fl.yelpcdn.com/bphoto/howYvOKNPXU9A5KUahEXLA/o.jpg",
				URL:          "https://www.yelp.com/biz/north-india-restaurant-san-francisco",
				IsClosed:     false,
				ReviewCount:  615,
				Transactions: []string{"pickup", "restaurant_reservation"},
			},
		},
	}

	client.BaseURI = ts.URL

	params := yelp.BusinessTransactionReq{
		Latitude:  37.787789124691,
		Longitude: -122.399305736113,
	}

	// Act
	res, err := client.TransactionSearch(params)
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	assert.Equal(t, res, expected)
}

func TestBusinessTransactionSearchError(t *testing.T) {
	// Arrange
	client := setup()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
	}))

	defer ts.Close()

	client.BaseURI = ts.URL

	params := yelp.BusinessTransactionReq{
		Latitude:  37.787789124691,
		Longitude: -122.399305736113,
	}

	// Act
	_, err := client.TransactionSearch(params)

	// Assert
	assert.Error(t, err, "500 Internal Server Error")
}

func TestAutocompleteSuccess(t *testing.T) {
	// Arrange
	client := setup()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprint(w, AUTOCOMPLETE_RESPONSE)
	}))

	defer ts.Close()

	expected := yelp.BusinessAutocompleteRes{
		Terms:      []yelp.Term{{Text: "Delivery"}},
		Categories: []yelp.Category{{Alias: "delis", Title: "Delis"}},
		Businesses: []yelp.AutocompleteBusiness{{Name: "Delfina", ID: "YqvoyaNvtoC8N5dA8pD2JA"}},
	}

	client.BaseURI = ts.URL

	params := yelp.BusinessAutocompleteReq{
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
	assert.Equal(t, res, expected)
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

	params := yelp.BusinessAutocompleteReq{
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
