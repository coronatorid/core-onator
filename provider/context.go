package provider

import "context"

// Context flow generated from API
type Context interface {
	Ctx() context.Context

	// Get retrieves data from the context.
	Get(key string) interface{}

	// Set saves data in the context.
	Set(key string, val interface{})
}
