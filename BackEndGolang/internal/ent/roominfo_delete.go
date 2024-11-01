// Code generated by ent, DO NOT EDIT.

package ent

import (
	"chat-backend/internal/ent/predicate"
	"chat-backend/internal/ent/roominfo"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoomInfoDelete is the builder for deleting a RoomInfo entity.
type RoomInfoDelete struct {
	config
	hooks    []Hook
	mutation *RoomInfoMutation
}

// Where appends a list predicates to the RoomInfoDelete builder.
func (rid *RoomInfoDelete) Where(ps ...predicate.RoomInfo) *RoomInfoDelete {
	rid.mutation.Where(ps...)
	return rid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rid *RoomInfoDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rid.sqlExec, rid.mutation, rid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rid *RoomInfoDelete) ExecX(ctx context.Context) int {
	n, err := rid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rid *RoomInfoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(roominfo.Table, sqlgraph.NewFieldSpec(roominfo.FieldID, field.TypeUUID))
	if ps := rid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rid.mutation.done = true
	return affected, err
}

// RoomInfoDeleteOne is the builder for deleting a single RoomInfo entity.
type RoomInfoDeleteOne struct {
	rid *RoomInfoDelete
}

// Where appends a list predicates to the RoomInfoDelete builder.
func (rido *RoomInfoDeleteOne) Where(ps ...predicate.RoomInfo) *RoomInfoDeleteOne {
	rido.rid.mutation.Where(ps...)
	return rido
}

// Exec executes the deletion query.
func (rido *RoomInfoDeleteOne) Exec(ctx context.Context) error {
	n, err := rido.rid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{roominfo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rido *RoomInfoDeleteOne) ExecX(ctx context.Context) {
	if err := rido.Exec(ctx); err != nil {
		panic(err)
	}
}