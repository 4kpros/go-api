package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Establishes a connection to the database.
func ConnectDatabase() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		Env.PostgresHost,
		Env.PostgresUsername,
		Env.PostgresPassword,
		Env.PostgresDatabase,
		Env.PostgresPort,
		Env.PostgresSslMode,
		Env.PostgresTimeZone,
	)
	var err error
	DB, err = gorm.Open(
		postgres.New(postgres.Config{
			DSN: dsn,
		}),
		&gorm.Config{},
	)
	return err
}
