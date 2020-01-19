package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shikanon/IoTOrbHub/config"
	"log"
)

// 连接数据库
func DbConn() *gorm.DB {
	user := config.MysqlConfig.User
	pwd := config.MysqlConfig.Pwd
	host := config.MysqlConfig.Host
	port := config.MysqlConfig.Port
	dbname := config.MysqlConfig.Db

	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, host, port, dbname)
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		log.Fatal(err)
	}
	db.SingularTable(true) // 表名为结构体变量名复数。true，不开启复数。
	return db
}

// 插入一条数据，返回记录id
func MysqlInsertOneData(data interface{}) (id int) {
	db := DbConn()
	defer db.Close()
	var result []int
	db.Create(data)
	db.Raw("select LAST_INSERT_ID() as id").Pluck("id", &result)
	return result[0]
}
