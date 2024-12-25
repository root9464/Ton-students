package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Service описывает таблицу услуг
type Service struct {
	ent.Schema
}

// Fields определяют поля модели Service
func (Service) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.Int64("user_id"), // Внешний ключ для связи с User
		field.Text("title").NotEmpty(),
		field.JSON("description", map[string]interface{}{}).
			Default(map[string]interface{}{
				"information": "",
			}),
		field.Int16("price"),
	}
}

// Edges определяют связи с другими таблицами
func (Service) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("services").
			Field("user_id").
			Unique().
			Required(),

		edge.From("service_tags", ServiceTag.Type).
			Ref("service"),
	}
}
