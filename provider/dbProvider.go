package provider

import (
	"fmt"

	. "C-SquaredService/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	UserName string = "root"
	Password string = "meowl870706"
	Network  string = "tcp"
	Addr     string = "66.42.44.23"
	Port     int    = 3306
	Database string = "c_squared"
)

var dbconn *gorm.DB
var connErr error

func init() {
	addr := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", UserName, Password, Network, Addr, Port, Database)

	dbconn, connErr = gorm.Open(mysql.Open(addr), &gorm.Config{})
	if connErr != nil {
		fmt.Println("connection to mysql failed", connErr)
		return
	}
}

func CreateUser(userData User) error {
	result := dbconn.Table("users").Create(&userData)
	return result.Error
}

// enp1s0
