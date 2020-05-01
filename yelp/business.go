package yelp

type BusinessSearch struct {
	Term       string  `json:"term"`
	Location   string  `json:"location"`
	Latitude   float32 `json:"latitude"`
	Longitude  float32 `json:"longitude"`
	Radius     int     `json:"radius"`
	Categories string  `json:"categories"`
	Locale     string  `json:"locale"`
	Limit      int     `json:"limit"`
	Offset     int     `json:"offset"`
	SortBy     string  `json:"sort_by"`
	Price      string  `json:"price"`
	OpenNow    bool    `json:"open_now"`
	Attributes string  `json:"attributes"`
}

type BusinessSearchResponse struct {
	Region     Region     `json:"region"`   // Suggested area in a map to display results in.
	Total      int        `json:"total"`    // Total number of business results
	Businesses []Business `json:"businesses"` // The list of business entries (see Business)
}

type Region struct {
	Center Center `json:"center"` // Center position of map area
}

type Center struct {
	Latitude  float32 `json:"latitude"` // Latitude position of map bounds center
	Longitude float32 `json:"longitude"` // Longitude position of map bounds center
}

type Coordinates struct {
	Latitude  float32 `json:"latitude"` // Latitude of this business
	Longitude float32 `json:"longitude"` // Longitude of this business
}

type Location struct {
	City     string `json:"city"` // City of this business
	Country  string `json:"country"` // ISO 3166-1 alpha-2 country code of this business
	Address1 string `json:"address1"` // Street address of this business
	Address2 string `json:"address2"` // Street address of this business, continued
	Address3 string `json:"address3"` // Street address of this business, continued
	State    string `json:"state"` // ISO 3166-2 state code of this business
	ZipCode  string `json:"zip_code"` // Zip code of this business
}

type Category struct {
	Alias string `json:"alias"` // When searching for business in certain categories, use alias rather than the title
	Title string `json:"title"` // For display purposes
}

// Business is a list of businesses Yelp finds based on the search criteria.
type Business struct {
	ID           string      `json:"id"` // Unique Yelp ID of this business
	Rating       float32      `json:"rating"`// Rating for this business (value ranges from 1, 1.5, ... 4.5, 5)
	Price        string      `json:"price"`// Price level of this business. Value is one of $, $$, $$$, and $$$$
	Phone        string      `json:"phone"`// Phone number of this business
	Alias        string      `json:"alias"`// Unique Yelp alias of this business. Can contain unicode characters
	IsClosed     bool        `json:"is_closed"`// Whether business has been (permanently) closed
	Categories   []Category  `json:"categories"`// List of category title and alias pairs associated with this business
	ReviewCount  int         `json:"review_count"`// Number of reviews for this business
	Name         string      `json:"name"`// Name of this business
	URL          string      `json:"url"`// URL for business page on Yelp
	Coordinates  Coordinates `json:"coordinates"`// Coordinates of this business
	ImageURL     string      `json:"image_url"`// URL of photo for this business
	Location     Location    `json:"location"`// Location of this business, including address, city, state, zip code, and country
	Distance     float32     `json:"distance"`// Distance in meters from the search location
	Transactions []string    `json:"transactions"`// List of Yelp transactions that the business is registered for. Current supported values are pickup, delivery, and restaurant_reservation
}
