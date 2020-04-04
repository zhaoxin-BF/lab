package main

import (
	"github.com/astaxie/beego/plugins/cors"
	_ "lab_system/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	//设置跨域
	beego.InsertFilter("*",
		beego.BeforeRouter,
		cors.Allow(&cors.Options{
			AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders: []string{
				"Origin",
				"Authorization",
				"Access-Control-Allow-Origin",
				"content-type",
				"Cookie"},
			ExposeHeaders: []string{
				"Content-Length",
				"Access-Control-Allow-Origin"},
			AllowCredentials: true,
			AllowAllOrigins: true,
			//AllowOrigins: []string{
			//	"https://devops.ucloudadmin.com",
			//	"https://devops-pre.ucloudadmin.com",
			//	"http://devops.pre.ucloudadmin.com",
			//	"https://ops-udisk.ucloudadmin.com",
			//	"https://ussov2.ucloudadmin.com",
			//	"https://mikasa.ucloudadmin.com",
			//	"http://172.18.178.200:8080",
			//	"http://172.18.178.201:8080",
			//	"http://localhost:3000",
			//	"http://localhost:8080",
			//},
		}))

	beego.Run()
}

