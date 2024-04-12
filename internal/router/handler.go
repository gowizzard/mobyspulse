// Package router is to define the different router for the http server.
package router

import (
	"github.com/gowizzard/mobyspulse/internal/handler"
	"github.com/gowizzard/mobyspulse/internal/middleware"
	"github.com/gowizzard/mobyspulse/internal/redirect"
	"net/http"
)

// Handler are to register the different handlers.
func Handler(mux *http.ServeMux) {
	mux.HandleFunc(http.MethodGet+" /", redirect.Metrics)
	mux.HandleFunc(http.MethodGet+" /favicon.ico", handler.Favicon)
	mux.HandleFunc(http.MethodGet+" /metrics", middleware.Authentication(handler.Metrics))
}
