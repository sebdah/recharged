package database

func TableIdTagWhitelist() *Table {
	table := new(Table)
	table.TableName = "idtag_whitelist"
	table.ReadCapacityUnits = 1
	table.WriteCapacityUnits = 1
	table.HashKeyName = "idtag"
	table.HashKeyType = "S"

	return table
}
