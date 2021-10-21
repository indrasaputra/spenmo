package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserWallet holds the schema definition for the UserWallet entity.
type UserWallet struct {
	ent.Schema
}

// Fields of the UserWallet.
func (UserWallet) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Immutable().
			Unique(),
		field.Float("balance"),
		field.Int64("user_id").
			Optional(),
	}
}

// Edges of the UserWallet.
func (UserWallet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("user_wallets").
			Field("user_id").
			Unique(),
		edge.To("user_wallet_cards", UserCard.Type),
	}
}

// Mixin defines mixin schemas used in UserWallet.
func (UserWallet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
