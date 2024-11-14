package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Document holds the schema definition for the Document entity.
type Document struct {
	ent.Schema
}

// Fields of the Document.
func (Document) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Unique().Default(""),
	}
}

// Edges of the Document.
func (Document) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("editors", User.Type),
		edge.From("owner", User.Type).Ref("owned_documents").Unique(),
		edge.From("allowed_users", User.Type).Ref("allowed_documents"),
	}
}
