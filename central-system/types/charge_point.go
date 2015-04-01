package types

type ChargePoint struct {
	model        string `max_length:20`
	serialNumber string `max_length:25`
	vendor       string `max_length:20`
	evses        []*Evse
}
