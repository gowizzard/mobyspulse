// Package handler provides the handler functions for the http server.
package handler

import (
	"github.com/gowizzard/mobyspulse/files"
	"github.com/gowizzard/mobyspulse/internal/write"
	"net/http"
)

// Favicon is to return the favicon.ico file.
func Favicon(w http.ResponseWriter, _ *http.Request) {

	w.Header().Set("Content-Type", "image/x-icon")

	file, err := files.Assets.ReadFile("assets/favicon.ico")
	if err != nil {
		write.Logger.Error("Read the favicon.ico file.", "err", err)
	}

	_, err = w.Write(file)
	if err != nil {
		write.Logger.Error("Write the favicon.ico file.", "err", err)
	}

}
