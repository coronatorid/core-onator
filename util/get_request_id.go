package util

import "context"

// GetRequestID from a context
func GetRequestID(ctx context.Context) string {
	if requestID, ok := ctx.Value("request-id").(string); ok {
		return requestID
	}

	return "missing-request-id"
}
