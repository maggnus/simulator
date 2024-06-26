// Code generated by ent, DO NOT EDIT.

package record

import (
	"simulator/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Record {
	return predicate.Record(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Record {
	return predicate.Record(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Record {
	return predicate.Record(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Record {
	return predicate.Record(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Record {
	return predicate.Record(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Record {
	return predicate.Record(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Record {
	return predicate.Record(sql.FieldLTE(FieldID, id))
}

// IPAddress applies equality check predicate on the "ip_address" field. It's identical to IPAddressEQ.
func IPAddress(v string) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldIPAddress, v))
}

// Instruction applies equality check predicate on the "instruction" field. It's identical to InstructionEQ.
func Instruction(v string) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldInstruction, v))
}

// Timestamp applies equality check predicate on the "timestamp" field. It's identical to TimestampEQ.
func Timestamp(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldTimestamp, v))
}

// HTTPMethodEQ applies the EQ predicate on the "http_method" field.
func HTTPMethodEQ(v HTTPMethod) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldHTTPMethod, v))
}

// HTTPMethodNEQ applies the NEQ predicate on the "http_method" field.
func HTTPMethodNEQ(v HTTPMethod) predicate.Record {
	return predicate.Record(sql.FieldNEQ(FieldHTTPMethod, v))
}

// HTTPMethodIn applies the In predicate on the "http_method" field.
func HTTPMethodIn(vs ...HTTPMethod) predicate.Record {
	return predicate.Record(sql.FieldIn(FieldHTTPMethod, vs...))
}

// HTTPMethodNotIn applies the NotIn predicate on the "http_method" field.
func HTTPMethodNotIn(vs ...HTTPMethod) predicate.Record {
	return predicate.Record(sql.FieldNotIn(FieldHTTPMethod, vs...))
}

// IPAddressEQ applies the EQ predicate on the "ip_address" field.
func IPAddressEQ(v string) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldIPAddress, v))
}

// IPAddressNEQ applies the NEQ predicate on the "ip_address" field.
func IPAddressNEQ(v string) predicate.Record {
	return predicate.Record(sql.FieldNEQ(FieldIPAddress, v))
}

// IPAddressIn applies the In predicate on the "ip_address" field.
func IPAddressIn(vs ...string) predicate.Record {
	return predicate.Record(sql.FieldIn(FieldIPAddress, vs...))
}

// IPAddressNotIn applies the NotIn predicate on the "ip_address" field.
func IPAddressNotIn(vs ...string) predicate.Record {
	return predicate.Record(sql.FieldNotIn(FieldIPAddress, vs...))
}

// IPAddressGT applies the GT predicate on the "ip_address" field.
func IPAddressGT(v string) predicate.Record {
	return predicate.Record(sql.FieldGT(FieldIPAddress, v))
}

// IPAddressGTE applies the GTE predicate on the "ip_address" field.
func IPAddressGTE(v string) predicate.Record {
	return predicate.Record(sql.FieldGTE(FieldIPAddress, v))
}

// IPAddressLT applies the LT predicate on the "ip_address" field.
func IPAddressLT(v string) predicate.Record {
	return predicate.Record(sql.FieldLT(FieldIPAddress, v))
}

// IPAddressLTE applies the LTE predicate on the "ip_address" field.
func IPAddressLTE(v string) predicate.Record {
	return predicate.Record(sql.FieldLTE(FieldIPAddress, v))
}

// IPAddressContains applies the Contains predicate on the "ip_address" field.
func IPAddressContains(v string) predicate.Record {
	return predicate.Record(sql.FieldContains(FieldIPAddress, v))
}

// IPAddressHasPrefix applies the HasPrefix predicate on the "ip_address" field.
func IPAddressHasPrefix(v string) predicate.Record {
	return predicate.Record(sql.FieldHasPrefix(FieldIPAddress, v))
}

// IPAddressHasSuffix applies the HasSuffix predicate on the "ip_address" field.
func IPAddressHasSuffix(v string) predicate.Record {
	return predicate.Record(sql.FieldHasSuffix(FieldIPAddress, v))
}

// IPAddressEqualFold applies the EqualFold predicate on the "ip_address" field.
func IPAddressEqualFold(v string) predicate.Record {
	return predicate.Record(sql.FieldEqualFold(FieldIPAddress, v))
}

