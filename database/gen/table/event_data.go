//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/sqlite"
)

var EventData = newEventDataTable("", "event_data", "")

type eventDataTable struct {
	sqlite.Table

	// Columns
	ID       sqlite.ColumnInteger
	EventID  sqlite.ColumnInteger
	Language sqlite.ColumnString
	Name     sqlite.ColumnString
	Content  sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type EventDataTable struct {
	eventDataTable

	EXCLUDED eventDataTable
}

// AS creates new EventDataTable with assigned alias
func (a EventDataTable) AS(alias string) *EventDataTable {
	return newEventDataTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new EventDataTable with assigned schema name
func (a EventDataTable) FromSchema(schemaName string) *EventDataTable {
	return newEventDataTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new EventDataTable with assigned table prefix
func (a EventDataTable) WithPrefix(prefix string) *EventDataTable {
	return newEventDataTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new EventDataTable with assigned table suffix
func (a EventDataTable) WithSuffix(suffix string) *EventDataTable {
	return newEventDataTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newEventDataTable(schemaName, tableName, alias string) *EventDataTable {
	return &EventDataTable{
		eventDataTable: newEventDataTableImpl(schemaName, tableName, alias),
		EXCLUDED:       newEventDataTableImpl("", "excluded", ""),
	}
}

func newEventDataTableImpl(schemaName, tableName, alias string) eventDataTable {
	var (
		IDColumn       = sqlite.IntegerColumn("id")
		EventIDColumn  = sqlite.IntegerColumn("event_id")
		LanguageColumn = sqlite.StringColumn("language")
		NameColumn     = sqlite.StringColumn("name")
		ContentColumn  = sqlite.StringColumn("content")
		allColumns     = sqlite.ColumnList{IDColumn, EventIDColumn, LanguageColumn, NameColumn, ContentColumn}
		mutableColumns = sqlite.ColumnList{EventIDColumn, LanguageColumn, NameColumn, ContentColumn}
	)

	return eventDataTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:       IDColumn,
		EventID:  EventIDColumn,
		Language: LanguageColumn,
		Name:     NameColumn,
		Content:  ContentColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
