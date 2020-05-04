package yelp

// BusinessSearch is the request payload for Business search API
type BusinessSearch struct {
	Term       string  `json:"term"`       // Optional. Search term, for example "food" or "restaurants". The term may also be business names, such as "Starbucks". If term is not included the endpoint will default to searching across businesses from a small number of popular categories
	Location   string  `json:"location"`   // Required if either latitude or longitude is not provided. This string indicates the geographic area to be used when searching for businesses
	Latitude   float32 `json:"latitude"`   // Required if location is not provided. Latitude of the location you want to search nearby
	Longitude  float32 `json:"longitude"`  // Required if location is not provided. Longitude of the location you want to search nearby
	Radius     int     `json:"radius"`     // Optional. A suggested search radius in meters. This field is used as a suggestion to the search
	Categories string  `json:"categories"` // Optional. Categories to filter the search results with
	Locale     string  `json:"locale"`     // Optional. Specify the locale into which to localize the business information. See the list of supported locales. https://www.yelp.ca/developers/documentation/v3/supported_locales. Defaults to en_US
	Limit      int     `json:"limit"`      // Optional. Number of business results to return. By default, it will return 20. Maximum is 50
	Offset     int     `json:"offset"`     // Optional. Offset the list of returned business results by this amount
	SortBy     string  `json:"sort_by"`    // Optional. Suggestion to the search algorithm that the results be sorted by one of the these modes: best_match, rating, review_count or distance. The default is best_match
	Price      string  `json:"price"`      // Optional. Pricing levels to filter the search result with: 1 = $, 2 = $$, 3 = $$$, 4 = $$$$. The price filter can be a list of comma delimited pricing levels. For example, "1, 2, 3" will filter the results to show the ones that are $, $$, or $$$
	OpenNow    bool    `json:"open_now"`   // Optional. Default to false. When set to true, only return the businesses open now
	OpenAt     int     `json:"open_at"`    // Optional. An integer represending the Unix time in the same timezone of the search location
	Attributes string  `json:"attributes"` // Optional. See list of attributes to try out here. https://www.yelp.ca/developers/documentation/v3/business_search
}

// BusinessSearchResponse is the response payload for Business Search API
type BusinessSearchResponse struct {
	Region     Region     `json:"region"`     // Suggested area in a map to display results in.
	Total      int        `json:"total"`      // Total number of business results
	Businesses []Business `json:"businesses"` // The list of business entries (see Business)
}

// BusinessDetailsResponse is the response payload for Business Details API
type BusinessDetailsResponse struct {
	ID           string                  `json:"id"`            // Unique Yelp ID of this business
	Alias        string                  `json:"alias"`         // Alias of a category
	Name         string                  `json:"name"`          // Name of this business
	ImageURL     string                  `json:"image_url"`     // URL of photo for this business
	IsClaimed    bool                    `json:"is_claimed"`    // Whether business has been claimed by a business owner
	IsClosed     bool                    `json:"is_closed"`     // Whether business has been (permanently) closed
	URL          string                  `json:"url"`           // URL for business page on Yelp
	Phone        string                  `json:"phone"`         // Phone number of the business
	DisplayPhone string                  `json:"display_phone"` // Phone number of the business formatted nicely to be displayed to users
	ReviewCount  int                     `json:"review_count"`  // Number of reviews for this business
	Categories   []Category              `json:"categories"`    // A list of category title and alias pairs associated with this business
	Rating       float32                 `json:"rating"`        // Rating for this business (value ranges from 1, 1.5, ... 4.5, 5)
	Location     LocationBusinessDetails `json:"location"`      // The location of this business, including address, city, state, zip code and country
	Coordinates  Coordinates             `json:"coordinates"`   // The coordinates of this business
	Photos       []string                `json:"photos"`        // URLs of up to three photos of the business
	Price        string                  `json:"price"`         // Price level of the business. Value is one of $, $$, $$$ and $$$$
	Hours        []Hours                 `json:"hours"`         // Opening hours of the business
	Transactions []string                `json:"transactions"`  // A list of Yelp transactions that the business is registered for. Current supported values are "pickup", "delivery", and "restaurant_reservation"
	SpecialHours []SpecialHours          `json:"special_hours"` // Out of the ordinary hours for the business that apply on certain dates. Whenever these are set, they will override the regular business hours found in the 'hours' field
	Messaging    Messaging               `json:"messaging"`     // Contains Business Messaging / Request a Quote information for this business. This field only appears in the response for businesses that have messaging enabled
}

// BusinessPhoneSearchResponse is the response payload for Business Phone Search API
type BusinessPhoneSearchResponse struct {
	Total      int        `json:"total"`
	Businesses []Business `json:"businesses"`
}

// BusinessReviewsResponse is the response payload for Business Reviews API
type BusinessReviewsResponse struct {
	Reviews           []Review `json:"reviews"`            // A list of up to three reviews of this business
	Total             int      `json:"total"`              // The total number of reviews that the business has
	PossibleLanguages []string `json:"possible_languages"` // A list of languages for which the business has at least one review.
}

