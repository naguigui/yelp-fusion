package main

import (
	"fmt"
	"github.com/naguigui/yelp-fusion/yelp"
	"os"
)

func main() {
	// Create client using access token from environment variables
	client, err := yelp.Init(&yelp.ClientOptions{APIKey: os.Getenv("YELP_API_KEY")})

	if err != nil {
		fmt.Printf("Oh noes, error: %v\n", err)
		return
	}

	params := yelp.BusinessAutocompleteReq{
		Coordinates: yelp.Coordinates{
			Latitude:  43.64784,
			Longitude: -79.38872,
		},
		Text: "thai",
	}

	res, err := client.Autocomplete(params)

	if err != nil {
		fmt.Printf("Oh noes, error: %v", err)
		return
	}

	fmt.Printf("Terms: %v\n", res.Terms)
	fmt.Printf("Businesses: %v\n", res.Businesses)
	fmt.Printf("Categories: %v\n", res.Categories)
}
