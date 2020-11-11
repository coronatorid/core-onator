package infrastructure

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	"github.com/coronatorid/core-onator/provider/infrastructure/command"

	"github.com/coronatorid/core-onator/provider"
	_ "github.com/go-sql-driver/mysql" // Import mysql driver
)

// Infrastructure provide infrastructure interface
type Infrastructure struct {
	mysqlMutex  *sync.Once
	mysqlDB     *sql.DB
	mysqlConfig struct {
		username string
		password string
		host     string
		port     string
		name     string
	}
}

// Fabricate infrastructure interface for coronator
func Fabricate() *Infrastructure {
	i := &Infrastructure{
		mysqlMutex: &sync.Once{},
	}

	i.mysqlConfig.username = os.Getenv("DATABASE_USERNAME")
	i.mysqlConfig.password = os.Getenv("DATABASE_PASSWORD")
	i.mysqlConfig.host = os.Getenv("DATABASE_HOST")
	i.mysqlConfig.port = os.Getenv("DATABASE_PORT")
	i.mysqlConfig.name = os.Getenv("DATABASE_NAME")
	return i
}

// FabricateCommand fabricate all infrastructure related commands
func (i *Infrastructure) FabricateCommand(cmd provider.Command) error {
	db, _ := i.MYSQL()

	cmd.InjectCommand(
		command.NewPingMYSQL(db),
		command.NewDBMigrate(db, i.mysqlConfig.name, "file://migration"),
		command.NewDBReset(db, i.mysqlConfig.name, "file://migration"),
	)

	return nil
}

// Close all initiated connection
func (i *Infrastructure) Close() {
	if i.mysqlDB != nil {
		_ = i.mysqlDB.Close()
	}
}

// MYSQL provide mysql interface
func (i *Infrastructure) MYSQL() (*sql.DB, error) {
	i.mysqlMutex.Do(func() {
		// Currently there are no possible error while fabricating this so the error handling is ignored
		db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&interpolateParams=true", i.mysqlConfig.username, i.mysqlConfig.password, i.mysqlConfig.host, i.mysqlConfig.port, i.mysqlConfig.name))
		i.mysqlDB = db
	})

	return i.mysqlDB, nil
}
