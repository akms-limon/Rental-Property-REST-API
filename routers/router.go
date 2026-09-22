// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"Rental-Property-REST-API/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	// Route for our API endpoints
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/properties",
			beego.NSRouter("", &controllers.PropertyController{}, "get:GetProperties"),
			beego.NSRouter("/:id", &controllers.PropertyController{}, "get:GetProperty"),
		),
	)

	beego.AddNamespace(ns)
}