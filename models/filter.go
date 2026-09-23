package models

// PropertyFilters is a struct that represents the filters for querying properties.
type PropertyFilters struct {
    MinPrice       *float64
    MaxPrice       *float64
	MinReviewScore *float64
    MinStarRating  *int
    MinReviews     *int
	Limit		   *int
	Feed           *int
    MinBedroom     *int
    Published      *bool
    PropertyType   string
    Amenities      []string
}