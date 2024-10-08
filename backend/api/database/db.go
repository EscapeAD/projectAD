package database

import (
	"escapead/projectAD/backend/api/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Connect - DB loader
func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(config.DBDRIVER, config.DBURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
