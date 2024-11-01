package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.Uint64("date_send"),
		field.Enum("type_message").Values("text", "image", "audio", "video", "location", "contact"),
		field.String("content"),
		field.UUID("id_room", uuid.UUID{}),
		field.UUID("id_user_send", uuid.UUID{}),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("rooms", Room.Type).Ref("messages").Unique().Required(),
		edge.From("users", User.Type).Ref("messages").Unique().Required(),
	}
}
