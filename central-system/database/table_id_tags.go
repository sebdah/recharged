package database

func TableIdTags() *Table {
	table := new(Table)
	table.TableName = "id_tags"
	table.ReadCapacityUnits = 1
	table.WriteCapacityUnits = 1
	table.HashKeyName = "id"
	table.HashKeyType = "S"

	return table
}
