// Code generated by ent, DO NOT EDIT.

package ent

import (
	"chat-backend/internal/ent/member"
	"chat-backend/internal/ent/room"
	"chat-backend/internal/ent/user"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// MemberCreate is the builder for creating a Member entity.
type MemberCreate struct {
	config
	mutation *MemberMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (mc *MemberCreate) SetUserID(u uuid.UUID) *MemberCreate {
	mc.mutation.SetUserID(u)
	return mc
}

// SetRoomID sets the "room_id" field.
func (mc *MemberCreate) SetRoomID(u uuid.UUID) *MemberCreate {
	mc.mutation.SetRoomID(u)
	return mc
}

// SetRole sets the "role" field.
func (mc *MemberCreate) SetRole(m member.Role) *MemberCreate {
	mc.mutation.SetRole(m)
	return mc
}

// SetNickName sets the "nick_name" field.
func (mc *MemberCreate) SetNickName(s string) *MemberCreate {
	mc.mutation.SetNickName(s)
	return mc
}

// SetNillableNickName sets the "nick_name" field if the given value is not nil.
func (mc *MemberCreate) SetNillableNickName(s *string) *MemberCreate {
	if s != nil {
		mc.SetNickName(*s)
	}
	return mc
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (mc *MemberCreate) SetUsersID(id uuid.UUID) *MemberCreate {
	mc.mutation.SetUsersID(id)
	return mc
}

// SetUsers sets the "users" edge to the User entity.
func (mc *MemberCreate) SetUsers(u *User) *MemberCreate {
	return mc.SetUsersID(u.ID)
}

// SetRoomsID sets the "rooms" edge to the Room entity by ID.
func (mc *MemberCreate) SetRoomsID(id uuid.UUID) *MemberCreate {
	mc.mutation.SetRoomsID(id)
	return mc
}

// SetRooms sets the "rooms" edge to the Room entity.
func (mc *MemberCreate) SetRooms(r *Room) *MemberCreate {
	return mc.SetRoomsID(r.ID)
}

// Mutation returns the MemberMutation object of the builder.
func (mc *MemberCreate) Mutation() *MemberMutation {
	return mc.mutation
}

// Save creates the Member in the database.
func (mc *MemberCreate) Save(ctx context.Context) (*Member, error) {
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MemberCreate) SaveX(ctx context.Context) *Member {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MemberCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MemberCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MemberCreate) check() error {
	if _, ok := mc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Member.user_id"`)}
	}
	if _, ok := mc.mutation.RoomID(); !ok {
		return &ValidationError{Name: "room_id", err: errors.New(`ent: missing required field "Member.room_id"`)}
	}
	if _, ok := mc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required field "Member.role"`)}
	}
	if v, ok := mc.mutation.Role(); ok {
		if err := member.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "Member.role": %w`, err)}
		}
	}
	if v, ok := mc.mutation.NickName(); ok {
		if err := member.NickNameValidator(v); err != nil {
			return &ValidationError{Name: "nick_name", err: fmt.Errorf(`ent: validator failed for field "Member.nick_name": %w`, err)}
		}
	}
	if len(mc.mutation.UsersIDs()) == 0 {
		return &ValidationError{Name: "users", err: errors.New(`ent: missing required edge "Member.users"`)}
	}
	if len(mc.mutation.RoomsIDs()) == 0 {
		return &ValidationError{Name: "rooms", err: errors.New(`ent: missing required edge "Member.rooms"`)}
	}
	return nil
}

func (mc *MemberCreate) sqlSave(ctx context.Context) (*Member, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MemberCreate) createSpec() (*Member, *sqlgraph.CreateSpec) {
	var (
		_node = &Member{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(member.Table, sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt))
	)
	if value, ok := mc.mutation.Role(); ok {
		_spec.SetField(member.FieldRole, field.TypeEnum, value)
		_node.Role = value
	}
	if value, ok := mc.mutation.NickName(); ok {
		_spec.SetField(member.FieldNickName, field.TypeString, value)
		_node.NickName = value
	}
	if nodes := mc.mutation.UsersIDs(); len(nodes) > 0 {
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
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.RoomsIDs(); len(nodes) > 0 {
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
		_node.RoomID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MemberCreateBulk is the builder for creating many Member entities in bulk.
type MemberCreateBulk struct {
	config
	err      error
	builders []*MemberCreate
}

// Save creates the Member entities in the database.
func (mcb *MemberCreateBulk) Save(ctx context.Context) ([]*Member, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Member, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MemberMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MemberCreateBulk) SaveX(ctx context.Context) []*Member {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MemberCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MemberCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}