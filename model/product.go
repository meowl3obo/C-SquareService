package model

type Classify struct {
	Id    string
	Name  string
	Path  string
	Child []ChildClassify
}

type InputProduct struct {
	Id             string `gorm:"primaryKey"`
	Name           string
	MainImg        string
	OtherImg       string
	Intro          string
	Illustrate     string
	ParentClassify string
	ChildClassify  string
	Status         string
	Price          string
	IsNew          bool
	IsSale         bool
}

type ProductInventory struct {
	Id             string
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
	Inventorys     []Inventory
	OtherImgs      []string
}
