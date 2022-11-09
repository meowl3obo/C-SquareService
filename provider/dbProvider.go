package provider

import (
	"encoding/json"
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
	result := dbconn.Table("parent_classify").Order("sort").Find(&parentClassify)
	return result.Error, parentClassify
}

func GetChildClassify() (error, []ChildClassify) {
	var childClassify []ChildClassify
	result := dbconn.Table("child_classify").Order("sort").Find(&childClassify)
	return result.Error, childClassify
}

func CreateProduct(productData Product) error {
	result := dbconn.Table("product").Create(&productData)
	var newError GormErr
	if result.Error != nil {
		byteErr, _ := json.Marshal(result.Error)
		json.Unmarshal((byteErr), &newError)
		if newError.Number != 1062 {
			return result.Error
		}
	}
	return nil
}

func CreateInventory(inventoryData Inventory) error {
	result := dbconn.Table("product_inventory").Create(&inventoryData)
	return result.Error
}

func GetSaleProduct() (error, []Product) {
	var products []Product
	result := dbconn.Table("product").Where("is_sale = ?", "1").Find(&products)
	return result.Error, products
}

func GetNewProduct() (error, []Product) {
	var products []Product
	result := dbconn.Table("product").Where("is_new = ?", "1").Find(&products)
	return result.Error, products
}

func GetProduct(parentID string, childID string) (error, []Product) {
	var products []Product
	if childID == "" {
		result := dbconn.Table("product").Where("parent_classify = ?", parentID).Find(&products)
		return result.Error, products
	} else {
		result := dbconn.Table("product").Where("parent_classify = ? AND child_classify = ?", parentID, childID).Find(&products)
		return result.Error, products
	}
}
