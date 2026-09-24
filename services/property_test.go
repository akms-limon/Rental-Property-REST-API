package services

import (
	"testing"

	"Rental-Property-REST-API/models"
)

// test function for TransformProperty
func TestTransformProperty(t *testing.T) {
	tests := []struct {
		name     string
		property models.SourceProperty
	}{
		{
			name: "transform property test",
			property: models.SourceProperty{
				ID:                   "TEST-001",
				Feed:                 11,
				City:                 "Dhaka",
				Country:              "Bangladesh",
				CountryCode:          "BD",
				PropertyName:         "Test Property",
				LocationID:           "LOC-001",
				USDPrice:             100,
				PropertyTypeCategory: "Hotel",
				AmenityCategories:    []string{"Internet", "Parking"},
				Images:               []string{"image1.jpg", "image2.jpg"},
				LonLat: models.SourceLonLat{
					Coordinates: []float64{90.4125, 23.8103},
				},
				Categories: `[{"LocationID":"LOC-001","Name":"Dhaka","Type":"city","Slug":"dhaka","Display":["Dhaka"]}]`,
			},
		},
	}

	// Test runner
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, err := TransformProperty(test.property)
			if err != nil {
				t.Fatalf("TransformProperty() returned an error: %v", err)
			}

			if response.ID != test.property.ID {
				t.Errorf("ID = %v, want %v", response.ID, test.property.ID)
			}
			if response.Property.Image.Count != len(test.property.Images) {
				t.Errorf("Image.Count = %v, want %v", response.Property.Image.Count, len(test.property.Images))
			}
			if response.GeoInfo.Lon != test.property.LonLat.Coordinates[0] {
				t.Errorf("Lon = %v, want %v", response.GeoInfo.Lon, test.property.LonLat.Coordinates[0])
			}
			if response.GeoInfo.Lat != test.property.LonLat.Coordinates[1] {
				t.Errorf("Lat = %v, want %v", response.GeoInfo.Lat, test.property.LonLat.Coordinates[1])
			}
			if len(response.GeoInfo.Breadcrumbs) != 1 {
				t.Errorf("Breadcrumbs length = %v, want 1", len(response.GeoInfo.Breadcrumbs))
			}
			if response.GeoInfo.Breadcrumbs[0].Name != "Dhaka" {
				t.Errorf("Breadcrumb name = %v, want Dhaka", response.GeoInfo.Breadcrumbs[0].Name)
			}
		})
	}
}


// Helper functions to create pointers for basic types
func intPtr(value int) *int {
	return &value
}

func boolPtr(value bool) *bool {
	return &value
}

func floatPtr(value float64) *float64 {
	return &value
}

// Test function for FilterProperties
func TestFilterProperties(t *testing.T) {
	properties := []models.SourceProperty{
		{
			ID:                   "P1",
			Feed:                 11,
			Published:            false,
			USDPrice:             100,
			StarRating:           4,
			ReviewScoreGeneral:   4.5,
			NumberOfReview:       100,
			PropertyTypeCategory: "Hotel",
			BedroomCount:         2,
			AmenityCategories:    []string{"Internet", "Parking"},
		},
		{
			ID:                   "P2",
			Feed:                 11,
			Published:            true,
			USDPrice:             200,
			StarRating:           5,
			ReviewScoreGeneral:   4.8,
			NumberOfReview:       200,
			PropertyTypeCategory: "Villa",
			BedroomCount:         3,
			AmenityCategories:    []string{"Pool"},
		},
		{
			ID:                   "P3",
			Feed:                 12,
			Published:            false,
			USDPrice:             150,
			StarRating:           3,
			ReviewScoreGeneral:   3.5,
			NumberOfReview:       50,
			PropertyTypeCategory: "Apartment",
			BedroomCount:         1,
			AmenityCategories:    []string{"Parking"},
		},
	}

	tests := []struct {
		name     string
		filters  models.PropertyFilters
		expected []string
	}{
		{
			name: "feed and published filters",
			filters: models.PropertyFilters{
				Feed:      intPtr(11),
				Published: boolPtr(false),
			},
			expected: []string{"P1"},
		},
		{
			name: "minimum price and property type",
			filters: models.PropertyFilters{
				MinPrice:    floatPtr(150),
				PropertyType: "Villa",
			},
			expected: []string{"P2"},
		},
		{
			name: "amenities OR filter",
			filters: models.PropertyFilters{
				Amenities: []string{"Internet", "Pool"},
			},
			expected: []string{"P1", "P2"},
		},
		{
			name: "combined feed and amenities filters",
			filters: models.PropertyFilters{
				Feed:     intPtr(11),
				Amenities: []string{"Internet", "Parking"},
			},
			expected: []string{"P1"},
		},
		{
			name: "empty result",
			filters: models.PropertyFilters{
				Feed: intPtr(24),
			},
			expected: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FilterProperties(properties, test.filters)

			if len(result) != len(test.expected) {
				t.Fatalf("result length = %v, want %v", len(result), len(test.expected))
			}

			for i, property := range result {
				if property.ID != test.expected[i] {
					t.Errorf("result[%d].ID = %v, want %v", i, property.ID, test.expected[i])
				}
			}
		})
	}
}



