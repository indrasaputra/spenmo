package entity

import "github.com/indrasaputra/hashids"

// User defines logical data for user.
type User struct {
	// ID represents a unique identifier of user.
	// This attribute is shown to the user using hashids.
	ID hashids.ID
	// Name represents user's name.
	Name string
	// Email represents user's email.
	Email string
}

// UserWallet defines logical data for user's wallet.
type UserWallet struct {
	// ID represents a unique identifier of user.
	// This attribute is shown to the user using hashids.
	ID hashids.ID
	// UserID represents to whom the wallet belongs to.
	// This attribute is shown to the user using hashids.
	UserID hashids.ID
	// Balance represents the balance.
	Balance float64
}

// UserCard defines logical data for user's card.
type UserCard struct {
	// ID represents a unique identifier of user.
	// This attribute is shown to the user using hashids.
	ID hashids.ID
	// UserID represents to whom the card belongs to.
	// This attribute is shown to the user using hashids.
	UserID hashids.ID
	// WalletID represents to which the card belongs to.
	// This attribute is shown to the user using hashids.
	WalletID hashids.ID
	// LimitDaily represents daily limit usage for the card.
	LimitDaily float64
	// LimitDaily represents monthly limit usage for the card.
	LimitMonthly float64
}
