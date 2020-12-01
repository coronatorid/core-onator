package main

import (
	"fmt"

	"github.com/coronatorid/core-onator/provider/tracker"

	"github.com/coronatorid/core-onator/provider/altair"
	"github.com/coronatorid/core-onator/provider/api"
	"github.com/coronatorid/core-onator/provider/auth"
	"github.com/coronatorid/core-onator/provider/command"
	"github.com/coronatorid/core-onator/provider/infrastructure"
	"github.com/coronatorid/core-onator/provider/user"
	"github.com/subosito/gotenv"
)

func main() {
	_ = gotenv.Load()

	// Command
	cmd := command.Fabricate()

	// Infra
	infra, err := infrastructure.Fabricate()
	if err != nil {
		panic(err)
	}
	defer infra.Close()

	memcached := infra.Memcached()
	whatsappTextPublisher, err := infra.WhatsappTextPublisher()
	if err != nil && err.Error() == "whatsapp not initiated yet" {
		fmt.Println("WHATSAPP IS NOT INITIATED. PLEASE INITIATE WHATSAPP CONNECTION TO MAKE CORE-ONATOR WORK!")
	} else if err != nil {
		panic(err)
	}

	db, err := infra.DB()
	if err != nil {
		panic(err)
	}

	network := infra.Network()

	if err := infra.FabricateCommand(cmd); err != nil {
		panic(err)
	}

	// API
	apiEngine := api.Fabricate()
	apiEngine.FabricateCommand(cmd)

	// User
	user := user.Fabricate(db)

	// Altair
	altair, err := altair.Fabricate(network)
	if err != nil {
		panic(err)
	}

	// Auth
	auth, err := auth.Fabricate(memcached, whatsappTextPublisher, user, altair)
	if err != nil {
		panic(err)
	}
	auth.FabricateAPI(apiEngine)

	// Tracker
	tracker := tracker.Fabricate(db)
	tracker.FabricateAPI(apiEngine)

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
