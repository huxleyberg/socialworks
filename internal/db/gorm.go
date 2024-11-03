package db

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect provides functionality for connecting to the database
func Connect(dialector gorm.Dialector) (*gorm.DB, error) {
	var db *gorm.DB
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Close provides functionality for closing the database connection
func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
