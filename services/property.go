package services

import (
	"encoding/json"
	"os"

	"Rental-Property-REST-API/models"
	"github.com/beego/beego/v2/server/web"
)



//PropertiesStore akta PropertySlice er wrapper struct. jeta JSON file ke struct e
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

func NewPropertyService(PropertyData *PropertyData) *PropertyService {
	return &PropertyService{
		PropertyData: PropertyData,
	}
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
			Lon: 		 property.LonLat.Coordinates[0],
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
func (service *PropertyService) GetAllResponseProperties() ([]models.ResponseProperty, error) {
    var responseProperties []models.ResponseProperty

    for _, property := range service.PropertyData.Properties {
        responseProperty, err := TransformProperty(property)
        if err != nil {
            return nil, err
        }
        responseProperties = append(responseProperties, responseProperty)
    }
    return responseProperties, nil
}
