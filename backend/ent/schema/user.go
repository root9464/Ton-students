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
	return []ent.Field{
		field.Int64("id").Immutable(),
		field.Text("username").NotEmpty().Unique(),
		field.Text("firstname").Default(""),
		field.Text("lastname").Default(""),
		field.Text("nickname").Default(""),
		field.Enum("selectedName").Values("firstname", "lastname", "nickname", "username").Default("username"),
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

// Edges определяют связи с другими таблицами (если они есть)
func (User) Edges() []ent.Edge {
	return nil
}