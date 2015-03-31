package models

type IdToken struct {
	id     string `type:"string" required="true" max_length="50"`
	idType string `type:"idType" required="false"`
}
