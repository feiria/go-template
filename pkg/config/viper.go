package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type _viper struct {
}

var Viper = new(_viper)

const (
	defaultConfigFilePath = "/conf/"
	defaultConfigFilename = "template"
	defaultConfigFileType = "yaml"
)

func (*_viper) InitViper(env string) {
	filename := fmt.Sprintf("%s-%s", defaultConfigFilename, env)

	viper.AddConfigPath(TemplatePath + defaultConfigFilePath)
	viper.SetConfigFile(TemplatePath + defaultConfigFilePath + filename + "." + defaultConfigFileType)
	viper.SetConfigType(defaultConfigFileType)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&Config{}); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	fmt.Println("viper init")
}

// Config 总配置
type Config struct {
	DB     dbConfig     `mapstructure:"db"`
	Server serverConfig `mapstructure:"server"`
	Jwt    jwtConfig    `mapstructure:"jwt"`
	Log    logConfig    `mapstructure:"zap"`
}

// ServerConfig 服务配置
type serverConfig struct {
	Debug bool   `mapstructure:"debug"` //代码中逻辑使用debug控制不同环境
	Log   bool   `mapstructure:"log"`   //日志单独控制，方便线上随时开启日志，定位问题
	Addr  string `mapstructure:"addr"`  //服务地址
	Udp   int    `mapstructure:"udp"`
}

// LogConfig 日志配置
type logConfig struct {
	Level         string `mapstructure:"level"` //日志级别
	Format        string `mapstructure:"format"`
	Prefix        string `mapstructure:"prefix"`
	Director      string `mapstructure:"director"`
	MaxAge        int    `mapstructure:"max-age"`
	ShowLine      bool   `mapstructure:"show-line"`
	EncoderLevel  string `mapstructure:"encoder-level"`
	StackTraceKey string `mapstructure:"stacktrace-key"`
	LogInConsole  bool   `mapstructure:"log-in-console"`
}

// 数据库配置
type dbConfig struct {
	Mysql mysqlConfig `mapstructure:"mysql"`
}

// mysql配置
type mysqlConfig struct {
	Host            string `mapstructure:"host"`              //主机地址
	Port            int    `mapstructure:"port"`              //端口
	Username        string `mapstructure:"username"`          //用户名
	Password        string `mapstructure:"password"`          //密码
	Schema          string `mapstructure:"schema"`            //约束(数据库)
	Charset         string `mapstructure:"charset"`           //字符集
	Location        string `mapstructure:"location"`          //位置
	Timeout         int    `mapstructure:"timeout"`           //连接超时时间
	ParseTime       bool   `mapstructure:"parse_time"`        //是否解析时间
	MaxOpenConn     int    `mapstructure:"max_open_conn"`     //最大连接数
	MaxIdleConn     int    `mapstructure:"max_idle_conn"`     //最大闲置连接数
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"` //连接最长维持时间
}

// 鉴权配置
type jwtConfig struct {
	SigningKey  string `mapstructure:"signing-key"`
	ExpiresTime string `mapstructure:"expires-time"`
	BufferTime  string `mapstructure:"buffer-time"`
	Issuer      string `mapstructure:"issuer"`
}
