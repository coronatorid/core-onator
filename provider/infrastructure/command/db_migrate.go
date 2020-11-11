package command

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file" // importing file path for mysql migrator
)

// DBMigrate is a command to migrating mysql database to a newer version
type DBMigrate struct {
	db            *sql.DB
	databaseName  string
	migrationPath string
}

// NewDBMigrate return CLI to migrating database version
func NewDBMigrate(db *sql.DB, databaseName, migrationPath string) *DBMigrate {
	return &DBMigrate{db: db, databaseName: databaseName, migrationPath: migrationPath}
}

// Use return how the command used
func (d *DBMigrate) Use() string {
	return "db:migrate"
}

// Example of the command
func (d *DBMigrate) Example() string {
	return "db:migrate"
}

// Short description about the command
func (d *DBMigrate) Short() string {
	return "Migrate coronator database to a newer version"
}

// Run the command with the args given by the caller
func (d *DBMigrate) Run(args []string) {
	m, err := migrator(d.db, d.databaseName, d.migrationPath)
	if err != nil {
		fmt.Printf("Migration error because of: %v (%s)\n", err, err.Error())
		return
	}

	if err := m.Up(); err != nil {
		fmt.Printf("Migration error because of: %v (%s)\n", err, err.Error())
		return
	}

	fmt.Println("Migration process success")
}

func migrator(db *sql.DB, databaseName, migrationPath string) (*migrate.Migrate, error) {
	driver, err := mysql.WithInstance(db, &mysql.Config{
		MigrationsTable: "db_versions",
		DatabaseName:    databaseName,
	})
	if err != nil {
		return nil, err
	}

	migrator, err := migrate.NewWithDatabaseInstance(migrationPath, "mysql", driver)
	if err != nil {
		return nil, err
	}

	return migrator, nil
}
