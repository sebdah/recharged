package idtag

type IdTag struct {
	idTag IdToken
}

// Constructor
func NewIdTag(idTag IdToken) *IdTag {
	tag := new(IdTag)
	tag.idTag = idTag

	return tag
}
