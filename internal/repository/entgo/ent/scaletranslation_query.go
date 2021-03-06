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
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/scale"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/scaletranslation"
)

// ScaleTranslationQuery is the builder for querying ScaleTranslation entities.
type ScaleTranslationQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ScaleTranslation
	// eager-loading edges.
	withScale *ScaleQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ScaleTranslationQuery builder.
func (stq *ScaleTranslationQuery) Where(ps ...predicate.ScaleTranslation) *ScaleTranslationQuery {
	stq.predicates = append(stq.predicates, ps...)
	return stq
}

// Limit adds a limit step to the query.
func (stq *ScaleTranslationQuery) Limit(limit int) *ScaleTranslationQuery {
	stq.limit = &limit
	return stq
}

// Offset adds an offset step to the query.
func (stq *ScaleTranslationQuery) Offset(offset int) *ScaleTranslationQuery {
	stq.offset = &offset
	return stq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (stq *ScaleTranslationQuery) Unique(unique bool) *ScaleTranslationQuery {
	stq.unique = &unique
	return stq
}

// Order adds an order step to the query.
func (stq *ScaleTranslationQuery) Order(o ...OrderFunc) *ScaleTranslationQuery {
	stq.order = append(stq.order, o...)
	return stq
}

// QueryScale chains the current query on the "scale" edge.
func (stq *ScaleTranslationQuery) QueryScale() *ScaleQuery {
	query := &ScaleQuery{config: stq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := stq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := stq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scaletranslation.Table, scaletranslation.FieldID, selector),
			sqlgraph.To(scale.Table, scale.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, scaletranslation.ScaleTable, scaletranslation.ScaleColumn),
		)
		fromU = sqlgraph.SetNeighbors(stq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ScaleTranslation entity from the query.
// Returns a *NotFoundError when no ScaleTranslation was found.
func (stq *ScaleTranslationQuery) First(ctx context.Context) (*ScaleTranslation, error) {
	nodes, err := stq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{scaletranslation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (stq *ScaleTranslationQuery) FirstX(ctx context.Context) *ScaleTranslation {
	node, err := stq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ScaleTranslation ID from the query.
// Returns a *NotFoundError when no ScaleTranslation ID was found.
func (stq *ScaleTranslationQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = stq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{scaletranslation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (stq *ScaleTranslationQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := stq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ScaleTranslation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ScaleTranslation entity is found.
// Returns a *NotFoundError when no ScaleTranslation entities are found.
func (stq *ScaleTranslationQuery) Only(ctx context.Context) (*ScaleTranslation, error) {
	nodes, err := stq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{scaletranslation.Label}
	default:
		return nil, &NotSingularError{scaletranslation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (stq *ScaleTranslationQuery) OnlyX(ctx context.Context) *ScaleTranslation {
	node, err := stq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ScaleTranslation ID in the query.
// Returns a *NotSingularError when more than one ScaleTranslation ID is found.
// Returns a *NotFoundError when no entities are found.
func (stq *ScaleTranslationQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = stq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{scaletranslation.Label}
	default:
		err = &NotSingularError{scaletranslation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (stq *ScaleTranslationQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := stq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ScaleTranslations.
func (stq *ScaleTranslationQuery) All(ctx context.Context) ([]*ScaleTranslation, error) {
	if err := stq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return stq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (stq *ScaleTranslationQuery) AllX(ctx context.Context) []*ScaleTranslation {
	nodes, err := stq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ScaleTranslation IDs.
func (stq *ScaleTranslationQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := stq.Select(scaletranslation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (stq *ScaleTranslationQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := stq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (stq *ScaleTranslationQuery) Count(ctx context.Context) (int, error) {
	if err := stq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return stq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (stq *ScaleTranslationQuery) CountX(ctx context.Context) int {
	count, err := stq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (stq *ScaleTranslationQuery) Exist(ctx context.Context) (bool, error) {
	if err := stq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return stq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (stq *ScaleTranslationQuery) ExistX(ctx context.Context) bool {
	exist, err := stq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ScaleTranslationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (stq *ScaleTranslationQuery) Clone() *ScaleTranslationQuery {
	if stq == nil {
		return nil
	}
	return &ScaleTranslationQuery{
		config:     stq.config,
		limit:      stq.limit,
		offset:     stq.offset,
		order:      append([]OrderFunc{}, stq.order...),
		predicates: append([]predicate.ScaleTranslation{}, stq.predicates...),
		withScale:  stq.withScale.Clone(),
		// clone intermediate query.
		sql:    stq.sql.Clone(),
		path:   stq.path,
		unique: stq.unique,
	}
}

// WithScale tells the query-builder to eager-load the nodes that are connected to
// the "scale" edge. The optional arguments are used to configure the query builder of the edge.
func (stq *ScaleTranslationQuery) WithScale(opts ...func(*ScaleQuery)) *ScaleTranslationQuery {
	query := &ScaleQuery{config: stq.config}
	for _, opt := range opts {
		opt(query)
	}
	stq.withScale = query
	return stq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Locale scaletranslation.Locale `json:"locale,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ScaleTranslation.Query().
//		GroupBy(scaletranslation.FieldLocale).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (stq *ScaleTranslationQuery) GroupBy(field string, fields ...string) *ScaleTranslationGroupBy {
	grbuild := &ScaleTranslationGroupBy{config: stq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := stq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return stq.sqlQuery(ctx), nil
	}
	grbuild.label = scaletranslation.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Locale scaletranslation.Locale `json:"locale,omitempty"`
//	}
//
//	client.ScaleTranslation.Query().
//		Select(scaletranslation.FieldLocale).
//		Scan(ctx, &v)
//
func (stq *ScaleTranslationQuery) Select(fields ...string) *ScaleTranslationSelect {
	stq.fields = append(stq.fields, fields...)
	selbuild := &ScaleTranslationSelect{ScaleTranslationQuery: stq}
	selbuild.label = scaletranslation.Label
	selbuild.flds, selbuild.scan = &stq.fields, selbuild.Scan
	return selbuild
}

func (stq *ScaleTranslationQuery) prepareQuery(ctx context.Context) error {
	for _, f := range stq.fields {
		if !scaletranslation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if stq.path != nil {
		prev, err := stq.path(ctx)
		if err != nil {
			return err
		}
		stq.sql = prev
	}
	return nil
}

func (stq *ScaleTranslationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ScaleTranslation, error) {
	var (
		nodes       = []*ScaleTranslation{}
		withFKs     = stq.withFKs
		_spec       = stq.querySpec()
		loadedTypes = [1]bool{
			stq.withScale != nil,
		}
	)
	if stq.withScale != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, scaletranslation.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*ScaleTranslation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &ScaleTranslation{config: stq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, stq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := stq.withScale; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*ScaleTranslation)
		for i := range nodes {
			if nodes[i].scale_translations == nil {
				continue
			}
			fk := *nodes[i].scale_translations
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(scale.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "scale_translations" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Scale = n
			}
		}
	}

	return nodes, nil
}

func (stq *ScaleTranslationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := stq.querySpec()
	_spec.Node.Columns = stq.fields
	if len(stq.fields) > 0 {
		_spec.Unique = stq.unique != nil && *stq.unique
	}
	return sqlgraph.CountNodes(ctx, stq.driver, _spec)
}

func (stq *ScaleTranslationQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := stq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (stq *ScaleTranslationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   scaletranslation.Table,
			Columns: scaletranslation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: scaletranslation.FieldID,
			},
		},
		From:   stq.sql,
		Unique: true,
	}
	if unique := stq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := stq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, scaletranslation.FieldID)
		for i := range fields {
			if fields[i] != scaletranslation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := stq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := stq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := stq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := stq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (stq *ScaleTranslationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(stq.driver.Dialect())
	t1 := builder.Table(scaletranslation.Table)
	columns := stq.fields
	if len(columns) == 0 {
		columns = scaletranslation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if stq.sql != nil {
		selector = stq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if stq.unique != nil && *stq.unique {
		selector.Distinct()
	}
	for _, p := range stq.predicates {
		p(selector)
	}
	for _, p := range stq.order {
		p(selector)
	}
	if offset := stq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := stq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ScaleTranslationGroupBy is the group-by builder for ScaleTranslation entities.
type ScaleTranslationGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (stgb *ScaleTranslationGroupBy) Aggregate(fns ...AggregateFunc) *ScaleTranslationGroupBy {
	stgb.fns = append(stgb.fns, fns...)
	return stgb
}

// Scan applies the group-by query and scans the result into the given value.
func (stgb *ScaleTranslationGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := stgb.path(ctx)
	if err != nil {
		return err
	}
	stgb.sql = query
	return stgb.sqlScan(ctx, v)
}

func (stgb *ScaleTranslationGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range stgb.fields {
		if !scaletranslation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := stgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := stgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (stgb *ScaleTranslationGroupBy) sqlQuery() *sql.Selector {
	selector := stgb.sql.Select()
	aggregation := make([]string, 0, len(stgb.fns))
	for _, fn := range stgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(stgb.fields)+len(stgb.fns))
		for _, f := range stgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(stgb.fields...)...)
}

// ScaleTranslationSelect is the builder for selecting fields of ScaleTranslation entities.
type ScaleTranslationSelect struct {
	*ScaleTranslationQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sts *ScaleTranslationSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sts.prepareQuery(ctx); err != nil {
		return err
	}
	sts.sql = sts.ScaleTranslationQuery.sqlQuery(ctx)
	return sts.sqlScan(ctx, v)
}

func (sts *ScaleTranslationSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sts.sql.Query()
	if err := sts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
