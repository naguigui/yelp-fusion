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

	params := &yelp.BusinessTransactionRequest{
		Location: "	1 Hacker Way East Palo Alto, California",
	}

	res, err := client.TransactionSearch(params)

	if err != nil {
		fmt.Printf("Oh noes, error: %v", err)
		return
	}

	fmt.Printf("Total results: %v\n", res.Total)

	for _, business := range res.Businesses {
		fmt.Printf("ID: %v\n", business.ID)
		fmt.Printf("Name: %v\n", business.Name)
		fmt.Printf("Location: %v\n", business.Location)
		fmt.Printf("Phone number: %v\n", business.Phone)
	}
}
