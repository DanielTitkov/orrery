// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/item"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/question"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/questiontranslation"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/test"
	"github.com/google/uuid"
)

// QuestionUpdate is the builder for updating Question entities.
type QuestionUpdate struct {
	config
	hooks    []Hook
	mutation *QuestionMutation
}

// Where appends a list predicates to the QuestionUpdate builder.
func (qu *QuestionUpdate) Where(ps ...predicate.Question) *QuestionUpdate {
	qu.mutation.Where(ps...)
	return qu
}

// SetUpdateTime sets the "update_time" field.
func (qu *QuestionUpdate) SetUpdateTime(t time.Time) *QuestionUpdate {
	qu.mutation.SetUpdateTime(t)
	return qu
}

// SetOrder sets the "order" field.
func (qu *QuestionUpdate) SetOrder(i int) *QuestionUpdate {
	qu.mutation.ResetOrder()
	qu.mutation.SetOrder(i)
	return qu
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (qu *QuestionUpdate) SetNillableOrder(i *int) *QuestionUpdate {
	if i != nil {
		qu.SetOrder(*i)
	}
	return qu
}

// AddOrder adds i to the "order" field.
func (qu *QuestionUpdate) AddOrder(i int) *QuestionUpdate {
	qu.mutation.AddOrder(i)
	return qu
}

// SetType sets the "type" field.
func (qu *QuestionUpdate) SetType(q question.Type) *QuestionUpdate {
	qu.mutation.SetType(q)
	return qu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (qu *QuestionUpdate) SetNillableType(q *question.Type) *QuestionUpdate {
	if q != nil {
		qu.SetType(*q)
	}
	return qu
}

// AddItemIDs adds the "items" edge to the Item entity by IDs.
func (qu *QuestionUpdate) AddItemIDs(ids ...uuid.UUID) *QuestionUpdate {
	qu.mutation.AddItemIDs(ids...)
	return qu
}

// AddItems adds the "items" edges to the Item entity.
func (qu *QuestionUpdate) AddItems(i ...*Item) *QuestionUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return qu.AddItemIDs(ids...)
}

// AddTranslationIDs adds the "translations" edge to the QuestionTranslation entity by IDs.
func (qu *QuestionUpdate) AddTranslationIDs(ids ...uuid.UUID) *QuestionUpdate {
	qu.mutation.AddTranslationIDs(ids...)
	return qu
}

// AddTranslations adds the "translations" edges to the QuestionTranslation entity.
func (qu *QuestionUpdate) AddTranslations(q ...*QuestionTranslation) *QuestionUpdate {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return qu.AddTranslationIDs(ids...)
}

// AddTestIDs adds the "test" edge to the Test entity by IDs.
func (qu *QuestionUpdate) AddTestIDs(ids ...uuid.UUID) *QuestionUpdate {
	qu.mutation.AddTestIDs(ids...)
	return qu
}

// AddTest adds the "test" edges to the Test entity.
func (qu *QuestionUpdate) AddTest(t ...*Test) *QuestionUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return qu.AddTestIDs(ids...)
}

// Mutation returns the QuestionMutation object of the builder.
func (qu *QuestionUpdate) Mutation() *QuestionMutation {
	return qu.mutation
}

// ClearItems clears all "items" edges to the Item entity.
func (qu *QuestionUpdate) ClearItems() *QuestionUpdate {
	qu.mutation.ClearItems()
	return qu
}

// RemoveItemIDs removes the "items" edge to Item entities by IDs.
func (qu *QuestionUpdate) RemoveItemIDs(ids ...uuid.UUID) *QuestionUpdate {
	qu.mutation.RemoveItemIDs(ids...)
	return qu
}

// RemoveItems removes "items" edges to Item entities.
func (qu *QuestionUpdate) RemoveItems(i ...*Item) *QuestionUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return qu.RemoveItemIDs(ids...)
}

