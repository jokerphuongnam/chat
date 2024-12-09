package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// RoomInfo holds the schema definition for the RoomInfo entity.
type RoomInfo struct {
	ent.Schema
}

// Fields of the RoomInfo.
func (RoomInfo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.String("room_image_url").Optional(),
        field.String("name").Optional(),
	}
}

// Edges of the RoomInfo.
func (RoomInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("rooms", Room.Type).Unique(),
	}
}
