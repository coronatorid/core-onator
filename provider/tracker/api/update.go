package api

import (
	"encoding/json"
	"net/http"

	"github.com/coronatorid/core-onator/entity"

	"github.com/coronatorid/core-onator/provider"
)

// Track api handler
type Track struct {
	trackProvider provider.Tracker
}

// NewTrack create new request otp handler object
func NewTrack(trackProvider provider.Tracker) *Track {
	return &Track{trackProvider: trackProvider}
}

// Path return api path
func (r *Track) Path() string {
	return "/tracker/update"
}

// Method return api method
func (r *Track) Method() string {
	return "POST"
}

// Handle request otp
func (r *Track) Handle(context provider.APIContext) {
	userID := context.Get("user-id").(int)
	if userID <= 0 {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{"bad request given by client"},
			"message": "Bad request",
		})
		return
	}

	var request entity.TrackRequest
	if err := json.NewDecoder(context.Request().Body).Decode(&request); err != nil {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{"bad request given by client"},
			"message": "Bad request",
		})
		return
	}

	response, err := r.trackProvider.Track(context.Request().Context(), 0, request)
	if err != nil {
		_ = context.JSON(err.HTTPStatus, map[string]interface{}{
			"errors":  err.ErrorString(),
			"message": err.Error(),
		})
		return
	}

	_ = context.JSON(http.StatusOK, map[string]interface{}{
		"data": response,
	})
}
