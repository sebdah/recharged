package types

type IdTagInfo struct {
	ExpiryDate string   `type:"string" required="false"`
	GroupTagId *IdToken `type:"idToken" required="false"`
	Status     string   `type:"string" required:"true"`
	Language   string   `type="LanguageCode" required="false"`
}
