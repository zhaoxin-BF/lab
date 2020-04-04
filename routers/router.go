// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"lab_system/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/courses_table",
			beego.NSInclude(
				&controllers.CoursesTableController{},
			),
		),

		beego.NSNamespace("/labs_table",
			beego.NSInclude(
				&controllers.LabsTableController{},
			),
		),

		beego.NSNamespace("/users_table",
			beego.NSInclude(
				&controllers.UsersTableController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
