package types

type KeyValues struct {
	// Mandatory, max 50 chars
	key string

	// Mandatory
	readonly bool

	// Optional, max 500 chars
	value string
}
