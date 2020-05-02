package main

import (
	"fmt"
	"github.com/naguigui/yelp-fusion/yelp"
)

func main() {
	// Create client using access token from environment variables
	client, err := yelp.Init("0E65hDwshpbMVa4KYP5a7GRBOldYaJk7fABEgG8DBt5fRrp4wZlIx7iT4F59liIGXj-PoBTqZ4fF7s6qbNfyFtP5jQvGiGaIMdGKs3eAdX3kTc-ZP7S2cFYRm2CnXnYx")

	if err != nil {
		fmt.Printf("Oh noes, error: %v\n", err)
		return
	}

	// Create business search params
	params := yelp.BusinessSearch{
		Term: "restaurants",
		Location: "220 Yonge St, Toronto, ON",
		Limit: 10,
		Radius: 39,
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