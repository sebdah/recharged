package handlers

import (
	"fmt"
	"net/http"

	"github.com/awslabs/aws-sdk-go/gen/dynamodb"
	"github.com/sebdah/recharged/central-system/database"
)

var Db *dynamodb.DynamoDB = database.GetDb()

func AuthorizeHandler(rw http.ResponseWriter, r *http.Request) {
	//token := idtag.NewIdToken(mux.Vars(r)["id"])
	//tag := idtag.NewIdTag(token)

	fmt.Println(Db.ListTables(new(dynamodb.ListTablesInput)))
}
