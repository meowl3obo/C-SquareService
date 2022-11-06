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
