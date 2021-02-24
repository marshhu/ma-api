package main

import (
	"github.com/marshhu/ma-api/config"
	"github.com/marshhu/ma-api/router"
	"github.com/marshhu/ma-api/src/interface/ioc"
	"github.com/marshhu/ma-frame/app"
	"github.com/marshhu/ma-frame/orm"
	"github.com/spf13/viper"
)

func main() {
	config.Init()
	app := app.NewApp(app.AppConfig{
		HostPort:     viper.GetInt("app.hostPort"),
		ReadTimeout:  viper.GetInt("app.readTimeout"),
		WriteTimeout: viper.GetInt("app.writeTimeout")})
	app.RegisterRouter(router.Register())
	initDb()
	ioc.InitIoc()
	app.Run()
}

//初始化数据库
func initDb() {
	var config = orm.DbSettings{
		Dialect:         viper.GetString("db.dialect"),
		DbName:          viper.GetString("db.dbName"),
		Host:            viper.GetString("db.host"),
		User:            viper.GetString("db.user"),
		Password:        viper.GetString("db.password"),
		Port:            viper.GetInt("db.port"),
		MaxIdleConns:    viper.GetInt("db.maxIdleConns"),
		MaxOpenConns:    viper.GetInt("db.maxOpenConns"),
		ConnMaxLifetime: viper.GetInt("db.connMaxLifetime"),
	}
	orm.Init(&config)
}
