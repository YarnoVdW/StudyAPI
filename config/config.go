package config

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("postgres",
		"port=5432 user=postgres dbname=Study password=yarno")
	if err != nil {
		panic(err.Error())
	}
	DB = db

	return DB
}

func GetDb() *gorm.DB {
	return DB
}

const (
	IdentityKey = "id"
	Key         = "secret_key_1911"
)
