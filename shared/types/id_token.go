package types

type IdToken struct {
	Id     string `json:"id" type:"string" required:"true" max_length:"50"`
	IdType string `json:"idType" type:"idType" required:"false"`
}
