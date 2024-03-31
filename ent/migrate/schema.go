// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// InstructionsLogColumns holds the columns for the "instructions_log" table.
	InstructionsLogColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "http_method", Type: field.TypeEnum, Enums: []string{"GET", "POST"}},
		{Name: "ip_address", Type: field.TypeString},
		{Name: "instruction", Type: field.TypeString},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// InstructionsLogTable holds the schema information for the "instructions_log" table.
	InstructionsLogTable = &schema.Table{
		Name:       "instructions_log",
		Columns:    InstructionsLogColumns,
		PrimaryKey: []*schema.Column{InstructionsLogColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		InstructionsLogTable,
	}
)

func init() {
	InstructionsLogTable.Annotation = &entsql.Annotation{
		Table: "instructions_log",
	}
}
