package registTable

import (
	"github.com/jinzhu/gorm"
	"wenku/model/dbModel"
)

func RegistTable(db *gorm.DB) {
	db.AutoMigrate(dbModel.User{})
}
