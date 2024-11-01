// Code generated by ent, DO NOT EDIT.

package user

import (
	"chat-backend/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// AvatarURL applies equality check predicate on the "avatar_url" field. It's identical to AvatarURLEQ.
func AvatarURL(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAvatarURL, v))
}

// IDAuthorize applies equality check predicate on the "id_authorize" field. It's identical to IDAuthorizeEQ.
func IDAuthorize(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIDAuthorize, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldName, v))
}

// AvatarURLEQ applies the EQ predicate on the "avatar_url" field.
func AvatarURLEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAvatarURL, v))
}

// AvatarURLNEQ applies the NEQ predicate on the "avatar_url" field.
func AvatarURLNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAvatarURL, v))
}

// AvatarURLIn applies the In predicate on the "avatar_url" field.
func AvatarURLIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldAvatarURL, vs...))
}

// AvatarURLNotIn applies the NotIn predicate on the "avatar_url" field.
func AvatarURLNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldAvatarURL, vs...))
}

// AvatarURLGT applies the GT predicate on the "avatar_url" field.
func AvatarURLGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldAvatarURL, v))
}

// AvatarURLGTE applies the GTE predicate on the "avatar_url" field.
func AvatarURLGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldAvatarURL, v))
}

// AvatarURLLT applies the LT predicate on the "avatar_url" field.
func AvatarURLLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldAvatarURL, v))
}

// AvatarURLLTE applies the LTE predicate on the "avatar_url" field.
func AvatarURLLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldAvatarURL, v))
}

// AvatarURLContains applies the Contains predicate on the "avatar_url" field.
func AvatarURLContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldAvatarURL, v))
}

// AvatarURLHasPrefix applies the HasPrefix predicate on the "avatar_url" field.
func AvatarURLHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldAvatarURL, v))
}

// AvatarURLHasSuffix applies the HasSuffix predicate on the "avatar_url" field.
func AvatarURLHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldAvatarURL, v))
}

// AvatarURLIsNil applies the IsNil predicate on the "avatar_url" field.
func AvatarURLIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldAvatarURL))
}

// AvatarURLNotNil applies the NotNil predicate on the "avatar_url" field.
func AvatarURLNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldAvatarURL))
}

// AvatarURLEqualFold applies the EqualFold predicate on the "avatar_url" field.
func AvatarURLEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldAvatarURL, v))
}

// AvatarURLContainsFold applies the ContainsFold predicate on the "avatar_url" field.
func AvatarURLContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldAvatarURL, v))
}

// IDAuthorizeEQ applies the EQ predicate on the "id_authorize" field.
func IDAuthorizeEQ(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIDAuthorize, v))
}

// IDAuthorizeNEQ applies the NEQ predicate on the "id_authorize" field.
func IDAuthorizeNEQ(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldIDAuthorize, v))
}

// IDAuthorizeIn applies the In predicate on the "id_authorize" field.
func IDAuthorizeIn(vs ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldIn(FieldIDAuthorize, vs...))
}

// IDAuthorizeNotIn applies the NotIn predicate on the "id_authorize" field.
func IDAuthorizeNotIn(vs ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldIDAuthorize, vs...))
}

// IDAuthorizeGT applies the GT predicate on the "id_authorize" field.
func IDAuthorizeGT(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGT(FieldIDAuthorize, v))
}

// IDAuthorizeGTE applies the GTE predicate on the "id_authorize" field.
func IDAuthorizeGTE(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGTE(FieldIDAuthorize, v))
}

// IDAuthorizeLT applies the LT predicate on the "id_authorize" field.
func IDAuthorizeLT(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLT(FieldIDAuthorize, v))
}

// IDAuthorizeLTE applies the LTE predicate on the "id_authorize" field.
func IDAuthorizeLTE(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLTE(FieldIDAuthorize, v))
}

// HasMembers applies the HasEdge predicate on the "members" edge.
func HasMembers() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MembersTable, MembersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMembersWith applies the HasEdge predicate on the "members" edge with a given conditions (other predicates).
func HasMembersWith(preds ...predicate.Member) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newMembersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMessages applies the HasEdge predicate on the "messages" edge.
func HasMessages() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MessagesTable, MessagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMessagesWith applies the HasEdge predicate on the "messages" edge with a given conditions (other predicates).
func HasMessagesWith(preds ...predicate.Message) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newMessagesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
