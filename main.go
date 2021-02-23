package main

import (
	"github.com/marshhu/ma-api/interface/ioc"
	"github.com/marshhu/ma-api/router"
	"github.com/marshhu/ma-frame/app"
	"github.com/marshhu/ma-frame/orm"
)

func main() {
	app := app.NewApp(8081)
	app.RegisterRouter(router.Register())
	initDb()
	ioc.InitIoc()
	app.Run()
}

//初始化数据库
func initDb() {
	var config = orm.DbSettings{
		Dialect:         "mysql",
		DbName:          "ma_db",
		Host:            "127.0.0.1",
		User:            "root",
		Password:        "123456",
		Port:            3306,
		MaxIdleConns:    5,
		MaxOpenConns:    20,
		ConnMaxLifetime: 5,
	}
	orm.Init(&config)
}
