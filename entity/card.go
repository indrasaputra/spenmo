package entity

// User defines logical data for user.
type User struct {
	// ID represents a unique identifier of user.
	ID int64
	// Name represents user's name.
	Name string
	// Email represents user's email.
	Email string
}

// UserWallet defines logical data for user's wallet.
type UserWallet struct {
	// ID represents a unique identifier of user.
	ID int64
	// UserID represents to whom the wallet belongs to.
	UserID int64
	// Balance represents the balance.
	Balance float64
}

// UserCard defines logical data for user's card.
type UserCard struct {
	// ID represents a unique identifier of user.
	ID int64
	// UserID represents to whom the card belongs to.
	UserID int64
	// WalletID represents to which the card belongs to.
	WalletID int64
	// LimitDaily represents daily limit usage for the card.
	LimitDaily float64
	// LimitDaily represents monthly limit usage for the card.
	LimitMonthly float64
}
