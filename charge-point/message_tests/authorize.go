package message_tests

import (
	"github.com/sebdah/recharged/shared/rpc"
	"github.com/sebdah/recharged/shared/transports"
)

type AuthorizeSuite struct {
	wsClient *transports.WsClient
}

// Constructor
func NewAuthorizeSuite(wsClient *transports.WsClient) (suite AuthorizeSuite) {
	suite = new(AuthorizeSuite)
	suite.wsClient = wsClient

	return
}

// Run all tests
func (this *AuthorizeSuite) RunAll() {
	this.AuthorizeBasic()
}

// Basic test of Authorize
func (this *AuthorizeSuite) AuthorizeBasic() {
	// Create call
	call := rpc.NewCall()
	call.UniqueId = "test"
	call.Action = "Authorize"
	call.Payload = `{ "idTag": { "id": "1234" } }`
}
