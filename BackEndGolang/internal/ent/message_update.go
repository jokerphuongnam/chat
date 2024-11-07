// Code generated by ent, DO NOT EDIT.

package ent

import (
	"chat-backend/internal/ent/message"
	"chat-backend/internal/ent/predicate"
	"chat-backend/internal/ent/room"
	"chat-backend/internal/ent/user"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// MessageUpdate is the builder for updating Message entities.
type MessageUpdate struct {
	config
	hooks    []Hook
	mutation *MessageMutation
}

// Where appends a list predicates to the MessageUpdate builder.
func (mu *MessageUpdate) Where(ps ...predicate.Message) *MessageUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetDateSend sets the "date_send" field.
func (mu *MessageUpdate) SetDateSend(u uint64) *MessageUpdate {
	mu.mutation.ResetDateSend()
	mu.mutation.SetDateSend(u)
	return mu
}

// SetNillableDateSend sets the "date_send" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableDateSend(u *uint64) *MessageUpdate {
	if u != nil {
		mu.SetDateSend(*u)
	}
	return mu
}

// AddDateSend adds u to the "date_send" field.
func (mu *MessageUpdate) AddDateSend(u int64) *MessageUpdate {
	mu.mutation.AddDateSend(u)
	return mu
}

// SetTypeMessage sets the "type_message" field.
func (mu *MessageUpdate) SetTypeMessage(mm message.TypeMessage) *MessageUpdate {
	mu.mutation.SetTypeMessage(mm)
	return mu
}

// SetNillableTypeMessage sets the "type_message" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableTypeMessage(mm *message.TypeMessage) *MessageUpdate {
	if mm != nil {
		mu.SetTypeMessage(*mm)
	}
	return mu
}

// SetContent sets the "content" field.
func (mu *MessageUpdate) SetContent(s string) *MessageUpdate {
	mu.mutation.SetContent(s)
	return mu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableContent(s *string) *MessageUpdate {
	if s != nil {
		mu.SetContent(*s)
	}
	return mu
}

// SetIDRoom sets the "id_room" field.
func (mu *MessageUpdate) SetIDRoom(u uuid.UUID) *MessageUpdate {
	mu.mutation.SetIDRoom(u)
	return mu
}

// SetNillableIDRoom sets the "id_room" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableIDRoom(u *uuid.UUID) *MessageUpdate {
	if u != nil {
		mu.SetIDRoom(*u)
	}
	return mu
}

// SetIDUserSend sets the "id_user_send" field.
func (mu *MessageUpdate) SetIDUserSend(u uuid.UUID) *MessageUpdate {
	mu.mutation.SetIDUserSend(u)
	return mu
}

// SetNillableIDUserSend sets the "id_user_send" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableIDUserSend(u *uuid.UUID) *MessageUpdate {
	if u != nil {
		mu.SetIDUserSend(*u)
	}
	return mu
}

// SetRoomsID sets the "rooms" edge to the Room entity by ID.
func (mu *MessageUpdate) SetRoomsID(id uuid.UUID) *MessageUpdate {
	mu.mutation.SetRoomsID(id)
	return mu
}

