package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Room holds the schema definition for the Room entity.
type Room struct {
	ent.Schema
}

// Fields of the Room.
func (Room) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.String("color").MaxLen(255),
		field.UUID("id_info", uuid.UUID{}).Optional().Unique(),
	}
}

// Edges of the Room.
func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("members", Member.Type),
		edge.From("room_info", RoomInfo.Type).Ref("rooms").Field("id_info").Unique(),
		edge.To("messages", Message.Type),
	}
}
