package provider

import (
	"github.com/coronatorid/core-onator/provider/command"
	"github.com/coronatorid/core-onator/provider/infrastructure"
)

// Provider ...
type Provider struct {
	cmd *command.Command
}

// Fabricate new provider
func Fabricate() *Provider {
	return &Provider{
		cmd: command.Fabricate(),
	}
}

// InjectCommand inject other provider command
func (p *Provider) InjectCommand(scaffold command.Scaffold) {
	p.cmd.InjectCommand(scaffold)
}

// Infrastructure return Infrastructure interface
func (p *Provider) Infrastructure() Infrastructure {
	return infrastructure.Fabricate()
}
