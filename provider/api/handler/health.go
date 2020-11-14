package handler

import "github.com/coronatorid/core-onator/provider"

// Health api handler
type Health struct{}

// NewHealth create new health object
func NewHealth() *Health {
	return &Health{}
}

// Path return api path
func (h *Health) Path() string {
	return "/health"
}

// Method return api method
func (h *Health) Method() string {
	return "GET"
}

// Handle health which always return 200
func (h *Health) Handle(context provider.APIContext) error {

	return nil
}
