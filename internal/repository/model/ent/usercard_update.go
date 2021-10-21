// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/indrasaputra/spenmo/internal/repository/model/ent/predicate"
	"github.com/indrasaputra/spenmo/internal/repository/model/ent/user"
	"github.com/indrasaputra/spenmo/internal/repository/model/ent/usercard"
	"github.com/indrasaputra/spenmo/internal/repository/model/ent/userwallet"
)

// UserCardUpdate is the builder for updating UserCard entities.
type UserCardUpdate struct {
	config
	hooks    []Hook
	mutation *UserCardMutation
}

// Where appends a list predicates to the UserCardUpdate builder.
func (ucu *UserCardUpdate) Where(ps ...predicate.UserCard) *UserCardUpdate {
	ucu.mutation.Where(ps...)
	return ucu
}

// SetUpdatedAt sets the "updated_at" field.
func (ucu *UserCardUpdate) SetUpdatedAt(t time.Time) *UserCardUpdate {
	ucu.mutation.SetUpdatedAt(t)
	return ucu
}

// SetDeletedAt sets the "deleted_at" field.
func (ucu *UserCardUpdate) SetDeletedAt(t time.Time) *UserCardUpdate {
	ucu.mutation.SetDeletedAt(t)
	return ucu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ucu *UserCardUpdate) SetNillableDeletedAt(t *time.Time) *UserCardUpdate {
	if t != nil {
		ucu.SetDeletedAt(*t)
	}
	return ucu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ucu *UserCardUpdate) ClearDeletedAt() *UserCardUpdate {
	ucu.mutation.ClearDeletedAt()
	return ucu
}

// SetUserID sets the "user_id" field.
func (ucu *UserCardUpdate) SetUserID(i int64) *UserCardUpdate {
	ucu.mutation.SetUserID(i)
	return ucu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ucu *UserCardUpdate) SetNillableUserID(i *int64) *UserCardUpdate {
	if i != nil {
		ucu.SetUserID(*i)
	}
	return ucu
}

// ClearUserID clears the value of the "user_id" field.
func (ucu *UserCardUpdate) ClearUserID() *UserCardUpdate {
	ucu.mutation.ClearUserID()
	return ucu
}

// SetWalletID sets the "wallet_id" field.
func (ucu *UserCardUpdate) SetWalletID(i int64) *UserCardUpdate {
	ucu.mutation.SetWalletID(i)
	return ucu
}

// SetNillableWalletID sets the "wallet_id" field if the given value is not nil.
func (ucu *UserCardUpdate) SetNillableWalletID(i *int64) *UserCardUpdate {
	if i != nil {
		ucu.SetWalletID(*i)
	}
	return ucu
}

// ClearWalletID clears the value of the "wallet_id" field.
func (ucu *UserCardUpdate) ClearWalletID() *UserCardUpdate {
	ucu.mutation.ClearWalletID()
	return ucu
}

// SetLimitDaily sets the "limit_daily" field.
func (ucu *UserCardUpdate) SetLimitDaily(f float64) *UserCardUpdate {
	ucu.mutation.ResetLimitDaily()
	ucu.mutation.SetLimitDaily(f)
	return ucu
}

// AddLimitDaily adds f to the "limit_daily" field.
func (ucu *UserCardUpdate) AddLimitDaily(f float64) *UserCardUpdate {
	ucu.mutation.AddLimitDaily(f)
	return ucu
}

// SetLimitMonthly sets the "limit_monthly" field.
func (ucu *UserCardUpdate) SetLimitMonthly(f float64) *UserCardUpdate {
	ucu.mutation.ResetLimitMonthly()
	ucu.mutation.SetLimitMonthly(f)
	return ucu
}

// AddLimitMonthly adds f to the "limit_monthly" field.
func (ucu *UserCardUpdate) AddLimitMonthly(f float64) *UserCardUpdate {
	ucu.mutation.AddLimitMonthly(f)
	return ucu
}

// SetUser sets the "user" edge to the User entity.
func (ucu *UserCardUpdate) SetUser(u *User) *UserCardUpdate {
	return ucu.SetUserID(u.ID)
}

// SetWallet sets the "wallet" edge to the UserWallet entity.
func (ucu *UserCardUpdate) SetWallet(u *UserWallet) *UserCardUpdate {
	return ucu.SetWalletID(u.ID)
}

// Mutation returns the UserCardMutation object of the builder.
func (ucu *UserCardUpdate) Mutation() *UserCardMutation {
	return ucu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ucu *UserCardUpdate) ClearUser() *UserCardUpdate {
	ucu.mutation.ClearUser()
	return ucu
}

