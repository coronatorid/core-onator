package infrastructure

import (
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/dghubble/sling"

	"github.com/Rhymen/go-whatsapp"
	"github.com/coronatorid/core-onator/provider/infrastructure/adapter"

	"github.com/bradfitz/gomemcache/memcache"
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

	memcachedMutex  *sync.Once
	memcached       *adapter.Memcached
	memcachedConfig struct {
		host string
	}

	whatsappMutex   *sync.Once
	whatsapp        *whatsapp.Conn
	whatsappSession whatsapp.Session
}

// Fabricate infrastructure interface for coronator
func Fabricate() (*Infrastructure, error) {
	i := &Infrastructure{
		mysqlMutex:     &sync.Once{},
		memcachedMutex: &sync.Once{},
		whatsappMutex:  &sync.Once{},
	}

	i.mysqlConfig.username = os.Getenv("DATABASE_USERNAME")
	i.mysqlConfig.password = os.Getenv("DATABASE_PASSWORD")
	i.mysqlConfig.host = os.Getenv("DATABASE_HOST")
	i.mysqlConfig.port = os.Getenv("DATABASE_PORT")
	i.mysqlConfig.name = os.Getenv("DATABASE_NAME")

	i.memcachedConfig.host = os.Getenv("MEMCACHE_HOST")

	decodedEncKey, err := hex.DecodeString(os.Getenv("WHATSAPP_ENC_KEY"))
	if err != nil {
		return nil, err
	}

	decodedMacKey, err := hex.DecodeString(os.Getenv("WHATSAPP_MAC_KEY"))
	if err != nil {
		return nil, err
	}

	i.whatsappSession.ClientId = os.Getenv("WHATSAPP_CLIENT_ID")
	i.whatsappSession.ClientToken = os.Getenv("WHATSAPP_CLIENT_TOKEN")
	i.whatsappSession.ServerToken = os.Getenv("WHATSAPP_SERVER_TOKEN")
	i.whatsappSession.EncKey = decodedEncKey
	i.whatsappSession.MacKey = decodedMacKey
	i.whatsappSession.Wid = os.Getenv("WHATSAPP_WID")

	return i, nil
}

// FabricateCommand fabricate all infrastructure related commands
func (i *Infrastructure) FabricateCommand(cmd provider.Command) error {
	db, _ := i.MYSQL()

	cmd.InjectCommand(
		command.NewPingMYSQL(db),
		command.NewDBMigrate(db, i.mysqlConfig.name, "file://migration"),
		command.NewDBReset(db, i.mysqlConfig.name, "file://migration"),
		command.NewWhatsappLogin(i.Whatsapp()),
	)

	return nil
}

// Close all initiated connection
func (i *Infrastructure) Close() {
	if i.mysqlDB != nil {
		_ = i.mysqlDB.Close()
	}
}

// DB wrapper for default golang sql
func (i *Infrastructure) DB() (provider.DB, error) {
	db, err := i.MYSQL()
	if err != nil {
		return nil, err
	}

	return adapter.AdaptSQL(db), nil
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

// Memcached provide cache interface
func (i *Infrastructure) Memcached() provider.Cache {
	i.memcachedMutex.Do(func() {
		i.memcached = adapter.AdaptMemcache(memcache.New(i.memcachedConfig.host))
	})

	return i.memcached
}

// WhatsappTextPublisher return adapted whatsapp client
func (i *Infrastructure) WhatsappTextPublisher() (provider.TextPublisher, error) {
	client, err := i.WhatsappOldSession()
	if err != nil {
		return nil, err
	}

	return adapter.AdaptWhatsapp(client), nil
}

// WhatsappOldSession return old whatsapp session
func (i *Infrastructure) WhatsappOldSession() (*whatsapp.Conn, error) {
	if i.whatsappSession.ClientId == "" || i.whatsappSession.ClientToken == "" || i.whatsappSession.ServerToken == "" {
		return nil, errors.New("whatsapp not initiated yet")
	}

	var err error
	i.whatsappMutex.Do(func() {
		i.whatsapp = i.Whatsapp()
		_, err = i.whatsapp.RestoreWithSession(i.whatsappSession)
	})

	if err != nil {
		i.whatsappMutex = &sync.Once{}
		return nil, err
	}

	return i.whatsapp, nil
}

// Whatsapp create new whatsapp connection
func (i *Infrastructure) Whatsapp() *whatsapp.Conn {
	wac, _ := whatsapp.NewConn(30 * time.Second)
	return wac
}

// Network provider
func (i *Infrastructure) Network() provider.Network {
	return adapter.AdaptNetwork(i.Sling())
}

// Sling return sling library used for network feature
func (i *Infrastructure) Sling() *sling.Sling {
	return sling.New()
}
