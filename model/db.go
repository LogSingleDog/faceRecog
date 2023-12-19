package model

import (
	"fmt"
	//"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func Init() {
	DB = open()
}
func open() *gorm.DB {
	name := viper.GetString("db.name")
	pwd := viper.GetString("db.password")
	dataBase := viper.GetString("db.dataBase")
	config := fmt.Sprintf("%s:%s@/%s", name, pwd, dataBase)
	db, err := gorm.Open("mysql", config)
	if err != nil {
		fmt.Printf("init database error:%v\n", err)
		panic(err)
	} else {
		fmt.Printf("init success!\n")
	}
	return db
}
func InitTables() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&UserTmp{})
}
