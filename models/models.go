package models

import (
	"fmt"

	"github.com/eaglexpf/GoAdminInit/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	var (
		err                                                     error
		dbType, dbName, user, password, host, port, tablePrefix string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		fmt.Println("Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	port = sec.Key("PORT").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	DB, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbName))

	if err != nil {
		fmt.Println(err)
	}

	gorm.DefaultTableNameHandler = func(DB *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer DB.Close()
}
