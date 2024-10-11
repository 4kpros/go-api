package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Establishes a connection to the database.
func ConnectDatabase() error {
	var dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		Env.PostGresHost,
		Env.PostGresUserName,
		Env.PostGresPassword,
		Env.PostGresDatabase,
		Env.PostGresPort,
		Env.PostGresSslMode,
		Env.PostGresTimeZone,
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
