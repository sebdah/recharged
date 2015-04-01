package types

type LocalizedText struct {
	language string `type:"LanguageCodeType" required="false"`
	text     string `type="string" required="true" max_length="200"`
}
