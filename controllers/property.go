package controllers

import (
	"Rental-Property-REST-API/services"
	"Rental-Property-REST-API/models"
	"Rental-Property-REST-API/validators"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/core/logs"
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
// @Success 200 {object} models.PropertyListResponse
// @Failure 400 {object} models.ErrorResponse
// @router /properties [get]
func (controller *PropertyController) GetProperties() {
	limitText := controller.GetString("limit")
	limit, err := validators.ValidateLimit(limitText)
	limit = limit
	if err != nil {
		logs.Error("Error validating limit: ", err)
		controller.Ctx.ResponseWriter.WriteHeader(400)
		controller.Data["json"] = models.ErrorResponse{Error: err.Error()}
		controller.ServeJSON()
		return
	}

	properties, err := PropertyService.GetAllResponseProperties()
	if err != nil {
		controller.Ctx.ResponseWriter.WriteHeader(500)
		controller.Data["json"] = models.ErrorResponse{Error: err.Error()}
		controller.ServeJSON()
		return
	}

	controller.Data["json"] = properties
	controller.ServeJSON()
}

// Method to handle GET requests for a single property by ID
func (controller *PropertyController) GetProperty() {
	if PropertyService == nil {
		controller.Ctx.WriteString("PropertyService is nill or not initialized.")
		return
	}
	controller.Ctx.WriteString("Property is available.")
}