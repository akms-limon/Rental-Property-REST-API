package main

import (
	_ "Rental-Property-REST-API/routers"
	"Rental-Property-REST-API/services"
	"Rental-Property-REST-API/controllers"
	
	beego "github.com/beego/beego/v2/server/web"
)

func main() {

	// Load the data from JSON file once at the start of the application and store it in memory for later use.
	propertyData, err := services.LoadPropertiesFromFile()
	if err != nil {
		panic(err)
	}

	// Create a new instance of PropertyService with the loaded data
	propertyService := services.NewPropertyService(propertyData)

	// Assign the created PropertyService instance to the global variable
	controllers.PropertyService = propertyService

	beego.Run()
}
