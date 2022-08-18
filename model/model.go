package model

import "time"

type StudyItem struct {
	Base
	ID            uint   `gorm:"primary_key" json:"id"`
	Description   string `json:"description"`
	numberOfLikes uint   `json:"numberOfLikes`
}

type Base struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at`
	UpdatedAt time.Time `json:"updated_at`
	DeletedAt time.Time `json:"deleted_at`
}

type User struct {
	Base
	Username string `json:"username"`
	Password string `json:"password"`
}
