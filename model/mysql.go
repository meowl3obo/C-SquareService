package model

type User struct {
	Email      string `gorm:"primaryKey"`
	Password   string
	GoogleId   string
	FacebookId string
	Name       string
	Phone      int
	Address    string
	Status     int
	CreateAt   string
	UpdateAt   string
}

type Product struct {
	Id             string `gorm:"primaryKey"`
	Name           string
	MainImg        string
	OtherImg       string
	Intro          string
	Illustrate     string
	ParentClassify int
	ChildClassify  int
	Status         int
	Price          int
}

type ParentClassify struct {
	Id   int `gorm:"primaryKey"`
	Name string
	Path string
}

type ChildClassify struct {
	Id       int `gorm:"primaryKey"`
	ParentId int
	Name     string
	Path     string
}
