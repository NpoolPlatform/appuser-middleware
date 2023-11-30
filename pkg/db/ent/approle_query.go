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
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/approle"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/predicate"
)

// AppRoleQuery is the builder for querying AppRole entities.
type AppRoleQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppRole
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppRoleQuery builder.
func (arq *AppRoleQuery) Where(ps ...predicate.AppRole) *AppRoleQuery {
	arq.predicates = append(arq.predicates, ps...)
	return arq
}

// Limit adds a limit step to the query.
func (arq *AppRoleQuery) Limit(limit int) *AppRoleQuery {
	arq.limit = &limit
	return arq
}

// Offset adds an offset step to the query.
func (arq *AppRoleQuery) Offset(offset int) *AppRoleQuery {
	arq.offset = &offset
	return arq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (arq *AppRoleQuery) Unique(unique bool) *AppRoleQuery {
	arq.unique = &unique
	return arq
}

// Order adds an order step to the query.
func (arq *AppRoleQuery) Order(o ...OrderFunc) *AppRoleQuery {
	arq.order = append(arq.order, o...)
	return arq
}

// First returns the first AppRole entity from the query.
// Returns a *NotFoundError when no AppRole was found.
func (arq *AppRoleQuery) First(ctx context.Context) (*AppRole, error) {
	nodes, err := arq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{approle.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (arq *AppRoleQuery) FirstX(ctx context.Context) *AppRole {
	node, err := arq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppRole ID from the query.
// Returns a *NotFoundError when no AppRole ID was found.
func (arq *AppRoleQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = arq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{approle.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (arq *AppRoleQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := arq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppRole entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppRole entity is found.
// Returns a *NotFoundError when no AppRole entities are found.
func (arq *AppRoleQuery) Only(ctx context.Context) (*AppRole, error) {
	nodes, err := arq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{approle.Label}
	default:
		return nil, &NotSingularError{approle.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (arq *AppRoleQuery) OnlyX(ctx context.Context) *AppRole {
	node, err := arq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppRole ID in the query.
// Returns a *NotSingularError when more than one AppRole ID is found.
// Returns a *NotFoundError when no entities are found.
func (arq *AppRoleQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = arq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{approle.Label}
	default:
		err = &NotSingularError{approle.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (arq *AppRoleQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := arq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppRoles.
func (arq *AppRoleQuery) All(ctx context.Context) ([]*AppRole, error) {
	if err := arq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return arq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (arq *AppRoleQuery) AllX(ctx context.Context) []*AppRole {
	nodes, err := arq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppRole IDs.
func (arq *AppRoleQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := arq.Select(approle.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (arq *AppRoleQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := arq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (arq *AppRoleQuery) Count(ctx context.Context) (int, error) {
	if err := arq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return arq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (arq *AppRoleQuery) CountX(ctx context.Context) int {
	count, err := arq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (arq *AppRoleQuery) Exist(ctx context.Context) (bool, error) {
	if err := arq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return arq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (arq *AppRoleQuery) ExistX(ctx context.Context) bool {
	exist, err := arq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppRoleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (arq *AppRoleQuery) Clone() *AppRoleQuery {
	if arq == nil {
		return nil
	}
	return &AppRoleQuery{
		config:     arq.config,
		limit:      arq.limit,
		offset:     arq.offset,
		order:      append([]OrderFunc{}, arq.order...),
		predicates: append([]predicate.AppRole{}, arq.predicates...),
		// clone intermediate query.
		sql:    arq.sql.Clone(),
		path:   arq.path,
		unique: arq.unique,
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
//	client.AppRole.Query().
//		GroupBy(approle.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (arq *AppRoleQuery) GroupBy(field string, fields ...string) *AppRoleGroupBy {
	grbuild := &AppRoleGroupBy{config: arq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := arq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return arq.sqlQuery(ctx), nil
	}
	grbuild.label = approle.Label
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
//	client.AppRole.Query().
//		Select(approle.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (arq *AppRoleQuery) Select(fields ...string) *AppRoleSelect {
	arq.fields = append(arq.fields, fields...)
	selbuild := &AppRoleSelect{AppRoleQuery: arq}
	selbuild.label = approle.Label
	selbuild.flds, selbuild.scan = &arq.fields, selbuild.Scan
	return selbuild
}

func (arq *AppRoleQuery) prepareQuery(ctx context.Context) error {
	for _, f := range arq.fields {
		if !approle.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if arq.path != nil {
		prev, err := arq.path(ctx)
		if err != nil {
			return err
		}
		arq.sql = prev
	}
	if approle.Policy == nil {
		return errors.New("ent: uninitialized approle.Policy (forgotten import ent/runtime?)")
	}
	if err := approle.Policy.EvalQuery(ctx, arq); err != nil {
		return err
	}
	return nil
}

func (arq *AppRoleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppRole, error) {
	var (
		nodes = []*AppRole{}
		_spec = arq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*AppRole).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &AppRole{config: arq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(arq.modifiers) > 0 {
		_spec.Modifiers = arq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, arq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (arq *AppRoleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := arq.querySpec()
	if len(arq.modifiers) > 0 {
		_spec.Modifiers = arq.modifiers
	}
	_spec.Node.Columns = arq.fields
	if len(arq.fields) > 0 {
		_spec.Unique = arq.unique != nil && *arq.unique
	}
	return sqlgraph.CountNodes(ctx, arq.driver, _spec)
}

func (arq *AppRoleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := arq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (arq *AppRoleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   approle.Table,
			Columns: approle.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: approle.FieldID,
			},
		},
		From:   arq.sql,
		Unique: true,
	}
	if unique := arq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := arq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, approle.FieldID)
		for i := range fields {
			if fields[i] != approle.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := arq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := arq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := arq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := arq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (arq *AppRoleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(arq.driver.Dialect())
	t1 := builder.Table(approle.Table)
	columns := arq.fields
	if len(columns) == 0 {
		columns = approle.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if arq.sql != nil {
		selector = arq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if arq.unique != nil && *arq.unique {
		selector.Distinct()
	}
	for _, m := range arq.modifiers {
		m(selector)
	}
	for _, p := range arq.predicates {
		p(selector)
	}
	for _, p := range arq.order {
		p(selector)
	}
	if offset := arq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := arq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (arq *AppRoleQuery) ForUpdate(opts ...sql.LockOption) *AppRoleQuery {
	if arq.driver.Dialect() == dialect.Postgres {
		arq.Unique(false)
	}
	arq.modifiers = append(arq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return arq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (arq *AppRoleQuery) ForShare(opts ...sql.LockOption) *AppRoleQuery {
	if arq.driver.Dialect() == dialect.Postgres {
		arq.Unique(false)
	}
	arq.modifiers = append(arq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return arq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (arq *AppRoleQuery) Modify(modifiers ...func(s *sql.Selector)) *AppRoleSelect {
	arq.modifiers = append(arq.modifiers, modifiers...)
	return arq.Select()
}

// AppRoleGroupBy is the group-by builder for AppRole entities.
type AppRoleGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (argb *AppRoleGroupBy) Aggregate(fns ...AggregateFunc) *AppRoleGroupBy {
	argb.fns = append(argb.fns, fns...)
	return argb
}

// Scan applies the group-by query and scans the result into the given value.
func (argb *AppRoleGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := argb.path(ctx)
	if err != nil {
		return err
	}
	argb.sql = query
	return argb.sqlScan(ctx, v)
}

func (argb *AppRoleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range argb.fields {
		if !approle.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := argb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := argb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (argb *AppRoleGroupBy) sqlQuery() *sql.Selector {
	selector := argb.sql.Select()
	aggregation := make([]string, 0, len(argb.fns))
	for _, fn := range argb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(argb.fields)+len(argb.fns))
		for _, f := range argb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(argb.fields...)...)
}

// AppRoleSelect is the builder for selecting fields of AppRole entities.
type AppRoleSelect struct {
	*AppRoleQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ars *AppRoleSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ars.prepareQuery(ctx); err != nil {
		return err
	}
	ars.sql = ars.AppRoleQuery.sqlQuery(ctx)
	return ars.sqlScan(ctx, v)
}

func (ars *AppRoleSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ars.sql.Query()
	if err := ars.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ars *AppRoleSelect) Modify(modifiers ...func(s *sql.Selector)) *AppRoleSelect {
	ars.modifiers = append(ars.modifiers, modifiers...)
	return ars
}
