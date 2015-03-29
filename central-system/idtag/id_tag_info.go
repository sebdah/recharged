package idtag

type IdTagInfo struct {
	// Optional. Unix datetime object
	expiryDate int

	// Optional
	parentIdToken IdToken

	// Mandatory. AuthorizationStatus
	status int
}
