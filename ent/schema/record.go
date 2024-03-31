package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// Record holds the schema definition for the Record entity.
type Record struct {
	ent.Schema
}

func (Record) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "instructions_log"},
	}
}

// Fields of the Record.
func (Record) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("http_method").
			Values("GET", "POST"),
		field.String("ip_address"),
		field.String("instruction"),
		field.Time("timestamp").
			Default(time.Now()),
	}
}

// Edges of the Record.
func (Record) Edges() []ent.Edge {
	return nil
}
