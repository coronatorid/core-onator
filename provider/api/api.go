package api

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/provider/api/command"
	"github.com/coronatorid/core-onator/provider/api/handler"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	logEcho "github.com/labstack/gommon/log"
)

// API ...
type API struct {
	engine *echo.Echo
	port   int
}

// Fabricate API
func Fabricate() *API {
	engine := echo.New()
	engine.Logger.SetLevel(logEcho.OFF)

	// We will do this one later on
	// engine.Logger.SetLevel(log.OFF)

	return &API{
		engine: engine,
		port:   2019,
	}
}

// FabricateCommand insert api related command
func (a *API) FabricateCommand(cmd provider.Command) {
	cmd.InjectCommand(
		command.NewRun(a),
	)
}

// InjectAPI inject new API into coronator
func (a *API) InjectAPI(handler provider.APIHandler) {
	a.engine.Add(handler.Method(), handler.Path(), func(context echo.Context) error {
		startTime := time.Now()
		req := context.Request()
		reqID := ""

		if reqID = req.Header.Get("X-Request-ID"); reqID != "" {
			context.Set("request-id", reqID)
		} else {
			reqID = uuid.New().String()
			context.Set("request-id", reqID)
		}

		if userID := req.Header.Get("Resource-Owner-ID"); userID != "" {
			convertedUserID, err := strconv.Atoi(userID)
			if err == nil {
				context.Set("user-id", convertedUserID)
			}
		}

		handler.Handle(FabricateContext(context))

		log.Info().
			Str("request_id", reqID).
			Str("remote_ip", context.RealIP()).
			Str("method", context.Request().Method).
			Str("uri", context.Request().URL.String()).
			Str("host", context.Request().Host).
			Dur("duration", time.Since(startTime)).
			Str("user_agent", context.Request().Header.Get("User-Agent")).
			Str("user_id", context.Request().Header.Get("Resource-Owner-ID")).
			Str("path", context.Path()).
			Int("response_status", context.Response().Status).
			Int("response_status_group", (context.Response().Status/100)*100).
			Msg(fmt.Sprintf("coronator api watcher: %s %s", context.Request().Method, context.Path()))
		return nil
	})
}

// Run api engine
func (a *API) Run() error {
	a.engine.Use(middleware.Logger())
	a.InjectAPI(handler.NewHealth())
	return a.engine.Start(fmt.Sprintf(":%d", a.port))
}

// Shutdown api engine
func (a *API) Shutdown(ctx context.Context) error {
	return a.engine.Shutdown(ctx)
}
