package types

type IdTagInfo struct {
	ExpiryDate string   `json:"expiryDate" type:"string" required="false"`
	GroupTagId *IdToken `json:"groupTagId" type:"idToken" required="false"`
	Status     string   `json:"status" type:"string" required:"true"`
	Language   string   `json:"language" type="LanguageCode" required="false"`
}

// Constructor
func NewIdTagInfo() (idTagInfo *IdTagInfo) {
	idTagInfo = new(IdTagInfo)
	idTagInfo.Language = "en"
	idTagInfo.Status = AuthorizationStatusInvalid

	return
}
