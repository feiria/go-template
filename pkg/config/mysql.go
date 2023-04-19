package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go-template/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type _mysql struct {
}

var Mysql = new(_mysql)

func (*_mysql) InitMySQL() {
	username := viper.Get("db.mysql.username")
	password := viper.Get("db.mysql.password")
	host := viper.Get("db.mysql.host")
	port := viper.Get("db.mysql.port")
	Dbname := viper.Get("db.mysql.schema")
	timeout := "10s"

	//拼接下dsn参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		username, password, host, port, Dbname, timeout)

	//自定义日志模板 打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //彩色
		},
	)

	var err error
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	sqlDB, _ := global.DB.DB()

	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。

	fmt.Println("MySQL init")
}
