package db

import (
	"log"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// ProvidePostgres returns a function that provides a Postgres client
func ProvidePostgres(postgresDbUrl string) *gorm.DB {

	// Connect to postgres
	postgresProviderDB, err := gorm.Open(postgres.Open(postgresDbUrl), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "public.", // perform operations on the "public" schema
		}})
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to connect to postgres"))
	}

	// Get the underlying sql.DB object to set connection pool configurations
	sqlDB, err := postgresProviderDB.DB()
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to get sql.DB from gorm"))
	}

	// Set connection pool settings
	sqlDB.SetMaxOpenConns(30)                  // Maximum number of open connections to the database
	sqlDB.SetMaxIdleConns(30)                  // Maximum number of idle connections
	sqlDB.SetConnMaxIdleTime(15 * time.Minute) // Maximum amount of time a connection may be idle

	// Run the migrations
	if err := Migrate(postgresDbUrl); err != nil {
		log.Fatal(errors.Wrap(err, "failed to run migrations").Error())
	}

	return postgresProviderDB
}
