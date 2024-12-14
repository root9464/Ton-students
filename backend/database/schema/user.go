package schema

import "entgo.io/ent"

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return nil
}

func (User) Edges() []ent.Edge {
	return nil
}
