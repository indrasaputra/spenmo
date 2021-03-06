// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// EdgeUserWallets holds the string denoting the user_wallets edge name in mutations.
	EdgeUserWallets = "user_wallets"
	// EdgeUserCards holds the string denoting the user_cards edge name in mutations.
	EdgeUserCards = "user_cards"
	// Table holds the table name of the user in the database.
	Table = "users"
	// UserWalletsTable is the table that holds the user_wallets relation/edge.
	UserWalletsTable = "user_wallets"
	// UserWalletsInverseTable is the table name for the UserWallet entity.
	// It exists in this package in order to avoid circular dependency with the "userwallet" package.
	UserWalletsInverseTable = "user_wallets"
	// UserWalletsColumn is the table column denoting the user_wallets relation/edge.
	UserWalletsColumn = "user_id"
	// UserCardsTable is the table that holds the user_cards relation/edge.
	UserCardsTable = "user_cards"
	// UserCardsInverseTable is the table name for the UserCard entity.
	// It exists in this package in order to avoid circular dependency with the "usercard" package.
	UserCardsInverseTable = "user_cards"
	// UserCardsColumn is the table column denoting the user_cards relation/edge.
	UserCardsColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldEmail,
	FieldPassword,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
