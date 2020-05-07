package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	Id   int
	Name string
}

func RegisterDB() {

	mysqlhost := beego.AppConfig.String("mysql::host")
	mysqlport := beego.AppConfig.String("mysql::port")
	mysqluser := beego.AppConfig.String("mysql::user")
	mysqlpwd := beego.AppConfig.String("mysql::pwd")
	mysqldb := beego.AppConfig.String("mysql::dbserv")
	maxIdle, _ := beego.AppConfig.Int("mysql::maxIdle")
	maxConn, _ := beego.AppConfig.Int("mysql::maxConn")

	connet := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", mysqluser, mysqlpwd, mysqlhost, mysqlport, mysqldb)

	fmt.Print("### ", connet, "\n")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", connet, maxIdle, maxConn)
	orm.RegisterModel(new(Users))
	orm.RunSyncdb("default", false, true)
}
