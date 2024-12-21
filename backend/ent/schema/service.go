package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Service описывает таблицу услуг
type Service struct {
	ent.Schema
}

// Fields определяют поля модели Service
func (Service) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Text("userName").NotEmpty(),
		field.Text("title").NotEmpty(),
		field.JSON("description", map[string]interface{}{}).Default(map[string]interface{}{
			"information": "",
		}),
		field.JSON("tags", []string{}),
		field.Int16("price"),
	}
}

// Edges определяют связи с другими таблицами (если они есть)
func (Service) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("services").Unique().Required(),
	}
}
