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
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/user"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/usersession"
)

// UserSessionQuery is the builder for querying UserSession entities.
type UserSessionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UserSession
	// eager-loading edges.
	withUser *UserQuery
	withFKs  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserSessionQuery builder.
func (usq *UserSessionQuery) Where(ps ...predicate.UserSession) *UserSessionQuery {
	usq.predicates = append(usq.predicates, ps...)
	return usq
}

// Limit adds a limit step to the query.
func (usq *UserSessionQuery) Limit(limit int) *UserSessionQuery {
	usq.limit = &limit
	return usq
}

// Offset adds an offset step to the query.
func (usq *UserSessionQuery) Offset(offset int) *UserSessionQuery {
	usq.offset = &offset
	return usq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (usq *UserSessionQuery) Unique(unique bool) *UserSessionQuery {
	usq.unique = &unique
	return usq
}

// Order adds an order step to the query.
func (usq *UserSessionQuery) Order(o ...OrderFunc) *UserSessionQuery {
	usq.order = append(usq.order, o...)
	return usq
}

// QueryUser chains the current query on the "user" edge.
func (usq *UserSessionQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: usq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := usq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := usq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usersession.Table, usersession.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, usersession.UserTable, usersession.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(usq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserSession entity from the query.
// Returns a *NotFoundError when no UserSession was found.
func (usq *UserSessionQuery) First(ctx context.Context) (*UserSession, error) {
	nodes, err := usq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usersession.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (usq *UserSessionQuery) FirstX(ctx context.Context) *UserSession {
	node, err := usq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserSession ID from the query.
// Returns a *NotFoundError when no UserSession ID was found.
func (usq *UserSessionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = usq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usersession.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (usq *UserSessionQuery) FirstIDX(ctx context.Context) int {
	id, err := usq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserSession entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserSession entity is found.
// Returns a *NotFoundError when no UserSession entities are found.
func (usq *UserSessionQuery) Only(ctx context.Context) (*UserSession, error) {
	nodes, err := usq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usersession.Label}
	default:
		return nil, &NotSingularError{usersession.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (usq *UserSessionQuery) OnlyX(ctx context.Context) *UserSession {
	node, err := usq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserSession ID in the query.
// Returns a *NotSingularError when more than one UserSession ID is found.
// Returns a *NotFoundError when no entities are found.
func (usq *UserSessionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = usq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usersession.Label}
	default:
		err = &NotSingularError{usersession.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (usq *UserSessionQuery) OnlyIDX(ctx context.Context) int {
	id, err := usq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserSessions.
func (usq *UserSessionQuery) All(ctx context.Context) ([]*UserSession, error) {
	if err := usq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return usq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (usq *UserSessionQuery) AllX(ctx context.Context) []*UserSession {
	nodes, err := usq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserSession IDs.
func (usq *UserSessionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := usq.Select(usersession.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (usq *UserSessionQuery) IDsX(ctx context.Context) []int {
	ids, err := usq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (usq *UserSessionQuery) Count(ctx context.Context) (int, error) {
	if err := usq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return usq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (usq *UserSessionQuery) CountX(ctx context.Context) int {
	count, err := usq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (usq *UserSessionQuery) Exist(ctx context.Context) (bool, error) {
	if err := usq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return usq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (usq *UserSessionQuery) ExistX(ctx context.Context) bool {
	exist, err := usq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserSessionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (usq *UserSessionQuery) Clone() *UserSessionQuery {
	if usq == nil {
		return nil
	}
	return &UserSessionQuery{
		config:     usq.config,
		limit:      usq.limit,
		offset:     usq.offset,
		order:      append([]OrderFunc{}, usq.order...),
		predicates: append([]predicate.UserSession{}, usq.predicates...),
		withUser:   usq.withUser.Clone(),
		// clone intermediate query.
		sql:    usq.sql.Clone(),
		path:   usq.path,
		unique: usq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (usq *UserSessionQuery) WithUser(opts ...func(*UserQuery)) *UserSessionQuery {
	query := &UserQuery{config: usq.config}
	for _, opt := range opts {
		opt(query)
	}
	usq.withUser = query
	return usq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserSession.Query().
//		GroupBy(usersession.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (usq *UserSessionQuery) GroupBy(field string, fields ...string) *UserSessionGroupBy {
	grbuild := &UserSessionGroupBy{config: usq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := usq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return usq.sqlQuery(ctx), nil
	}
	grbuild.label = usersession.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.UserSession.Query().
//		Select(usersession.FieldCreateTime).
//		Scan(ctx, &v)
//
func (usq *UserSessionQuery) Select(fields ...string) *UserSessionSelect {
	usq.fields = append(usq.fields, fields...)
	selbuild := &UserSessionSelect{UserSessionQuery: usq}
	selbuild.label = usersession.Label
	selbuild.flds, selbuild.scan = &usq.fields, selbuild.Scan
	return selbuild
}

func (usq *UserSessionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range usq.fields {
		if !usersession.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if usq.path != nil {
		prev, err := usq.path(ctx)
		if err != nil {
			return err
		}
		usq.sql = prev
	}
	return nil
}

func (usq *UserSessionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserSession, error) {
	var (
		nodes       = []*UserSession{}
		withFKs     = usq.withFKs
		_spec       = usq.querySpec()
		loadedTypes = [1]bool{
			usq.withUser != nil,
		}
	)
	if usq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, usersession.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*UserSession).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &UserSession{config: usq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, usq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := usq.withUser; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*UserSession)
		for i := range nodes {
			if nodes[i].user_sessions == nil {
				continue
			}
			fk := *nodes[i].user_sessions
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
				return nil, fmt.Errorf(`unexpected foreign-key "user_sessions" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	return nodes, nil
}

func (usq *UserSessionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := usq.querySpec()
	_spec.Node.Columns = usq.fields
	if len(usq.fields) > 0 {
		_spec.Unique = usq.unique != nil && *usq.unique
	}
	return sqlgraph.CountNodes(ctx, usq.driver, _spec)
}

func (usq *UserSessionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := usq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (usq *UserSessionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usersession.Table,
			Columns: usersession.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usersession.FieldID,
			},
		},
		From:   usq.sql,
		Unique: true,
	}
	if unique := usq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := usq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usersession.FieldID)
		for i := range fields {
			if fields[i] != usersession.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := usq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := usq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := usq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := usq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (usq *UserSessionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(usq.driver.Dialect())
	t1 := builder.Table(usersession.Table)
	columns := usq.fields
	if len(columns) == 0 {
		columns = usersession.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if usq.sql != nil {
		selector = usq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if usq.unique != nil && *usq.unique {
		selector.Distinct()
	}
	for _, p := range usq.predicates {
		p(selector)
	}
	for _, p := range usq.order {
		p(selector)
	}
	if offset := usq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := usq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserSessionGroupBy is the group-by builder for UserSession entities.
type UserSessionGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (usgb *UserSessionGroupBy) Aggregate(fns ...AggregateFunc) *UserSessionGroupBy {
	usgb.fns = append(usgb.fns, fns...)
	return usgb
}

// Scan applies the group-by query and scans the result into the given value.
func (usgb *UserSessionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := usgb.path(ctx)
	if err != nil {
		return err
	}
	usgb.sql = query
	return usgb.sqlScan(ctx, v)
}

func (usgb *UserSessionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range usgb.fields {
		if !usersession.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := usgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := usgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (usgb *UserSessionGroupBy) sqlQuery() *sql.Selector {
	selector := usgb.sql.Select()
	aggregation := make([]string, 0, len(usgb.fns))
	for _, fn := range usgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(usgb.fields)+len(usgb.fns))
		for _, f := range usgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(usgb.fields...)...)
}

// UserSessionSelect is the builder for selecting fields of UserSession entities.
type UserSessionSelect struct {
	*UserSessionQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (uss *UserSessionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := uss.prepareQuery(ctx); err != nil {
		return err
	}
	uss.sql = uss.UserSessionQuery.sqlQuery(ctx)
	return uss.sqlScan(ctx, v)
}

func (uss *UserSessionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := uss.sql.Query()
	if err := uss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
