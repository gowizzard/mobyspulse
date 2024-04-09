// Package main provides the entry point for the application. It initializes the service and starts the http server.
package main

import (
	"context"
	request "github.com/gowizzard/mobyspulse/internal/requests"
	"github.com/gowizzard/mobyspulse/internal/router"
	"github.com/gowizzard/mobyspulse/internal/write"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"
)

// init is to log the start of the service to the console. It also pings the docker unix socket to check if it is available.
// If the ping fails, the service will exit with an error. If the ping is successful, the service will log the status to the console.
func init() {

	write.Logger.Info("Start the mobyspulse service.")

	status, err := request.Ping()
	if err != nil {
		write.Logger.Error("Ping the docker unix socket.", "err", err)
		os.Exit(1)
	}

	if !reflect.DeepEqual(status, http.StatusOK) {
		write.Logger.Error("Ping the docker unix socket.", "status", status)
		os.Exit(1)
	}

	write.Logger.Info("Ping the docker unix socket.", "status", status)

}

// main is to start the http server and listen for signals to shut down the service. If the service receives a signal, it will shutdown the http server.
// If the shutdown is successful, the service will log the status to the console. If the shutdown fails, the service will log the error to the console.
func main() {

	mux := http.NewServeMux()
	router.Handler(mux)

	server := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			write.Logger.Error("Listen and serve the http requests.", "err", err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	select {
	case <-signals:

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			write.Logger.Error("Shutdown the http server.", "err", err)
		}

	}

}
