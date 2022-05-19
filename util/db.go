package util

import (
	"fmt"
	"github.com/go-ini/ini"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

type DbConf struct {
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
}

var Db *gorm.DB
var dbConf *DbConf
var once sync.Once

func init() {
	once.Do(func() {
		dbConf = &DbConf{}
		iniFile, _ := ini.Load("conf.ini")
		err := iniFile.Section("db").MapTo(dbConf)
		if err != nil {
			panic("conf.ini error, err=" + err.Error())
			return
		}
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			dbConf.Username,
			dbConf.Password,
			dbConf.Host,
			dbConf.Port,
			dbConf.Dbname)
		Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("db connection error, err=" + err.Error())
		}
		sqlDB, _ := Db.DB()
		// 设置数据库连接池参数
		sqlDB.SetMaxOpenConns(100) // 设置数据库连接池最大连接数
		sqlDB.SetMaxIdleConns(20)  // 连接池最大允许的空闲连接数，超过的连接会被连接池关闭
	})
}
