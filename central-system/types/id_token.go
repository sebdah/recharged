package types

type IdToken struct {
	// Mandatory
	idToken string
}

func NewIdToken(idToken string) *IdToken {
	token := new(IdToken)
	token.idToken = idToken
	return token
}
