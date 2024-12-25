package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ServiceTag holds the schema definition for the intermediate table between Service and Tags
type ServiceTag struct {
	ent.Schema
}

// Fields for the ServiceTag.
func (ServiceTag) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
	}
}

// Edges for the ServiceTag.
func (ServiceTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("service", Service.Type).
			Unique().
			Required(),
		edge.To("tag", Tags.Type).
			Unique().
			Required(),
	}
}
