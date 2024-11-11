package schema

import "entgo.io/ent"

// Document holds the schema definition for the Document entity.
type Document struct {
	ent.Schema
}

// Fields of the Document.
func (Document) Fields() []ent.Field {
	return nil
}

// Edges of the Document.
func (Document) Edges() []ent.Edge {
	return nil
}
