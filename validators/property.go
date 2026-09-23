package validators

import (
	"errors"
	"strconv"
	"net/url"

	"Rental-Property-REST-API/models"
)

func ValidateLimit(value string) (int, error) {
	if value == "" {
		return 0, nil
	}
	limit, err := strconv.Atoi(value)
	if err != nil || limit <= 0 {
		return 0, errors.New("limit must be a positive integer")
	}
	return limit, nil
}

func ParsePropertyFilters(query url.Values) (models.PropertyFilters, error) {
    filters := models.PropertyFilters{}
	if values, exists := query["min_price"]; exists {
		if values[0] == "" {
			return models.PropertyFilters{}, errors.New("min_price cannot be empty")
		}

		value, err := strconv.ParseFloat(values[0], 64)
		if err != nil {
			return models.PropertyFilters{}, errors.New("min_price must be a valid number")
		}

		filters.MinPrice = &value
	}

	return filters, nil
}