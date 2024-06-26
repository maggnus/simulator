// Code generated by ent, DO NOT EDIT.

package record

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the record type in the database.
	Label = "record"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHTTPMethod holds the string denoting the http_method field in the database.
	FieldHTTPMethod = "http_method"
	// FieldIPAddress holds the string denoting the ip_address field in the database.
	FieldIPAddress = "ip_address"
	// FieldInstruction holds the string denoting the instruction field in the database.
	FieldInstruction = "instruction"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"
	// Table holds the table name of the record in the database.
	Table = "instructions_log"
)

// Columns holds all SQL columns for record fields.
var Columns = []string{
	FieldID,
	FieldHTTPMethod,
	FieldIPAddress,
	FieldInstruction,
	FieldTimestamp,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultTimestamp holds the default value on creation for the "timestamp" field.
	DefaultTimestamp time.Time
)

// HTTPMethod defines the type for the "http_method" enum field.
type HTTPMethod string

// HTTPMethod values.
const (
	HTTPMethodGET  HTTPMethod = "GET"
	HTTPMethodPOST HTTPMethod = "POST"
)

func (hm HTTPMethod) String() string {
	return string(hm)
}

// HTTPMethodValidator is a validator for the "http_method" field enum values. It is called by the builders before save.
func HTTPMethodValidator(hm HTTPMethod) error {
	switch hm {
	case HTTPMethodGET, HTTPMethodPOST:
		return nil
	default:
		return fmt.Errorf("record: invalid enum value for http_method field: %q", hm)
	}
}

// OrderOption defines the ordering options for the Record queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByHTTPMethod orders the results by the http_method field.
func ByHTTPMethod(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHTTPMethod, opts...).ToFunc()
}

// ByIPAddress orders the results by the ip_address field.
func ByIPAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIPAddress, opts...).ToFunc()
}

// ByInstruction orders the results by the instruction field.
func ByInstruction(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInstruction, opts...).ToFunc()
}

// ByTimestamp orders the results by the timestamp field.
func ByTimestamp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimestamp, opts...).ToFunc()
}