// ClearTranslations clears all "translations" edges to the QuestionTranslation entity.
func (qu *QuestionUpdate) ClearTranslations() *QuestionUpdate {
	qu.mutation.ClearTranslations()
	return qu
}

// RemoveTranslationIDs removes the "translations" edge to QuestionTranslation entities by IDs.
func (qu *QuestionUpdate) RemoveTranslationIDs(ids ...uuid.UUID) *QuestionUpdate {
	qu.mutation.RemoveTranslationIDs(ids...)
	return qu
}

// RemoveTranslations removes "translations" edges to QuestionTranslation entities.
func (qu *QuestionUpdate) RemoveTranslations(q ...*QuestionTranslation) *QuestionUpdate {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return qu.RemoveTranslationIDs(ids...)
}

// ClearTest clears all "test" edges to the Test entity.
func (qu *QuestionUpdate) ClearTest() *QuestionUpdate {
	qu.mutation.ClearTest()
	return qu
}

// RemoveTestIDs removes the "test" edge to Test entities by IDs.
func (qu *QuestionUpdate) RemoveTestIDs(ids ...uuid.UUID) *QuestionUpdate {
	qu.mutation.RemoveTestIDs(ids...)
	return qu
}

// RemoveTest removes "test" edges to Test entities.
func (qu *QuestionUpdate) RemoveTest(t ...*Test) *QuestionUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return qu.RemoveTestIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (qu *QuestionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	qu.defaults()
	if len(qu.hooks) == 0 {
		if err = qu.check(); err != nil {
			return 0, err
		}
		affected, err = qu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QuestionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = qu.check(); err != nil {
				return 0, err
			}
			qu.mutation = mutation
			affected, err = qu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(qu.hooks) - 1; i >= 0; i-- {
			if qu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = qu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (qu *QuestionUpdate) SaveX(ctx context.Context) int {
	affected, err := qu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (qu *QuestionUpdate) Exec(ctx context.Context) error {
	_, err := qu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qu *QuestionUpdate) ExecX(ctx context.Context) {
	if err := qu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (qu *QuestionUpdate) defaults() {
	if _, ok := qu.mutation.UpdateTime(); !ok {
		v := question.UpdateDefaultUpdateTime()
		qu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (qu *QuestionUpdate) check() error {
	if v, ok := qu.mutation.GetType(); ok {
		if err := question.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Question.type": %w`, err)}
		}
	}
	return nil
}

func (qu *QuestionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   question.Table,
			Columns: question.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: question.FieldID,
			},
		},
	}
	if ps := qu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := qu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: question.FieldUpdateTime,
		})
	}
	if value, ok := qu.mutation.Order(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: question.FieldOrder,
		})
	}
	if value, ok := qu.mutation.AddedOrder(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: question.FieldOrder,
		})
	}
	if value, ok := qu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: question.FieldType,
		})
	}
	if qu.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.ItemsTable,
			Columns: question.ItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: item.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.RemovedItemsIDs(); len(nodes) > 0 && !qu.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.ItemsTable,
			Columns: question.ItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: item.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.ItemsTable,
			Columns: question.ItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: item.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qu.mutation.TranslationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.TranslationsTable,
			Columns: []string{question.TranslationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: questiontranslation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.RemovedTranslationsIDs(); len(nodes) > 0 && !qu.mutation.TranslationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.TranslationsTable,
			Columns: []string{question.TranslationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: questiontranslation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.TranslationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.TranslationsTable,
			Columns: []string{question.TranslationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: questiontranslation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qu.mutation.TestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   question.TestTable,
			Columns: question.TestPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: test.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.RemovedTestIDs(); len(nodes) > 0 && !qu.mutation.TestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   question.TestTable,
			Columns: question.TestPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: test.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.TestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   question.TestTable,
			Columns: question.TestPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: test.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, qu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{question.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// QuestionUpdateOne is the builder for updating a single Question entity.
type QuestionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *QuestionMutation
}

// SetUpdateTime sets the "update_time" field.
func (quo *QuestionUpdateOne) SetUpdateTime(t time.Time) *QuestionUpdateOne {
	quo.mutation.SetUpdateTime(t)
	return quo
}

// SetOrder sets the "order" field.
func (quo *QuestionUpdateOne) SetOrder(i int) *QuestionUpdateOne {
	quo.mutation.ResetOrder()
	quo.mutation.SetOrder(i)
	return quo
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (quo *QuestionUpdateOne) SetNillableOrder(i *int) *QuestionUpdateOne {
	if i != nil {
		quo.SetOrder(*i)
	}
	return quo
}

// AddOrder adds i to the "order" field.
func (quo *QuestionUpdateOne) AddOrder(i int) *QuestionUpdateOne {
	quo.mutation.AddOrder(i)
	return quo
}

// SetType sets the "type" field.
func (quo *QuestionUpdateOne) SetType(q question.Type) *QuestionUpdateOne {
	quo.mutation.SetType(q)
	return quo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (quo *QuestionUpdateOne) SetNillableType(q *question.Type) *QuestionUpdateOne {
	if q != nil {
		quo.SetType(*q)
	}
	return quo
}

// AddItemIDs adds the "items" edge to the Item entity by IDs.
func (quo *QuestionUpdateOne) AddItemIDs(ids ...uuid.UUID) *QuestionUpdateOne {
	quo.mutation.AddItemIDs(ids...)
	return quo
}

// AddItems adds the "items" edges to the Item entity.
func (quo *QuestionUpdateOne) AddItems(i ...*Item) *QuestionUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return quo.AddItemIDs(ids...)
}

// AddTranslationIDs adds the "translations" edge to the QuestionTranslation entity by IDs.
func (quo *QuestionUpdateOne) AddTranslationIDs(ids ...uuid.UUID) *QuestionUpdateOne {
	quo.mutation.AddTranslationIDs(ids...)
	return quo
}

// AddTranslations adds the "translations" edges to the QuestionTranslation entity.
func (quo *QuestionUpdateOne) AddTranslations(q ...*QuestionTranslation) *QuestionUpdateOne {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return quo.AddTranslationIDs(ids...)
}

// AddTestIDs adds the "test" edge to the Test entity by IDs.
func (quo *QuestionUpdateOne) AddTestIDs(ids ...uuid.UUID) *QuestionUpdateOne {
	quo.mutation.AddTestIDs(ids...)
	return quo
}

// AddTest adds the "test" edges to the Test entity.
func (quo *QuestionUpdateOne) AddTest(t ...*Test) *QuestionUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return quo.AddTestIDs(ids...)
}

// Mutation returns the QuestionMutation object of the builder.
func (quo *QuestionUpdateOne) Mutation() *QuestionMutation {
	return quo.mutation
}

// ClearItems clears all "items" edges to the Item entity.
func (quo *QuestionUpdateOne) ClearItems() *QuestionUpdateOne {
	quo.mutation.ClearItems()
	return quo
}

// RemoveItemIDs removes the "items" edge to Item entities by IDs.
func (quo *QuestionUpdateOne) RemoveItemIDs(ids ...uuid.UUID) *QuestionUpdateOne {
	quo.mutation.RemoveItemIDs(ids...)
	return quo
}

// RemoveItems removes "items" edges to Item entities.
func (quo *QuestionUpdateOne) RemoveItems(i ...*Item) *QuestionUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return quo.RemoveItemIDs(ids...)
}

// ClearTranslations clears all "translations" edges to the QuestionTranslation entity.
func (quo *QuestionUpdateOne) ClearTranslations() *QuestionUpdateOne {
	quo.mutation.ClearTranslations()
	return quo
}

// RemoveTranslationIDs removes the "translations" edge to QuestionTranslation entities by IDs.
func (quo *QuestionUpdateOne) RemoveTranslationIDs(ids ...uuid.UUID) *QuestionUpdateOne {
	quo.mutation.RemoveTranslationIDs(ids...)
	return quo
}

// RemoveTranslations removes "translations" edges to QuestionTranslation entities.
func (quo *QuestionUpdateOne) RemoveTranslations(q ...*QuestionTranslation) *QuestionUpdateOne {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return quo.RemoveTranslationIDs(ids...)
}

// ClearTest clears all "test" edges to the Test entity.
func (quo *QuestionUpdateOne) ClearTest() *QuestionUpdateOne {
	quo.mutation.ClearTest()
	return quo
}

// RemoveTestIDs removes the "test" edge to Test entities by IDs.
func (quo *QuestionUpdateOne) RemoveTestIDs(ids ...uuid.UUID) *QuestionUpdateOne {
	quo.mutation.RemoveTestIDs(ids...)
	return quo
}

// RemoveTest removes "test" edges to Test entities.
func (quo *QuestionUpdateOne) RemoveTest(t ...*Test) *QuestionUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return quo.RemoveTestIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (quo *QuestionUpdateOne) Select(field string, fields ...string) *QuestionUpdateOne {
	quo.fields = append([]string{field}, fields...)
	return quo
}

// Save executes the query and returns the updated Question entity.
func (quo *QuestionUpdateOne) Save(ctx context.Context) (*Question, error) {
	var (
		err  error
		node *Question
	)
	quo.defaults()
	if len(quo.hooks) == 0 {
		if err = quo.check(); err != nil {
			return nil, err
		}
		node, err = quo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QuestionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = quo.check(); err != nil {
				return nil, err
			}
			quo.mutation = mutation
			node, err = quo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(quo.hooks) - 1; i >= 0; i-- {
			if quo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = quo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, quo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Question)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from QuestionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (quo *QuestionUpdateOne) SaveX(ctx context.Context) *Question {
	node, err := quo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (quo *QuestionUpdateOne) Exec(ctx context.Context) error {
	_, err := quo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (quo *QuestionUpdateOne) ExecX(ctx context.Context) {
	if err := quo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (quo *QuestionUpdateOne) defaults() {
	if _, ok := quo.mutation.UpdateTime(); !ok {
		v := question.UpdateDefaultUpdateTime()
		quo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (quo *QuestionUpdateOne) check() error {
	if v, ok := quo.mutation.GetType(); ok {
		if err := question.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Question.type": %w`, err)}
		}
	}
	return nil
}

func (quo *QuestionUpdateOne) sqlSave(ctx context.Context) (_node *Question, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   question.Table,
			Columns: question.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: question.FieldID,
			},
		},
	}
	id, ok := quo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Question.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := quo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, question.FieldID)
		for _, f := range fields {
			if !question.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != question.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := quo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := quo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: question.FieldUpdateTime,
		})
	}
	if value, ok := quo.mutation.Order(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: question.FieldOrder,
		})
	}
	if value, ok := quo.mutation.AddedOrder(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: question.FieldOrder,
		})
	}
	if value, ok := quo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: question.FieldType,
		})
	}
	if quo.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.ItemsTable,
			Columns: question.ItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: item.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.RemovedItemsIDs(); len(nodes) > 0 && !quo.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.ItemsTable,
			Columns: question.ItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: item.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.ItemsTable,
			Columns: question.ItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: item.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if quo.mutation.TranslationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.TranslationsTable,
			Columns: []string{question.TranslationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: questiontranslation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.RemovedTranslationsIDs(); len(nodes) > 0 && !quo.mutation.TranslationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.TranslationsTable,
			Columns: []string{question.TranslationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: questiontranslation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.TranslationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.TranslationsTable,
			Columns: []string{question.TranslationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: questiontranslation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if quo.mutation.TestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   question.TestTable,
			Columns: question.TestPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: test.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.RemovedTestIDs(); len(nodes) > 0 && !quo.mutation.TestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   question.TestTable,
			Columns: question.TestPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: test.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.TestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   question.TestTable,
			Columns: question.TestPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: test.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Question{config: quo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, quo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{question.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
