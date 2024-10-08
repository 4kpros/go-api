package migrate

import (
	roleMigrate "github.com/4kpros/go-api/services/role/migrate"
	userMigrate "github.com/4kpros/go-api/services/user/migrate"
)

func Start() {
	roleMigrate.Migrate() // User migrations
	userMigrate.Migrate() // User migrations
}
