package validators

import (
	"Rental-Property-REST-API/models"
	"Rental-Property-REST-API/utils"
	"errors"
	"net/url"
)

// Parameter validator function
func ValidatePropertyParameters(query url.Values) error {
	supportedParameters := map[string]bool{
		"min_price":        true,
		"max_price":        true,
		"min_star_rating":  true,
		"min_review_score": true,
		"min_reviews":      true,
		"published":        true,
		"property_type":    true,
		"feed":             true,
		"min_bedroom":      true,
		"amenities":        true,
		"limit":            true,
	}

	for parameter := range query {
		if !supportedParameters[parameter] {
			return errors.New("Unsupported query parameter: " + parameter)
		}
	}

	return nil
}

/* For better understanding
	query := url.Values{
    	"min_price": {"50"},
    	"limit":     {"10"},
    	"amenities":      {"wifi", "pool"},
	}

	so if,
	values, ok := query[amenities]
	then,
	values[0] = {"wifi", "pool"} for amenities
*/

// Validator for filter parameter's values
func ParsePropertyFilters(query url.Values) (models.PropertyFilters, error) {
	var filters models.PropertyFilters

	for key, values := range query {
		if len(values) == 0 {
			return filters, errors.New(key + " cannot be empty")
		}

		switch key {
		case "min_price":
			minPrice, err := utils.ParseFloat(values[0], "min_price")
			if err != nil {
				return filters, err
			}
			filters.MinPrice = &minPrice

		case "max_price":
			maxPrice, err := utils.ParseFloat(values[0], "max_price")
			if err != nil {
				return filters, err
			}
			filters.MaxPrice = &maxPrice

		case "min_star_rating":
			minStarRating, err := utils.ParseInt(values[0], "min_star_rating")
			if err != nil {
				return filters, err
			}
			filters.MinStarRating = &minStarRating

		case "min_review_score":
			minReviewScore, err := utils.ParseFloat(values[0], "min_review_score")
			if err != nil {
				return filters, err
			}
			filters.MinReviewScore = &minReviewScore

		case "min_reviews":
			minReviews, err := utils.ParseInt(values[0], "min_reviews")
			if err != nil {
				return filters, err
			}
			filters.MinReviews = &minReviews

		case "published":
			published, err := utils.ParseBool(values[0], "published")
			if err != nil {
				return filters, err
			}
			filters.Published = &published

		case "property_type":
			switch values[0] {
			case "Hotel", "House", "Apartment", "Villa", "Resort", "Hostel":
				filters.PropertyType = values[0]
			default:
				return filters, errors.New("property_type must be one of: Hotel, House, Apartment, Villa, Resort, Hostel")
			}

		case "feed":
			feed, err := utils.ParseInt(values[0], "feed")
			if err != nil {
				return filters, err
			}

			switch feed {
			case 11, 12, 22, 24:
				filters.Feed = &feed
			default:
				return filters, errors.New("feed must be one of: 11, 12, 22, 24")
			}

		case "min_bedroom":
			minBedroom, err := utils.ParseInt(values[0], "min_bedroom")
			if err != nil {
				return filters, err
			}
			filters.MinBedroom = &minBedroom

		case "amenities":
			amenities, err := utils.ParseStringList(values, "amenities")
			if err != nil {
				return filters, err
			}
			filters.Amenities = amenities

		case "limit":
			limit, err := utils.ParseInt(values[0], "limit")
			if err != nil {
				return filters, err
			}
			if limit == 0 {
				return filters, errors.New("limit must be greater than zero")
			}
			filters.Limit = &limit
		}
	}

	return filters, nil
}