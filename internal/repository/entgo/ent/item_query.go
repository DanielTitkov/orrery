// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/item"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/itemtranslation"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/predicate"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/question"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/response"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/scale"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/scaleitem"
)

// ItemQuery is the builder for querying Item entities.
type ItemQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Item
	// eager-loading edges.
	withResponses    *ResponseQuery
	withTranslations *ItemTranslationQuery
	withScale        *ScaleQuery
	withQuestion     *QuestionQuery
	withScaleItem    *ScaleItemQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ItemQuery builder.
func (iq *ItemQuery) Where(ps ...predicate.Item) *ItemQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit adds a limit step to the query.
func (iq *ItemQuery) Limit(limit int) *ItemQuery {
	iq.limit = &limit
	return iq
}

// Offset adds an offset step to the query.
func (iq *ItemQuery) Offset(offset int) *ItemQuery {
	iq.offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *ItemQuery) Unique(unique bool) *ItemQuery {
	iq.unique = &unique
	return iq
}

// Order adds an order step to the query.
func (iq *ItemQuery) Order(o ...OrderFunc) *ItemQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryResponses chains the current query on the "responses" edge.
func (iq *ItemQuery) QueryResponses() *ResponseQuery {
	query := &ResponseQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, selector),
			sqlgraph.To(response.Table, response.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, item.ResponsesTable, item.ResponsesColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTranslations chains the current query on the "translations" edge.
func (iq *ItemQuery) QueryTranslations() *ItemTranslationQuery {
	query := &ItemTranslationQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, selector),
			sqlgraph.To(itemtranslation.Table, itemtranslation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, item.TranslationsTable, item.TranslationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryScale chains the current query on the "scale" edge.
func (iq *ItemQuery) QueryScale() *ScaleQuery {
	query := &ScaleQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, selector),
			sqlgraph.To(scale.Table, scale.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, item.ScaleTable, item.ScalePrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryQuestion chains the current query on the "question" edge.
func (iq *ItemQuery) QueryQuestion() *QuestionQuery {
	query := &QuestionQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, selector),
			sqlgraph.To(question.Table, question.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, item.QuestionTable, item.QuestionPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryScaleItem chains the current query on the "scale_item" edge.
func (iq *ItemQuery) QueryScaleItem() *ScaleItemQuery {
	query := &ScaleItemQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, selector),
			sqlgraph.To(scaleitem.Table, scaleitem.ItemColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, item.ScaleItemTable, item.ScaleItemColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Item entity from the query.
// Returns a *NotFoundError when no Item was found.
func (iq *ItemQuery) First(ctx context.Context) (*Item, error) {
	nodes, err := iq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{item.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *ItemQuery) FirstX(ctx context.Context) *Item {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Item ID from the query.
// Returns a *NotFoundError when no Item ID was found.
func (iq *ItemQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{item.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *ItemQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Item entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Item entity is found.
// Returns a *NotFoundError when no Item entities are found.
func (iq *ItemQuery) Only(ctx context.Context) (*Item, error) {
	nodes, err := iq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{item.Label}
	default:
		return nil, &NotSingularError{item.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *ItemQuery) OnlyX(ctx context.Context) *Item {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Item ID in the query.
// Returns a *NotSingularError when more than one Item ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *ItemQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{item.Label}
	default:
		err = &NotSingularError{item.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *ItemQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Items.
func (iq *ItemQuery) All(ctx context.Context) ([]*Item, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return iq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (iq *ItemQuery) AllX(ctx context.Context) []*Item {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Item IDs.
func (iq *ItemQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := iq.Select(item.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *ItemQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *ItemQuery) Count(ctx context.Context) (int, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return iq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (iq *ItemQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *ItemQuery) Exist(ctx context.Context) (bool, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return iq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *ItemQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *ItemQuery) Clone() *ItemQuery {
	if iq == nil {
		return nil
	}
	return &ItemQuery{
		config:           iq.config,
		limit:            iq.limit,
		offset:           iq.offset,
		order:            append([]OrderFunc{}, iq.order...),
		predicates:       append([]predicate.Item{}, iq.predicates...),
		withResponses:    iq.withResponses.Clone(),
		withTranslations: iq.withTranslations.Clone(),
		withScale:        iq.withScale.Clone(),
		withQuestion:     iq.withQuestion.Clone(),
		withScaleItem:    iq.withScaleItem.Clone(),
		// clone intermediate query.
		sql:    iq.sql.Clone(),
		path:   iq.path,
		unique: iq.unique,
	}
}

// WithResponses tells the query-builder to eager-load the nodes that are connected to
// the "responses" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *ItemQuery) WithResponses(opts ...func(*ResponseQuery)) *ItemQuery {
	query := &ResponseQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withResponses = query
	return iq
}

// WithTranslations tells the query-builder to eager-load the nodes that are connected to
// the "translations" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *ItemQuery) WithTranslations(opts ...func(*ItemTranslationQuery)) *ItemQuery {
	query := &ItemTranslationQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withTranslations = query
	return iq
}

// WithScale tells the query-builder to eager-load the nodes that are connected to
// the "scale" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *ItemQuery) WithScale(opts ...func(*ScaleQuery)) *ItemQuery {
	query := &ScaleQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withScale = query
	return iq
}

// WithQuestion tells the query-builder to eager-load the nodes that are connected to
// the "question" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *ItemQuery) WithQuestion(opts ...func(*QuestionQuery)) *ItemQuery {
	query := &QuestionQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withQuestion = query
	return iq
}

// WithScaleItem tells the query-builder to eager-load the nodes that are connected to
// the "scale_item" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *ItemQuery) WithScaleItem(opts ...func(*ScaleItemQuery)) *ItemQuery {
	query := &ScaleItemQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withScaleItem = query
	return iq
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
//	client.Item.Query().
//		GroupBy(item.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (iq *ItemQuery) GroupBy(field string, fields ...string) *ItemGroupBy {
	grbuild := &ItemGroupBy{config: iq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return iq.sqlQuery(ctx), nil
	}
	grbuild.label = item.Label
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
//	client.Item.Query().
//		Select(item.FieldCreateTime).
//		Scan(ctx, &v)
//
func (iq *ItemQuery) Select(fields ...string) *ItemSelect {
	iq.fields = append(iq.fields, fields...)
	selbuild := &ItemSelect{ItemQuery: iq}
	selbuild.label = item.Label
	selbuild.flds, selbuild.scan = &iq.fields, selbuild.Scan
	return selbuild
}

func (iq *ItemQuery) prepareQuery(ctx context.Context) error {
	for _, f := range iq.fields {
		if !item.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *ItemQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Item, error) {
	var (
		nodes       = []*Item{}
		_spec       = iq.querySpec()
		loadedTypes = [5]bool{
			iq.withResponses != nil,
			iq.withTranslations != nil,
			iq.withScale != nil,
			iq.withQuestion != nil,
			iq.withScaleItem != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Item).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Item{config: iq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := iq.withResponses; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Item)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Responses = []*Response{}
		}
		query.withFKs = true
		query.Where(predicate.Response(func(s *sql.Selector) {
			s.Where(sql.InValues(item.ResponsesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.item_responses
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "item_responses" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "item_responses" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Responses = append(node.Edges.Responses, n)
		}
	}

	if query := iq.withTranslations; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Item)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Translations = []*ItemTranslation{}
		}
		query.withFKs = true
		query.Where(predicate.ItemTranslation(func(s *sql.Selector) {
			s.Where(sql.InValues(item.TranslationsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.item_translations
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "item_translations" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "item_translations" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Translations = append(node.Edges.Translations, n)
		}
	}

	if query := iq.withScale; query != nil {
		edgeids := make([]driver.Value, len(nodes))
		byid := make(map[uuid.UUID]*Item)
		nids := make(map[uuid.UUID]map[*Item]struct{})
		for i, node := range nodes {
			edgeids[i] = node.ID
			byid[node.ID] = node
			node.Edges.Scale = []*Scale{}
		}
		query.Where(func(s *sql.Selector) {
			joinT := sql.Table(item.ScaleTable)
			s.Join(joinT).On(s.C(scale.FieldID), joinT.C(item.ScalePrimaryKey[0]))
			s.Where(sql.InValues(joinT.C(item.ScalePrimaryKey[1]), edgeids...))
			columns := s.SelectedColumns()
			s.Select(joinT.C(item.ScalePrimaryKey[1]))
			s.AppendSelect(columns...)
			s.SetDistinct(false)
		})
		neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]interface{}, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]interface{}{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []interface{}) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Item]struct{}{byid[outValue]: struct{}{}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byid[outValue]] = struct{}{}
				return nil
			}
		})
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "scale" node returned %v`, n.ID)
			}
			for kn := range nodes {
				kn.Edges.Scale = append(kn.Edges.Scale, n)
			}
		}
	}

	if query := iq.withQuestion; query != nil {
		edgeids := make([]driver.Value, len(nodes))
		byid := make(map[uuid.UUID]*Item)
		nids := make(map[uuid.UUID]map[*Item]struct{})
		for i, node := range nodes {
			edgeids[i] = node.ID
			byid[node.ID] = node
			node.Edges.Question = []*Question{}
		}
		query.Where(func(s *sql.Selector) {
			joinT := sql.Table(item.QuestionTable)
			s.Join(joinT).On(s.C(question.FieldID), joinT.C(item.QuestionPrimaryKey[0]))
			s.Where(sql.InValues(joinT.C(item.QuestionPrimaryKey[1]), edgeids...))
			columns := s.SelectedColumns()
			s.Select(joinT.C(item.QuestionPrimaryKey[1]))
			s.AppendSelect(columns...)
			s.SetDistinct(false)
		})
		neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]interface{}, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]interface{}{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []interface{}) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Item]struct{}{byid[outValue]: struct{}{}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byid[outValue]] = struct{}{}
				return nil
			}
		})
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "question" node returned %v`, n.ID)
			}
			for kn := range nodes {
				kn.Edges.Question = append(kn.Edges.Question, n)
			}
		}
	}

	if query := iq.withScaleItem; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Item)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.ScaleItem = []*ScaleItem{}
		}
		query.Where(predicate.ScaleItem(func(s *sql.Selector) {
			s.Where(sql.InValues(item.ScaleItemColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.ItemID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "item_id" returned %v for node %v`, fk, n)
			}
			node.Edges.ScaleItem = append(node.Edges.ScaleItem, n)
		}
	}

	return nodes, nil
}

func (iq *ItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Columns = iq.fields
	if len(iq.fields) > 0 {
		_spec.Unique = iq.unique != nil && *iq.unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *ItemQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := iq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (iq *ItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   item.Table,
			Columns: item.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: item.FieldID,
			},
		},
		From:   iq.sql,
		Unique: true,
	}
	if unique := iq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := iq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, item.FieldID)
		for i := range fields {
			if fields[i] != item.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *ItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(item.Table)
	columns := iq.fields
	if len(columns) == 0 {
		columns = item.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.unique != nil && *iq.unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ItemGroupBy is the group-by builder for Item entities.
type ItemGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *ItemGroupBy) Aggregate(fns ...AggregateFunc) *ItemGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the group-by query and scans the result into the given value.
func (igb *ItemGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := igb.path(ctx)
	if err != nil {
		return err
	}
	igb.sql = query
	return igb.sqlScan(ctx, v)
}

func (igb *ItemGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range igb.fields {
		if !item.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := igb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (igb *ItemGroupBy) sqlQuery() *sql.Selector {
	selector := igb.sql.Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(igb.fields)+len(igb.fns))
		for _, f := range igb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(igb.fields...)...)
}

// ItemSelect is the builder for selecting fields of Item entities.
type ItemSelect struct {
	*ItemQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (is *ItemSelect) Scan(ctx context.Context, v interface{}) error {
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	is.sql = is.ItemQuery.sqlQuery(ctx)
	return is.sqlScan(ctx, v)
}

func (is *ItemSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := is.sql.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
