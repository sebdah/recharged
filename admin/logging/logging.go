package logging

import (
	"os"

	"github.com/op/go-logging"
)

// Configure logging
func Setup() {
	// Create a logging backend
	backend := logging.NewLogBackend(os.Stderr, "", 0)

	// Set formatting
	format := logging.MustStringFormatter("%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}")
	backendFormatter := logging.NewBackendFormatter(backend, format)

	// Use the backends
	logging.SetBackend(backendFormatter)
}
