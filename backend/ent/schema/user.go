package schema

import (
	"errors"
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User описывает таблицу пользователей
type User struct {
	ent.Schema
}

// Fields определяют поля модели User
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable(),
		field.Text("username").NotEmpty().Unique(),
		field.Text("firstname").Default(""),
		field.Text("lastname").Default(""),
		field.Enum("role").Values("user", "creator", "moderator", "administrator").Default("user"),
		field.JSON("info", map[string]interface{}{}).Default(map[string]interface{}{
			"information": "",
			"description": "",
		}),
		field.Bool("isPremium").Default(false),
		field.String("hash").NotEmpty().Validate(func(hash string) error {
			matched, err := regexp.MatchString("^[a-fA-F0-9]{64}$", hash)
			if err != nil || !matched {
				return errors.New("invalid hash")
			}
			return nil
		}),
	}
}

// Edges определяют связи с другими таблицами
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("services", Service.Type),
	}
}
