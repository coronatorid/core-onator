package main

import (
	"github.com/coronatorid/core-onator/provider"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	p := provider.Fabricate()
	_ = p.Command().Execute()
}
