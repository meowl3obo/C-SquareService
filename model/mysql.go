package model

type Users struct {
	Email    string `gorm:"primaryKey"`
	Password string
	Status   int
	CreateAt string
	UpdateAt string
}
