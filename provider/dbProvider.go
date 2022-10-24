package provider

import (
	"fmt"

	. "C-SquaredService/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	UserName string = "meowl"
	Password string = "meowl870706"
	Network  string = "tcp"
	Addr     string = "66.42.44.23"
	Port     int    = 3306
	Database string = "c_squared"
)

var dbconn *gorm.DB
var ConnErr error

func init() {
	addr := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", UserName, Password, Network, Addr, Port, Database)

	dbconn, ConnErr = gorm.Open(mysql.Open(addr), &gorm.Config{})
	if ConnErr != nil {
		fmt.Println("Error！connection to mysql failed, ", ConnErr)
		return
	}
}

func CreateUser(userData User) error {
	result := dbconn.Table("users").Create(&userData)
	return result.Error
}

func GetParentClassify() (error, []ParentClassify) {
	var parentClassify []ParentClassify
	result := dbconn.Table("parent_classify").Find(&parentClassify)
	return result.Error, parentClassify
}

func GetChildClassify() (error, []ChildClassify) {
	var childClassify []ChildClassify
	result := dbconn.Table("child_classify").Find(&childClassify)
	return result.Error, childClassify
}
