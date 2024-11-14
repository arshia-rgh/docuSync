package schema

import (
	"context"
	gen "docuSync/ent"
	"docuSync/ent/hook"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"fmt"
	"time"
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
			Optional(),
		field.Text("text").
			Optional(),
		field.Time("created_at").
			Default(time.Now()),
		field.Time("updated_at").
			Default(time.Now()).
			UpdateDefault(func() time.Time {
				return time.Now()
			}),
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
						createdAt, _ := m.CreatedAt()
						ownerID, _ := m.OwnerID()
						m.SetTitle(fmt.Sprintf("doc-%v-owner-%v", createdAt, ownerID))
					}
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate,
		),
	}
}
