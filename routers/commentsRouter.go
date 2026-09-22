package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["Rental-Property-REST-API/controllers:PropertyController"] = append(beego.GlobalControllerRouter["Rental-Property-REST-API/controllers:PropertyController"],
        beego.ControllerComments{
            Method: "GetProperties",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Rental-Property-REST-API/controllers:PropertyController"] = append(beego.GlobalControllerRouter["Rental-Property-REST-API/controllers:PropertyController"],
        beego.ControllerComments{
            Method: "GetProperty",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
