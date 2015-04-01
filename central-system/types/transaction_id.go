package types

type TransactionId struct {
	transactionId       int    `type="int"`
	transactionIdString string `type="string" required="true" max_length="50"`
}