// SetRooms sets the "rooms" edge to the Room entity.
func (mu *MessageUpdate) SetRooms(r *Room) *MessageUpdate {
	return mu.SetRoomsID(r.ID)
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (mu *MessageUpdate) SetUsersID(id uuid.UUID) *MessageUpdate {
	mu.mutation.SetUsersID(id)
	return mu
}

// SetUsers sets the "users" edge to the User entity.
func (mu *MessageUpdate) SetUsers(u *User) *MessageUpdate {
	return mu.SetUsersID(u.ID)
}

// Mutation returns the MessageMutation object of the builder.
func (mu *MessageUpdate) Mutation() *MessageMutation {
	return mu.mutation
}

// ClearRooms clears the "rooms" edge to the Room entity.
func (mu *MessageUpdate) ClearRooms() *MessageUpdate {
	mu.mutation.ClearRooms()
	return mu
}

// ClearUsers clears the "users" edge to the User entity.
func (mu *MessageUpdate) ClearUsers() *MessageUpdate {
	mu.mutation.ClearUsers()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MessageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MessageUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MessageUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MessageUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MessageUpdate) check() error {
	if v, ok := mu.mutation.TypeMessage(); ok {
		if err := message.TypeMessageValidator(v); err != nil {
			return &ValidationError{Name: "type_message", err: fmt.Errorf(`ent: validator failed for field "Message.type_message": %w`, err)}
		}
	}
	if mu.mutation.RoomsCleared() && len(mu.mutation.RoomsIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Message.rooms"`)
	}
	if mu.mutation.UsersCleared() && len(mu.mutation.UsersIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Message.users"`)
	}
	return nil
}

func (mu *MessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(message.Table, message.Columns, sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.DateSend(); ok {
		_spec.SetField(message.FieldDateSend, field.TypeUint64, value)
	}
	if value, ok := mu.mutation.AddedDateSend(); ok {
		_spec.AddField(message.FieldDateSend, field.TypeUint64, value)
	}
	if value, ok := mu.mutation.TypeMessage(); ok {
		_spec.SetField(message.FieldTypeMessage, field.TypeEnum, value)
	}
	if value, ok := mu.mutation.Content(); ok {
		_spec.SetField(message.FieldContent, field.TypeString, value)
	}
	if mu.mutation.RoomsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.RoomsTable,
			Columns: []string{message.RoomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RoomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.RoomsTable,
			Columns: []string{message.RoomsColumn},
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
	if mu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.UsersTable,
			Columns: []string{message.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.UsersTable,
			Columns: []string{message.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MessageUpdateOne is the builder for updating a single Message entity.
type MessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageMutation
}

// SetDateSend sets the "date_send" field.
func (muo *MessageUpdateOne) SetDateSend(u uint64) *MessageUpdateOne {
	muo.mutation.ResetDateSend()
	muo.mutation.SetDateSend(u)
	return muo
}

// SetNillableDateSend sets the "date_send" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableDateSend(u *uint64) *MessageUpdateOne {
	if u != nil {
		muo.SetDateSend(*u)
	}
	return muo
}

// AddDateSend adds u to the "date_send" field.
func (muo *MessageUpdateOne) AddDateSend(u int64) *MessageUpdateOne {
	muo.mutation.AddDateSend(u)
	return muo
}

// SetTypeMessage sets the "type_message" field.
func (muo *MessageUpdateOne) SetTypeMessage(mm message.TypeMessage) *MessageUpdateOne {
	muo.mutation.SetTypeMessage(mm)
	return muo
}

// SetNillableTypeMessage sets the "type_message" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableTypeMessage(mm *message.TypeMessage) *MessageUpdateOne {
	if mm != nil {
		muo.SetTypeMessage(*mm)
	}
	return muo
}

// SetContent sets the "content" field.
func (muo *MessageUpdateOne) SetContent(s string) *MessageUpdateOne {
	muo.mutation.SetContent(s)
	return muo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableContent(s *string) *MessageUpdateOne {
	if s != nil {
		muo.SetContent(*s)
	}
	return muo
}

// SetIDRoom sets the "id_room" field.
func (muo *MessageUpdateOne) SetIDRoom(u uuid.UUID) *MessageUpdateOne {
	muo.mutation.SetIDRoom(u)
	return muo
}

// SetNillableIDRoom sets the "id_room" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableIDRoom(u *uuid.UUID) *MessageUpdateOne {
	if u != nil {
		muo.SetIDRoom(*u)
	}
	return muo
}

// SetIDUserSend sets the "id_user_send" field.
func (muo *MessageUpdateOne) SetIDUserSend(u uuid.UUID) *MessageUpdateOne {
	muo.mutation.SetIDUserSend(u)
	return muo
}

// SetNillableIDUserSend sets the "id_user_send" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableIDUserSend(u *uuid.UUID) *MessageUpdateOne {
	if u != nil {
		muo.SetIDUserSend(*u)
	}
	return muo
}

// SetRoomsID sets the "rooms" edge to the Room entity by ID.
func (muo *MessageUpdateOne) SetRoomsID(id uuid.UUID) *MessageUpdateOne {
	muo.mutation.SetRoomsID(id)
	return muo
}

// SetRooms sets the "rooms" edge to the Room entity.
func (muo *MessageUpdateOne) SetRooms(r *Room) *MessageUpdateOne {
	return muo.SetRoomsID(r.ID)
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (muo *MessageUpdateOne) SetUsersID(id uuid.UUID) *MessageUpdateOne {
	muo.mutation.SetUsersID(id)
	return muo
}

// SetUsers sets the "users" edge to the User entity.
func (muo *MessageUpdateOne) SetUsers(u *User) *MessageUpdateOne {
	return muo.SetUsersID(u.ID)
}

// Mutation returns the MessageMutation object of the builder.
func (muo *MessageUpdateOne) Mutation() *MessageMutation {
	return muo.mutation
}

// ClearRooms clears the "rooms" edge to the Room entity.
func (muo *MessageUpdateOne) ClearRooms() *MessageUpdateOne {
	muo.mutation.ClearRooms()
	return muo
}

// ClearUsers clears the "users" edge to the User entity.
func (muo *MessageUpdateOne) ClearUsers() *MessageUpdateOne {
	muo.mutation.ClearUsers()
	return muo
}

// Where appends a list predicates to the MessageUpdate builder.
func (muo *MessageUpdateOne) Where(ps ...predicate.Message) *MessageUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MessageUpdateOne) Select(field string, fields ...string) *MessageUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Message entity.
func (muo *MessageUpdateOne) Save(ctx context.Context) (*Message, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MessageUpdateOne) SaveX(ctx context.Context) *Message {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MessageUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MessageUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MessageUpdateOne) check() error {
	if v, ok := muo.mutation.TypeMessage(); ok {
		if err := message.TypeMessageValidator(v); err != nil {
			return &ValidationError{Name: "type_message", err: fmt.Errorf(`ent: validator failed for field "Message.type_message": %w`, err)}
		}
	}
	if muo.mutation.RoomsCleared() && len(muo.mutation.RoomsIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Message.rooms"`)
	}
	if muo.mutation.UsersCleared() && len(muo.mutation.UsersIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Message.users"`)
	}
	return nil
}

func (muo *MessageUpdateOne) sqlSave(ctx context.Context) (_node *Message, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(message.Table, message.Columns, sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Message.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, message.FieldID)
		for _, f := range fields {
			if !message.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != message.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.DateSend(); ok {
		_spec.SetField(message.FieldDateSend, field.TypeUint64, value)
	}
	if value, ok := muo.mutation.AddedDateSend(); ok {
		_spec.AddField(message.FieldDateSend, field.TypeUint64, value)
	}
	if value, ok := muo.mutation.TypeMessage(); ok {
		_spec.SetField(message.FieldTypeMessage, field.TypeEnum, value)
	}
	if value, ok := muo.mutation.Content(); ok {
		_spec.SetField(message.FieldContent, field.TypeString, value)
	}
	if muo.mutation.RoomsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.RoomsTable,
			Columns: []string{message.RoomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RoomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.RoomsTable,
			Columns: []string{message.RoomsColumn},
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
	if muo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.UsersTable,
			Columns: []string{message.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.UsersTable,
			Columns: []string{message.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Message{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
