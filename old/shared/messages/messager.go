package messages

type Messager interface {
	GetMessageType() string
	ParseJson(string) error
	String() string
}
