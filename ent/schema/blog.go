package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Blog holds the schema definition for the Blog entity.
type Blog struct {
	ent.Schema
}

// Fields of the Blog.
func (Blog) Fields() []ent.Field {
	return []ent.Field{field.UUID("blogId", uuid.UUID{}).Immutable().Unique(), field.String("blogTitle"), field.String("blogContent")}

}

// Edges of the Blog.
func (Blog) Edges() []ent.Edge {
	return nil
}
