// Code generated by ent, DO NOT EDIT.

package ent

import (
	"chat-backend/internal/ent/predicate"
	"chat-backend/internal/ent/usernamepassword"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UsernamePasswordDelete is the builder for deleting a UsernamePassword entity.
type UsernamePasswordDelete struct {
	config
	hooks    []Hook
	mutation *UsernamePasswordMutation
}

// Where appends a list predicates to the UsernamePasswordDelete builder.
func (upd *UsernamePasswordDelete) Where(ps ...predicate.UsernamePassword) *UsernamePasswordDelete {
	upd.mutation.Where(ps...)
	return upd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (upd *UsernamePasswordDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, upd.sqlExec, upd.mutation, upd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (upd *UsernamePasswordDelete) ExecX(ctx context.Context) int {
	n, err := upd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (upd *UsernamePasswordDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(usernamepassword.Table, sqlgraph.NewFieldSpec(usernamepassword.FieldID, field.TypeUUID))
	if ps := upd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, upd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	upd.mutation.done = true
	return affected, err
}

// UsernamePasswordDeleteOne is the builder for deleting a single UsernamePassword entity.
type UsernamePasswordDeleteOne struct {
	upd *UsernamePasswordDelete
}

// Where appends a list predicates to the UsernamePasswordDelete builder.
func (updo *UsernamePasswordDeleteOne) Where(ps ...predicate.UsernamePassword) *UsernamePasswordDeleteOne {
	updo.upd.mutation.Where(ps...)
	return updo
}

// Exec executes the deletion query.
func (updo *UsernamePasswordDeleteOne) Exec(ctx context.Context) error {
	n, err := updo.upd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{usernamepassword.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (updo *UsernamePasswordDeleteOne) ExecX(ctx context.Context) {
	if err := updo.Exec(ctx); err != nil {
		panic(err)
	}
}
