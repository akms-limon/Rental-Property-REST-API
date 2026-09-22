// @APIVersion 1.0.0
// @Title Rental Property API
// @Description REST API for searching and retrieving rental properties
package routers

import (
	"Rental-Property-REST-API/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/properties",
			beego.NSRouter("", &controllers.PropertyController{}, "get:GetProperties"),
			beego.NSRouter("/:id", &controllers.PropertyController{}, "get:GetProperty"),
		),
	)

	beego.AddNamespace(ns)
}