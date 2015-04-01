package messages

import "github.com/sebdah/recharged/central-system/types"

type AuthorizeReq struct {
	IdTag types.IdToken `json:"idTag"`
}

type AuthorizeConf struct {
	IdTagInfo *types.IdTagInfo `json:"idTagInfo"`
	// PriceScheme, Not yet implemented
}
