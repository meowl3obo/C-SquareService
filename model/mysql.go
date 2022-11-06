package model

type GormErr struct {
	Number  int    `json:"Number"`
	Message string `json:"Message"`
}

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
	ParentClassify string
	ChildClassify  string
	Status         int
	Price          int
	IsNew          bool
	IsSale         bool
}

type Inventory struct {
	Id             string `gorm:"primaryKey"`
	ProductId      string
	Color          string
	Size           float64
	Unit           string
	PreOrderAmount int
	NowAmount      int
}

type ParentClassify struct {
	Id   string `gorm:"primaryKey"`
	Name string
	Path string
	Sort int
}

type ChildClassify struct {
	Id       string `gorm:"primaryKey"`
	ParentId string
	Name     string
	Path     string
	Sort     int
}
