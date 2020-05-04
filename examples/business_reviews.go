package main

import (
	"fmt"
	"github.com/naguigui/yelp-fusion/yelp"
	"os"
)

func main() {
	// Create client using access token from environment variables
	client, err := yelp.Init(&yelp.Options{APIKey: os.Getenv("YELP_API_KEY")})

	if err != nil {
		fmt.Printf("Oh noes, error: %v\n", err)
		return
	}

	businessID := "saijdv-vXJrvsCfvr7SZOw"
	canadaLocale := "en_CA"

	res, err := client.BusinessReviews(businessID, canadaLocale)

	if err != nil {
		fmt.Printf("Oh noes, error: %v", err)
		return
	}

	for _, v := range res.Reviews {
		fmt.Println("ID:", v.ID)
		fmt.Println("Rating:", v.Rating)
		fmt.Println("Text:", v.Text)
		fmt.Println("Time Created:", v.TimeCreated)
		fmt.Println("User:", v.User)
		fmt.Println("Url:", v.URL)
	}
}
