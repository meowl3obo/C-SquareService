package model

type User struct {
	Email    string `gorm:"primaryKey"`
	Password string
	Status   int
	CreateAt string
	UpdateAt string
}
