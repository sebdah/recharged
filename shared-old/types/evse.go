package types

type Evse struct {
	id         int
	connectors []*Connector
}
