package main

import (
	"fmt"
	"go-serv/models"
	_ "go-serv/routers"

	"github.com/astaxie/beego"
)

func init() {
	models.RegisterDB()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	fmt.Print(beego.AppConfig.String("mysql::host"))
	beego.Run("0.0.0.0:9000")
}
