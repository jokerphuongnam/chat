// Code generated by ent, DO NOT EDIT.

package room

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the room type in the database.
	Label = "room"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldColor holds the string denoting the color field in the database.
	FieldColor = "color"
	// FieldIDInfo holds the string denoting the id_info field in the database.
	FieldIDInfo = "id_info"
	// EdgeMembers holds the string denoting the members edge name in mutations.
	EdgeMembers = "members"
	// EdgeRoomInfo holds the string denoting the room_info edge name in mutations.
	EdgeRoomInfo = "room_info"
	// EdgeMessages holds the string denoting the messages edge name in mutations.
	EdgeMessages = "messages"
	// Table holds the table name of the room in the database.
	Table = "rooms"
	// MembersTable is the table that holds the members relation/edge.
	MembersTable = "members"
	// MembersInverseTable is the table name for the Member entity.
	// It exists in this package in order to avoid circular dependency with the "member" package.
	MembersInverseTable = "members"
	// MembersColumn is the table column denoting the members relation/edge.
	MembersColumn = "room_id"
	// RoomInfoTable is the table that holds the room_info relation/edge.
	RoomInfoTable = "rooms"
	// RoomInfoInverseTable is the table name for the RoomInfo entity.
	// It exists in this package in order to avoid circular dependency with the "roominfo" package.
	RoomInfoInverseTable = "room_infos"
	// RoomInfoColumn is the table column denoting the room_info relation/edge.
	RoomInfoColumn = "id_info"
	// MessagesTable is the table that holds the messages relation/edge.
	MessagesTable = "messages"
	// MessagesInverseTable is the table name for the Message entity.
	// It exists in this package in order to avoid circular dependency with the "message" package.
	MessagesInverseTable = "messages"
	// MessagesColumn is the table column denoting the messages relation/edge.
	MessagesColumn = "room_messages"
)

// Columns holds all SQL columns for room fields.
var Columns = []string{
	FieldID,
	FieldColor,
	FieldIDInfo,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// ColorValidator is a validator for the "color" field. It is called by the builders before save.
	ColorValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Room queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByColor orders the results by the color field.
func ByColor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldColor, opts...).ToFunc()
}

// ByIDInfo orders the results by the id_info field.
func ByIDInfo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIDInfo, opts...).ToFunc()
}

// ByMembersCount orders the results by members count.
func ByMembersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMembersStep(), opts...)
	}
}

// ByMembers orders the results by members terms.
func ByMembers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMembersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRoomInfoField orders the results by room_info field.
func ByRoomInfoField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoomInfoStep(), sql.OrderByField(field, opts...))
	}
}

// ByMessagesCount orders the results by messages count.
func ByMessagesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMessagesStep(), opts...)
	}
}

// ByMessages orders the results by messages terms.
func ByMessages(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMessagesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newMembersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MembersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MembersTable, MembersColumn),
	)
}
func newRoomInfoStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoomInfoInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, RoomInfoTable, RoomInfoColumn),
	)
}
func newMessagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MessagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MessagesTable, MessagesColumn),
	)
}
