package models

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Positive().
			Unique(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now()),
		field.String("email").
			Unique(),
		field.String("password_hash").
			Optional(),
		field.String("confirm_selector").
			Optional(),
		field.String("confirm_verifier").
			Optional(),
		field.Bool("confirmed").
			Optional(),
		field.Int("attempt_count").
			Optional(),
		field.Time("last_attempt").
			Optional(),
		field.Time("locked").
			Optional(),
		field.String("recover_selector").
			Optional(),
		field.String("recover_verifier").
			Optional(),
		field.Time("recover_token_expiry").
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