// IPAddressContainsFold applies the ContainsFold predicate on the "ip_address" field.
func IPAddressContainsFold(v string) predicate.Record {
	return predicate.Record(sql.FieldContainsFold(FieldIPAddress, v))
}

// InstructionEQ applies the EQ predicate on the "instruction" field.
func InstructionEQ(v string) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldInstruction, v))
}

// InstructionNEQ applies the NEQ predicate on the "instruction" field.
func InstructionNEQ(v string) predicate.Record {
	return predicate.Record(sql.FieldNEQ(FieldInstruction, v))
}

// InstructionIn applies the In predicate on the "instruction" field.
func InstructionIn(vs ...string) predicate.Record {
	return predicate.Record(sql.FieldIn(FieldInstruction, vs...))
}

// InstructionNotIn applies the NotIn predicate on the "instruction" field.
func InstructionNotIn(vs ...string) predicate.Record {
	return predicate.Record(sql.FieldNotIn(FieldInstruction, vs...))
}

// InstructionGT applies the GT predicate on the "instruction" field.
func InstructionGT(v string) predicate.Record {
	return predicate.Record(sql.FieldGT(FieldInstruction, v))
}

// InstructionGTE applies the GTE predicate on the "instruction" field.
func InstructionGTE(v string) predicate.Record {
	return predicate.Record(sql.FieldGTE(FieldInstruction, v))
}

// InstructionLT applies the LT predicate on the "instruction" field.
func InstructionLT(v string) predicate.Record {
	return predicate.Record(sql.FieldLT(FieldInstruction, v))
}

// InstructionLTE applies the LTE predicate on the "instruction" field.
func InstructionLTE(v string) predicate.Record {
	return predicate.Record(sql.FieldLTE(FieldInstruction, v))
}

// InstructionContains applies the Contains predicate on the "instruction" field.
func InstructionContains(v string) predicate.Record {
	return predicate.Record(sql.FieldContains(FieldInstruction, v))
}

// InstructionHasPrefix applies the HasPrefix predicate on the "instruction" field.
func InstructionHasPrefix(v string) predicate.Record {
	return predicate.Record(sql.FieldHasPrefix(FieldInstruction, v))
}

// InstructionHasSuffix applies the HasSuffix predicate on the "instruction" field.
func InstructionHasSuffix(v string) predicate.Record {
	return predicate.Record(sql.FieldHasSuffix(FieldInstruction, v))
}

// InstructionEqualFold applies the EqualFold predicate on the "instruction" field.
func InstructionEqualFold(v string) predicate.Record {
	return predicate.Record(sql.FieldEqualFold(FieldInstruction, v))
}

// InstructionContainsFold applies the ContainsFold predicate on the "instruction" field.
func InstructionContainsFold(v string) predicate.Record {
	return predicate.Record(sql.FieldContainsFold(FieldInstruction, v))
}

// TimestampEQ applies the EQ predicate on the "timestamp" field.
func TimestampEQ(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldTimestamp, v))
}

// TimestampNEQ applies the NEQ predicate on the "timestamp" field.
func TimestampNEQ(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldNEQ(FieldTimestamp, v))
}

// TimestampIn applies the In predicate on the "timestamp" field.
func TimestampIn(vs ...time.Time) predicate.Record {
	return predicate.Record(sql.FieldIn(FieldTimestamp, vs...))
}

// TimestampNotIn applies the NotIn predicate on the "timestamp" field.
func TimestampNotIn(vs ...time.Time) predicate.Record {
	return predicate.Record(sql.FieldNotIn(FieldTimestamp, vs...))
}

// TimestampGT applies the GT predicate on the "timestamp" field.
func TimestampGT(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldGT(FieldTimestamp, v))
}

// TimestampGTE applies the GTE predicate on the "timestamp" field.
func TimestampGTE(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldGTE(FieldTimestamp, v))
}

// TimestampLT applies the LT predicate on the "timestamp" field.
func TimestampLT(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldLT(FieldTimestamp, v))
}

// TimestampLTE applies the LTE predicate on the "timestamp" field.
func TimestampLTE(v time.Time) predicate.Record {
	return predicate.Record(sql.FieldLTE(FieldTimestamp, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Record) predicate.Record {
	return predicate.Record(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Record) predicate.Record {
	return predicate.Record(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Record) predicate.Record {
	return predicate.Record(sql.NotPredicates(p))
}
