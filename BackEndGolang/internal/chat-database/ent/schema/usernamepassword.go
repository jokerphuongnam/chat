package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UsernamePassword holds the schema definition for the UsernamePassword entity.
type UsernamePassword struct {
	ent.Schema
}

// Fields of the UsernamePassword.
func (UsernamePassword) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.String("username").MinLen(4).MaxLen(16).Unique().NotEmpty(),
		field.String("password"),
	}
}

// Edges of the UsernamePassword.
func (UsernamePassword) Edges() []ent.Edge {
	return nil
}
