package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Source holds the schema definition for the Source entity.
type Plugin struct {
	ent.Schema
}

// Fields of the Source.
func (Plugin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").Unique(),
		field.UUID("owner_id", uuid.UUID{}),
		field.Int("source_id"),
	}
}

func (Plugin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("owner", User.Type).
			Field("owner_id").
			Unique().
			Required(),
		edge.To("source", Source.Type).
			Field("source_id").
			Unique().
			Required(),
		edge.To("publications", Publication.Type),
	}
}
