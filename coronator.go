package main

import (
	"github.com/coronatorid/core-onator/provider"
)

func main() {
	p := provider.Fabricate()
	_ = p.Command().Execute()
}
