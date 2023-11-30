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
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/banapp"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/predicate"
)

// BanAppQuery is the builder for querying BanApp entities.
type BanAppQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BanApp
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BanAppQuery builder.
func (baq *BanAppQuery) Where(ps ...predicate.BanApp) *BanAppQuery {
	baq.predicates = append(baq.predicates, ps...)
	return baq
}

// Limit adds a limit step to the query.
func (baq *BanAppQuery) Limit(limit int) *BanAppQuery {
	baq.limit = &limit
	return baq
}

// Offset adds an offset step to the query.
func (baq *BanAppQuery) Offset(offset int) *BanAppQuery {
	baq.offset = &offset
	return baq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (baq *BanAppQuery) Unique(unique bool) *BanAppQuery {
	baq.unique = &unique
	return baq
}

// Order adds an order step to the query.
func (baq *BanAppQuery) Order(o ...OrderFunc) *BanAppQuery {
	baq.order = append(baq.order, o...)
	return baq
}

// First returns the first BanApp entity from the query.
// Returns a *NotFoundError when no BanApp was found.
func (baq *BanAppQuery) First(ctx context.Context) (*BanApp, error) {
	nodes, err := baq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{banapp.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (baq *BanAppQuery) FirstX(ctx context.Context) *BanApp {
	node, err := baq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BanApp ID from the query.
// Returns a *NotFoundError when no BanApp ID was found.
func (baq *BanAppQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = baq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{banapp.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (baq *BanAppQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := baq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BanApp entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BanApp entity is found.
// Returns a *NotFoundError when no BanApp entities are found.
func (baq *BanAppQuery) Only(ctx context.Context) (*BanApp, error) {
	nodes, err := baq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{banapp.Label}
	default:
		return nil, &NotSingularError{banapp.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (baq *BanAppQuery) OnlyX(ctx context.Context) *BanApp {
	node, err := baq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BanApp ID in the query.
// Returns a *NotSingularError when more than one BanApp ID is found.
// Returns a *NotFoundError when no entities are found.
func (baq *BanAppQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = baq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{banapp.Label}
	default:
		err = &NotSingularError{banapp.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (baq *BanAppQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := baq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BanApps.
func (baq *BanAppQuery) All(ctx context.Context) ([]*BanApp, error) {
	if err := baq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return baq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (baq *BanAppQuery) AllX(ctx context.Context) []*BanApp {
	nodes, err := baq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BanApp IDs.
func (baq *BanAppQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := baq.Select(banapp.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (baq *BanAppQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := baq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (baq *BanAppQuery) Count(ctx context.Context) (int, error) {
	if err := baq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return baq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (baq *BanAppQuery) CountX(ctx context.Context) int {
	count, err := baq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (baq *BanAppQuery) Exist(ctx context.Context) (bool, error) {
	if err := baq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return baq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (baq *BanAppQuery) ExistX(ctx context.Context) bool {
	exist, err := baq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BanAppQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (baq *BanAppQuery) Clone() *BanAppQuery {
	if baq == nil {
		return nil
	}
	return &BanAppQuery{
		config:     baq.config,
		limit:      baq.limit,
		offset:     baq.offset,
		order:      append([]OrderFunc{}, baq.order...),
		predicates: append([]predicate.BanApp{}, baq.predicates...),
		// clone intermediate query.
		sql:    baq.sql.Clone(),
		path:   baq.path,
		unique: baq.unique,
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
//	client.BanApp.Query().
//		GroupBy(banapp.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (baq *BanAppQuery) GroupBy(field string, fields ...string) *BanAppGroupBy {
	grbuild := &BanAppGroupBy{config: baq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := baq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return baq.sqlQuery(ctx), nil
	}
	grbuild.label = banapp.Label
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
//	client.BanApp.Query().
//		Select(banapp.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (baq *BanAppQuery) Select(fields ...string) *BanAppSelect {
	baq.fields = append(baq.fields, fields...)
	selbuild := &BanAppSelect{BanAppQuery: baq}
	selbuild.label = banapp.Label
	selbuild.flds, selbuild.scan = &baq.fields, selbuild.Scan
	return selbuild
}

func (baq *BanAppQuery) prepareQuery(ctx context.Context) error {
	for _, f := range baq.fields {
		if !banapp.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if baq.path != nil {
		prev, err := baq.path(ctx)
		if err != nil {
			return err
		}
		baq.sql = prev
	}
	if banapp.Policy == nil {
		return errors.New("ent: uninitialized banapp.Policy (forgotten import ent/runtime?)")
	}
	if err := banapp.Policy.EvalQuery(ctx, baq); err != nil {
		return err
	}
	return nil
}

func (baq *BanAppQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BanApp, error) {
	var (
		nodes = []*BanApp{}
		_spec = baq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*BanApp).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &BanApp{config: baq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(baq.modifiers) > 0 {
		_spec.Modifiers = baq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, baq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (baq *BanAppQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := baq.querySpec()
	if len(baq.modifiers) > 0 {
		_spec.Modifiers = baq.modifiers
	}
	_spec.Node.Columns = baq.fields
	if len(baq.fields) > 0 {
		_spec.Unique = baq.unique != nil && *baq.unique
	}
	return sqlgraph.CountNodes(ctx, baq.driver, _spec)
}

func (baq *BanAppQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := baq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (baq *BanAppQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   banapp.Table,
			Columns: banapp.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: banapp.FieldID,
			},
		},
		From:   baq.sql,
		Unique: true,
	}
	if unique := baq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := baq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, banapp.FieldID)
		for i := range fields {
			if fields[i] != banapp.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := baq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := baq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := baq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := baq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (baq *BanAppQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(baq.driver.Dialect())
	t1 := builder.Table(banapp.Table)
	columns := baq.fields
	if len(columns) == 0 {
		columns = banapp.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if baq.sql != nil {
		selector = baq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if baq.unique != nil && *baq.unique {
		selector.Distinct()
	}
	for _, m := range baq.modifiers {
		m(selector)
	}
	for _, p := range baq.predicates {
		p(selector)
	}
	for _, p := range baq.order {
		p(selector)
	}
	if offset := baq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := baq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (baq *BanAppQuery) ForUpdate(opts ...sql.LockOption) *BanAppQuery {
	if baq.driver.Dialect() == dialect.Postgres {
		baq.Unique(false)
	}
	baq.modifiers = append(baq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return baq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (baq *BanAppQuery) ForShare(opts ...sql.LockOption) *BanAppQuery {
	if baq.driver.Dialect() == dialect.Postgres {
		baq.Unique(false)
	}
	baq.modifiers = append(baq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return baq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (baq *BanAppQuery) Modify(modifiers ...func(s *sql.Selector)) *BanAppSelect {
	baq.modifiers = append(baq.modifiers, modifiers...)
	return baq.Select()
}

// BanAppGroupBy is the group-by builder for BanApp entities.
type BanAppGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bagb *BanAppGroupBy) Aggregate(fns ...AggregateFunc) *BanAppGroupBy {
	bagb.fns = append(bagb.fns, fns...)
	return bagb
}

// Scan applies the group-by query and scans the result into the given value.
func (bagb *BanAppGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bagb.path(ctx)
	if err != nil {
		return err
	}
	bagb.sql = query
	return bagb.sqlScan(ctx, v)
}

func (bagb *BanAppGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bagb.fields {
		if !banapp.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bagb *BanAppGroupBy) sqlQuery() *sql.Selector {
	selector := bagb.sql.Select()
	aggregation := make([]string, 0, len(bagb.fns))
	for _, fn := range bagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bagb.fields)+len(bagb.fns))
		for _, f := range bagb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bagb.fields...)...)
}

// BanAppSelect is the builder for selecting fields of BanApp entities.
type BanAppSelect struct {
	*BanAppQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bas *BanAppSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bas.prepareQuery(ctx); err != nil {
		return err
	}
	bas.sql = bas.BanAppQuery.sqlQuery(ctx)
	return bas.sqlScan(ctx, v)
}

func (bas *BanAppSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bas.sql.Query()
	if err := bas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (bas *BanAppSelect) Modify(modifiers ...func(s *sql.Selector)) *BanAppSelect {
	bas.modifiers = append(bas.modifiers, modifiers...)
	return bas
}
