package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/marshhu/ma-frame/utils"
	"github.com/spf13/viper"
)

var (
	defaultCfgName = "config" // 默认配置文件名
	defaultCfgEnv  = ""       // 默认环境
	defaultCfgType = "yaml"   // 默认配置文件格式
)

func Init() {
	configFile := getConfigName()
	viper.SetConfigName(configFile)     // name of config file (without extension)
	viper.SetConfigType(defaultCfgType) // REQUIRED if the config file does not have the extension in the name
	configDirs := getConfigDirs()
	for _, dir := range configDirs {
		if utils.IsEmpty(dir) {
			continue
		}
		viper.AddConfigPath(dir) // path to look for the config file in
	}

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("加载配置文件失败: %s \n", err))
	}
	log.Println("使用的配置文件位置：" + viper.ConfigFileUsed())
}

func getConfigName() string {
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

	return cfgFileName
}

func getConfigDirs() []string {
	// 添加多个配置文件路径，以先找到的为准
	rootDir := utils.RootDir()
	configDir := filepath.FromSlash(path.Join(rootDir, "config"))
	envDir := os.Getenv("CFG_DIR")

	etcDirs := []string{envDir, configDir, rootDir}
	return etcDirs
}
