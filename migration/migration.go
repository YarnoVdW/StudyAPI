package migration

import (
	"github.com/YarnoVdW/StudyAPI/go-service/cmd/service/model"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.StudyItem{})
}
