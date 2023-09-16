package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Movie holds the schema definition for the Movie entity.
type Movie struct {
	ent.Schema
}

// Fields of the Movie.
func (Movie) Fields() []ent.Field {
	return []ent.Field{

		field.String("name").
			Default("unknown"),
		field.Int("year"),
		field.String("genre"),
		field.String("language"),
		field.String("country"),
		field.String("director"),
		field.String("actors"),
	}
}

// Edges of the Movie.
func (Movie) Edges() []ent.Edge {
	return nil
}
