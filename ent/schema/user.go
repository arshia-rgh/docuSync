package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"regexp"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Optional(),
		field.String("last_name").
			Optional(),
		field.String("username").
			Unique(),
		field.String("password").
			Match(regexp.MustCompile(`[A-Za-z\d@$!%*?&]{8,}`)),
		field.String("email").
			Unique().
			Match(regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)),
		field.Time("created_at").
			Default(time.Now()),
		field.Time("updated_at").
			Default(time.Now()).
			UpdateDefault(func() time.Time {
				return time.Now()
			}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("allowed_documents", Document.Type),
		edge.To("owned_documents", Document.Type),
		edge.From("edited_documents", Document.Type).Ref("editors"),
	}
}
