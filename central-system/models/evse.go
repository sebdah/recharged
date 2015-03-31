package models

type Evse struct {
	id         int
	connectors []*Connector
}
