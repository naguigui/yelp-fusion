package main

import (
	"errors"
	"github.com/naguigui/yelp-fusion/yelp"
	"log"
	"os"
)

func getReviewsForRestaurant(c *yelp.Client, name string) (*yelp.BusinessReviewsRes, error) {
	params := yelp.BusinessSearchReq{
		Term:     name,
		Location: "220 Yonge St, Toronto, ON",
		Limit:    10,
		Radius:   39,
	}

	// Make the request with created params
	res, err := c.BusinessSearch(params)

	if err != nil {
		return &yelp.BusinessReviewsRes{}, err
	}

	for _, business := range res.Businesses {
		if business.Name == name {
			res, err := c.BusinessReviews(business.ID, "en_CA")
			if err != nil {
				log.Fatal(err)
				break
			}

			return &res, nil
		}
	}
	return &yelp.BusinessReviewsRes{}, errors.New("could not find restaurant")
}

func main() {
	// Create client using access token from environment variables
	client, err := yelp.Init(&yelp.ClientOptions{APIKey: os.Getenv("YELP_API_KEY")})

	if err != nil {
		log.Fatalf("Oh noes, error: %v\n", err)
		return
	}

	reviews, err := getReviewsForRestaurant(client, "Katsuya")
	if err != nil {
		log.Fatalf("Uh oh, big error coming through: %v", err)
	}

	log.Println(reviews)
}
