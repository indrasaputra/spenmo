// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/indrasaputra/spenmo/internal/repository/model/ent/user"
	"github.com/indrasaputra/spenmo/internal/repository/model/ent/usercard"
	"github.com/indrasaputra/spenmo/internal/repository/model/ent/userwallet"
)

// UserWalletCreate is the builder for creating a UserWallet entity.
type UserWalletCreate struct {
	config
	mutation *UserWalletMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (uwc *UserWalletCreate) SetCreatedAt(t time.Time) *UserWalletCreate {
	uwc.mutation.SetCreatedAt(t)
	return uwc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uwc *UserWalletCreate) SetNillableCreatedAt(t *time.Time) *UserWalletCreate {
	if t != nil {
		uwc.SetCreatedAt(*t)
	}
	return uwc
}

// SetUpdatedAt sets the "updated_at" field.
func (uwc *UserWalletCreate) SetUpdatedAt(t time.Time) *UserWalletCreate {
	uwc.mutation.SetUpdatedAt(t)
	return uwc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uwc *UserWalletCreate) SetNillableUpdatedAt(t *time.Time) *UserWalletCreate {
	if t != nil {
		uwc.SetUpdatedAt(*t)
	}
	return uwc
}

// SetDeletedAt sets the "deleted_at" field.
func (uwc *UserWalletCreate) SetDeletedAt(t time.Time) *UserWalletCreate {
	uwc.mutation.SetDeletedAt(t)
	return uwc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uwc *UserWalletCreate) SetNillableDeletedAt(t *time.Time) *UserWalletCreate {
	if t != nil {
		uwc.SetDeletedAt(*t)
	}
	return uwc
}

// SetBalance sets the "balance" field.
func (uwc *UserWalletCreate) SetBalance(f float64) *UserWalletCreate {
	uwc.mutation.SetBalance(f)
	return uwc
}

// SetUserID sets the "user_id" field.
func (uwc *UserWalletCreate) SetUserID(i int64) *UserWalletCreate {
	uwc.mutation.SetUserID(i)
	return uwc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (uwc *UserWalletCreate) SetNillableUserID(i *int64) *UserWalletCreate {
	if i != nil {
		uwc.SetUserID(*i)
	}
	return uwc
}

// SetID sets the "id" field.
func (uwc *UserWalletCreate) SetID(i int64) *UserWalletCreate {
	uwc.mutation.SetID(i)
	return uwc
}

// SetUser sets the "user" edge to the User entity.
func (uwc *UserWalletCreate) SetUser(u *User) *UserWalletCreate {
	return uwc.SetUserID(u.ID)
}

// AddUserWalletCardIDs adds the "user_wallet_cards" edge to the UserCard entity by IDs.
func (uwc *UserWalletCreate) AddUserWalletCardIDs(ids ...int64) *UserWalletCreate {
	uwc.mutation.AddUserWalletCardIDs(ids...)
	return uwc
}

// AddUserWalletCards adds the "user_wallet_cards" edges to the UserCard entity.
func (uwc *UserWalletCreate) AddUserWalletCards(u ...*UserCard) *UserWalletCreate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uwc.AddUserWalletCardIDs(ids...)
}

// Mutation returns the UserWalletMutation object of the builder.
func (uwc *UserWalletCreate) Mutation() *UserWalletMutation {
	return uwc.mutation
}

// Save creates the UserWallet in the database.
func (uwc *UserWalletCreate) Save(ctx context.Context) (*UserWallet, error) {
	var (
		err  error
		node *UserWallet
	)
	uwc.defaults()
	if len(uwc.hooks) == 0 {
		if err = uwc.check(); err != nil {
			return nil, err
		}
		node, err = uwc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserWalletMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uwc.check(); err != nil {
				return nil, err
			}
			uwc.mutation = mutation
			if node, err = uwc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uwc.hooks) - 1; i >= 0; i-- {
			if uwc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uwc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uwc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uwc *UserWalletCreate) SaveX(ctx context.Context) *UserWallet {
	v, err := uwc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uwc *UserWalletCreate) Exec(ctx context.Context) error {
	_, err := uwc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uwc *UserWalletCreate) ExecX(ctx context.Context) {
	if err := uwc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uwc *UserWalletCreate) defaults() {
	if _, ok := uwc.mutation.CreatedAt(); !ok {
		v := userwallet.DefaultCreatedAt()
		uwc.mutation.SetCreatedAt(v)
	}
	if _, ok := uwc.mutation.UpdatedAt(); !ok {
		v := userwallet.DefaultUpdatedAt()
		uwc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uwc *UserWalletCreate) check() error {
	if _, ok := uwc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := uwc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := uwc.mutation.Balance(); !ok {
		return &ValidationError{Name: "balance", err: errors.New(`ent: missing required field "balance"`)}
	}
	return nil
}

func (uwc *UserWalletCreate) sqlSave(ctx context.Context) (*UserWallet, error) {
	_node, _spec := uwc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uwc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (uwc *UserWalletCreate) createSpec() (*UserWallet, *sqlgraph.CreateSpec) {
	var (
		_node = &UserWallet{config: uwc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: userwallet.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: userwallet.FieldID,
			},
		}
	)
	if id, ok := uwc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uwc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userwallet.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := uwc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userwallet.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := uwc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userwallet.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := uwc.mutation.Balance(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: userwallet.FieldBalance,
		})
		_node.Balance = value
	}
	if nodes := uwc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userwallet.UserTable,
			Columns: []string{userwallet.UserColumn},
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
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uwc.mutation.UserWalletCardsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userwallet.UserWalletCardsTable,
			Columns: []string{userwallet.UserWalletCardsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: usercard.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserWalletCreateBulk is the builder for creating many UserWallet entities in bulk.
type UserWalletCreateBulk struct {
	config
	builders []*UserWalletCreate
}

// Save creates the UserWallet entities in the database.
func (uwcb *UserWalletCreateBulk) Save(ctx context.Context) ([]*UserWallet, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uwcb.builders))
	nodes := make([]*UserWallet, len(uwcb.builders))
	mutators := make([]Mutator, len(uwcb.builders))
	for i := range uwcb.builders {
		func(i int, root context.Context) {
			builder := uwcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserWalletMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uwcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uwcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, uwcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uwcb *UserWalletCreateBulk) SaveX(ctx context.Context) []*UserWallet {
	v, err := uwcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uwcb *UserWalletCreateBulk) Exec(ctx context.Context) error {
	_, err := uwcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uwcb *UserWalletCreateBulk) ExecX(ctx context.Context) {
	if err := uwcb.Exec(ctx); err != nil {
		panic(err)
	}
}
