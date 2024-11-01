package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Member holds the schema definition for the Member entity.
type Member struct {
	ent.Schema
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("room_id", uuid.UUID{}),
		field.Enum("role").Values("OWNER", "ADMIN", "USER"),
		field.String("nick_name").NotEmpty().MaxLen(24).Optional(),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("members").Field("user_id").Unique().Required(),
		edge.From("rooms", Room.Type).Ref("members").Field("room_id").Unique().Required(),
	}
}

func (Member) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "room_id").Unique(),
	}
}
