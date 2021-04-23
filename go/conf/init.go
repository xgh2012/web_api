package conf

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"log"
)

var (
	Engine *xorm.Engine
)

func init() {
	args := "root:123456@tcp(127.0.0.1:3306)/admin?charset=utf8"
	db, err := xorm.NewEngine("mysql", args)
	if err != nil || db == nil {
		log.Fatalln("DataBase init error:", err)
	}
	Engine = db
	Engine.ShowSQL(true)
}
