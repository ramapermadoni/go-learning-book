package migration

import (
	"embed"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/gorm"
)

//go:embed sql_migrations/*.sql
var dbMigrations embed.FS

func Initiator(dbParam *gorm.DB) {
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "sql_migrations",
	}

	// You need to get the underlying *sql.DB from GORM for sql-migrate
	sqlDB, err := dbParam.DB()
	if err != nil {
		panic(err)
	}

	n, errs := migrate.Exec(sqlDB, "postgres", migrations, migrate.Up)
	if errs != nil {
		panic(errs)
	}

	fmt.Println("Migration success, applied", n, "migrations!")
}
