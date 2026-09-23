package services

import (
	"encoding/json"
	"errors"
	"os"

	"Rental-Property-REST-API/models"
	"github.com/beego/beego/v2/server/web"
)

// PropertiesStore akta PropertySlice er wrapper struct. jeta JSON file ke struct e
// convert korar por propertySlice er vitor store kore rakhe.
type PropertyData struct {
	Properties []models.SourceProperty
}

func LoadPropertiesFromFile() (*PropertyData, error) {
	dataFilePath, err := web.AppConfig.String("property_data_path")
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(dataFilePath)
	if err != nil {
		return nil, err
	}

	var properties []models.SourceProperty
	err = json.Unmarshal(data, &properties)
	if err != nil {
		return nil, err
	}

	return &PropertyData{
		Properties: properties,
	}, nil
}

type PropertyService struct {
	PropertyData *PropertyData
}

func NewPropertyService(propertyData *PropertyData) *PropertyService {
	return &PropertyService{
		PropertyData: propertyData,
	}
}

// Filter all properties
func FilterProperties(properties []models.SourceProperty, filters models.PropertyFilters) []models.SourceProperty {
	var filtered []models.SourceProperty
	for _, property := range properties {
		if filters.MinPrice != nil && property.USDPrice < *filters.MinPrice {
			continue
		}
		if filters.MaxPrice != nil && property.USDPrice > *filters.MaxPrice {
			continue
		}
		if filters.MinStarRating != nil && property.StarRating < *filters.MinStarRating {
			continue
		}
		if filters.MinReviewScore != nil && property.ReviewScoreGeneral < *filters.MinReviewScore {
			continue
		}
		if filters.MinReviews != nil && property.NumberOfReview < *filters.MinReviews {
			continue
		}
		if filters.Published != nil && property.Published != *filters.Published {
			continue
		}
		if filters.PropertyType != "" && property.PropertyTypeCategory != filters.PropertyType {
			continue
		}
		if filters.Feed != nil && property.Feed != *filters.Feed {
			continue
		}
		if filters.MinBedroom != nil && property.BedroomCount < *filters.MinBedroom {
			continue
		}
		if len(filters.Amenities) > 0 {
			matched := false
			for _, requestedAmenity := range filters.Amenities {
				for _, propertyAmenity := range property.AmenityCategories {
					if requestedAmenity == propertyAmenity {
						matched = true
						break
					}
				}
				if matched {
					break
				}
			}
			if !matched {
				continue
			}
		}
		filtered = append(filtered, property)
	}
	return filtered
}

// TransformProperty transforms a SourceProperty to a ResponseProperty.
func TransformProperty(property models.SourceProperty) (models.ResponseProperty, error) {
	var breadcrumbs []models.Breadcrumb

	err := json.Unmarshal([]byte(property.Categories), &breadcrumbs)
	if err != nil {
		return models.ResponseProperty{}, err
	}

	responseProperty := models.ResponseProperty{
		ID:        property.ID,
		Feed:      property.Feed,
		Published: property.Published,

		GeoInfo: models.GeoInfo{
			Breadcrumbs: breadcrumbs,
			City:        property.City,
			Country:     property.Country,
			CountryCode: property.CountryCode,
			Name:        property.PropertyName,
			LocationID:  property.LocationID,
			Lon:         property.LonLat.Coordinates[0],
			Lat:         property.LonLat.Coordinates[1],
			State:       property.State,
			StateAbbr:   property.StateAbbr,
		},

		Property: models.Property{
			Amenities:    property.AmenityCategories,
			Name:         property.PropertyName,
			Slug:         property.PropertySlug,
			PropertyType: property.PropertyTypeCategory,
			Price:        property.USDPrice,
			ReviewScore:  property.ReviewScoreGeneral,
			StarRating:   property.StarRating,

			Counts: models.Counts{
				Bathroom:  property.BathroomCount,
				Bedroom:   property.BedroomCount,
				Reviews:   property.NumberOfReview,
				Occupancy: property.Occupancy,
			},

			Image: models.Image{
				Count:  len(property.Images),
				Images: property.Images,
			},
		},
	}

	return responseProperty, nil
}

// GetAllProperties returns all properties from the PropertyData.
func (service *PropertyService) GetAllProperties(filters models.PropertyFilters) (models.PropertyListResponse, error) {
	filteredProperties := FilterProperties(service.PropertyData.Properties, filters)

	responseProperties := make([]models.ResponseProperty, 0)
	for _, property := range filteredProperties {
		responseProperty, err := TransformProperty(property)
		if err != nil {
			return models.PropertyListResponse{}, err
		}
		responseProperties = append(responseProperties, responseProperty)
	}

	if filters.Limit != nil && *filters.Limit < len(responseProperties) {
		responseProperties = responseProperties[:*filters.Limit]
	}

	return models.PropertyListResponse{
		Result: models.PropertyListResult{
			Count: len(responseProperties),
			Items: responseProperties,
		},
	}, nil
}

// GetPropertyByID returns a property by its ID.
func (service *PropertyService) GetPropertyByID(id string) (models.ResponseProperty, error) {
	for _, property := range service.PropertyData.Properties {
		if property.ID == id {
			return TransformProperty(property)
		}
	}
	return models.ResponseProperty{}, errors.New("property not found")
}
