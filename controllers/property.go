package controllers

import (
	"Rental-Property-REST-API/services"

	beego "github.com/beego/beego/v2/server/web"
)

// Keeping a global variable for PropertyService to be used in the controller methods
// and the assignment operation is done in main.go
var PropertyService  *services.PropertyService


type PropertyController struct {
	beego.Controller  // Beego gives us request/response information through the embedded controller
}

func (controller *PropertyController) GetProperties() {
	controller.Ctx.WriteString("GetProperties called");
}

func (controller *PropertyController) GetProperty() {
	controller.Ctx.WriteString("GetProperty called");
}