package validators

import (
	"errors"
	"strconv"
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