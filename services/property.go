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

// GetAllProperties returns all properties from the PropertyData.
func (service *PropertyService) GetAllProperties() []models.SourceProperty {
	return service.PropertyData.Properties
}
