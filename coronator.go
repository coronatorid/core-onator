package main

import (
	"fmt"

	"github.com/coronatorid/core-onator/provider/admin"
	"github.com/coronatorid/core-onator/provider/inappcron"
	"github.com/coronatorid/core-onator/provider/report"
	"github.com/coronatorid/core-onator/provider/storage"
	"github.com/coronatorid/core-onator/provider/tracker"
	"github.com/rs/zerolog"

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

	// Initiate zero log
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

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

	// InAppCron
	inAppCron := inappcron.Fabricate()

	// API
	apiEngine := api.Fabricate(inAppCron)
	apiEngine.FabricateCommand(cmd)

	// User
	user := user.Fabricate(db)

	// Altair
	altair, err := altair.Fabricate(network)
	if err != nil {
		panic(err)
	}

	// Auth
	auth, err := auth.Fabricate(memcached, whatsappTextPublisher, user, altair, infra.WhatsappTextPublisher)
	if err != nil {
		panic(err)
	}
	auth.FabricateAPI(apiEngine)
	auth.FabricateInAppCronjob(inAppCron)

	// Tracker
	tracker := tracker.Fabricate(db)
	tracker.FabricateAPI(apiEngine)

	// Report
	report := report.Fabricate(db)
	report.FabricateAPI(apiEngine)

	// Storage
	storageProvider := storage.Fabricate()
	storageProvider.FabricateAPI(apiEngine)

	// Admin
	adminProvider := admin.Fabricate(altair, auth, user, report)
	adminProvider.FabricateAPI(apiEngine)

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