// Business is the full data of a specific business from the Yelp Fusion Business API consisting of its ID, Rating, Price, Phone Number, Opening Hours, and etc.
type Business struct {
	ID           string      `json:"id"`                 // Unique Yelp ID of this business
	Rating       float32     `json:"rating"`             // Rating for this business (value ranges from 1, 1.5, ... 4.5, 5)
	Price        string      `json:"price"`              // Price level of this business. Value is one of $, $$, $$$, and $$$$
	Phone        string      `json:"phone"`              // Phone number of this business
	Alias        string      `json:"alias"`              // Alias of a category
	IsClosed     bool        `json:"is_closed"`          // Whether business has been (permanently) closed
	Categories   []Category  `json:"categories"`         // List of category title and alias pairs associated with this business
	ReviewCount  int         `json:"review_count"`       // Number of reviews for this business
	Name         string      `json:"name"`               // Name of this business
	URL          string      `json:"url"`                // URL for business page on Yelp
	Coordinates  Coordinates `json:"coordinates"`        // Coordinates of this business
	ImageURL     string      `json:"image_url"`          // URL of photo for this business
	Location     Location    `json:"location"`           // Location of this business, including address, city, state, zip code, and country
	Distance     float32     `json:"distance,omitempty"` // Distance in meters from the search location
	Transactions []string    `json:"transactions"`       // List of Yelp transactions that the business is registered for. Current supported values are pickup, delivery, and restaurant_reservation
}

// Region is the suggested area in a map to display results in.
type Region struct {
	Center Center `json:"center"` // Center position of map area
}

// Center is the position of map area
type Center struct {
	Latitude  float32 `json:"latitude"`  // Latitude position of map bounds center
	Longitude float32 `json:"longitude"` // Longitude position of map bounds center
}

// Coordinates is the coordinates of the business, consisting of latitude/longitude
type Coordinates struct {
	Latitude  float32 `json:"latitude"`  // Latitude of this business
	Longitude float32 `json:"longitude"` // Longitude of this business
}

// Location indicates the geographic area to be used when searching for businesses. Examples: "New York City", "NYC", "350 5th Ave, New York, NY 10118"
type Location struct {
	City     string `json:"city"`     // City of this business
	Country  string `json:"country"`  // ISO 3166-1 alpha-2 country code of this business
	Address1 string `json:"address1"` // Street address of this business
	Address2 string `json:"address2"` // Street address of this business, continued
	Address3 string `json:"address3"` // Street address of this business, continued
	State    string `json:"state"`    // ISO 3166-2 state code of this business
	ZipCode  string `json:"zip_code"` // Zip code of this business
}

// LocationBusinessDetails is the location of this business, including address, city, state, zip code and country
type LocationBusinessDetails struct {
	Location
	DisplayAddress []string `json:"display_address"` // Array of strings that if organized vertically give an address that is in the standard address format for the business's country
	CrossStreets   string   `json:"cross_streets"`   // Cross streets for this business
}

// Category defines the alias and title pair associated with the business
type Category struct {
	Alias string `json:"alias"` // When searching for business in certain categories, use alias rather than the title
	Title string `json:"title"` // For display purposes
}

// Open is the detailed opening hours of each day in a week
type Open struct {
	IsOvernight bool   `json:"is_overnight"` // Whether the business opens overnight or not. When this is true, the end time will be lower than the start time
	Start       string `json:"start"`        // Start of the opening hours in a day, in 24-hour clock notation, like 1000 means 10 AM
	End         string `json:"end"`          // End of the opening hours in a day, in 24-hour clock notation, like 2130 means 9:30 PM
	Day         int    `json:"day"`          // From 0 to 6, representing day of the week from Monday to Sunday. Notice that you may get the same day of the week more than once if the business has more than one opening time slots
}

// Hours is the opening hours of the business
type Hours struct {
	Open      []Open `json:"open"`        // The detailed opening hours of each day in a week
	HoursType string `json:"hours_type"`  // The type of the opening hours information. Right now, this is always REGULAR
	IsOpenNow bool   `json:"is_open_now"` // Whether the business is currently open or not
}

// SpecialHours is out of the ordinary hours for the business that apply on certain dates. Whenever these are set, they will override the regular business hours found in the 'hours' field
type SpecialHours struct {
	Date        string `json:"date"`         // An ISO8601 date string representing the date for which these special hours apply
	IsClosed    bool   `json:"is_closed"`    // Whether this particular special hour represents a date where the business is closed
	Start       string `json:"start"`        // Start of the opening hours in a day, in 24-hour clock notation, like 1000 means 10 AM
	End         string `json:"end"`          // End of the opening hours in a day, in 24-hour clock notation, like 2130 means 9:30 PM
	IsOvernight bool   `json:"is_overnight"` // Whether the special hours time range spans across midnight or not. When this is true, the end time will be lower than the start time
}

// Messaging data contains business quote information for this business. This field only appears in the response for businesses that have messaging enabled.
type Messaging struct {
	URL         string `json:"url"`           // Action Link URL that drops user directly in to the messaging flow for this business.
	UseCaseText string `json:"use_case_text"` // Indicates what kind of messaging can be done with the business, for example "Request a Quote" for a home services business. This text will be localized based on the "locale" input parameter
}

// Review data consists of a list of user reviews for a business
type Review struct {
	ID          string `json:"id"`           // A unique identifier for this review
	Rating      int    `json:"rating"`       // Rating of this review
	User        User   `json:"user"`         // The user who wrote the review
	Text        string `json:"text"`         // Text excerpt of this review
	TimeCreated string `json:"time_created"` // The time that the review was created in PST
	URL         string `json:"url"`          // URL of this review
}

// User data from business reviews
type User struct {
	ID         string `json:"id"`          // ID of the user
	ProfileURL string `json:"profile_url"` // URL of the user's profile
	ImageURL   string `json:"image_url"`   // URL of the user's profile photo
	Name       string `json:"name"`        // User screen name (first name and first initial of last name)
}
