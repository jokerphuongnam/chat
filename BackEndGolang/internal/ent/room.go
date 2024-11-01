// Code generated by ent, DO NOT EDIT.

package ent

import (
	"chat-backend/internal/ent/room"
	"chat-backend/internal/ent/roominfo"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Room is the model entity for the Room schema.
type Room struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Color holds the value of the "color" field.
	Color string `json:"color,omitempty"`
	// IDInfo holds the value of the "id_info" field.
	IDInfo uuid.UUID `json:"id_info,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoomQuery when eager-loading is set.
	Edges        RoomEdges `json:"edges"`
	selectValues sql.SelectValues
}

// RoomEdges holds the relations/edges for other nodes in the graph.
type RoomEdges struct {
	// Members holds the value of the members edge.
	Members []*Member `json:"members,omitempty"`
	// RoomInfo holds the value of the room_info edge.
	RoomInfo *RoomInfo `json:"room_info,omitempty"`
	// Messages holds the value of the messages edge.
	Messages []*Message `json:"messages,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// MembersOrErr returns the Members value or an error if the edge
// was not loaded in eager-loading.
func (e RoomEdges) MembersOrErr() ([]*Member, error) {
	if e.loadedTypes[0] {
		return e.Members, nil
	}
	return nil, &NotLoadedError{edge: "members"}
}

// RoomInfoOrErr returns the RoomInfo value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoomEdges) RoomInfoOrErr() (*RoomInfo, error) {
	if e.RoomInfo != nil {
		return e.RoomInfo, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: roominfo.Label}
	}
	return nil, &NotLoadedError{edge: "room_info"}
}

// MessagesOrErr returns the Messages value or an error if the edge
// was not loaded in eager-loading.
func (e RoomEdges) MessagesOrErr() ([]*Message, error) {
	if e.loadedTypes[2] {
		return e.Messages, nil
	}
	return nil, &NotLoadedError{edge: "messages"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Room) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case room.FieldColor:
			values[i] = new(sql.NullString)
		case room.FieldID, room.FieldIDInfo:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Room fields.
func (r *Room) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case room.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				r.ID = *value
			}
		case room.FieldColor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field color", values[i])
			} else if value.Valid {
				r.Color = value.String
			}
		case room.FieldIDInfo:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id_info", values[i])
			} else if value != nil {
				r.IDInfo = *value
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Room.
// This includes values selected through modifiers, order, etc.
func (r *Room) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryMembers queries the "members" edge of the Room entity.
func (r *Room) QueryMembers() *MemberQuery {
	return NewRoomClient(r.config).QueryMembers(r)
}

// QueryRoomInfo queries the "room_info" edge of the Room entity.
func (r *Room) QueryRoomInfo() *RoomInfoQuery {
	return NewRoomClient(r.config).QueryRoomInfo(r)
}

// QueryMessages queries the "messages" edge of the Room entity.
func (r *Room) QueryMessages() *MessageQuery {
	return NewRoomClient(r.config).QueryMessages(r)
}

// Update returns a builder for updating this Room.
// Note that you need to call Room.Unwrap() before calling this method if this Room
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Room) Update() *RoomUpdateOne {
	return NewRoomClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Room entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Room) Unwrap() *Room {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Room is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Room) String() string {
	var builder strings.Builder
	builder.WriteString("Room(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("color=")
	builder.WriteString(r.Color)
	builder.WriteString(", ")
	builder.WriteString("id_info=")
	builder.WriteString(fmt.Sprintf("%v", r.IDInfo))
	builder.WriteByte(')')
	return builder.String()
}

// Rooms is a parsable slice of Room.
type Rooms []*Room
