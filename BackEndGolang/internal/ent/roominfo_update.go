// Code generated by ent, DO NOT EDIT.

package ent

import (
	"chat-backend/internal/ent/predicate"
	"chat-backend/internal/ent/room"
	"chat-backend/internal/ent/roominfo"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// RoomInfoUpdate is the builder for updating RoomInfo entities.
type RoomInfoUpdate struct {
	config
	hooks    []Hook
	mutation *RoomInfoMutation
}

// Where appends a list predicates to the RoomInfoUpdate builder.
func (riu *RoomInfoUpdate) Where(ps ...predicate.RoomInfo) *RoomInfoUpdate {
	riu.mutation.Where(ps...)
	return riu
}

// SetRoomImageURL sets the "room_image_url" field.
func (riu *RoomInfoUpdate) SetRoomImageURL(s string) *RoomInfoUpdate {
	riu.mutation.SetRoomImageURL(s)
	return riu
}

// SetNillableRoomImageURL sets the "room_image_url" field if the given value is not nil.
func (riu *RoomInfoUpdate) SetNillableRoomImageURL(s *string) *RoomInfoUpdate {
	if s != nil {
		riu.SetRoomImageURL(*s)
	}
	return riu
}

// ClearRoomImageURL clears the value of the "room_image_url" field.
func (riu *RoomInfoUpdate) ClearRoomImageURL() *RoomInfoUpdate {
	riu.mutation.ClearRoomImageURL()
	return riu
}

// SetName sets the "name" field.
func (riu *RoomInfoUpdate) SetName(s string) *RoomInfoUpdate {
	riu.mutation.SetName(s)
	return riu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (riu *RoomInfoUpdate) SetNillableName(s *string) *RoomInfoUpdate {
	if s != nil {
		riu.SetName(*s)
	}
	return riu
}

// ClearName clears the value of the "name" field.
func (riu *RoomInfoUpdate) ClearName() *RoomInfoUpdate {
	riu.mutation.ClearName()
	return riu
}

// SetRoomsID sets the "rooms" edge to the Room entity by ID.
func (riu *RoomInfoUpdate) SetRoomsID(id uuid.UUID) *RoomInfoUpdate {
	riu.mutation.SetRoomsID(id)
	return riu
}

// SetNillableRoomsID sets the "rooms" edge to the Room entity by ID if the given value is not nil.
func (riu *RoomInfoUpdate) SetNillableRoomsID(id *uuid.UUID) *RoomInfoUpdate {
	if id != nil {
		riu = riu.SetRoomsID(*id)
	}
	return riu
}

// SetRooms sets the "rooms" edge to the Room entity.
func (riu *RoomInfoUpdate) SetRooms(r *Room) *RoomInfoUpdate {
	return riu.SetRoomsID(r.ID)
}

// Mutation returns the RoomInfoMutation object of the builder.
func (riu *RoomInfoUpdate) Mutation() *RoomInfoMutation {
	return riu.mutation
}

// ClearRooms clears the "rooms" edge to the Room entity.
func (riu *RoomInfoUpdate) ClearRooms() *RoomInfoUpdate {
	riu.mutation.ClearRooms()
	return riu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (riu *RoomInfoUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, riu.sqlSave, riu.mutation, riu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (riu *RoomInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := riu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (riu *RoomInfoUpdate) Exec(ctx context.Context) error {
	_, err := riu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (riu *RoomInfoUpdate) ExecX(ctx context.Context) {
	if err := riu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (riu *RoomInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(roominfo.Table, roominfo.Columns, sqlgraph.NewFieldSpec(roominfo.FieldID, field.TypeUUID))
	if ps := riu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := riu.mutation.RoomImageURL(); ok {
		_spec.SetField(roominfo.FieldRoomImageURL, field.TypeString, value)
	}
	if riu.mutation.RoomImageURLCleared() {
		_spec.ClearField(roominfo.FieldRoomImageURL, field.TypeString)
	}
	if value, ok := riu.mutation.Name(); ok {
		_spec.SetField(roominfo.FieldName, field.TypeString, value)
	}
	if riu.mutation.NameCleared() {
		_spec.ClearField(roominfo.FieldName, field.TypeString)
	}
	if riu.mutation.RoomsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   roominfo.RoomsTable,
			Columns: []string{roominfo.RoomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := riu.mutation.RoomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   roominfo.RoomsTable,
			Columns: []string{roominfo.RoomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, riu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roominfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	riu.mutation.done = true
	return n, nil
}

// RoomInfoUpdateOne is the builder for updating a single RoomInfo entity.
type RoomInfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoomInfoMutation
}

// SetRoomImageURL sets the "room_image_url" field.
func (riuo *RoomInfoUpdateOne) SetRoomImageURL(s string) *RoomInfoUpdateOne {
	riuo.mutation.SetRoomImageURL(s)
	return riuo
}

// SetNillableRoomImageURL sets the "room_image_url" field if the given value is not nil.
func (riuo *RoomInfoUpdateOne) SetNillableRoomImageURL(s *string) *RoomInfoUpdateOne {
	if s != nil {
		riuo.SetRoomImageURL(*s)
	}
	return riuo
}

// ClearRoomImageURL clears the value of the "room_image_url" field.
func (riuo *RoomInfoUpdateOne) ClearRoomImageURL() *RoomInfoUpdateOne {
	riuo.mutation.ClearRoomImageURL()
	return riuo
}

// SetName sets the "name" field.
func (riuo *RoomInfoUpdateOne) SetName(s string) *RoomInfoUpdateOne {
	riuo.mutation.SetName(s)
	return riuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (riuo *RoomInfoUpdateOne) SetNillableName(s *string) *RoomInfoUpdateOne {
	if s != nil {
		riuo.SetName(*s)
	}
	return riuo
}

// ClearName clears the value of the "name" field.
func (riuo *RoomInfoUpdateOne) ClearName() *RoomInfoUpdateOne {
	riuo.mutation.ClearName()
	return riuo
}

// SetRoomsID sets the "rooms" edge to the Room entity by ID.
func (riuo *RoomInfoUpdateOne) SetRoomsID(id uuid.UUID) *RoomInfoUpdateOne {
	riuo.mutation.SetRoomsID(id)
	return riuo
}

// SetNillableRoomsID sets the "rooms" edge to the Room entity by ID if the given value is not nil.
func (riuo *RoomInfoUpdateOne) SetNillableRoomsID(id *uuid.UUID) *RoomInfoUpdateOne {
	if id != nil {
		riuo = riuo.SetRoomsID(*id)
	}
	return riuo
}

// SetRooms sets the "rooms" edge to the Room entity.
func (riuo *RoomInfoUpdateOne) SetRooms(r *Room) *RoomInfoUpdateOne {
	return riuo.SetRoomsID(r.ID)
}

// Mutation returns the RoomInfoMutation object of the builder.
func (riuo *RoomInfoUpdateOne) Mutation() *RoomInfoMutation {
	return riuo.mutation
}

// ClearRooms clears the "rooms" edge to the Room entity.
func (riuo *RoomInfoUpdateOne) ClearRooms() *RoomInfoUpdateOne {
	riuo.mutation.ClearRooms()
	return riuo
}

// Where appends a list predicates to the RoomInfoUpdate builder.
func (riuo *RoomInfoUpdateOne) Where(ps ...predicate.RoomInfo) *RoomInfoUpdateOne {
	riuo.mutation.Where(ps...)
	return riuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (riuo *RoomInfoUpdateOne) Select(field string, fields ...string) *RoomInfoUpdateOne {
	riuo.fields = append([]string{field}, fields...)
	return riuo
}

// Save executes the query and returns the updated RoomInfo entity.
func (riuo *RoomInfoUpdateOne) Save(ctx context.Context) (*RoomInfo, error) {
	return withHooks(ctx, riuo.sqlSave, riuo.mutation, riuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (riuo *RoomInfoUpdateOne) SaveX(ctx context.Context) *RoomInfo {
	node, err := riuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (riuo *RoomInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := riuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (riuo *RoomInfoUpdateOne) ExecX(ctx context.Context) {
	if err := riuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (riuo *RoomInfoUpdateOne) sqlSave(ctx context.Context) (_node *RoomInfo, err error) {
	_spec := sqlgraph.NewUpdateSpec(roominfo.Table, roominfo.Columns, sqlgraph.NewFieldSpec(roominfo.FieldID, field.TypeUUID))
	id, ok := riuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RoomInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := riuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, roominfo.FieldID)
		for _, f := range fields {
			if !roominfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != roominfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := riuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := riuo.mutation.RoomImageURL(); ok {
		_spec.SetField(roominfo.FieldRoomImageURL, field.TypeString, value)
	}
	if riuo.mutation.RoomImageURLCleared() {
		_spec.ClearField(roominfo.FieldRoomImageURL, field.TypeString)
	}
	if value, ok := riuo.mutation.Name(); ok {
		_spec.SetField(roominfo.FieldName, field.TypeString, value)
	}
	if riuo.mutation.NameCleared() {
		_spec.ClearField(roominfo.FieldName, field.TypeString)
	}
	if riuo.mutation.RoomsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   roominfo.RoomsTable,
			Columns: []string{roominfo.RoomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := riuo.mutation.RoomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   roominfo.RoomsTable,
			Columns: []string{roominfo.RoomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &RoomInfo{config: riuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, riuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roominfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	riuo.mutation.done = true
	return _node, nil
}
