package controllers

import (
	"Rental-Property-REST-API/models"
	"Rental-Property-REST-API/services"
	"Rental-Property-REST-API/validators"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Keeping a global variable for PropertyService to be used in the controller methods
// and the assignment operation is done in main.go
var PropertyService *services.PropertyService

type PropertyController struct {
	beego.Controller // Beego gives us request/response information through the embedded controller
}

// @Title Get Properties
// @Description Get rental properties with optional filters.
// @Param min_price query number false "Minimum USD price"
// @Param max_price query number false "Maximum USD price"
// @Param min_star_rating query int false "Minimum star rating"
// @Param min_review_score query number false "Minimum review score"
// @Param min_reviews query int false "Minimum number of reviews"
// @Param published query boolean false "Published status"
// @Param property_type query string false "Property type: Hotel, House, Apartment, Villa, Resort, Hostel"
// @Param feed query int false "Feed: 11, 12, 22, or 24"
// @Param min_bedroom query int false "Minimum number of bedrooms"
// @Param amenities query string false "Comma-separated amenities, e.g. Internet,Parking"
// @Param limit query int false "Maximum number of properties to return"
// @Success 200 {object} models.PropertyListResponse
// @Failure 400 {object} models.ErrorResponse
// @router / [get]
func (controller *PropertyController) GetProperties() {
	query := controller.Ctx.Request.URL.Query()
	var filters models.PropertyFilters

	if len(query) > 0 {
		if err := validators.ValidatePropertyParameters(query); err != nil {
			logs.Error(err)
			controller.CustomAbort(400, err.Error())
			return
		}
		var err error
		filters, err = validators.ParsePropertyFilters(query)
		if err != nil {
			logs.Error(err)
			controller.CustomAbort(400, err.Error())
			return
		}
	}

	properties, err := PropertyService.GetAllProperties(filters)
	if err != nil {
		logs.Error(err)
		controller.CustomAbort(500, err.Error())
		return
	}

	controller.Data["json"] = properties
	controller.ServeJSON()
}

// @Title Get Property
// @Description Get a rental property by ID.
// @Param id path string true "Property ID"
// @Success 200 {object} models.ResponseProperty
// @Failure 404 {object} models.ErrorResponse
// @router /:id [get]
func (controller *PropertyController) GetProperty() {
	propertyID := controller.Ctx.Input.Param(":id")

	property, err := PropertyService.GetPropertyByID(propertyID)
	if err != nil {
		logs.Error("Property not found: %v", err)

		controller.Ctx.ResponseWriter.WriteHeader(404)
		controller.Data["json"] = models.ErrorResponse{
			Error: "Property not found",
		}
		controller.ServeJSON()
		return
	}
	controller.Data["json"] = property
	controller.ServeJSON()
}
