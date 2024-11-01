// Code generated by ent, DO NOT EDIT.

package roominfo

import (
	"chat-backend/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldLTE(FieldID, id))
}

// RoomImageURL applies equality check predicate on the "room_image_url" field. It's identical to RoomImageURLEQ.
func RoomImageURL(v string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldEQ(FieldRoomImageURL, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldEQ(FieldName, v))
}

// RoomImageURLEQ applies the EQ predicate on the "room_image_url" field.
func RoomImageURLEQ(v string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldEQ(FieldRoomImageURL, v))
}

// RoomImageURLNEQ applies the NEQ predicate on the "room_image_url" field.
func RoomImageURLNEQ(v string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldNEQ(FieldRoomImageURL, v))
}

// RoomImageURLIn applies the In predicate on the "room_image_url" field.
func RoomImageURLIn(vs ...string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldIn(FieldRoomImageURL, vs...))
}

// RoomImageURLNotIn applies the NotIn predicate on the "room_image_url" field.
func RoomImageURLNotIn(vs ...string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldNotIn(FieldRoomImageURL, vs...))
}

// RoomImageURLGT applies the GT predicate on the "room_image_url" field.
func RoomImageURLGT(v string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldGT(FieldRoomImageURL, v))
}

// RoomImageURLGTE applies the GTE predicate on the "room_image_url" field.
func RoomImageURLGTE(v string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldGTE(FieldRoomImageURL, v))
}

// RoomImageURLLT applies the LT predicate on the "room_image_url" field.
func RoomImageURLLT(v string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldLT(FieldRoomImageURL, v))
}

// RoomImageURLLTE applies the LTE predicate on the "room_image_url" field.
func RoomImageURLLTE(v string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldLTE(FieldRoomImageURL, v))
}

// RoomImageURLContains applies the Contains predicate on the "room_image_url" field.
func RoomImageURLContains(v string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldContains(FieldRoomImageURL, v))
}

// RoomImageURLHasPrefix applies the HasPrefix predicate on the "room_image_url" field.
func RoomImageURLHasPrefix(v string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldHasPrefix(FieldRoomImageURL, v))
}

// RoomImageURLHasSuffix applies the HasSuffix predicate on the "room_image_url" field.
func RoomImageURLHasSuffix(v string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldHasSuffix(FieldRoomImageURL, v))
}

// RoomImageURLIsNil applies the IsNil predicate on the "room_image_url" field.
func RoomImageURLIsNil() predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldIsNull(FieldRoomImageURL))
}

// RoomImageURLNotNil applies the NotNil predicate on the "room_image_url" field.
func RoomImageURLNotNil() predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldNotNull(FieldRoomImageURL))
}

// RoomImageURLEqualFold applies the EqualFold predicate on the "room_image_url" field.
func RoomImageURLEqualFold(v string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldEqualFold(FieldRoomImageURL, v))
}

// RoomImageURLContainsFold applies the ContainsFold predicate on the "room_image_url" field.
func RoomImageURLContainsFold(v string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldContainsFold(FieldRoomImageURL, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.RoomInfo {
	return predicate.RoomInfo(sql.FieldContainsFold(FieldName, v))
}

// HasRooms applies the HasEdge predicate on the "rooms" edge.
func HasRooms() predicate.RoomInfo {
	return predicate.RoomInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, RoomsTable, RoomsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoomsWith applies the HasEdge predicate on the "rooms" edge with a given conditions (other predicates).
func HasRoomsWith(preds ...predicate.Room) predicate.RoomInfo {
	return predicate.RoomInfo(func(s *sql.Selector) {
		step := newRoomsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RoomInfo) predicate.RoomInfo {
	return predicate.RoomInfo(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RoomInfo) predicate.RoomInfo {
	return predicate.RoomInfo(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RoomInfo) predicate.RoomInfo {
	return predicate.RoomInfo(sql.NotPredicates(p))
}
