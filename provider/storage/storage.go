package storage

import (
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/provider/storage/api"
)

// Storage provide function for managing storage
type Storage struct {
}

// Fabricate storage provider
func Fabricate() *Storage {
	return &Storage{}
}

// FabricateAPI related to storage
func (r *Storage) FabricateAPI(engine provider.APIEngine) {
	engine.InjectAPI(api.NewRetrieve())
}
