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
	mysqlMutex *sync.Once
	mysqlDB    *sql.DB
}

// Fabricate infrastructure interface for coronator
func Fabricate() *Infrastructure {
	return &Infrastructure{
		mysqlMutex: &sync.Once{},
	}
}

// FabricateCommand fabricate all infrastructure related commands
func (i *Infrastructure) FabricateCommand(cmd provider.Command) error {
	db, err := i.MYSQL()
	if err != nil {
		return err
	}

	cmd.InjectCommand(
		command.NewPingMYSQL(db),
	)

	return nil
}

// MYSQL provide mysql interface
func (i *Infrastructure) MYSQL() (*sql.DB, error) {
	i.mysqlMutex.Do(func() {
		// Currently there are no possible error while fabricating this so the error handling is ignored
		db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&interpolateParams=true", os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_NAME")))
		i.mysqlDB = db
	})

	return i.mysqlDB, nil
}
