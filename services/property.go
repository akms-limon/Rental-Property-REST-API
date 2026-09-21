package services

import (
	"encoding/json"
	"os"

	"Rental-Property-REST-API/models"
	"github.com/beego/beego/v2/server/web"
)


//PropertiesStore akta PropertySlice er wrapper struct. jeta JSON file ke struct e
// convert korar por propertySlice er vitor store kore rakhe.

type PropertiesStore struct {
	PropertySlice []models.SourceProperty
}

func PropertiesReader() (*PropertiesStore, error) {
	path, err := web.AppConfig.String("property_data_path")
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var propertySlice []models.SourceProperty

	err = json.Unmarshal(data, &propertySlice)
	if err != nil {
		return nil, err
	}

	return &PropertiesStore{
		PropertySlice: propertySlice,
	}, nil
}