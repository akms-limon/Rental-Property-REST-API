// @APIVersion 1.0.0
// @Title Rental Property API
// @Description REST API for searching and retrieving rental properties
// @Contact Your Name
package routers

import (
	"Rental-Property-REST-API/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/properties",
			beego.NSInclude(&controllers.PropertyController{}),
		),
	)

	beego.AddNamespace(ns)
}
