package models

// PropertyResponse is the response structure for any response from the server side.
// from server side we got a JSON and then we need to unmarshal the JSON to PropertyResponse structure to use the data in our program
type ResponseProperty struct {
	ID        string   `json:"ID"`
	Feed      int      `json:"Feed"`
	Published bool     `json:"Published"`
	GeoInfo   GeoInfo  `json:"GeoInfo"`
	Property  Property `json:"Property"`
}

type GeoInfo struct {
	Breadcrumbs []Breadcrumb `json:"Breadcrumbs"`
	City        string       `json:"City"`
	Country     string       `json:"Country"`
	CountryCode string       `json:"CountryCode"`
	Name        string       `json:"Name"`
	LocationID  string       `json:"LocationID"`
	Lat         float64      `json:"Lat"`
	Lon         float64      `json:"Lon"`
	State       string       `json:"State"`
	StateAbbr   string       `json:"StateAbbr"`
}

type Breadcrumb struct {
	LocationID string   `json:"LocationID"`
	Name       string   `json:"Name"`
	Type       string   `json:"Type"`
	Slug       string   `json:"Slug"`
	Display    []string `json:"Display"`
}

type Property struct {
	Amenities    []string `json:"Amenities"`
	Name         string   `json:"Name"`
	Slug         string   `json:"Slug"`
	PropertyType string   `json:"PropertyType"`
	Price        float64  `json:"Price"`
	ReviewScore  float64  `json:"ReviewScore"`
	StarRating   int      `json:"StarRating"`
	Counts       Counts   `json:"Counts"`
	Image        Image    `json:"Image"`
}

type Counts struct {
	Bathroom  int `json:"Bathroom"`
	Bedroom   int `json:"Bedroom"`
	Reviews   int `json:"Reviews"`
	Occupancy int `json:"Occupancy"`
}

type Image struct {
	Count  int      `json:"Count"`
	Images []string `json:"Images"`
}

type PropertyListResponse struct {
	Result PropertyListResult `json:"Result"`
}

type PropertyListResult struct {
	Count int                `json:"Count"`
	Items []ResponseProperty `json:"Items"`
}

type ErrorResponse struct {
	Error string `json:"Error"`
}