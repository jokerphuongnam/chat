// Code generated by ent, DO NOT EDIT.

package ent

import (
	"chat-service/internal/ent/predicate"
	"chat-service/internal/ent/usernamepassword"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UsernamePasswordUpdate is the builder for updating UsernamePassword entities.
type UsernamePasswordUpdate struct {
	config
	hooks    []Hook
	mutation *UsernamePasswordMutation
}

// Where appends a list predicates to the UsernamePasswordUpdate builder.
func (upu *UsernamePasswordUpdate) Where(ps ...predicate.UsernamePassword) *UsernamePasswordUpdate {
	upu.mutation.Where(ps...)
	return upu
}

// SetUsername sets the "username" field.
func (upu *UsernamePasswordUpdate) SetUsername(s string) *UsernamePasswordUpdate {
	upu.mutation.SetUsername(s)
	return upu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (upu *UsernamePasswordUpdate) SetNillableUsername(s *string) *UsernamePasswordUpdate {
	if s != nil {
		upu.SetUsername(*s)
	}
	return upu
}

// SetPassword sets the "password" field.
func (upu *UsernamePasswordUpdate) SetPassword(s string) *UsernamePasswordUpdate {
	upu.mutation.SetPassword(s)
	return upu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (upu *UsernamePasswordUpdate) SetNillablePassword(s *string) *UsernamePasswordUpdate {
	if s != nil {
		upu.SetPassword(*s)
	}
	return upu
}

// Mutation returns the UsernamePasswordMutation object of the builder.
func (upu *UsernamePasswordUpdate) Mutation() *UsernamePasswordMutation {
	return upu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (upu *UsernamePasswordUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, upu.sqlSave, upu.mutation, upu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (upu *UsernamePasswordUpdate) SaveX(ctx context.Context) int {
	affected, err := upu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (upu *UsernamePasswordUpdate) Exec(ctx context.Context) error {
	_, err := upu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upu *UsernamePasswordUpdate) ExecX(ctx context.Context) {
	if err := upu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upu *UsernamePasswordUpdate) check() error {
	if v, ok := upu.mutation.Username(); ok {
		if err := usernamepassword.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "UsernamePassword.username": %w`, err)}
		}
	}
	return nil
}

func (upu *UsernamePasswordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := upu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(usernamepassword.Table, usernamepassword.Columns, sqlgraph.NewFieldSpec(usernamepassword.FieldID, field.TypeUUID))
	if ps := upu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := upu.mutation.Username(); ok {
		_spec.SetField(usernamepassword.FieldUsername, field.TypeString, value)
	}
	if value, ok := upu.mutation.Password(); ok {
		_spec.SetField(usernamepassword.FieldPassword, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, upu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usernamepassword.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	upu.mutation.done = true
	return n, nil
}

// UsernamePasswordUpdateOne is the builder for updating a single UsernamePassword entity.
type UsernamePasswordUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UsernamePasswordMutation
}

// SetUsername sets the "username" field.
func (upuo *UsernamePasswordUpdateOne) SetUsername(s string) *UsernamePasswordUpdateOne {
	upuo.mutation.SetUsername(s)
	return upuo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (upuo *UsernamePasswordUpdateOne) SetNillableUsername(s *string) *UsernamePasswordUpdateOne {
	if s != nil {
		upuo.SetUsername(*s)
	}
	return upuo
}

// SetPassword sets the "password" field.
func (upuo *UsernamePasswordUpdateOne) SetPassword(s string) *UsernamePasswordUpdateOne {
	upuo.mutation.SetPassword(s)
	return upuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (upuo *UsernamePasswordUpdateOne) SetNillablePassword(s *string) *UsernamePasswordUpdateOne {
	if s != nil {
		upuo.SetPassword(*s)
	}
	return upuo
}

// Mutation returns the UsernamePasswordMutation object of the builder.
func (upuo *UsernamePasswordUpdateOne) Mutation() *UsernamePasswordMutation {
	return upuo.mutation
}

// Where appends a list predicates to the UsernamePasswordUpdate builder.
func (upuo *UsernamePasswordUpdateOne) Where(ps ...predicate.UsernamePassword) *UsernamePasswordUpdateOne {
	upuo.mutation.Where(ps...)
	return upuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (upuo *UsernamePasswordUpdateOne) Select(field string, fields ...string) *UsernamePasswordUpdateOne {
	upuo.fields = append([]string{field}, fields...)
	return upuo
}

// Save executes the query and returns the updated UsernamePassword entity.
func (upuo *UsernamePasswordUpdateOne) Save(ctx context.Context) (*UsernamePassword, error) {
	return withHooks(ctx, upuo.sqlSave, upuo.mutation, upuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (upuo *UsernamePasswordUpdateOne) SaveX(ctx context.Context) *UsernamePassword {
	node, err := upuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (upuo *UsernamePasswordUpdateOne) Exec(ctx context.Context) error {
	_, err := upuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upuo *UsernamePasswordUpdateOne) ExecX(ctx context.Context) {
	if err := upuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upuo *UsernamePasswordUpdateOne) check() error {
	if v, ok := upuo.mutation.Username(); ok {
		if err := usernamepassword.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "UsernamePassword.username": %w`, err)}
		}
	}
	return nil
}

func (upuo *UsernamePasswordUpdateOne) sqlSave(ctx context.Context) (_node *UsernamePassword, err error) {
	if err := upuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(usernamepassword.Table, usernamepassword.Columns, sqlgraph.NewFieldSpec(usernamepassword.FieldID, field.TypeUUID))
	id, ok := upuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UsernamePassword.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := upuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usernamepassword.FieldID)
		for _, f := range fields {
			if !usernamepassword.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usernamepassword.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := upuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := upuo.mutation.Username(); ok {
		_spec.SetField(usernamepassword.FieldUsername, field.TypeString, value)
	}
	if value, ok := upuo.mutation.Password(); ok {
		_spec.SetField(usernamepassword.FieldPassword, field.TypeString, value)
	}
	_node = &UsernamePassword{config: upuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, upuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usernamepassword.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	upuo.mutation.done = true
	return _node, nil
}
