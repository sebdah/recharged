package models

import "github.com/op/go-logging"

var log logging.Logger

// Ensure all indexes
func EnsureAllIndexes() {
	log.Info("Ensuring indexes")
	EnsureIndexes(new(IdTag))
	EnsureIndexes(new(ChargePoint))
}
