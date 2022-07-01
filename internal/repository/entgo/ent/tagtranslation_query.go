// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/tag"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/tagtranslation"
	"github.com/google/uuid"
)

// TagTranslationQuery is the builder for querying TagTranslation entities.
type TagTranslationQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TagTranslation
	// eager-loading edges.
	withTag *TagQuery
	withFKs bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TagTranslationQuery builder.
func (ttq *TagTranslationQuery) Where(ps ...predicate.TagTranslation) *TagTranslationQuery {
	ttq.predicates = append(ttq.predicates, ps...)
	return ttq
}

// Limit adds a limit step to the query.
func (ttq *TagTranslationQuery) Limit(limit int) *TagTranslationQuery {
	ttq.limit = &limit
	return ttq
}

// Offset adds an offset step to the query.
func (ttq *TagTranslationQuery) Offset(offset int) *TagTranslationQuery {
	ttq.offset = &offset
	return ttq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ttq *TagTranslationQuery) Unique(unique bool) *TagTranslationQuery {
	ttq.unique = &unique
	return ttq
}

// Order adds an order step to the query.
func (ttq *TagTranslationQuery) Order(o ...OrderFunc) *TagTranslationQuery {
	ttq.order = append(ttq.order, o...)
	return ttq
}

// QueryTag chains the current query on the "tag" edge.
func (ttq *TagTranslationQuery) QueryTag() *TagQuery {
	query := &TagQuery{config: ttq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ttq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tagtranslation.Table, tagtranslation.FieldID, selector),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, tagtranslation.TagTable, tagtranslation.TagColumn),
		)
		fromU = sqlgraph.SetNeighbors(ttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TagTranslation entity from the query.
// Returns a *NotFoundError when no TagTranslation was found.
func (ttq *TagTranslationQuery) First(ctx context.Context) (*TagTranslation, error) {
	nodes, err := ttq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tagtranslation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ttq *TagTranslationQuery) FirstX(ctx context.Context) *TagTranslation {
	node, err := ttq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TagTranslation ID from the query.
// Returns a *NotFoundError when no TagTranslation ID was found.
func (ttq *TagTranslationQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ttq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tagtranslation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ttq *TagTranslationQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ttq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TagTranslation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TagTranslation entity is found.
// Returns a *NotFoundError when no TagTranslation entities are found.
func (ttq *TagTranslationQuery) Only(ctx context.Context) (*TagTranslation, error) {
	nodes, err := ttq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tagtranslation.Label}
	default:
		return nil, &NotSingularError{tagtranslation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ttq *TagTranslationQuery) OnlyX(ctx context.Context) *TagTranslation {
	node, err := ttq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TagTranslation ID in the query.
// Returns a *NotSingularError when more than one TagTranslation ID is found.
// Returns a *NotFoundError when no entities are found.
func (ttq *TagTranslationQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ttq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tagtranslation.Label}
	default:
		err = &NotSingularError{tagtranslation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ttq *TagTranslationQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ttq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TagTranslations.
func (ttq *TagTranslationQuery) All(ctx context.Context) ([]*TagTranslation, error) {
	if err := ttq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ttq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ttq *TagTranslationQuery) AllX(ctx context.Context) []*TagTranslation {
	nodes, err := ttq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TagTranslation IDs.
func (ttq *TagTranslationQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := ttq.Select(tagtranslation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ttq *TagTranslationQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ttq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ttq *TagTranslationQuery) Count(ctx context.Context) (int, error) {
	if err := ttq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ttq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ttq *TagTranslationQuery) CountX(ctx context.Context) int {
	count, err := ttq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ttq *TagTranslationQuery) Exist(ctx context.Context) (bool, error) {
	if err := ttq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ttq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ttq *TagTranslationQuery) ExistX(ctx context.Context) bool {
	exist, err := ttq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TagTranslationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ttq *TagTranslationQuery) Clone() *TagTranslationQuery {
	if ttq == nil {
		return nil
	}
	return &TagTranslationQuery{
		config:     ttq.config,
		limit:      ttq.limit,
		offset:     ttq.offset,
		order:      append([]OrderFunc{}, ttq.order...),
		predicates: append([]predicate.TagTranslation{}, ttq.predicates...),
		withTag:    ttq.withTag.Clone(),
		// clone intermediate query.
		sql:    ttq.sql.Clone(),
		path:   ttq.path,
		unique: ttq.unique,
	}
}

// WithTag tells the query-builder to eager-load the nodes that are connected to
// the "tag" edge. The optional arguments are used to configure the query builder of the edge.
func (ttq *TagTranslationQuery) WithTag(opts ...func(*TagQuery)) *TagTranslationQuery {
	query := &TagQuery{config: ttq.config}
	for _, opt := range opts {
		opt(query)
	}
	ttq.withTag = query
	return ttq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Locale tagtranslation.Locale `json:"locale,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TagTranslation.Query().
//		GroupBy(tagtranslation.FieldLocale).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ttq *TagTranslationQuery) GroupBy(field string, fields ...string) *TagTranslationGroupBy {
	grbuild := &TagTranslationGroupBy{config: ttq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ttq.sqlQuery(ctx), nil
	}
	grbuild.label = tagtranslation.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Locale tagtranslation.Locale `json:"locale,omitempty"`
//	}
//
//	client.TagTranslation.Query().
//		Select(tagtranslation.FieldLocale).
//		Scan(ctx, &v)
//
func (ttq *TagTranslationQuery) Select(fields ...string) *TagTranslationSelect {
	ttq.fields = append(ttq.fields, fields...)
	selbuild := &TagTranslationSelect{TagTranslationQuery: ttq}
	selbuild.label = tagtranslation.Label
	selbuild.flds, selbuild.scan = &ttq.fields, selbuild.Scan
	return selbuild
}

func (ttq *TagTranslationQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ttq.fields {
		if !tagtranslation.ValidColumn(f) {
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

func (ttq *TagTranslationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TagTranslation, error) {
	var (
		nodes       = []*TagTranslation{}
		withFKs     = ttq.withFKs
		_spec       = ttq.querySpec()
		loadedTypes = [1]bool{
			ttq.withTag != nil,
		}
	)
	if ttq.withTag != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, tagtranslation.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*TagTranslation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &TagTranslation{config: ttq.config}
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

	if query := ttq.withTag; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*TagTranslation)
		for i := range nodes {
			if nodes[i].tag_translations == nil {
				continue
			}
			fk := *nodes[i].tag_translations
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(tag.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "tag_translations" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Tag = n
			}
		}
	}

	return nodes, nil
}

func (ttq *TagTranslationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ttq.querySpec()
	_spec.Node.Columns = ttq.fields
	if len(ttq.fields) > 0 {
		_spec.Unique = ttq.unique != nil && *ttq.unique
	}
	return sqlgraph.CountNodes(ctx, ttq.driver, _spec)
}

func (ttq *TagTranslationQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ttq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ttq *TagTranslationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tagtranslation.Table,
			Columns: tagtranslation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: tagtranslation.FieldID,
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
		_spec.Node.Columns = append(_spec.Node.Columns, tagtranslation.FieldID)
		for i := range fields {
			if fields[i] != tagtranslation.FieldID {
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

func (ttq *TagTranslationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ttq.driver.Dialect())
	t1 := builder.Table(tagtranslation.Table)
	columns := ttq.fields
	if len(columns) == 0 {
		columns = tagtranslation.Columns
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

// TagTranslationGroupBy is the group-by builder for TagTranslation entities.
type TagTranslationGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ttgb *TagTranslationGroupBy) Aggregate(fns ...AggregateFunc) *TagTranslationGroupBy {
	ttgb.fns = append(ttgb.fns, fns...)
	return ttgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ttgb *TagTranslationGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ttgb.path(ctx)
	if err != nil {
		return err
	}
	ttgb.sql = query
	return ttgb.sqlScan(ctx, v)
}

func (ttgb *TagTranslationGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ttgb.fields {
		if !tagtranslation.ValidColumn(f) {
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

func (ttgb *TagTranslationGroupBy) sqlQuery() *sql.Selector {
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

// TagTranslationSelect is the builder for selecting fields of TagTranslation entities.
type TagTranslationSelect struct {
	*TagTranslationQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tts *TagTranslationSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tts.prepareQuery(ctx); err != nil {
		return err
	}
	tts.sql = tts.TagTranslationQuery.sqlQuery(ctx)
	return tts.sqlScan(ctx, v)
}

func (tts *TagTranslationSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tts.sql.Query()
	if err := tts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}