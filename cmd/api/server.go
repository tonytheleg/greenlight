package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (app *application) serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)

	go func() {
		// quit channel to carry os.Signal values
		quit := make(chan os.Signal, 1)

		// signal.Notify listens for incoming SIGINT/SIGTERM and relays to
		// quit channel
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		// reads signal from quit channel
		// this blocks until a signal is received
		s := <-quit

		app.logger.PrintInfo("caught signal", map[string]string{
			"signal": s.String(),
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// calls shutdown on server, passing the context with a timeout of 5 seconds
		// should return nil or an error if failed or timeout hit
		shutdownError <- srv.Shutdown(ctx)
	}()

	app.logger.PrintInfo("starting server", map[string]string{
		"addr": srv.Addr,
		"env":  app.config.env,
	})

	// calling shutdown immediately returns ErrServerClosed
	// this checks for graceful shutdown
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	// waits to receive return value from shutdown
	// if its an error there was issue with graceful shutdown
	err = <-shutdownError
	if err != nil {
		return err
	}

	app.logger.PrintInfo("stopped server", map[string]string{
		"addr": srv.Addr,
	})

	return nil
}
