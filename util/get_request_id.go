package util

import "github.com/coronatorid/core-onator/provider"

// GetRequestID from a context
func GetRequestID(ctx provider.Context) string {
	if requestID, ok := ctx.Get("request-id").(string); ok {
		return requestID
	}

	return "missing-request-id"
}
