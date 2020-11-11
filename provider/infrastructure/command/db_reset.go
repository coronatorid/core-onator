package command

import (
	"database/sql"
	"fmt"
)

// DBReset is a command to migrating mysql database all way down
type DBReset struct {
	db            *sql.DB
	databaseName  string
	migrationPath string
}

// NewDBReset return CLI to migrating database version
func NewDBReset(db *sql.DB, databaseName, migrationPath string) *DBReset {
	return &DBReset{db: db, databaseName: databaseName, migrationPath: migrationPath}
}

// Use return how the command used
func (d *DBReset) Use() string {
	return "db:reset"
}

// Example of the command
func (d *DBReset) Example() string {
	return "db:reset"
}

// Short description about the command
func (d *DBReset) Short() string {
	return "Migrate coronator database all way down"
}

// Run the command with the args given by the caller
func (d *DBReset) Run(args []string) {
	m, err := migrator(d.db, d.databaseName, d.migrationPath)
	if err != nil {
		fmt.Printf("Migration error because of: %v (%s)\n", err, err.Error())
		return
	}

	if err := m.Down(); err != nil {
		fmt.Printf("Migration error because of: %v (%s)\n", err, err.Error())
		return
	}

	fmt.Println("Migration process success")
}
