package messages

import "github.com/sebdah/recharged/central-system/models"

type AuthorizeReq struct {
	IdTag models.IdToken `json:"idTag"`
}

type AuthorizeConf struct {
	IdTagInfo *models.IdTagInfo `json:"idTagInfo"`
	// PriceScheme, Not yet implemented
}
