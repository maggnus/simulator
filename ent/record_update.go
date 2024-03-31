// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"simulator/ent/predicate"
	"simulator/ent/record"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RecordUpdate is the builder for updating Record entities.
type RecordUpdate struct {
	config
	hooks     []Hook
	mutation  *RecordMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the RecordUpdate builder.
func (ru *RecordUpdate) Where(ps ...predicate.Record) *RecordUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetHTTPMethod sets the "http_method" field.
func (ru *RecordUpdate) SetHTTPMethod(rm record.HTTPMethod) *RecordUpdate {
	ru.mutation.SetHTTPMethod(rm)
	return ru
}

// SetNillableHTTPMethod sets the "http_method" field if the given value is not nil.
func (ru *RecordUpdate) SetNillableHTTPMethod(rm *record.HTTPMethod) *RecordUpdate {
	if rm != nil {
		ru.SetHTTPMethod(*rm)
	}
	return ru
}

// SetIPAddress sets the "ip_address" field.
func (ru *RecordUpdate) SetIPAddress(s string) *RecordUpdate {
	ru.mutation.SetIPAddress(s)
	return ru
}

// SetNillableIPAddress sets the "ip_address" field if the given value is not nil.
func (ru *RecordUpdate) SetNillableIPAddress(s *string) *RecordUpdate {
	if s != nil {
		ru.SetIPAddress(*s)
	}
	return ru
}

// SetInstruction sets the "instruction" field.
func (ru *RecordUpdate) SetInstruction(s string) *RecordUpdate {
	ru.mutation.SetInstruction(s)
	return ru
}

// SetNillableInstruction sets the "instruction" field if the given value is not nil.
func (ru *RecordUpdate) SetNillableInstruction(s *string) *RecordUpdate {
	if s != nil {
		ru.SetInstruction(*s)
	}
	return ru
}

// SetTimestamp sets the "timestamp" field.
func (ru *RecordUpdate) SetTimestamp(t time.Time) *RecordUpdate {
	ru.mutation.SetTimestamp(t)
	return ru
}

// SetNillableTimestamp sets the "timestamp" field if the given value is not nil.
func (ru *RecordUpdate) SetNillableTimestamp(t *time.Time) *RecordUpdate {
	if t != nil {
		ru.SetTimestamp(*t)
	}
	return ru
}

// Mutation returns the RecordMutation object of the builder.
func (ru *RecordUpdate) Mutation() *RecordMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RecordUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RecordUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RecordUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RecordUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RecordUpdate) check() error {
	if v, ok := ru.mutation.HTTPMethod(); ok {
		if err := record.HTTPMethodValidator(v); err != nil {
			return &ValidationError{Name: "http_method", err: fmt.Errorf(`ent: validator failed for field "Record.http_method": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ru *RecordUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RecordUpdate {
	ru.modifiers = append(ru.modifiers, modifiers...)
	return ru
}

func (ru *RecordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(record.Table, record.Columns, sqlgraph.NewFieldSpec(record.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.HTTPMethod(); ok {
		_spec.SetField(record.FieldHTTPMethod, field.TypeEnum, value)
	}
	if value, ok := ru.mutation.IPAddress(); ok {
		_spec.SetField(record.FieldIPAddress, field.TypeString, value)
	}
	if value, ok := ru.mutation.Instruction(); ok {
		_spec.SetField(record.FieldInstruction, field.TypeString, value)
	}
	if value, ok := ru.mutation.Timestamp(); ok {
		_spec.SetField(record.FieldTimestamp, field.TypeTime, value)
	}
	_spec.AddModifiers(ru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{record.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RecordUpdateOne is the builder for updating a single Record entity.
type RecordUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *RecordMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetHTTPMethod sets the "http_method" field.
func (ruo *RecordUpdateOne) SetHTTPMethod(rm record.HTTPMethod) *RecordUpdateOne {
	ruo.mutation.SetHTTPMethod(rm)
	return ruo
}

// SetNillableHTTPMethod sets the "http_method" field if the given value is not nil.
func (ruo *RecordUpdateOne) SetNillableHTTPMethod(rm *record.HTTPMethod) *RecordUpdateOne {
	if rm != nil {
		ruo.SetHTTPMethod(*rm)
	}
	return ruo
}

// SetIPAddress sets the "ip_address" field.
func (ruo *RecordUpdateOne) SetIPAddress(s string) *RecordUpdateOne {
	ruo.mutation.SetIPAddress(s)
	return ruo
}

// SetNillableIPAddress sets the "ip_address" field if the given value is not nil.
func (ruo *RecordUpdateOne) SetNillableIPAddress(s *string) *RecordUpdateOne {
	if s != nil {
		ruo.SetIPAddress(*s)
	}
	return ruo
}

// SetInstruction sets the "instruction" field.
func (ruo *RecordUpdateOne) SetInstruction(s string) *RecordUpdateOne {
	ruo.mutation.SetInstruction(s)
	return ruo
}

// SetNillableInstruction sets the "instruction" field if the given value is not nil.
func (ruo *RecordUpdateOne) SetNillableInstruction(s *string) *RecordUpdateOne {
	if s != nil {
		ruo.SetInstruction(*s)
	}
	return ruo
}

// SetTimestamp sets the "timestamp" field.
func (ruo *RecordUpdateOne) SetTimestamp(t time.Time) *RecordUpdateOne {
	ruo.mutation.SetTimestamp(t)
	return ruo
}

// SetNillableTimestamp sets the "timestamp" field if the given value is not nil.
func (ruo *RecordUpdateOne) SetNillableTimestamp(t *time.Time) *RecordUpdateOne {
	if t != nil {
		ruo.SetTimestamp(*t)
	}
	return ruo
}

// Mutation returns the RecordMutation object of the builder.
func (ruo *RecordUpdateOne) Mutation() *RecordMutation {
	return ruo.mutation
}

// Where appends a list predicates to the RecordUpdate builder.
func (ruo *RecordUpdateOne) Where(ps ...predicate.Record) *RecordUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RecordUpdateOne) Select(field string, fields ...string) *RecordUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Record entity.
func (ruo *RecordUpdateOne) Save(ctx context.Context) (*Record, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RecordUpdateOne) SaveX(ctx context.Context) *Record {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RecordUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RecordUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RecordUpdateOne) check() error {
	if v, ok := ruo.mutation.HTTPMethod(); ok {
		if err := record.HTTPMethodValidator(v); err != nil {
			return &ValidationError{Name: "http_method", err: fmt.Errorf(`ent: validator failed for field "Record.http_method": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ruo *RecordUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RecordUpdateOne {
	ruo.modifiers = append(ruo.modifiers, modifiers...)
	return ruo
}

func (ruo *RecordUpdateOne) sqlSave(ctx context.Context) (_node *Record, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(record.Table, record.Columns, sqlgraph.NewFieldSpec(record.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Record.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, record.FieldID)
		for _, f := range fields {
			if !record.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != record.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.HTTPMethod(); ok {
		_spec.SetField(record.FieldHTTPMethod, field.TypeEnum, value)
	}
	if value, ok := ruo.mutation.IPAddress(); ok {
		_spec.SetField(record.FieldIPAddress, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Instruction(); ok {
		_spec.SetField(record.FieldInstruction, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Timestamp(); ok {
		_spec.SetField(record.FieldTimestamp, field.TypeTime, value)
	}
	_spec.AddModifiers(ruo.modifiers...)
	_node = &Record{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{record.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}