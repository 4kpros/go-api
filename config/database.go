package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToPostgresDB() (err error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		AppEnv.PostGresHost,
		AppEnv.PostGresUserName,
		AppEnv.PostGresPassword,
		AppEnv.PostGresDatabase,
		AppEnv.PostGresPort,
		AppEnv.PostGresSslMode,
		AppEnv.PostGresTimeZone,
	)
	DB, err = gorm.Open(
		postgres.New(postgres.Config{
			DSN: dsn,
		}),
		&gorm.Config{},
	)
	return
}
