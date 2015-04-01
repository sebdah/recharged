package types

type MeterValue struct {
	timestamp string `type="string" required="true"`
	value     string `type="string" required="true" max_length="200"`
	context   string `type="ReadingContext" required="false"`
	format    string `type="ValueFormat" required="false"`
	measurand string `type="Measurand" required="false"`
	location  string `type="Location" required="false"`
	unit      string `type="UnitOfMeasure" required="false"`
}
