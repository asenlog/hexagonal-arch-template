package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"app/internal/config"
	"app/internal/infra/http"
	"app/internal/infra/repository/sql"
)

func main() {
	logger := log.New(os.Stdout, "SERVICE : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	if err := run(logger); err != nil {
		log.Println("main: error:", err)
		runtime.Goexit()
	}
}

func run(log *log.Logger) error {
	ctx := context.Background()

	var err error

	defer func() {
		if err != nil {
			log.Fatal("main", "%v", err)
		}
	}()

	// Init configuration (env vars)
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("main: API failed to instantiate config")

		return err
	}

	// Init database connection pool
	db, err := sql.New(ctx, cfg.DB, cfg.DB.WriterHost)
	if err != nil {
		log.Fatal("main: API failed to instantiate database")

		return err
	}

	// Bootstrap the HTTP component and create the router for our app.
	api := http.New().Server()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// Start the service listening for requests.
	go func() {
		log.Printf("main: API listening on %s", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error, %w", err)

	case sig := <-shutdown:
		log.Printf("main: %v : Start shutdown", sig)

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
		defer cancel()

		// Asking listener to shutdown and load shed.
		err := api.Shutdown(ctx)
		if err != nil {
			log.Printf("main: Graceful shutdown did not complete in %v : %v", cfg.ShutdownTimeout, err)
			err = api.Close()
		}

		// Log the status of this shutdown.
		switch {
		case sig == syscall.SIGSTOP:
			log.Printf("")
			// nolint
			return fmt.Errorf("integrity issue caused shutdown")
		case err != nil:
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}
