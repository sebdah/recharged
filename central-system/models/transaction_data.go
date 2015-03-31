package models

type TransactionData struct {
	values []*MeterValue `type="MeterValue" required="true"`
}
