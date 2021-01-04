package command

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/coronatorid/core-onator/provider"
	_ "github.com/golang-migrate/migrate/v4/source/file" // importing file path for mysql migrator
)

// Run is a command to run api engine
type Run struct {
	engine     provider.APIEngine
	inappCronn provider.InAppCron
}

// NewRun return CLI to run api engine
func NewRun(engine provider.APIEngine, inappCronn provider.InAppCron) *Run {
	return &Run{engine: engine, inappCronn: inappCronn}
}

// Use return how the command used
func (r *Run) Use() string {
	return "run:api"
}

// Example of the command
func (r *Run) Example() string {
	return "run:api"
}

// Short description about the command
func (r *Run) Short() string {
	return "Run API engine"
}

// Run the command with the args given by the caller
func (r *Run) Run(args []string) {
	log.Info().Msg("start coronator server")

	go func() {
		_ = r.engine.Run()
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 3 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// omit the error
	_ = r.engine.Shutdown(ctx)

	fmt.Println("\nGracefully shutdown the server...")
}