// Test function for GetPropertyByID
func TestGetPropertyByID(t *testing.T) {
	properties := []models.SourceProperty{
		{
			ID:                   "P1",
			Feed:                 11,
			City:                 "Dhaka",
			Country:              "Bangladesh",
			CountryCode:          "BD",
			PropertyName:         "Test Property",
			LocationID:           "LOC-001",
			PropertyTypeCategory: "Hotel",
			AmenityCategories:    []string{"Internet"},
			Images:               []string{"image1.jpg"},
			LonLat: models.SourceLonLat{
				Coordinates: []float64{90.4125, 23.8103},
			},
			Categories: `[{"LocationID":"LOC-001","Name":"Dhaka","Type":"city","Slug":"dhaka","Display":["Dhaka"]}]`,
		},
	}

	service := NewPropertyService(&PropertyData{
		Properties: properties,
	})

	tests := []struct {
		name       string
		id         string
		wantID     string
		wantError  bool
	}{
		{
			name:      "property found",
			id:        "P1",
			wantID:    "P1",
			wantError: false,
		},
		{
			name:      "property not found",
			id:        "P999",
			wantID:    "",
			wantError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := service.GetPropertyByID(test.id)

			if (err != nil) != test.wantError {
				t.Errorf("GetPropertyByID() error = %v, wantError %v", err, test.wantError)
				return
			}

			if result.ID != test.wantID {
				t.Errorf("GetPropertyByID() ID = %v, want %v", result.ID, test.wantID)
			}
		})
	}
}


// Test function for GetAllProperties
func TestGetAllProperties(t *testing.T) {
	properties := []models.SourceProperty{
		{
			ID:                   "P1",
			Feed:                 11,
			Published:            false,
			City:                 "Dhaka",
			Country:              "Bangladesh",
			CountryCode:          "BD",
			PropertyName:         "Property One",
			LocationID:           "LOC-001",
			PropertyTypeCategory: "Hotel",
			USDPrice:             100,
			AmenityCategories:    []string{"Internet"},
			Images:               []string{"image1.jpg"},
			LonLat: models.SourceLonLat{
				Coordinates: []float64{90.4125, 23.8103},
			},
			Categories: `[{"LocationID":"LOC-001","Name":"Dhaka","Type":"city","Slug":"dhaka","Display":["Dhaka"]}]`,
		},
		{
			ID:                   "P2",
			Feed:                 11,
			Published:            true,
			City:                 "Dhaka",
			Country:              "Bangladesh",
			CountryCode:          "BD",
			PropertyName:         "Property Two",
			LocationID:           "LOC-002",
			PropertyTypeCategory: "Villa",
			USDPrice:             200,
			AmenityCategories:    []string{"Pool"},
			Images:               []string{"image2.jpg"},
			LonLat: models.SourceLonLat{
				Coordinates: []float64{90.4200, 23.8200},
			},
			Categories: `[{"LocationID":"LOC-002","Name":"Dhaka","Type":"city","Slug":"dhaka","Display":["Dhaka"]}]`,
		},
	}

	service := NewPropertyService(&PropertyData{
		Properties: properties,
	})

	limit := 1

	filters := models.PropertyFilters{
		Feed:  intPtr(11),
		Limit: &limit,
	}

	result, err := service.GetAllProperties(filters)

	if err != nil {
		t.Fatalf("GetAllProperties() returned an error: %v", err)
	}
	if result.Result.Count != 1 {
		t.Errorf("Count = %v, want 1", result.Result.Count)
	}
	if len(result.Result.Items) != 1 {
		t.Errorf("Items length = %v, want 1", len(result.Result.Items))
	}
	if result.Result.Items[0].ID != "P1" {
		t.Errorf("Items[0].ID = %v, want P1", result.Result.Items[0].ID)
	}
}