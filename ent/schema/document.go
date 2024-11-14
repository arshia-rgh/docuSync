package schema

import (
	"context"
	gen "docuSync/ent"
	"docuSync/ent/hook"
	_ "docuSync/ent/runtime"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"fmt"
)

// Document holds the schema definition for the Document entity.
type Document struct {
	ent.Schema
}

// Fields of the Document.
func (Document) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			Unique().
			Default("").
			Optional(),
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

func (Document) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.DocumentFunc(func(ctx context.Context, m *gen.DocumentMutation) (gen.Value, error) {
					if title, exists := m.Title(); !exists || title == "" {
						iD, _ := m.ID()
						ownerID, _ := m.OwnerID()
						m.SetTitle(fmt.Sprintf("doc-%d-owner-%d", iD, ownerID))
					}
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate,
		),
	}
}
