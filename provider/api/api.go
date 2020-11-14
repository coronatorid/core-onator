package api

import (
	"context"
	"fmt"

	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/provider/api/command"
	"github.com/coronatorid/core-onator/provider/api/handler"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// API ...
type API struct {
	engine *echo.Echo
	port   int
}

// Fabricate API
func Fabricate() *API {
	return &API{
		engine: echo.New(),
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
		requestID := uuid.New()
		context.Set("request_id", requestID.String())
		handler.Handle(context)

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
