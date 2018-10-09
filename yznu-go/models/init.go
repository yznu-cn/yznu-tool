package models

import (
	"fmt"

	"github.com/yznu-cn/yznu-tool/yznu-go/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db0      *gorm.DB
	NotFound = gorm.ErrRecordNotFound
)

func GetDB() *gorm.DB {
	if db0 == nil {
		var err error
		db0, err = getdb()
		if err != nil {
			panic("get db err:" + err.Error())
		}
		return db0
	}
	return db0
}

func getdb() (*gorm.DB, error) {
	server := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", config.MySqlConf.UserName, config.MySqlConf.PassWd, config.MySqlConf.Host, config.MySqlConf.Port, config.MySqlConf.DbName)
	db, err := gorm.Open("mysql", server)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10) //连接池的空闲数大小
	db.DB().SetMaxOpenConns(50) //最大打开连接数
	db.SingularTable(true)
	return db, nil
}
