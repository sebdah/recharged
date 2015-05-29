package fixtures

import "github.com/op/go-logging"

var log logging.Logger

// Set up all fixtures
func Setup() {
	SetupChargePoint()
	SetupIdTag()
}
