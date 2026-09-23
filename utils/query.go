package utils

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

import (
	"errors"
	"strconv"
	"strings"
)

func ParseFloat(value string, key string) (float64, error) {
    if value == "" {
        return 0, errors.New(key + " cannot be empty")
    }
    number, err := strconv.ParseFloat(value, 64)
    if err != nil {
        return 0, errors.New(key + " must be a valid number")
    }
    if number < 0 {
        return 0, errors.New(key + " cannot be negative")
    }
    return number, nil
}

func ParseInt(value string, key string) (int, error) {
	if value == "" {
		return 0, errors.New(key + " cannot be empty")
	}
	number, err := strconv.Atoi(value)
	if err != nil {
		return 0, errors.New(key + " must be a valid integer")
	}
	if number < 0 {
		return 0, errors.New(key + " cannot be negative")
	}
	return number, nil
}

func ParseBool(value string, key string) (bool, error) {
	if value == "" {
		return false, errors.New(key + " cannot be empty")
	}
	boolean, err := strconv.ParseBool(value)
	if err != nil {
		return false, errors.New(key + " must be a valid boolean")
	}
	return boolean, nil
}

func ParseStringList(values []string, key string) ([]string, error) {
	if len(values) == 0 || values[0] == "" {
		return nil, errors.New(key + " cannot be empty")
	}
	list := strings.Split(values[0], ",")
	for i, item := range list {
		list[i] = strings.TrimSpace(item)
	}
	return list, nil
}