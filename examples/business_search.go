package main

import (
	"github.com/naguigui/yelp-fusion/yelp"
	"fmt"
	"os"
)

func main() {
	// Create client using access token from environment variables
	client, err := yelp.Init(os.Getenv("YELP_ACCESS_TOKEN"))

	if err != nil {
		fmt.Println("Oh noes, error: %v", err)
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
		fmt.Println("Oh noes, error: %v", err)
	}

	for _, business := range res.Businesses {
		fmt.Printf("ID: %v\n", business.ID)
		fmt.Printf("Name: %v\n", business.Name)
		fmt.Printf("Location: %v\n", business.Location)
		fmt.Printf("Phone number: %v\n", business.Phone)
	}
}