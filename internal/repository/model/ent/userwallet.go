// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"

	"github.com/indrasaputra/spenmo/internal/repository/model/ent/user"
	"github.com/indrasaputra/spenmo/internal/repository/model/ent/userwallet"
)

// UserWallet is the model entity for the UserWallet schema.
type UserWallet struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Balance holds the value of the "balance" field.
	Balance float64 `json:"balance,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int64 `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserWalletQuery when eager-loading is set.
	Edges UserWalletEdges `json:"edges"`
}

// UserWalletEdges holds the relations/edges for other nodes in the graph.
type UserWalletEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// UserWalletCards holds the value of the user_wallet_cards edge.
	UserWalletCards []*UserCard `json:"user_wallet_cards,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserWalletEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// UserWalletCardsOrErr returns the UserWalletCards value or an error if the edge
// was not loaded in eager-loading.
func (e UserWalletEdges) UserWalletCardsOrErr() ([]*UserCard, error) {
	if e.loadedTypes[1] {
		return e.UserWalletCards, nil
	}
	return nil, &NotLoadedError{edge: "user_wallet_cards"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserWallet) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case userwallet.FieldBalance:
			values[i] = new(sql.NullFloat64)
		case userwallet.FieldID, userwallet.FieldUserID:
			values[i] = new(sql.NullInt64)
		case userwallet.FieldCreatedAt, userwallet.FieldUpdatedAt, userwallet.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserWallet", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserWallet fields.
func (uw *UserWallet) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userwallet.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			uw.ID = int64(value.Int64)
		case userwallet.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				uw.CreatedAt = value.Time
			}
		case userwallet.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				uw.UpdatedAt = value.Time
			}
		case userwallet.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				uw.DeletedAt = new(time.Time)
				*uw.DeletedAt = value.Time
			}
		case userwallet.FieldBalance:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field balance", values[i])
			} else if value.Valid {
				uw.Balance = value.Float64
			}
		case userwallet.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				uw.UserID = value.Int64
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the UserWallet entity.
func (uw *UserWallet) QueryUser() *UserQuery {
	return (&UserWalletClient{config: uw.config}).QueryUser(uw)
}

// QueryUserWalletCards queries the "user_wallet_cards" edge of the UserWallet entity.
func (uw *UserWallet) QueryUserWalletCards() *UserCardQuery {
	return (&UserWalletClient{config: uw.config}).QueryUserWalletCards(uw)
}

// Update returns a builder for updating this UserWallet.
// Note that you need to call UserWallet.Unwrap() before calling this method if this UserWallet
// was returned from a transaction, and the transaction was committed or rolled back.
func (uw *UserWallet) Update() *UserWalletUpdateOne {
	return (&UserWalletClient{config: uw.config}).UpdateOne(uw)
}

// Unwrap unwraps the UserWallet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uw *UserWallet) Unwrap() *UserWallet {
	tx, ok := uw.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserWallet is not a transactional entity")
	}
	uw.config.driver = tx.drv
	return uw
}

// String implements the fmt.Stringer.
func (uw *UserWallet) String() string {
	var builder strings.Builder
	builder.WriteString("UserWallet(")
	builder.WriteString(fmt.Sprintf("id=%v", uw.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(uw.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(uw.UpdatedAt.Format(time.ANSIC))
	if v := uw.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", balance=")
	builder.WriteString(fmt.Sprintf("%v", uw.Balance))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", uw.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// UserWallets is a parsable slice of UserWallet.
type UserWallets []*UserWallet

func (uw UserWallets) config(cfg config) {
	for _i := range uw {
		uw[_i].config = cfg
	}
}
