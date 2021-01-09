package api

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/codefluence-x/aurelia"
	"github.com/coronatorid/core-onator/provider"
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
	req := context.Request()

	signature := req.URL.Query().Get("signature")
	expiresAt := req.URL.Query().Get("expires_at")

	if signature == "" || expiresAt == "" {
		context.NoContent(http.StatusBadRequest)
		return
	}

	expiresAtUnix, err := strconv.Atoi(expiresAt)
	if err != nil {
		context.NoContent(http.StatusBadRequest)
		return
	}

	path := req.URL.Path
	if aurelia.Authenticate(os.Getenv("APP_ENCRIPTION_KEY"), fmt.Sprintf("%d%s", expiresAtUnix, path[1:len(path)]), signature) == false {
		context.NoContent(http.StatusUnauthorized)
		return
	}

	if time.Now().After(time.Unix(int64(expiresAtUnix), 0)) {
		context.NoContent(http.StatusNotFound)
		return
	}
}
