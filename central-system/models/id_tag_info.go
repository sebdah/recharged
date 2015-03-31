package models

type IdTagInfo struct {
	expiryDate string   `type:"string" required="false"`
	groupTagId *IdToken `type:"idToken" required="false"`
	status     string   `type:"string" required:"true"`
	language   string   `type="LanguageCode" required="false"`
}
