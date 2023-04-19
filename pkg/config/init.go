package config

import "os"

var (
	TemplatePath string
)

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	TemplatePath = pwd
}

func Init(env string) {
	Viper.InitViper(env)
	Zap.InitZap()
	Mysql.InitMySQL()
}
