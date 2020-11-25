package main

import (
	"github.com/coronatorid/core-onator/provider/api"
	"github.com/coronatorid/core-onator/provider/auth"
	"github.com/coronatorid/core-onator/provider/command"
	"github.com/coronatorid/core-onator/provider/infrastructure"
	"github.com/subosito/gotenv"
)

func main() {
	_ = gotenv.Load()

	cmd := command.Fabricate()

	infra, err := infrastructure.Fabricate()
	if err != nil {
		panic(err)
	}
	defer infra.Close()

	memcached := infra.Memcached()
	whatsappTextPublisher, err := infra.WhatsappTextPublisher()
	if err != nil {
		panic(err)
	}

	if err := infra.FabricateCommand(cmd); err != nil {
		panic(err)
	}

	apiEngine := api.Fabricate()
	apiEngine.FabricateCommand(cmd)

	auth, err := auth.Fabricate(memcached, whatsappTextPublisher)
	if err != nil {
		panic(err)
	}
	auth.FabricateAPI(apiEngine)

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
