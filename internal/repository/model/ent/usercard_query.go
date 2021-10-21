// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/indrasaputra/spenmo/internal/repository/model/ent/predicate"
	"github.com/indrasaputra/spenmo/internal/repository/model/ent/user"
	"github.com/indrasaputra/spenmo/internal/repository/model/ent/usercard"
	"github.com/indrasaputra/spenmo/internal/repository/model/ent/userwallet"
)

// UserCardQuery is the builder for querying UserCard entities.
type UserCardQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UserCard
	// eager-loading edges.
	withUser   *UserQuery
	withWallet *UserWalletQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserCardQuery builder.
func (ucq *UserCardQuery) Where(ps ...predicate.UserCard) *UserCardQuery {
	ucq.predicates = append(ucq.predicates, ps...)
	return ucq
}

// Limit adds a limit step to the query.
func (ucq *UserCardQuery) Limit(limit int) *UserCardQuery {
	ucq.limit = &limit
	return ucq
}

// Offset adds an offset step to the query.
func (ucq *UserCardQuery) Offset(offset int) *UserCardQuery {
	ucq.offset = &offset
	return ucq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ucq *UserCardQuery) Unique(unique bool) *UserCardQuery {
	ucq.unique = &unique
	return ucq
}

// Order adds an order step to the query.
func (ucq *UserCardQuery) Order(o ...OrderFunc) *UserCardQuery {
	ucq.order = append(ucq.order, o...)
	return ucq
}

