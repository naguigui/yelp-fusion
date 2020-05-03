# Yelp Fusion Go Client

Yelp Fusion API client for Go

[![CircleCI](https://circleci.com/gh/naguigui/yelp-fusion/tree/master.svg?style=svg)](https://circleci.com/gh/naguigui/yelp-fusion/tree/master)

Refer to the official Yelp documentation for more information on the API: https://www.yelp.com/developers/documentation/v3 including how to authenticate to use the API.

## Installation

```
go get github.com/naguigui/yelp-fusion
```

<br/>

## Table of Contents

Business Endpoints:

- [Business Search](#business-search)
- [Business Details](#business-details)
- [Business Phone Search](#business-phone-search)

<br/>

## Business Endpoints

### Business Search

For more details on request/response payloads, refer to https://www.yelp.com/developers/documentation/v3/business_search

```go
  import (
      "github.com/naguigui/yelp-fusion"
  )

  func main() {
      // Create client using access token from environment variables
	client, err := yelp.Init(os.Getenv("YELP_API_KEY"), nil)

	if err != nil {
		fmt.Printf("Oh noes, error: %v\n", err)
		return
	}

	// Create business search params
	params := yelp.BusinessSearch{
		Term:     "restaurants",
		Location: "220 Yonge St, Toronto, ON",
		Limit:    10,
		Radius:   39,
	}

	// Make the request with created params
	res, err := client.BusinessSearch(params)

	if err != nil {
		fmt.Printf("Oh noes, error: %v\n", err)
		return
	}

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
import (
	"fmt"
	"github.com/naguigui/yelp-fusion/yelp"
	"os"
)

func main() {
	// Create client using access token from environment variables
	client, err := yelp.Init(os.Getenv("YELP_API_KEY"), nil)

	if err != nil {
		fmt.Printf("Oh noes, error: %v\n", err)
		return
	}

	businessId := "saijdv-vXJrvsCfvr7SZOw"
	canadaLocale := "en_CA"

	res, err := client.BusinessDetails(businessId, canadaLocale)

	if err != nil {
		fmt.Printf("Oh noes, error: %v", err)
		return
	}

	fmt.Printf("ID: %v\n", res.ID)
	fmt.Printf("Name: %v\n", res.Name)
	fmt.Printf("City: %v\n", res.Location.City)
	fmt.Printf("Country: %v\n", res.Location.Country)
	fmt.Printf("Address: %v\n", res.Location.DisplayAddress)
	fmt.Printf("Phone number: %v\n", res.Phone)
	fmt.Printf("Photos: %v", res.Photos)
}
```

## Business Phone Search

For more details on request/response payloads, refer to https://www.yelp.ca/developers/documentation/v3/business_search_phone

```go
import (
	"fmt"
	"github.com/naguigui/yelp-fusion/yelp"
	"os"
)

func main() {
	// Create client using access token from environment variables
	client, err := yelp.Init(os.Getenv("YELP_API_KEY"), nil)

	if err != nil {
		fmt.Printf("Oh noes, error: %v\n", err)
		return
	}

	phoneNumber := "+1somephonenumber"

	res, err := client.BusinessPhoneSearch(phoneNumber, "")
	if err != nil {
		fmt.Printf("Oh noes, error: %v", err)
		return
	}
	fmt.Printf("res: %v\n", res)
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
