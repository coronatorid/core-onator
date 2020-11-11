package command

import (
	"database/sql"
	"fmt"
)

// PingMYSQL is a command to ping mysql database
type PingMYSQL struct {
	db *sql.DB
}

// NewPingMYSQL return CLI to ping mysql database
func NewPingMYSQL(db *sql.DB) *PingMYSQL {
	return &PingMYSQL{db: db}
}

// Use return how the command used
func (p *PingMYSQL) Use() string {
	return "ping:mysql"
}

// Example of the command
func (p *PingMYSQL) Example() string {
	return "ping:mysql"
}

// Short description about the command
func (p *PingMYSQL) Short() string {
	return "Ping coronator mysql database"
}

// Run the command with the args given by the caller
func (p *PingMYSQL) Run(args []string) {
	err := p.db.Ping()
	if err != nil {
		fmt.Printf("Failed pingin mysql because of: %v (%s)\n", err, err.Error())
		return
	}

	fmt.Println("Successfully pinging mysql.")
}