// QueryUser chains the current query on the "user" edge.
func (ucq *UserCardQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: ucq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ucq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ucq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usercard.Table, usercard.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, usercard.UserTable, usercard.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ucq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWallet chains the current query on the "wallet" edge.
func (ucq *UserCardQuery) QueryWallet() *UserWalletQuery {
	query := &UserWalletQuery{config: ucq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ucq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ucq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usercard.Table, usercard.FieldID, selector),
			sqlgraph.To(userwallet.Table, userwallet.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, usercard.WalletTable, usercard.WalletColumn),
		)
		fromU = sqlgraph.SetNeighbors(ucq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserCard entity from the query.
// Returns a *NotFoundError when no UserCard was found.
func (ucq *UserCardQuery) First(ctx context.Context) (*UserCard, error) {
	nodes, err := ucq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usercard.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ucq *UserCardQuery) FirstX(ctx context.Context) *UserCard {
	node, err := ucq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserCard ID from the query.
// Returns a *NotFoundError when no UserCard ID was found.
func (ucq *UserCardQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ucq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usercard.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ucq *UserCardQuery) FirstIDX(ctx context.Context) int64 {
	id, err := ucq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserCard entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one UserCard entity is not found.
// Returns a *NotFoundError when no UserCard entities are found.
func (ucq *UserCardQuery) Only(ctx context.Context) (*UserCard, error) {
	nodes, err := ucq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usercard.Label}
	default:
		return nil, &NotSingularError{usercard.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ucq *UserCardQuery) OnlyX(ctx context.Context) *UserCard {
	node, err := ucq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserCard ID in the query.
// Returns a *NotSingularError when exactly one UserCard ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ucq *UserCardQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ucq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usercard.Label}
	default:
		err = &NotSingularError{usercard.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ucq *UserCardQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := ucq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserCards.
func (ucq *UserCardQuery) All(ctx context.Context) ([]*UserCard, error) {
	if err := ucq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ucq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ucq *UserCardQuery) AllX(ctx context.Context) []*UserCard {
	nodes, err := ucq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserCard IDs.
func (ucq *UserCardQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := ucq.Select(usercard.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ucq *UserCardQuery) IDsX(ctx context.Context) []int64 {
	ids, err := ucq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ucq *UserCardQuery) Count(ctx context.Context) (int, error) {
	if err := ucq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ucq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ucq *UserCardQuery) CountX(ctx context.Context) int {
	count, err := ucq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ucq *UserCardQuery) Exist(ctx context.Context) (bool, error) {
	if err := ucq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ucq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ucq *UserCardQuery) ExistX(ctx context.Context) bool {
	exist, err := ucq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserCardQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ucq *UserCardQuery) Clone() *UserCardQuery {
	if ucq == nil {
		return nil
	}
	return &UserCardQuery{
		config:     ucq.config,
		limit:      ucq.limit,
		offset:     ucq.offset,
		order:      append([]OrderFunc{}, ucq.order...),
		predicates: append([]predicate.UserCard{}, ucq.predicates...),
		withUser:   ucq.withUser.Clone(),
		withWallet: ucq.withWallet.Clone(),
		// clone intermediate query.
		sql:  ucq.sql.Clone(),
		path: ucq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ucq *UserCardQuery) WithUser(opts ...func(*UserQuery)) *UserCardQuery {
	query := &UserQuery{config: ucq.config}
	for _, opt := range opts {
		opt(query)
	}
	ucq.withUser = query
	return ucq
}

// WithWallet tells the query-builder to eager-load the nodes that are connected to
// the "wallet" edge. The optional arguments are used to configure the query builder of the edge.
func (ucq *UserCardQuery) WithWallet(opts ...func(*UserWalletQuery)) *UserCardQuery {
	query := &UserWalletQuery{config: ucq.config}
	for _, opt := range opts {
		opt(query)
	}
	ucq.withWallet = query
	return ucq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserCard.Query().
//		GroupBy(usercard.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ucq *UserCardQuery) GroupBy(field string, fields ...string) *UserCardGroupBy {
	group := &UserCardGroupBy{config: ucq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ucq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ucq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.UserCard.Query().
//		Select(usercard.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (ucq *UserCardQuery) Select(fields ...string) *UserCardSelect {
	ucq.fields = append(ucq.fields, fields...)
	return &UserCardSelect{UserCardQuery: ucq}
}

func (ucq *UserCardQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ucq.fields {
		if !usercard.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ucq.path != nil {
		prev, err := ucq.path(ctx)
		if err != nil {
			return err
		}
		ucq.sql = prev
	}
	return nil
}

func (ucq *UserCardQuery) sqlAll(ctx context.Context) ([]*UserCard, error) {
	var (
		nodes       = []*UserCard{}
		_spec       = ucq.querySpec()
		loadedTypes = [2]bool{
			ucq.withUser != nil,
			ucq.withWallet != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &UserCard{config: ucq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, ucq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ucq.withUser; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*UserCard)
		for i := range nodes {
			fk := nodes[i].UserID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	if query := ucq.withWallet; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*UserCard)
		for i := range nodes {
			fk := nodes[i].WalletID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(userwallet.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "wallet_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Wallet = n
			}
		}
	}

	return nodes, nil
}

func (ucq *UserCardQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ucq.querySpec()
	return sqlgraph.CountNodes(ctx, ucq.driver, _spec)
}

func (ucq *UserCardQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ucq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ucq *UserCardQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usercard.Table,
			Columns: usercard.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: usercard.FieldID,
			},
		},
		From:   ucq.sql,
		Unique: true,
	}
	if unique := ucq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ucq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usercard.FieldID)
		for i := range fields {
			if fields[i] != usercard.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ucq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ucq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ucq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ucq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ucq *UserCardQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ucq.driver.Dialect())
	t1 := builder.Table(usercard.Table)
	columns := ucq.fields
	if len(columns) == 0 {
		columns = usercard.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ucq.sql != nil {
		selector = ucq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range ucq.predicates {
		p(selector)
	}
	for _, p := range ucq.order {
		p(selector)
	}
	if offset := ucq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ucq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserCardGroupBy is the group-by builder for UserCard entities.
type UserCardGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ucgb *UserCardGroupBy) Aggregate(fns ...AggregateFunc) *UserCardGroupBy {
	ucgb.fns = append(ucgb.fns, fns...)
	return ucgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ucgb *UserCardGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ucgb.path(ctx)
	if err != nil {
		return err
	}
	ucgb.sql = query
	return ucgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ucgb *UserCardGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ucgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ucgb *UserCardGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ucgb.fields) > 1 {
		return nil, errors.New("ent: UserCardGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ucgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ucgb *UserCardGroupBy) StringsX(ctx context.Context) []string {
	v, err := ucgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ucgb *UserCardGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ucgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usercard.Label}
	default:
		err = fmt.Errorf("ent: UserCardGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ucgb *UserCardGroupBy) StringX(ctx context.Context) string {
	v, err := ucgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ucgb *UserCardGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ucgb.fields) > 1 {
		return nil, errors.New("ent: UserCardGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ucgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ucgb *UserCardGroupBy) IntsX(ctx context.Context) []int {
	v, err := ucgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ucgb *UserCardGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ucgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usercard.Label}
	default:
		err = fmt.Errorf("ent: UserCardGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ucgb *UserCardGroupBy) IntX(ctx context.Context) int {
	v, err := ucgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ucgb *UserCardGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ucgb.fields) > 1 {
		return nil, errors.New("ent: UserCardGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ucgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ucgb *UserCardGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ucgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ucgb *UserCardGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ucgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usercard.Label}
	default:
		err = fmt.Errorf("ent: UserCardGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ucgb *UserCardGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ucgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ucgb *UserCardGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ucgb.fields) > 1 {
		return nil, errors.New("ent: UserCardGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ucgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ucgb *UserCardGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ucgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ucgb *UserCardGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ucgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usercard.Label}
	default:
		err = fmt.Errorf("ent: UserCardGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ucgb *UserCardGroupBy) BoolX(ctx context.Context) bool {
	v, err := ucgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ucgb *UserCardGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ucgb.fields {
		if !usercard.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ucgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ucgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ucgb *UserCardGroupBy) sqlQuery() *sql.Selector {
	selector := ucgb.sql.Select()
	aggregation := make([]string, 0, len(ucgb.fns))
	for _, fn := range ucgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ucgb.fields)+len(ucgb.fns))
		for _, f := range ucgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ucgb.fields...)...)
}

// UserCardSelect is the builder for selecting fields of UserCard entities.
type UserCardSelect struct {
	*UserCardQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ucs *UserCardSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ucs.prepareQuery(ctx); err != nil {
		return err
	}
	ucs.sql = ucs.UserCardQuery.sqlQuery(ctx)
	return ucs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ucs *UserCardSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ucs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ucs *UserCardSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ucs.fields) > 1 {
		return nil, errors.New("ent: UserCardSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ucs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ucs *UserCardSelect) StringsX(ctx context.Context) []string {
	v, err := ucs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ucs *UserCardSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ucs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usercard.Label}
	default:
		err = fmt.Errorf("ent: UserCardSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ucs *UserCardSelect) StringX(ctx context.Context) string {
	v, err := ucs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ucs *UserCardSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ucs.fields) > 1 {
		return nil, errors.New("ent: UserCardSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ucs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ucs *UserCardSelect) IntsX(ctx context.Context) []int {
	v, err := ucs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ucs *UserCardSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ucs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usercard.Label}
	default:
		err = fmt.Errorf("ent: UserCardSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ucs *UserCardSelect) IntX(ctx context.Context) int {
	v, err := ucs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ucs *UserCardSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ucs.fields) > 1 {
		return nil, errors.New("ent: UserCardSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ucs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ucs *UserCardSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ucs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ucs *UserCardSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ucs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usercard.Label}
	default:
		err = fmt.Errorf("ent: UserCardSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ucs *UserCardSelect) Float64X(ctx context.Context) float64 {
	v, err := ucs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ucs *UserCardSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ucs.fields) > 1 {
		return nil, errors.New("ent: UserCardSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ucs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ucs *UserCardSelect) BoolsX(ctx context.Context) []bool {
	v, err := ucs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ucs *UserCardSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ucs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usercard.Label}
	default:
		err = fmt.Errorf("ent: UserCardSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ucs *UserCardSelect) BoolX(ctx context.Context) bool {
	v, err := ucs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ucs *UserCardSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ucs.sql.Query()
	if err := ucs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
