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

	katsuyaID := "saijdv-vXJrvsCfvr7SZOw"
	canadaLocale := "en_CA"

	res, err := client.BusinessDetails(katsuyaID, canadaLocale)

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
