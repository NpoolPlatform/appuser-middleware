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
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appuserthirdparty"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/predicate"
)

// AppUserThirdPartyQuery is the builder for querying AppUserThirdParty entities.
type AppUserThirdPartyQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppUserThirdParty
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppUserThirdPartyQuery builder.
func (autpq *AppUserThirdPartyQuery) Where(ps ...predicate.AppUserThirdParty) *AppUserThirdPartyQuery {
	autpq.predicates = append(autpq.predicates, ps...)
	return autpq
}

// Limit adds a limit step to the query.
func (autpq *AppUserThirdPartyQuery) Limit(limit int) *AppUserThirdPartyQuery {
	autpq.limit = &limit
	return autpq
}

// Offset adds an offset step to the query.
func (autpq *AppUserThirdPartyQuery) Offset(offset int) *AppUserThirdPartyQuery {
	autpq.offset = &offset
	return autpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (autpq *AppUserThirdPartyQuery) Unique(unique bool) *AppUserThirdPartyQuery {
	autpq.unique = &unique
	return autpq
}

// Order adds an order step to the query.
func (autpq *AppUserThirdPartyQuery) Order(o ...OrderFunc) *AppUserThirdPartyQuery {
	autpq.order = append(autpq.order, o...)
	return autpq
}

// First returns the first AppUserThirdParty entity from the query.
// Returns a *NotFoundError when no AppUserThirdParty was found.
func (autpq *AppUserThirdPartyQuery) First(ctx context.Context) (*AppUserThirdParty, error) {
	nodes, err := autpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appuserthirdparty.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (autpq *AppUserThirdPartyQuery) FirstX(ctx context.Context) *AppUserThirdParty {
	node, err := autpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppUserThirdParty ID from the query.
// Returns a *NotFoundError when no AppUserThirdParty ID was found.
func (autpq *AppUserThirdPartyQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = autpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appuserthirdparty.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (autpq *AppUserThirdPartyQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := autpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppUserThirdParty entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppUserThirdParty entity is found.
// Returns a *NotFoundError when no AppUserThirdParty entities are found.
func (autpq *AppUserThirdPartyQuery) Only(ctx context.Context) (*AppUserThirdParty, error) {
	nodes, err := autpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appuserthirdparty.Label}
	default:
		return nil, &NotSingularError{appuserthirdparty.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (autpq *AppUserThirdPartyQuery) OnlyX(ctx context.Context) *AppUserThirdParty {
	node, err := autpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppUserThirdParty ID in the query.
// Returns a *NotSingularError when more than one AppUserThirdParty ID is found.
// Returns a *NotFoundError when no entities are found.
func (autpq *AppUserThirdPartyQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = autpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appuserthirdparty.Label}
	default:
		err = &NotSingularError{appuserthirdparty.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (autpq *AppUserThirdPartyQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := autpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppUserThirdParties.
func (autpq *AppUserThirdPartyQuery) All(ctx context.Context) ([]*AppUserThirdParty, error) {
	if err := autpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return autpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (autpq *AppUserThirdPartyQuery) AllX(ctx context.Context) []*AppUserThirdParty {
	nodes, err := autpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppUserThirdParty IDs.
func (autpq *AppUserThirdPartyQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := autpq.Select(appuserthirdparty.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (autpq *AppUserThirdPartyQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := autpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (autpq *AppUserThirdPartyQuery) Count(ctx context.Context) (int, error) {
	if err := autpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return autpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (autpq *AppUserThirdPartyQuery) CountX(ctx context.Context) int {
	count, err := autpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (autpq *AppUserThirdPartyQuery) Exist(ctx context.Context) (bool, error) {
	if err := autpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return autpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (autpq *AppUserThirdPartyQuery) ExistX(ctx context.Context) bool {
	exist, err := autpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppUserThirdPartyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (autpq *AppUserThirdPartyQuery) Clone() *AppUserThirdPartyQuery {
	if autpq == nil {
		return nil
	}
	return &AppUserThirdPartyQuery{
		config:     autpq.config,
		limit:      autpq.limit,
		offset:     autpq.offset,
		order:      append([]OrderFunc{}, autpq.order...),
		predicates: append([]predicate.AppUserThirdParty{}, autpq.predicates...),
		// clone intermediate query.
		sql:    autpq.sql.Clone(),
		path:   autpq.path,
		unique: autpq.unique,
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
//	client.AppUserThirdParty.Query().
//		GroupBy(appuserthirdparty.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (autpq *AppUserThirdPartyQuery) GroupBy(field string, fields ...string) *AppUserThirdPartyGroupBy {
	grbuild := &AppUserThirdPartyGroupBy{config: autpq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := autpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return autpq.sqlQuery(ctx), nil
	}
	grbuild.label = appuserthirdparty.Label
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
//	client.AppUserThirdParty.Query().
//		Select(appuserthirdparty.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (autpq *AppUserThirdPartyQuery) Select(fields ...string) *AppUserThirdPartySelect {
	autpq.fields = append(autpq.fields, fields...)
	selbuild := &AppUserThirdPartySelect{AppUserThirdPartyQuery: autpq}
	selbuild.label = appuserthirdparty.Label
	selbuild.flds, selbuild.scan = &autpq.fields, selbuild.Scan
	return selbuild
}

func (autpq *AppUserThirdPartyQuery) prepareQuery(ctx context.Context) error {
	for _, f := range autpq.fields {
		if !appuserthirdparty.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if autpq.path != nil {
		prev, err := autpq.path(ctx)
		if err != nil {
			return err
		}
		autpq.sql = prev
	}
	if appuserthirdparty.Policy == nil {
		return errors.New("ent: uninitialized appuserthirdparty.Policy (forgotten import ent/runtime?)")
	}
	if err := appuserthirdparty.Policy.EvalQuery(ctx, autpq); err != nil {
		return err
	}
	return nil
}

func (autpq *AppUserThirdPartyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppUserThirdParty, error) {
	var (
		nodes = []*AppUserThirdParty{}
		_spec = autpq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*AppUserThirdParty).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &AppUserThirdParty{config: autpq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(autpq.modifiers) > 0 {
		_spec.Modifiers = autpq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, autpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (autpq *AppUserThirdPartyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := autpq.querySpec()
	if len(autpq.modifiers) > 0 {
		_spec.Modifiers = autpq.modifiers
	}
	_spec.Node.Columns = autpq.fields
	if len(autpq.fields) > 0 {
		_spec.Unique = autpq.unique != nil && *autpq.unique
	}
	return sqlgraph.CountNodes(ctx, autpq.driver, _spec)
}

func (autpq *AppUserThirdPartyQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := autpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (autpq *AppUserThirdPartyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appuserthirdparty.Table,
			Columns: appuserthirdparty.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appuserthirdparty.FieldID,
			},
		},
		From:   autpq.sql,
		Unique: true,
	}
	if unique := autpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := autpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appuserthirdparty.FieldID)
		for i := range fields {
			if fields[i] != appuserthirdparty.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := autpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := autpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := autpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := autpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (autpq *AppUserThirdPartyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(autpq.driver.Dialect())
	t1 := builder.Table(appuserthirdparty.Table)
	columns := autpq.fields
	if len(columns) == 0 {
		columns = appuserthirdparty.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if autpq.sql != nil {
		selector = autpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if autpq.unique != nil && *autpq.unique {
		selector.Distinct()
	}
	for _, m := range autpq.modifiers {
		m(selector)
	}
	for _, p := range autpq.predicates {
		p(selector)
	}
	for _, p := range autpq.order {
		p(selector)
	}
	if offset := autpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := autpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (autpq *AppUserThirdPartyQuery) ForUpdate(opts ...sql.LockOption) *AppUserThirdPartyQuery {
	if autpq.driver.Dialect() == dialect.Postgres {
		autpq.Unique(false)
	}
	autpq.modifiers = append(autpq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return autpq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (autpq *AppUserThirdPartyQuery) ForShare(opts ...sql.LockOption) *AppUserThirdPartyQuery {
	if autpq.driver.Dialect() == dialect.Postgres {
		autpq.Unique(false)
	}
	autpq.modifiers = append(autpq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return autpq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (autpq *AppUserThirdPartyQuery) Modify(modifiers ...func(s *sql.Selector)) *AppUserThirdPartySelect {
	autpq.modifiers = append(autpq.modifiers, modifiers...)
	return autpq.Select()
}

// AppUserThirdPartyGroupBy is the group-by builder for AppUserThirdParty entities.
type AppUserThirdPartyGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (autpgb *AppUserThirdPartyGroupBy) Aggregate(fns ...AggregateFunc) *AppUserThirdPartyGroupBy {
	autpgb.fns = append(autpgb.fns, fns...)
	return autpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (autpgb *AppUserThirdPartyGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := autpgb.path(ctx)
	if err != nil {
		return err
	}
	autpgb.sql = query
	return autpgb.sqlScan(ctx, v)
}

func (autpgb *AppUserThirdPartyGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range autpgb.fields {
		if !appuserthirdparty.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := autpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := autpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (autpgb *AppUserThirdPartyGroupBy) sqlQuery() *sql.Selector {
	selector := autpgb.sql.Select()
	aggregation := make([]string, 0, len(autpgb.fns))
	for _, fn := range autpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(autpgb.fields)+len(autpgb.fns))
		for _, f := range autpgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(autpgb.fields...)...)
}

// AppUserThirdPartySelect is the builder for selecting fields of AppUserThirdParty entities.
type AppUserThirdPartySelect struct {
	*AppUserThirdPartyQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (autps *AppUserThirdPartySelect) Scan(ctx context.Context, v interface{}) error {
	if err := autps.prepareQuery(ctx); err != nil {
		return err
	}
	autps.sql = autps.AppUserThirdPartyQuery.sqlQuery(ctx)
	return autps.sqlScan(ctx, v)
}

func (autps *AppUserThirdPartySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := autps.sql.Query()
	if err := autps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (autps *AppUserThirdPartySelect) Modify(modifiers ...func(s *sql.Selector)) *AppUserThirdPartySelect {
	autps.modifiers = append(autps.modifiers, modifiers...)
	return autps
}
