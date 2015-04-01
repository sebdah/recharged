package models

type IdTag struct {
	idToken   *IdToken   `type="*IdToken" required="true"`
	idTagInfo *IdTagInfo `type:="*IdTagInfo"`
}
