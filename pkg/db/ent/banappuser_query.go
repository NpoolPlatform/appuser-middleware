// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/banappuser"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/predicate"
)

// BanAppUserQuery is the builder for querying BanAppUser entities.
type BanAppUserQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BanAppUser
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BanAppUserQuery builder.
func (bauq *BanAppUserQuery) Where(ps ...predicate.BanAppUser) *BanAppUserQuery {
	bauq.predicates = append(bauq.predicates, ps...)
	return bauq
}

// Limit adds a limit step to the query.
func (bauq *BanAppUserQuery) Limit(limit int) *BanAppUserQuery {
	bauq.limit = &limit
	return bauq
}

// Offset adds an offset step to the query.
func (bauq *BanAppUserQuery) Offset(offset int) *BanAppUserQuery {
	bauq.offset = &offset
	return bauq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bauq *BanAppUserQuery) Unique(unique bool) *BanAppUserQuery {
	bauq.unique = &unique
	return bauq
}

// Order adds an order step to the query.
func (bauq *BanAppUserQuery) Order(o ...OrderFunc) *BanAppUserQuery {
	bauq.order = append(bauq.order, o...)
	return bauq
}

// First returns the first BanAppUser entity from the query.
// Returns a *NotFoundError when no BanAppUser was found.
func (bauq *BanAppUserQuery) First(ctx context.Context) (*BanAppUser, error) {
	nodes, err := bauq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{banappuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bauq *BanAppUserQuery) FirstX(ctx context.Context) *BanAppUser {
	node, err := bauq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BanAppUser ID from the query.
// Returns a *NotFoundError when no BanAppUser ID was found.
func (bauq *BanAppUserQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = bauq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{banappuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bauq *BanAppUserQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := bauq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BanAppUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BanAppUser entity is found.
// Returns a *NotFoundError when no BanAppUser entities are found.
func (bauq *BanAppUserQuery) Only(ctx context.Context) (*BanAppUser, error) {
	nodes, err := bauq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{banappuser.Label}
	default:
		return nil, &NotSingularError{banappuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bauq *BanAppUserQuery) OnlyX(ctx context.Context) *BanAppUser {
	node, err := bauq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BanAppUser ID in the query.
// Returns a *NotSingularError when more than one BanAppUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (bauq *BanAppUserQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = bauq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{banappuser.Label}
	default:
		err = &NotSingularError{banappuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bauq *BanAppUserQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := bauq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BanAppUsers.
func (bauq *BanAppUserQuery) All(ctx context.Context) ([]*BanAppUser, error) {
	if err := bauq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bauq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bauq *BanAppUserQuery) AllX(ctx context.Context) []*BanAppUser {
	nodes, err := bauq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BanAppUser IDs.
func (bauq *BanAppUserQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := bauq.Select(banappuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bauq *BanAppUserQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := bauq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bauq *BanAppUserQuery) Count(ctx context.Context) (int, error) {
	if err := bauq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bauq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bauq *BanAppUserQuery) CountX(ctx context.Context) int {
	count, err := bauq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bauq *BanAppUserQuery) Exist(ctx context.Context) (bool, error) {
	if err := bauq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bauq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bauq *BanAppUserQuery) ExistX(ctx context.Context) bool {
	exist, err := bauq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BanAppUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bauq *BanAppUserQuery) Clone() *BanAppUserQuery {
	if bauq == nil {
		return nil
	}
	return &BanAppUserQuery{
		config:     bauq.config,
		limit:      bauq.limit,
		offset:     bauq.offset,
		order:      append([]OrderFunc{}, bauq.order...),
		predicates: append([]predicate.BanAppUser{}, bauq.predicates...),
		// clone intermediate query.
		sql:    bauq.sql.Clone(),
		path:   bauq.path,
		unique: bauq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BanAppUser.Query().
//		GroupBy(banappuser.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (bauq *BanAppUserQuery) GroupBy(field string, fields ...string) *BanAppUserGroupBy {
	grbuild := &BanAppUserGroupBy{config: bauq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bauq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bauq.sqlQuery(ctx), nil
	}
	grbuild.label = banappuser.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.BanAppUser.Query().
//		Select(banappuser.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (bauq *BanAppUserQuery) Select(fields ...string) *BanAppUserSelect {
	bauq.fields = append(bauq.fields, fields...)
	selbuild := &BanAppUserSelect{BanAppUserQuery: bauq}
	selbuild.label = banappuser.Label
	selbuild.flds, selbuild.scan = &bauq.fields, selbuild.Scan
	return selbuild
}

func (bauq *BanAppUserQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bauq.fields {
		if !banappuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bauq.path != nil {
		prev, err := bauq.path(ctx)
		if err != nil {
			return err
		}
		bauq.sql = prev
	}
	if banappuser.Policy == nil {
		return errors.New("ent: uninitialized banappuser.Policy (forgotten import ent/runtime?)")
	}
	if err := banappuser.Policy.EvalQuery(ctx, bauq); err != nil {
		return err
	}
	return nil
}

func (bauq *BanAppUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BanAppUser, error) {
	var (
		nodes = []*BanAppUser{}
		_spec = bauq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*BanAppUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &BanAppUser{config: bauq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(bauq.modifiers) > 0 {
		_spec.Modifiers = bauq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bauq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (bauq *BanAppUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bauq.querySpec()
	if len(bauq.modifiers) > 0 {
		_spec.Modifiers = bauq.modifiers
	}
	_spec.Node.Columns = bauq.fields
	if len(bauq.fields) > 0 {
		_spec.Unique = bauq.unique != nil && *bauq.unique
	}
	return sqlgraph.CountNodes(ctx, bauq.driver, _spec)
}

func (bauq *BanAppUserQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bauq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (bauq *BanAppUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   banappuser.Table,
			Columns: banappuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: banappuser.FieldID,
			},
		},
		From:   bauq.sql,
		Unique: true,
	}
	if unique := bauq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bauq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, banappuser.FieldID)
		for i := range fields {
			if fields[i] != banappuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bauq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bauq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bauq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bauq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bauq *BanAppUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bauq.driver.Dialect())
	t1 := builder.Table(banappuser.Table)
	columns := bauq.fields
	if len(columns) == 0 {
		columns = banappuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bauq.sql != nil {
		selector = bauq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bauq.unique != nil && *bauq.unique {
		selector.Distinct()
	}
	for _, m := range bauq.modifiers {
		m(selector)
	}
	for _, p := range bauq.predicates {
		p(selector)
	}
	for _, p := range bauq.order {
		p(selector)
	}
	if offset := bauq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bauq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (bauq *BanAppUserQuery) ForUpdate(opts ...sql.LockOption) *BanAppUserQuery {
	if bauq.driver.Dialect() == dialect.Postgres {
		bauq.Unique(false)
	}
	bauq.modifiers = append(bauq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return bauq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (bauq *BanAppUserQuery) ForShare(opts ...sql.LockOption) *BanAppUserQuery {
	if bauq.driver.Dialect() == dialect.Postgres {
		bauq.Unique(false)
	}
	bauq.modifiers = append(bauq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return bauq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (bauq *BanAppUserQuery) Modify(modifiers ...func(s *sql.Selector)) *BanAppUserSelect {
	bauq.modifiers = append(bauq.modifiers, modifiers...)
	return bauq.Select()
}

// BanAppUserGroupBy is the group-by builder for BanAppUser entities.
type BanAppUserGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (baugb *BanAppUserGroupBy) Aggregate(fns ...AggregateFunc) *BanAppUserGroupBy {
	baugb.fns = append(baugb.fns, fns...)
	return baugb
}

// Scan applies the group-by query and scans the result into the given value.
func (baugb *BanAppUserGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := baugb.path(ctx)
	if err != nil {
		return err
	}
	baugb.sql = query
	return baugb.sqlScan(ctx, v)
}

func (baugb *BanAppUserGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range baugb.fields {
		if !banappuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := baugb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := baugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (baugb *BanAppUserGroupBy) sqlQuery() *sql.Selector {
	selector := baugb.sql.Select()
	aggregation := make([]string, 0, len(baugb.fns))
	for _, fn := range baugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(baugb.fields)+len(baugb.fns))
		for _, f := range baugb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(baugb.fields...)...)
}

// BanAppUserSelect is the builder for selecting fields of BanAppUser entities.
type BanAppUserSelect struct {
	*BanAppUserQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (baus *BanAppUserSelect) Scan(ctx context.Context, v interface{}) error {
	if err := baus.prepareQuery(ctx); err != nil {
		return err
	}
	baus.sql = baus.BanAppUserQuery.sqlQuery(ctx)
	return baus.sqlScan(ctx, v)
}

func (baus *BanAppUserSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := baus.sql.Query()
	if err := baus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (baus *BanAppUserSelect) Modify(modifiers ...func(s *sql.Selector)) *BanAppUserSelect {
	baus.modifiers = append(baus.modifiers, modifiers...)
	return baus
}
