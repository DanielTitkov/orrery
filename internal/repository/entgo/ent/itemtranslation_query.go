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
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/item"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/itemtranslation"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/predicate"
)

// ItemTranslationQuery is the builder for querying ItemTranslation entities.
type ItemTranslationQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ItemTranslation
	// eager-loading edges.
	withItem *ItemQuery
	withFKs  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ItemTranslationQuery builder.
func (itq *ItemTranslationQuery) Where(ps ...predicate.ItemTranslation) *ItemTranslationQuery {
	itq.predicates = append(itq.predicates, ps...)
	return itq
}

// Limit adds a limit step to the query.
func (itq *ItemTranslationQuery) Limit(limit int) *ItemTranslationQuery {
	itq.limit = &limit
	return itq
}

// Offset adds an offset step to the query.
func (itq *ItemTranslationQuery) Offset(offset int) *ItemTranslationQuery {
	itq.offset = &offset
	return itq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (itq *ItemTranslationQuery) Unique(unique bool) *ItemTranslationQuery {
	itq.unique = &unique
	return itq
}

// Order adds an order step to the query.
func (itq *ItemTranslationQuery) Order(o ...OrderFunc) *ItemTranslationQuery {
	itq.order = append(itq.order, o...)
	return itq
}

// QueryItem chains the current query on the "item" edge.
func (itq *ItemTranslationQuery) QueryItem() *ItemQuery {
	query := &ItemQuery{config: itq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := itq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := itq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(itemtranslation.Table, itemtranslation.FieldID, selector),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, itemtranslation.ItemTable, itemtranslation.ItemColumn),
		)
		fromU = sqlgraph.SetNeighbors(itq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ItemTranslation entity from the query.
// Returns a *NotFoundError when no ItemTranslation was found.
func (itq *ItemTranslationQuery) First(ctx context.Context) (*ItemTranslation, error) {
	nodes, err := itq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{itemtranslation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (itq *ItemTranslationQuery) FirstX(ctx context.Context) *ItemTranslation {
	node, err := itq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ItemTranslation ID from the query.
// Returns a *NotFoundError when no ItemTranslation ID was found.
func (itq *ItemTranslationQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = itq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{itemtranslation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (itq *ItemTranslationQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := itq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ItemTranslation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ItemTranslation entity is found.
// Returns a *NotFoundError when no ItemTranslation entities are found.
func (itq *ItemTranslationQuery) Only(ctx context.Context) (*ItemTranslation, error) {
	nodes, err := itq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{itemtranslation.Label}
	default:
		return nil, &NotSingularError{itemtranslation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (itq *ItemTranslationQuery) OnlyX(ctx context.Context) *ItemTranslation {
	node, err := itq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ItemTranslation ID in the query.
// Returns a *NotSingularError when more than one ItemTranslation ID is found.
// Returns a *NotFoundError when no entities are found.
func (itq *ItemTranslationQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = itq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{itemtranslation.Label}
	default:
		err = &NotSingularError{itemtranslation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (itq *ItemTranslationQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := itq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ItemTranslations.
func (itq *ItemTranslationQuery) All(ctx context.Context) ([]*ItemTranslation, error) {
	if err := itq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return itq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (itq *ItemTranslationQuery) AllX(ctx context.Context) []*ItemTranslation {
	nodes, err := itq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ItemTranslation IDs.
func (itq *ItemTranslationQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := itq.Select(itemtranslation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (itq *ItemTranslationQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := itq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (itq *ItemTranslationQuery) Count(ctx context.Context) (int, error) {
	if err := itq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return itq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (itq *ItemTranslationQuery) CountX(ctx context.Context) int {
	count, err := itq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (itq *ItemTranslationQuery) Exist(ctx context.Context) (bool, error) {
	if err := itq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return itq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (itq *ItemTranslationQuery) ExistX(ctx context.Context) bool {
	exist, err := itq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ItemTranslationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (itq *ItemTranslationQuery) Clone() *ItemTranslationQuery {
	if itq == nil {
		return nil
	}
	return &ItemTranslationQuery{
		config:     itq.config,
		limit:      itq.limit,
		offset:     itq.offset,
		order:      append([]OrderFunc{}, itq.order...),
		predicates: append([]predicate.ItemTranslation{}, itq.predicates...),
		withItem:   itq.withItem.Clone(),
		// clone intermediate query.
		sql:    itq.sql.Clone(),
		path:   itq.path,
		unique: itq.unique,
	}
}

// WithItem tells the query-builder to eager-load the nodes that are connected to
// the "item" edge. The optional arguments are used to configure the query builder of the edge.
func (itq *ItemTranslationQuery) WithItem(opts ...func(*ItemQuery)) *ItemTranslationQuery {
	query := &ItemQuery{config: itq.config}
	for _, opt := range opts {
		opt(query)
	}
	itq.withItem = query
	return itq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Locale itemtranslation.Locale `json:"locale,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ItemTranslation.Query().
//		GroupBy(itemtranslation.FieldLocale).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (itq *ItemTranslationQuery) GroupBy(field string, fields ...string) *ItemTranslationGroupBy {
	grbuild := &ItemTranslationGroupBy{config: itq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := itq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return itq.sqlQuery(ctx), nil
	}
	grbuild.label = itemtranslation.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Locale itemtranslation.Locale `json:"locale,omitempty"`
//	}
//
//	client.ItemTranslation.Query().
//		Select(itemtranslation.FieldLocale).
//		Scan(ctx, &v)
//
func (itq *ItemTranslationQuery) Select(fields ...string) *ItemTranslationSelect {
	itq.fields = append(itq.fields, fields...)
	selbuild := &ItemTranslationSelect{ItemTranslationQuery: itq}
	selbuild.label = itemtranslation.Label
	selbuild.flds, selbuild.scan = &itq.fields, selbuild.Scan
	return selbuild
}

func (itq *ItemTranslationQuery) prepareQuery(ctx context.Context) error {
	for _, f := range itq.fields {
		if !itemtranslation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if itq.path != nil {
		prev, err := itq.path(ctx)
		if err != nil {
			return err
		}
		itq.sql = prev
	}
	return nil
}

func (itq *ItemTranslationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ItemTranslation, error) {
	var (
		nodes       = []*ItemTranslation{}
		withFKs     = itq.withFKs
		_spec       = itq.querySpec()
		loadedTypes = [1]bool{
			itq.withItem != nil,
		}
	)
	if itq.withItem != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, itemtranslation.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*ItemTranslation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &ItemTranslation{config: itq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, itq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := itq.withItem; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*ItemTranslation)
		for i := range nodes {
			if nodes[i].item_translations == nil {
				continue
			}
			fk := *nodes[i].item_translations
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(item.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "item_translations" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Item = n
			}
		}
	}

	return nodes, nil
}

func (itq *ItemTranslationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := itq.querySpec()
	_spec.Node.Columns = itq.fields
	if len(itq.fields) > 0 {
		_spec.Unique = itq.unique != nil && *itq.unique
	}
	return sqlgraph.CountNodes(ctx, itq.driver, _spec)
}

func (itq *ItemTranslationQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := itq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (itq *ItemTranslationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   itemtranslation.Table,
			Columns: itemtranslation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: itemtranslation.FieldID,
			},
		},
		From:   itq.sql,
		Unique: true,
	}
	if unique := itq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := itq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, itemtranslation.FieldID)
		for i := range fields {
			if fields[i] != itemtranslation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := itq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := itq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := itq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := itq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (itq *ItemTranslationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(itq.driver.Dialect())
	t1 := builder.Table(itemtranslation.Table)
	columns := itq.fields
	if len(columns) == 0 {
		columns = itemtranslation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if itq.sql != nil {
		selector = itq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if itq.unique != nil && *itq.unique {
		selector.Distinct()
	}
	for _, p := range itq.predicates {
		p(selector)
	}
	for _, p := range itq.order {
		p(selector)
	}
	if offset := itq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := itq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ItemTranslationGroupBy is the group-by builder for ItemTranslation entities.
type ItemTranslationGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (itgb *ItemTranslationGroupBy) Aggregate(fns ...AggregateFunc) *ItemTranslationGroupBy {
	itgb.fns = append(itgb.fns, fns...)
	return itgb
}

// Scan applies the group-by query and scans the result into the given value.
func (itgb *ItemTranslationGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := itgb.path(ctx)
	if err != nil {
		return err
	}
	itgb.sql = query
	return itgb.sqlScan(ctx, v)
}

func (itgb *ItemTranslationGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range itgb.fields {
		if !itemtranslation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := itgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := itgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (itgb *ItemTranslationGroupBy) sqlQuery() *sql.Selector {
	selector := itgb.sql.Select()
	aggregation := make([]string, 0, len(itgb.fns))
	for _, fn := range itgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(itgb.fields)+len(itgb.fns))
		for _, f := range itgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(itgb.fields...)...)
}

// ItemTranslationSelect is the builder for selecting fields of ItemTranslation entities.
type ItemTranslationSelect struct {
	*ItemTranslationQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (its *ItemTranslationSelect) Scan(ctx context.Context, v interface{}) error {
	if err := its.prepareQuery(ctx); err != nil {
		return err
	}
	its.sql = its.ItemTranslationQuery.sqlQuery(ctx)
	return its.sqlScan(ctx, v)
}

func (its *ItemTranslationSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := its.sql.Query()
	if err := its.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
