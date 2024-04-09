// Package write is to write the log messages to the console.
package write

import (
	"log/slog"
	"os"
)

// handler is to create a handler for the logger.
// Logger is to create a logger to handle the errors.
var (
	handler = slog.NewTextHandler(os.Stdout, nil)
	Logger  = slog.New(handler)
)
