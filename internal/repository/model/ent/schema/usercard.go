package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// UserCard holds the schema definition for the UserCard entity.
type UserCard struct {
	ent.Schema
}

// Fields of the UserCard.
func (UserCard) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Immutable().
			Unique(),
		field.Int64("user_id").
			Optional(),
		field.Int64("wallet_id").
			Optional(),
		field.Float("limit_daily"),
		field.Float("limit_monthly"),
	}
}

// Edges of the UserCard.
func (UserCard) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("user_cards").
			Field("user_id").
			Unique(),
		edge.From("wallet", UserWallet.Type).
			Ref("user_wallet_cards").
			Field("wallet_id").
			Unique(),
	}
}

// Mixin defines mixin schemas used in UserCard.
func (UserCard) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Indexes defines indexes used in UserCard.
func (UserCard) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id", "user_id", "deleted_at"),
		index.Fields("user_id", "deleted_at"),
	}
}
