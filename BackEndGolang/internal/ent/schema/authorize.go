package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Authorize holds the schema definition for the Authorize entity.
type Authorize struct {
	ent.Schema
}

// Fields of the Authorize.
func (Authorize) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.String("jwt_token").MaxLen(255).Optional(),
		field.String("token").NotEmpty().Unique(),
	}
}

// Edges of the Authorize.
func (Authorize) Edges() []ent.Edge {
	return nil
}
