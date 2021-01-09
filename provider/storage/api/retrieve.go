package api

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/codefluence-x/aurelia"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Retrieve api handler
type Retrieve struct {
}

// NewRetrieve find previously created report
func NewRetrieve() *Retrieve {
	return &Retrieve{}
}

// Path return api path
func (r *Retrieve) Path() string {
	return "/storage/*"
}

// Method return api method
func (r *Retrieve) Method() string {
	return "GET"
}

// Handle request otp
func (r *Retrieve) Handle(context provider.APIContext) {

	signature := context.QueryParam("signature")
	expiresAt := context.QueryParam("expires_at")

	if signature == "" || expiresAt == "" {
		context.NoContent(http.StatusBadRequest)
		log.Error().
			Str("request_id", util.GetRequestID(context)).
			Array("tags", zerolog.Arr().Str("provider").Str("storage").Str("retrieve")).
			Msg("request is bad")
		return
	}

	expiresAtUnix, err := strconv.Atoi(expiresAt)
	if err != nil {
		log.Error().
			Str("request_id", util.GetRequestID(context)).
			Array("tags", zerolog.Arr().Str("provider").Str("storage").Str("retrieve")).
			Msg("expires at is not valid")
		context.NoContent(http.StatusBadRequest)
		return
	}

	path := context.Request().URL.Path
	if aurelia.Authenticate(os.Getenv("APP_ENCRIPTION_KEY"), fmt.Sprintf("%d%s", expiresAtUnix, path[1:len(path)]), signature) == false {
		log.Error().
			Str("request_id", util.GetRequestID(context)).
			Array("tags", zerolog.Arr().Str("provider").Str("storage").Str("retrieve")).
			Msg("signature is not valid")
		context.NoContent(http.StatusUnauthorized)
		return
	}

	if time.Now().After(time.Unix(int64(expiresAtUnix), 0)) {
		log.Error().
			Str("request_id", util.GetRequestID(context)).
			Array("tags", zerolog.Arr().Str("provider").Str("storage").Str("retrieve")).
			Msg("content already expired")
		context.NoContent(http.StatusNotFound)
		return
	}

	_ = context.File(path[1:len(path)])
}
