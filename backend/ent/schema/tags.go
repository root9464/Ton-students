package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Tags holds the schema definition for the Tags entity.
type Tags struct {
	ent.Schema
}

// Fields of the Tags.
func (Tags) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.String("tagName").NotEmpty(),
	}
}

// Edges of the Tags.
func (Tags) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("service_tags", ServiceTag.Type).
			Ref("tag"),
	}
}
