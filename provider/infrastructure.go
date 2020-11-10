package provider

import "database/sql"

// Infrastructure provide infrastructure interface
type Infrastructure interface {
	MYSQL() (*sql.DB, error)
}
