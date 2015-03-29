package types

type MeterValue struct {
	// Mandatory, unix timestamp
	timestamp int

	// Mandatory
	value string

	// Optional, ReadingContext
	context int

	// Optional, ValueFormat
	format int

	// Optional, Measurand
	measurand int

	// Optional, Location
	location int

	// Optional, UnitOfMeasure
	unit string
}
