package controllers

import (
	"Rental-Property-REST-API/models"
	"Rental-Property-REST-API/services"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Keeping a global variable for PropertyService to be used in the controller methods
// and the assignment operation is done in main.go
var PropertyService *services.PropertyService

type PropertyController struct {
	beego.Controller // Beego gives us request/response information through the embedded controller
}

// Method to handle GET requests for all properties
// GetProperties returns a list of rental properties.
// @Title Get Properties
// @Description Get rental properties with optional filters.
// @Param limit query int false "Maximum number of properties to return"
// @Success 200 {object} models.PropertyListResponse
// @Failure 400 {object} models.ErrorResponse
// @router / [get]
func (controller *PropertyController) GetProperties() {

	var limit = 10

	properties, err := PropertyService.GetAllResponseProperties(&limit)
	if err != nil {
		logs.Error("Failed to get properties: %v", err)

		controller.Ctx.ResponseWriter.WriteHeader(500)
		controller.Data["json"] = models.ErrorResponse{
			Error: err.Error(),
		}
		controller.ServeJSON()
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