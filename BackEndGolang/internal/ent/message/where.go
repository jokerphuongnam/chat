// Code generated by ent, DO NOT EDIT.

package message

import (
	"chat-backend/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldID, id))
}

// DateSend applies equality check predicate on the "date_send" field. It's identical to DateSendEQ.
func DateSend(v uint64) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldDateSend, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldContent, v))
}

// IDRoom applies equality check predicate on the "id_room" field. It's identical to IDRoomEQ.
func IDRoom(v uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldIDRoom, v))
}

// IDUserSend applies equality check predicate on the "id_user_send" field. It's identical to IDUserSendEQ.
func IDUserSend(v uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldIDUserSend, v))
}

// DateSendEQ applies the EQ predicate on the "date_send" field.
func DateSendEQ(v uint64) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldDateSend, v))
}

// DateSendNEQ applies the NEQ predicate on the "date_send" field.
func DateSendNEQ(v uint64) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldDateSend, v))
}

// DateSendIn applies the In predicate on the "date_send" field.
func DateSendIn(vs ...uint64) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldDateSend, vs...))
}

// DateSendNotIn applies the NotIn predicate on the "date_send" field.
func DateSendNotIn(vs ...uint64) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldDateSend, vs...))
}

// DateSendGT applies the GT predicate on the "date_send" field.
func DateSendGT(v uint64) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldDateSend, v))
}

// DateSendGTE applies the GTE predicate on the "date_send" field.
func DateSendGTE(v uint64) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldDateSend, v))
}

// DateSendLT applies the LT predicate on the "date_send" field.
func DateSendLT(v uint64) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldDateSend, v))
}

// DateSendLTE applies the LTE predicate on the "date_send" field.
func DateSendLTE(v uint64) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldDateSend, v))
}

// TypeMessageEQ applies the EQ predicate on the "type_message" field.
func TypeMessageEQ(v TypeMessage) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldTypeMessage, v))
}

// TypeMessageNEQ applies the NEQ predicate on the "type_message" field.
func TypeMessageNEQ(v TypeMessage) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldTypeMessage, v))
}

// TypeMessageIn applies the In predicate on the "type_message" field.
func TypeMessageIn(vs ...TypeMessage) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldTypeMessage, vs...))
}

// TypeMessageNotIn applies the NotIn predicate on the "type_message" field.
func TypeMessageNotIn(vs ...TypeMessage) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldTypeMessage, vs...))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Message {
	return predicate.Message(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Message {
	return predicate.Message(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Message {
	return predicate.Message(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Message {
	return predicate.Message(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Message {
	return predicate.Message(sql.FieldContainsFold(FieldContent, v))
}

// IDRoomEQ applies the EQ predicate on the "id_room" field.
func IDRoomEQ(v uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldIDRoom, v))
}

// IDRoomNEQ applies the NEQ predicate on the "id_room" field.
func IDRoomNEQ(v uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldIDRoom, v))
}

// IDRoomIn applies the In predicate on the "id_room" field.
func IDRoomIn(vs ...uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldIDRoom, vs...))
}

// IDRoomNotIn applies the NotIn predicate on the "id_room" field.
func IDRoomNotIn(vs ...uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldIDRoom, vs...))
}

// IDRoomGT applies the GT predicate on the "id_room" field.
func IDRoomGT(v uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldIDRoom, v))
}

// IDRoomGTE applies the GTE predicate on the "id_room" field.
func IDRoomGTE(v uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldIDRoom, v))
}

// IDRoomLT applies the LT predicate on the "id_room" field.
func IDRoomLT(v uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldIDRoom, v))
}

// IDRoomLTE applies the LTE predicate on the "id_room" field.
func IDRoomLTE(v uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldIDRoom, v))
}

// IDUserSendEQ applies the EQ predicate on the "id_user_send" field.
func IDUserSendEQ(v uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldIDUserSend, v))
}

// IDUserSendNEQ applies the NEQ predicate on the "id_user_send" field.
func IDUserSendNEQ(v uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldIDUserSend, v))
}

// IDUserSendIn applies the In predicate on the "id_user_send" field.
func IDUserSendIn(vs ...uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldIDUserSend, vs...))
}

// IDUserSendNotIn applies the NotIn predicate on the "id_user_send" field.
func IDUserSendNotIn(vs ...uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldIDUserSend, vs...))
}

// IDUserSendGT applies the GT predicate on the "id_user_send" field.
func IDUserSendGT(v uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldIDUserSend, v))
}

// IDUserSendGTE applies the GTE predicate on the "id_user_send" field.
func IDUserSendGTE(v uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldIDUserSend, v))
}

// IDUserSendLT applies the LT predicate on the "id_user_send" field.
func IDUserSendLT(v uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldIDUserSend, v))
}

// IDUserSendLTE applies the LTE predicate on the "id_user_send" field.
func IDUserSendLTE(v uuid.UUID) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldIDUserSend, v))
}

// HasRooms applies the HasEdge predicate on the "rooms" edge.
func HasRooms() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoomsTable, RoomsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoomsWith applies the HasEdge predicate on the "rooms" edge with a given conditions (other predicates).
func HasRoomsWith(preds ...predicate.Room) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := newRoomsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Message) predicate.Message {
	return predicate.Message(sql.NotPredicates(p))
}
