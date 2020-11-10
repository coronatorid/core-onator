package provider

import (
	"github.com/coronatorid/core-onator/provider/command"
)

// Provider ...
type Provider struct{}

// Fabricate new provider
func Fabricate() *Provider {
	return &Provider{}
}

// Command provide command for core-onator CLI
func (p *Provider) Command() Command {
	return command.Fabricate()
}
