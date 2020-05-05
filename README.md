# Yelp Fusion Go Client

Yelp Fusion API client for Go

[![CircleCI](https://circleci.com/gh/naguigui/yelp-fusion/tree/master.svg?style=svg)](https://circleci.com/gh/naguigui/yelp-fusion/tree/master)

Refer to the official Yelp documentation for more information on the API: https://www.yelp.com/developers/documentation/v3 including how to authenticate to use the API.

## Installation

```
go get github.com/naguigui/yelp-fusion
```

## Client Init

```go
import "github.com/naguigui/yelp-fusion/yelp"

func main() {
	client, err := yelp.Init(&yelp.Options{APIKey: os.Getenv("YELP_API_KEY")})
}
```

<br/>

## Table of Contents

Business Endpoints:

- [Business Search](#business-search)
- [Business Details](#business-details)
- [Business Phone Search](#business-phone-search)
- [Business Reviews](#business-reviews)
- [Business Transaction Search](#business-transaction-search)

<br/>

## Business Endpoints

### Business Search

For more details on request/response payloads, refer to https://www.yelp.com/developers/documentation/v3/business_search

```go
// Create client using access token from environment variables
client, err := yelp.Init(&yelp.Options{APIKey: os.Getenv("YELP_API_KEY")})

// Create business search params
params := yelp.BusinessSearch{
	Term:     "restaurants",
	Location: "220 Yonge St, Toronto, ON",
	Limit:    10,
	Radius:   39,
}

// Make the request with created params
res, err := client.BusinessSearch(params)

for _, business := range res.Businesses {
	fmt.Printf("ID: %v\n", business.ID)
	fmt.Printf("Name: %v\n", business.Name)
	fmt.Printf("Location: %v\n", business.Location)
	fmt.Printf("Phone number: %v\n", business.Phone)
}
}
```

## Business Details

For more details on request/response payloads, refer to https://www.yelp.com/developers/documentation/v3/business

```go
// Create client using access token from environment variables
client, err := yelp.Init(&yelp.Options{APIKey: os.Getenv("YELP_API_KEY")})

businessId := "saijdv-vXJrvsCfvr7SZOw"
canadaLocale := "en_CA"

res, err := client.BusinessDetails(businessId, canadaLocale)

fmt.Printf("ID: %v\n", res.ID)
fmt.Printf("Name: %v\n", res.Name)
fmt.Printf("City: %v\n", res.Location.City)
fmt.Printf("Country: %v\n", res.Location.Country)
fmt.Printf("Address: %v\n", res.Location.DisplayAddress)
fmt.Printf("Phone number: %v\n", res.Phone)
fmt.Printf("Photos: %v", res.Photos)
```

## Business Phone Search

For more details on request/response payloads, refer to https://www.yelp.ca/developers/documentation/v3/business_search_phone

```go
// Create client using access token from environment variables
client, _ := yelp.Init(&yelp.Options{APIKey: os.Getenv("YELP_API_KEY")})

phoneNumber := "+1somephonenumber"

res, _ := client.BusinessPhoneSearch(phoneNumber, "")

fmt.Printf("Total businesses: %v\n", res.Total)
fmt.Printf("Businesses: %v\n", res.Businesses)
```

## Business Reviews

For more details on request/response payloads, refer to https://www.yelp.com/developers/documentation/v3/business_reviews

```go
// Create client using access token from environment variables
client, err := yelp.Init(&yelp.Options{APIKey: os.Getenv("YELP_API_KEY")})

businessID := "saijdv-vXJrvsCfvr7SZOw"
canadaLocale := "en_CA"

res, err := client.BusinessReviews(businessID, canadaLocale)

for k, v := range res.Reviews {
	fmt.Println("ID:", v.ID)
	fmt.Println("Rating:", v.Rating)
	fmt.Println("Text:", v.Text)
	fmt.Println("Time Created:", v.TimeCreated)
	fmt.Println("User:", v.User)
	fmt.Println("Url:", v.URL)
}
```

## Business Transaction Search

For more details on request/response payloads, refer to https://www.yelp.ca/developers/documentation/v3/transaction_search

Note: at this time, the API does not return businesses without any reviews and only supports food delivery in the US.
This only supports "delivery" as a transaction_type as Yelp only supports food delivery.

```go
// Create client using access token from environment variables
client, err := yelp.Init(&yelp.Options{APIKey: os.Getenv("YELP_API_KEY")})

params := &yelp.BusinessTransactionRequest{
	Location: "1 Hacker Way East Palo Alto, California",
}

res, err := client.TransactionSearch(params)

fmt.Printf("Total results: %v\n", res.Total)

for _, business := range res.Businesses {
	fmt.Printf("ID: %v\n", business.ID)
	fmt.Printf("Name: %v\n", business.Name)
	fmt.Printf("Location: %v\n", business.Location)
	fmt.Printf("Phone number: %v\n", business.Phone)
}
```

### License

The source code is made available under the [MIT license](LICENSE)

### Contributing

1. Create an issue that describes the bugfix or feature you wish to implement.
2. Fork the repo.
3. Create a feature branch `git checkout -b my-feature`
4. Commit your changes and push up your branch
5. Create a PR (Please leave the issue in the PR body)
