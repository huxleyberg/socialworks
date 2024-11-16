package db

import (
	"gorm.io/gorm"
)

// Close provides functionality for closing the database connection
func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
