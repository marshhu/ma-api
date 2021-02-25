package main

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/marshhu/ma-api/src/interface/ioc"
	"github.com/marshhu/ma-api/src/router"
	"github.com/marshhu/ma-frame/app"
	"github.com/marshhu/ma-frame/config"
	"github.com/marshhu/ma-frame/orm"
	"github.com/marshhu/ma-frame/utils"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	initDb()
	ioc.InitIoc()
	app := app.NewApp(app.AppConfig{
		HostPort:     viper.GetInt("app.hostPort"),
		ReadTimeout:  viper.GetInt("app.readTimeout"),
		WriteTimeout: viper.GetInt("app.writeTimeout")})
	app.RegisterRouter(router.Register())
	app.Run()
}

//初始化配置文件
func initConfig() {
	defaultCfgName := "config"
	defaultCfgEnv := "debug" //调试
	defaultCfgType := "yaml"

	cfgName := strings.TrimSpace(os.Getenv("CFG_FILE"))
	if utils.IsEmpty(cfgName) {
		cfgName = defaultCfgName
	}

	// 根据环境变量加载不同的配置文件
	env := strings.TrimSpace(os.Getenv("ENV"))
	if utils.IsEmpty(env) {
		env = defaultCfgEnv
	}
	cfgFileName := cfgName
	if !utils.IsEmpty(env) {
		cfgFileName += "." + env
	}

	// 添加多个配置文件路径，以先找到的为准
	rootDir := utils.RootDir()
	configDir := filepath.FromSlash(path.Join(rootDir, "config"))
	envDir := os.Getenv("CFG_DIR")
	cfgDirs := []string{envDir, configDir, rootDir}

	config.Init(config.ConfigSetting{CfgFile: cfgFileName, CfgDirs: cfgDirs, CfgType: defaultCfgType})
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
