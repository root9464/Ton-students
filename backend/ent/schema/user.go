package schema

import (
	"errors"
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// UserSchema описывает таблицу пользователей в базе данныхtype User struct {
type User struct {
	ent.Schema
}

// Fields определяют поля модели User
func (User) Fields() []ent.Field {
	return []ent.Field{field.Int64("id").Unique().Immutable(),
		field.String("firstName").Default(""),
		field.String("lastName").Default(""),
		field.String("username").NotEmpty().Unique(),
		field.Bool("isPremium").Default(false),
		field.Enum("role").Values("user", "creator", "moderator", "administrator").Default("user"),
		field.String("hash").NotEmpty().Validate(func(hash string) error {
			matched, err := regexp.MatchString("^[a-fA-F0-9]{64}$", hash)
			if err != nil || !matched {
				return errors.New("invalid hash")
			}
			return nil
		}),
	}
}

// Edges определяют связи с другими таблицами (если они есть)
func (User) Edges() []ent.Edge {
	return nil
}