// ClearWallet clears the "wallet" edge to the UserWallet entity.
func (ucu *UserCardUpdate) ClearWallet() *UserCardUpdate {
	ucu.mutation.ClearWallet()
	return ucu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ucu *UserCardUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ucu.defaults()
	if len(ucu.hooks) == 0 {
		affected, err = ucu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserCardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ucu.mutation = mutation
			affected, err = ucu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ucu.hooks) - 1; i >= 0; i-- {
			if ucu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ucu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ucu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ucu *UserCardUpdate) SaveX(ctx context.Context) int {
	affected, err := ucu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ucu *UserCardUpdate) Exec(ctx context.Context) error {
	_, err := ucu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucu *UserCardUpdate) ExecX(ctx context.Context) {
	if err := ucu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ucu *UserCardUpdate) defaults() {
	if _, ok := ucu.mutation.UpdatedAt(); !ok {
		v := usercard.UpdateDefaultUpdatedAt()
		ucu.mutation.SetUpdatedAt(v)
	}
}

func (ucu *UserCardUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usercard.Table,
			Columns: usercard.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: usercard.FieldID,
			},
		},
	}
	if ps := ucu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usercard.FieldUpdatedAt,
		})
	}
	if value, ok := ucu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usercard.FieldDeletedAt,
		})
	}
	if ucu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: usercard.FieldDeletedAt,
		})
	}
	if value, ok := ucu.mutation.LimitDaily(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: usercard.FieldLimitDaily,
		})
	}
	if value, ok := ucu.mutation.AddedLimitDaily(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: usercard.FieldLimitDaily,
		})
	}
	if value, ok := ucu.mutation.LimitMonthly(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: usercard.FieldLimitMonthly,
		})
	}
	if value, ok := ucu.mutation.AddedLimitMonthly(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: usercard.FieldLimitMonthly,
		})
	}
	if ucu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usercard.UserTable,
			Columns: []string{usercard.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usercard.UserTable,
			Columns: []string{usercard.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ucu.mutation.WalletCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usercard.WalletTable,
			Columns: []string{usercard.WalletColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: userwallet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucu.mutation.WalletIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usercard.WalletTable,
			Columns: []string{usercard.WalletColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: userwallet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ucu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usercard.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserCardUpdateOne is the builder for updating a single UserCard entity.
type UserCardUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserCardMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ucuo *UserCardUpdateOne) SetUpdatedAt(t time.Time) *UserCardUpdateOne {
	ucuo.mutation.SetUpdatedAt(t)
	return ucuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ucuo *UserCardUpdateOne) SetDeletedAt(t time.Time) *UserCardUpdateOne {
	ucuo.mutation.SetDeletedAt(t)
	return ucuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ucuo *UserCardUpdateOne) SetNillableDeletedAt(t *time.Time) *UserCardUpdateOne {
	if t != nil {
		ucuo.SetDeletedAt(*t)
	}
	return ucuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ucuo *UserCardUpdateOne) ClearDeletedAt() *UserCardUpdateOne {
	ucuo.mutation.ClearDeletedAt()
	return ucuo
}

// SetUserID sets the "user_id" field.
func (ucuo *UserCardUpdateOne) SetUserID(i int64) *UserCardUpdateOne {
	ucuo.mutation.SetUserID(i)
	return ucuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ucuo *UserCardUpdateOne) SetNillableUserID(i *int64) *UserCardUpdateOne {
	if i != nil {
		ucuo.SetUserID(*i)
	}
	return ucuo
}

// ClearUserID clears the value of the "user_id" field.
func (ucuo *UserCardUpdateOne) ClearUserID() *UserCardUpdateOne {
	ucuo.mutation.ClearUserID()
	return ucuo
}

// SetWalletID sets the "wallet_id" field.
func (ucuo *UserCardUpdateOne) SetWalletID(i int64) *UserCardUpdateOne {
	ucuo.mutation.SetWalletID(i)
	return ucuo
}

// SetNillableWalletID sets the "wallet_id" field if the given value is not nil.
func (ucuo *UserCardUpdateOne) SetNillableWalletID(i *int64) *UserCardUpdateOne {
	if i != nil {
		ucuo.SetWalletID(*i)
	}
	return ucuo
}

// ClearWalletID clears the value of the "wallet_id" field.
func (ucuo *UserCardUpdateOne) ClearWalletID() *UserCardUpdateOne {
	ucuo.mutation.ClearWalletID()
	return ucuo
}

// SetLimitDaily sets the "limit_daily" field.
func (ucuo *UserCardUpdateOne) SetLimitDaily(f float64) *UserCardUpdateOne {
	ucuo.mutation.ResetLimitDaily()
	ucuo.mutation.SetLimitDaily(f)
	return ucuo
}

// AddLimitDaily adds f to the "limit_daily" field.
func (ucuo *UserCardUpdateOne) AddLimitDaily(f float64) *UserCardUpdateOne {
	ucuo.mutation.AddLimitDaily(f)
	return ucuo
}

// SetLimitMonthly sets the "limit_monthly" field.
func (ucuo *UserCardUpdateOne) SetLimitMonthly(f float64) *UserCardUpdateOne {
	ucuo.mutation.ResetLimitMonthly()
	ucuo.mutation.SetLimitMonthly(f)
	return ucuo
}

// AddLimitMonthly adds f to the "limit_monthly" field.
func (ucuo *UserCardUpdateOne) AddLimitMonthly(f float64) *UserCardUpdateOne {
	ucuo.mutation.AddLimitMonthly(f)
	return ucuo
}

// SetUser sets the "user" edge to the User entity.
func (ucuo *UserCardUpdateOne) SetUser(u *User) *UserCardUpdateOne {
	return ucuo.SetUserID(u.ID)
}

// SetWallet sets the "wallet" edge to the UserWallet entity.
func (ucuo *UserCardUpdateOne) SetWallet(u *UserWallet) *UserCardUpdateOne {
	return ucuo.SetWalletID(u.ID)
}

// Mutation returns the UserCardMutation object of the builder.
func (ucuo *UserCardUpdateOne) Mutation() *UserCardMutation {
	return ucuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ucuo *UserCardUpdateOne) ClearUser() *UserCardUpdateOne {
	ucuo.mutation.ClearUser()
	return ucuo
}

// ClearWallet clears the "wallet" edge to the UserWallet entity.
func (ucuo *UserCardUpdateOne) ClearWallet() *UserCardUpdateOne {
	ucuo.mutation.ClearWallet()
	return ucuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ucuo *UserCardUpdateOne) Select(field string, fields ...string) *UserCardUpdateOne {
	ucuo.fields = append([]string{field}, fields...)
	return ucuo
}

// Save executes the query and returns the updated UserCard entity.
func (ucuo *UserCardUpdateOne) Save(ctx context.Context) (*UserCard, error) {
	var (
		err  error
		node *UserCard
	)
	ucuo.defaults()
	if len(ucuo.hooks) == 0 {
		node, err = ucuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserCardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ucuo.mutation = mutation
			node, err = ucuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ucuo.hooks) - 1; i >= 0; i-- {
			if ucuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ucuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ucuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ucuo *UserCardUpdateOne) SaveX(ctx context.Context) *UserCard {
	node, err := ucuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ucuo *UserCardUpdateOne) Exec(ctx context.Context) error {
	_, err := ucuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucuo *UserCardUpdateOne) ExecX(ctx context.Context) {
	if err := ucuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ucuo *UserCardUpdateOne) defaults() {
	if _, ok := ucuo.mutation.UpdatedAt(); !ok {
		v := usercard.UpdateDefaultUpdatedAt()
		ucuo.mutation.SetUpdatedAt(v)
	}
}

func (ucuo *UserCardUpdateOne) sqlSave(ctx context.Context) (_node *UserCard, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usercard.Table,
			Columns: usercard.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: usercard.FieldID,
			},
		},
	}
	id, ok := ucuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing UserCard.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ucuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usercard.FieldID)
		for _, f := range fields {
			if !usercard.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usercard.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ucuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usercard.FieldUpdatedAt,
		})
	}
	if value, ok := ucuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usercard.FieldDeletedAt,
		})
	}
	if ucuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: usercard.FieldDeletedAt,
		})
	}
	if value, ok := ucuo.mutation.LimitDaily(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: usercard.FieldLimitDaily,
		})
	}
	if value, ok := ucuo.mutation.AddedLimitDaily(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: usercard.FieldLimitDaily,
		})
	}
	if value, ok := ucuo.mutation.LimitMonthly(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: usercard.FieldLimitMonthly,
		})
	}
	if value, ok := ucuo.mutation.AddedLimitMonthly(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: usercard.FieldLimitMonthly,
		})
	}
	if ucuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usercard.UserTable,
			Columns: []string{usercard.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usercard.UserTable,
			Columns: []string{usercard.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ucuo.mutation.WalletCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usercard.WalletTable,
			Columns: []string{usercard.WalletColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: userwallet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucuo.mutation.WalletIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usercard.WalletTable,
			Columns: []string{usercard.WalletColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: userwallet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserCard{config: ucuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ucuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usercard.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}