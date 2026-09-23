package validators

import (
	"errors"
	"net/url"

)

func ValidatePropertyParameters(query url.Values) error {
	supportedParameters := map[string]bool{
		"min_price":       true,
		"max_price":       true,
		"min_star_rating": true,
		"min_review_score": true,
		"min_reviews":     true,
		"published":       true,
		"property_type":   true,
		"feed":            true,
		"min_bedroom":     true,
		"amenities":       true,
		"limit":           true,
	}

	for parameter := range query {
		if !supportedParameters[parameter] {
			return errors.New("unsupported query parameter: " + parameter)
		}
	}

	return nil
}