// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/predicate"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/test"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/testtranslation"
)

// TestTranslationQuery is the builder for querying TestTranslation entities.
type TestTranslationQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TestTranslation
	// eager-loading edges.
	withTest *TestQuery
	withFKs  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TestTranslationQuery builder.
func (ttq *TestTranslationQuery) Where(ps ...predicate.TestTranslation) *TestTranslationQuery {
	ttq.predicates = append(ttq.predicates, ps...)
	return ttq
}

// Limit adds a limit step to the query.
func (ttq *TestTranslationQuery) Limit(limit int) *TestTranslationQuery {
	ttq.limit = &limit
	return ttq
}

// Offset adds an offset step to the query.
func (ttq *TestTranslationQuery) Offset(offset int) *TestTranslationQuery {
	ttq.offset = &offset
	return ttq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ttq *TestTranslationQuery) Unique(unique bool) *TestTranslationQuery {
	ttq.unique = &unique
	return ttq
}

// Order adds an order step to the query.
func (ttq *TestTranslationQuery) Order(o ...OrderFunc) *TestTranslationQuery {
	ttq.order = append(ttq.order, o...)
	return ttq
}

// QueryTest chains the current query on the "test" edge.
func (ttq *TestTranslationQuery) QueryTest() *TestQuery {
	query := &TestQuery{config: ttq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ttq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(testtranslation.Table, testtranslation.FieldID, selector),
			sqlgraph.To(test.Table, test.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, testtranslation.TestTable, testtranslation.TestColumn),
		)
		fromU = sqlgraph.SetNeighbors(ttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TestTranslation entity from the query.
// Returns a *NotFoundError when no TestTranslation was found.
func (ttq *TestTranslationQuery) First(ctx context.Context) (*TestTranslation, error) {
	nodes, err := ttq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{testtranslation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ttq *TestTranslationQuery) FirstX(ctx context.Context) *TestTranslation {
	node, err := ttq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TestTranslation ID from the query.
// Returns a *NotFoundError when no TestTranslation ID was found.
func (ttq *TestTranslationQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ttq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{testtranslation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ttq *TestTranslationQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ttq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TestTranslation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TestTranslation entity is found.
// Returns a *NotFoundError when no TestTranslation entities are found.
func (ttq *TestTranslationQuery) Only(ctx context.Context) (*TestTranslation, error) {
	nodes, err := ttq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{testtranslation.Label}
	default:
		return nil, &NotSingularError{testtranslation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ttq *TestTranslationQuery) OnlyX(ctx context.Context) *TestTranslation {
	node, err := ttq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TestTranslation ID in the query.
// Returns a *NotSingularError when more than one TestTranslation ID is found.
// Returns a *NotFoundError when no entities are found.
func (ttq *TestTranslationQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ttq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{testtranslation.Label}
	default:
		err = &NotSingularError{testtranslation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ttq *TestTranslationQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ttq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TestTranslations.
func (ttq *TestTranslationQuery) All(ctx context.Context) ([]*TestTranslation, error) {
	if err := ttq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ttq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ttq *TestTranslationQuery) AllX(ctx context.Context) []*TestTranslation {
	nodes, err := ttq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TestTranslation IDs.
func (ttq *TestTranslationQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := ttq.Select(testtranslation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ttq *TestTranslationQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ttq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ttq *TestTranslationQuery) Count(ctx context.Context) (int, error) {
	if err := ttq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ttq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ttq *TestTranslationQuery) CountX(ctx context.Context) int {
	count, err := ttq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ttq *TestTranslationQuery) Exist(ctx context.Context) (bool, error) {
	if err := ttq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ttq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ttq *TestTranslationQuery) ExistX(ctx context.Context) bool {
	exist, err := ttq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TestTranslationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ttq *TestTranslationQuery) Clone() *TestTranslationQuery {
	if ttq == nil {
		return nil
	}
	return &TestTranslationQuery{
		config:     ttq.config,
		limit:      ttq.limit,
		offset:     ttq.offset,
		order:      append([]OrderFunc{}, ttq.order...),
		predicates: append([]predicate.TestTranslation{}, ttq.predicates...),
		withTest:   ttq.withTest.Clone(),
		// clone intermediate query.
		sql:    ttq.sql.Clone(),
		path:   ttq.path,
		unique: ttq.unique,
	}
}

// WithTest tells the query-builder to eager-load the nodes that are connected to
// the "test" edge. The optional arguments are used to configure the query builder of the edge.
func (ttq *TestTranslationQuery) WithTest(opts ...func(*TestQuery)) *TestTranslationQuery {
	query := &TestQuery{config: ttq.config}
	for _, opt := range opts {
		opt(query)
	}
	ttq.withTest = query
	return ttq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Locale testtranslation.Locale `json:"locale,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TestTranslation.Query().
//		GroupBy(testtranslation.FieldLocale).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ttq *TestTranslationQuery) GroupBy(field string, fields ...string) *TestTranslationGroupBy {
	grbuild := &TestTranslationGroupBy{config: ttq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ttq.sqlQuery(ctx), nil
	}
	grbuild.label = testtranslation.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Locale testtranslation.Locale `json:"locale,omitempty"`
//	}
//
//	client.TestTranslation.Query().
//		Select(testtranslation.FieldLocale).
//		Scan(ctx, &v)
//
func (ttq *TestTranslationQuery) Select(fields ...string) *TestTranslationSelect {
	ttq.fields = append(ttq.fields, fields...)
	selbuild := &TestTranslationSelect{TestTranslationQuery: ttq}
	selbuild.label = testtranslation.Label
	selbuild.flds, selbuild.scan = &ttq.fields, selbuild.Scan
	return selbuild
}

func (ttq *TestTranslationQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ttq.fields {
		if !testtranslation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ttq.path != nil {
		prev, err := ttq.path(ctx)
		if err != nil {
			return err
		}
		ttq.sql = prev
	}
	return nil
}

func (ttq *TestTranslationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TestTranslation, error) {
	var (
		nodes       = []*TestTranslation{}
		withFKs     = ttq.withFKs
		_spec       = ttq.querySpec()
		loadedTypes = [1]bool{
			ttq.withTest != nil,
		}
	)
	if ttq.withTest != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, testtranslation.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*TestTranslation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &TestTranslation{config: ttq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ttq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ttq.withTest; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*TestTranslation)
		for i := range nodes {
			if nodes[i].test_translations == nil {
				continue
			}
			fk := *nodes[i].test_translations
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(test.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "test_translations" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Test = n
			}
		}
	}

	return nodes, nil
}

func (ttq *TestTranslationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ttq.querySpec()
	_spec.Node.Columns = ttq.fields
	if len(ttq.fields) > 0 {
		_spec.Unique = ttq.unique != nil && *ttq.unique
	}
	return sqlgraph.CountNodes(ctx, ttq.driver, _spec)
}

func (ttq *TestTranslationQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ttq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ttq *TestTranslationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   testtranslation.Table,
			Columns: testtranslation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: testtranslation.FieldID,
			},
		},
		From:   ttq.sql,
		Unique: true,
	}
	if unique := ttq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ttq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, testtranslation.FieldID)
		for i := range fields {
			if fields[i] != testtranslation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ttq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ttq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ttq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ttq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ttq *TestTranslationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ttq.driver.Dialect())
	t1 := builder.Table(testtranslation.Table)
	columns := ttq.fields
	if len(columns) == 0 {
		columns = testtranslation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ttq.sql != nil {
		selector = ttq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ttq.unique != nil && *ttq.unique {
		selector.Distinct()
	}
	for _, p := range ttq.predicates {
		p(selector)
	}
	for _, p := range ttq.order {
		p(selector)
	}
	if offset := ttq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ttq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TestTranslationGroupBy is the group-by builder for TestTranslation entities.
type TestTranslationGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ttgb *TestTranslationGroupBy) Aggregate(fns ...AggregateFunc) *TestTranslationGroupBy {
	ttgb.fns = append(ttgb.fns, fns...)
	return ttgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ttgb *TestTranslationGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ttgb.path(ctx)
	if err != nil {
		return err
	}
	ttgb.sql = query
	return ttgb.sqlScan(ctx, v)
}

func (ttgb *TestTranslationGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ttgb.fields {
		if !testtranslation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ttgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ttgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ttgb *TestTranslationGroupBy) sqlQuery() *sql.Selector {
	selector := ttgb.sql.Select()
	aggregation := make([]string, 0, len(ttgb.fns))
	for _, fn := range ttgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ttgb.fields)+len(ttgb.fns))
		for _, f := range ttgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ttgb.fields...)...)
}

// TestTranslationSelect is the builder for selecting fields of TestTranslation entities.
type TestTranslationSelect struct {
	*TestTranslationQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tts *TestTranslationSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tts.prepareQuery(ctx); err != nil {
		return err
	}
	tts.sql = tts.TestTranslationQuery.sqlQuery(ctx)
	return tts.sqlScan(ctx, v)
}

func (tts *TestTranslationSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tts.sql.Query()
	if err := tts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
