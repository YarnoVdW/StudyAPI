package model

type StudyItem struct {
	Base
	ID            uint   `gorm:"primary_key" json:"id"`
	Description   string `json:"description"`
	numberOfLikes uint   `json:"numberOfLikes`
}

type Base struct {
	ID uint `gorm:"primary_key" json:"id"`
}
