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

	phoneNumber := "+14169014724"

	res, err := client.BusinessPhoneSearch(phoneNumber, "")
	if err != nil {
		fmt.Printf("Oh noes, error: %v", err)
		return
	}
	fmt.Printf("Total businesses: %v\n", res.Total)
	fmt.Printf("Businesses: %v\n", res.Businesses)
}
