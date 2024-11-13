// Code generated by ent, DO NOT EDIT.

package ent

import (
	"chat-service/internal/ent/member"
	"chat-service/internal/ent/predicate"
	"chat-service/internal/ent/room"
	"chat-service/internal/ent/user"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// MemberUpdate is the builder for updating Member entities.
type MemberUpdate struct {
	config
	hooks    []Hook
	mutation *MemberMutation
}

// Where appends a list predicates to the MemberUpdate builder.
func (mu *MemberUpdate) Where(ps ...predicate.Member) *MemberUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetUserID sets the "user_id" field.
func (mu *MemberUpdate) SetUserID(u uuid.UUID) *MemberUpdate {
	mu.mutation.SetUserID(u)
	return mu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableUserID(u *uuid.UUID) *MemberUpdate {
	if u != nil {
		mu.SetUserID(*u)
	}
	return mu
}

// SetRoomID sets the "room_id" field.
func (mu *MemberUpdate) SetRoomID(u uuid.UUID) *MemberUpdate {
	mu.mutation.SetRoomID(u)
	return mu
}

// SetNillableRoomID sets the "room_id" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableRoomID(u *uuid.UUID) *MemberUpdate {
	if u != nil {
		mu.SetRoomID(*u)
	}
	return mu
}

// SetRole sets the "role" field.
func (mu *MemberUpdate) SetRole(m member.Role) *MemberUpdate {
	mu.mutation.SetRole(m)
	return mu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableRole(m *member.Role) *MemberUpdate {
	if m != nil {
		mu.SetRole(*m)
	}
	return mu
}

// SetNickName sets the "nick_name" field.
func (mu *MemberUpdate) SetNickName(s string) *MemberUpdate {
	mu.mutation.SetNickName(s)
	return mu
}

// SetNillableNickName sets the "nick_name" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableNickName(s *string) *MemberUpdate {
	if s != nil {
		mu.SetNickName(*s)
	}
	return mu
}

// ClearNickName clears the value of the "nick_name" field.
func (mu *MemberUpdate) ClearNickName() *MemberUpdate {
	mu.mutation.ClearNickName()
	return mu
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (mu *MemberUpdate) SetUsersID(id uuid.UUID) *MemberUpdate {
	mu.mutation.SetUsersID(id)
	return mu
}

// SetUsers sets the "users" edge to the User entity.
func (mu *MemberUpdate) SetUsers(u *User) *MemberUpdate {
	return mu.SetUsersID(u.ID)
}

// SetRoomsID sets the "rooms" edge to the Room entity by ID.
func (mu *MemberUpdate) SetRoomsID(id uuid.UUID) *MemberUpdate {
	mu.mutation.SetRoomsID(id)
	return mu
}

// SetRooms sets the "rooms" edge to the Room entity.
func (mu *MemberUpdate) SetRooms(r *Room) *MemberUpdate {
	return mu.SetRoomsID(r.ID)
}

// Mutation returns the MemberMutation object of the builder.
func (mu *MemberUpdate) Mutation() *MemberMutation {
	return mu.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (mu *MemberUpdate) ClearUsers() *MemberUpdate {
	mu.mutation.ClearUsers()
	return mu
}

// ClearRooms clears the "rooms" edge to the Room entity.
func (mu *MemberUpdate) ClearRooms() *MemberUpdate {
	mu.mutation.ClearRooms()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MemberUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MemberUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MemberUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MemberUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MemberUpdate) check() error {
	if v, ok := mu.mutation.Role(); ok {
		if err := member.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "Member.role": %w`, err)}
		}
	}
	if v, ok := mu.mutation.NickName(); ok {
		if err := member.NickNameValidator(v); err != nil {
			return &ValidationError{Name: "nick_name", err: fmt.Errorf(`ent: validator failed for field "Member.nick_name": %w`, err)}
		}
	}
	if mu.mutation.UsersCleared() && len(mu.mutation.UsersIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Member.users"`)
	}
	if mu.mutation.RoomsCleared() && len(mu.mutation.RoomsIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Member.rooms"`)
	}
	return nil
}

func (mu *MemberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(member.Table, member.Columns, sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Role(); ok {
		_spec.SetField(member.FieldRole, field.TypeEnum, value)
	}
	if value, ok := mu.mutation.NickName(); ok {
		_spec.SetField(member.FieldNickName, field.TypeString, value)
	}
	if mu.mutation.NickNameCleared() {
		_spec.ClearField(member.FieldNickName, field.TypeString)
	}
	if mu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   member.UsersTable,
			Columns: []string{member.UsersColumn},
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
			Table:   member.UsersTable,
			Columns: []string{member.UsersColumn},
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
	if mu.mutation.RoomsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   member.RoomsTable,
			Columns: []string{member.RoomsColumn},
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
			Table:   member.RoomsTable,
			Columns: []string{member.RoomsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MemberUpdateOne is the builder for updating a single Member entity.
type MemberUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MemberMutation
}

// SetUserID sets the "user_id" field.
func (muo *MemberUpdateOne) SetUserID(u uuid.UUID) *MemberUpdateOne {
	muo.mutation.SetUserID(u)
	return muo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableUserID(u *uuid.UUID) *MemberUpdateOne {
	if u != nil {
		muo.SetUserID(*u)
	}
	return muo
}

// SetRoomID sets the "room_id" field.
func (muo *MemberUpdateOne) SetRoomID(u uuid.UUID) *MemberUpdateOne {
	muo.mutation.SetRoomID(u)
	return muo
}

// SetNillableRoomID sets the "room_id" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableRoomID(u *uuid.UUID) *MemberUpdateOne {
	if u != nil {
		muo.SetRoomID(*u)
	}
	return muo
}

// SetRole sets the "role" field.
func (muo *MemberUpdateOne) SetRole(m member.Role) *MemberUpdateOne {
	muo.mutation.SetRole(m)
	return muo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableRole(m *member.Role) *MemberUpdateOne {
	if m != nil {
		muo.SetRole(*m)
	}
	return muo
}

// SetNickName sets the "nick_name" field.
func (muo *MemberUpdateOne) SetNickName(s string) *MemberUpdateOne {
	muo.mutation.SetNickName(s)
	return muo
}

// SetNillableNickName sets the "nick_name" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableNickName(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetNickName(*s)
	}
	return muo
}

// ClearNickName clears the value of the "nick_name" field.
func (muo *MemberUpdateOne) ClearNickName() *MemberUpdateOne {
	muo.mutation.ClearNickName()
	return muo
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (muo *MemberUpdateOne) SetUsersID(id uuid.UUID) *MemberUpdateOne {
	muo.mutation.SetUsersID(id)
	return muo
}

// SetUsers sets the "users" edge to the User entity.
func (muo *MemberUpdateOne) SetUsers(u *User) *MemberUpdateOne {
	return muo.SetUsersID(u.ID)
}

// SetRoomsID sets the "rooms" edge to the Room entity by ID.
func (muo *MemberUpdateOne) SetRoomsID(id uuid.UUID) *MemberUpdateOne {
	muo.mutation.SetRoomsID(id)
	return muo
}

// SetRooms sets the "rooms" edge to the Room entity.
func (muo *MemberUpdateOne) SetRooms(r *Room) *MemberUpdateOne {
	return muo.SetRoomsID(r.ID)
}

// Mutation returns the MemberMutation object of the builder.
func (muo *MemberUpdateOne) Mutation() *MemberMutation {
	return muo.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (muo *MemberUpdateOne) ClearUsers() *MemberUpdateOne {
	muo.mutation.ClearUsers()
	return muo
}

// ClearRooms clears the "rooms" edge to the Room entity.
func (muo *MemberUpdateOne) ClearRooms() *MemberUpdateOne {
	muo.mutation.ClearRooms()
	return muo
}

// Where appends a list predicates to the MemberUpdate builder.
func (muo *MemberUpdateOne) Where(ps ...predicate.Member) *MemberUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MemberUpdateOne) Select(field string, fields ...string) *MemberUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Member entity.
func (muo *MemberUpdateOne) Save(ctx context.Context) (*Member, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MemberUpdateOne) SaveX(ctx context.Context) *Member {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MemberUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MemberUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MemberUpdateOne) check() error {
	if v, ok := muo.mutation.Role(); ok {
		if err := member.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "Member.role": %w`, err)}
		}
	}
	if v, ok := muo.mutation.NickName(); ok {
		if err := member.NickNameValidator(v); err != nil {
			return &ValidationError{Name: "nick_name", err: fmt.Errorf(`ent: validator failed for field "Member.nick_name": %w`, err)}
		}
	}
	if muo.mutation.UsersCleared() && len(muo.mutation.UsersIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Member.users"`)
	}
	if muo.mutation.RoomsCleared() && len(muo.mutation.RoomsIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Member.rooms"`)
	}
	return nil
}

func (muo *MemberUpdateOne) sqlSave(ctx context.Context) (_node *Member, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(member.Table, member.Columns, sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Member.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, member.FieldID)
		for _, f := range fields {
			if !member.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != member.FieldID {
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
	if value, ok := muo.mutation.Role(); ok {
		_spec.SetField(member.FieldRole, field.TypeEnum, value)
	}
	if value, ok := muo.mutation.NickName(); ok {
		_spec.SetField(member.FieldNickName, field.TypeString, value)
	}
	if muo.mutation.NickNameCleared() {
		_spec.ClearField(member.FieldNickName, field.TypeString)
	}
	if muo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   member.UsersTable,
			Columns: []string{member.UsersColumn},
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
			Table:   member.UsersTable,
			Columns: []string{member.UsersColumn},
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
	if muo.mutation.RoomsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   member.RoomsTable,
			Columns: []string{member.RoomsColumn},
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
			Table:   member.RoomsTable,
			Columns: []string{member.RoomsColumn},
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
	_node = &Member{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
