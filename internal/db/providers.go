package db

import (
	"log"

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
			TablePrefix: "socialworks.", // perform operations on the "socialworks" schema
		}})
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to connect to postgres"))
	}

	// Run the migrations
	if err := Migrate(postgresDbUrl); err != nil {
		log.Fatal(errors.Wrap(err, "failed to run migrations").Error())
	}

	return postgresProviderDB
}
