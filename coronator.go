package main

import (
	"github.com/coronatorid/core-onator/provider/command"
	"github.com/coronatorid/core-onator/provider/infrastructure"
	"github.com/subosito/gotenv"
)

func main() {
	_ = gotenv.Load()

	cmd := command.Fabricate()

	infra := infrastructure.Fabricate()
	defer infra.Close()

	if err := infra.FabricateCommand(cmd); err != nil {
		panic(err)
	}

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
