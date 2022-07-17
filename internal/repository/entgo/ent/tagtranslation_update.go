// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/predicate"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/tag"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/tagtranslation"
)

// TagTranslationUpdate is the builder for updating TagTranslation entities.
type TagTranslationUpdate struct {
	config
	hooks    []Hook
	mutation *TagTranslationMutation
}

// Where appends a list predicates to the TagTranslationUpdate builder.
func (ttu *TagTranslationUpdate) Where(ps ...predicate.TagTranslation) *TagTranslationUpdate {
	ttu.mutation.Where(ps...)
	return ttu
}

// SetContent sets the "content" field.
func (ttu *TagTranslationUpdate) SetContent(s string) *TagTranslationUpdate {
	ttu.mutation.SetContent(s)
	return ttu
}

// SetTagID sets the "tag" edge to the Tag entity by ID.
func (ttu *TagTranslationUpdate) SetTagID(id uuid.UUID) *TagTranslationUpdate {
	ttu.mutation.SetTagID(id)
	return ttu
}

// SetNillableTagID sets the "tag" edge to the Tag entity by ID if the given value is not nil.
func (ttu *TagTranslationUpdate) SetNillableTagID(id *uuid.UUID) *TagTranslationUpdate {
	if id != nil {
		ttu = ttu.SetTagID(*id)
	}
	return ttu
}

// SetTag sets the "tag" edge to the Tag entity.
func (ttu *TagTranslationUpdate) SetTag(t *Tag) *TagTranslationUpdate {
	return ttu.SetTagID(t.ID)
}

// Mutation returns the TagTranslationMutation object of the builder.
func (ttu *TagTranslationUpdate) Mutation() *TagTranslationMutation {
	return ttu.mutation
}

// ClearTag clears the "tag" edge to the Tag entity.
func (ttu *TagTranslationUpdate) ClearTag() *TagTranslationUpdate {
	ttu.mutation.ClearTag()
	return ttu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ttu *TagTranslationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ttu.hooks) == 0 {
		if err = ttu.check(); err != nil {
			return 0, err
		}
		affected, err = ttu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TagTranslationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ttu.check(); err != nil {
				return 0, err
			}
			ttu.mutation = mutation
			affected, err = ttu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ttu.hooks) - 1; i >= 0; i-- {
			if ttu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ttu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ttu *TagTranslationUpdate) SaveX(ctx context.Context) int {
	affected, err := ttu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ttu *TagTranslationUpdate) Exec(ctx context.Context) error {
	_, err := ttu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttu *TagTranslationUpdate) ExecX(ctx context.Context) {
	if err := ttu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttu *TagTranslationUpdate) check() error {
	if v, ok := ttu.mutation.Content(); ok {
		if err := tagtranslation.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "TagTranslation.content": %w`, err)}
		}
	}
	return nil
}

func (ttu *TagTranslationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tagtranslation.Table,
			Columns: tagtranslation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: tagtranslation.FieldID,
			},
		},
	}
	if ps := ttu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ttu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tagtranslation.FieldContent,
		})
	}
	if ttu.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tagtranslation.TagTable,
			Columns: []string{tagtranslation.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttu.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tagtranslation.TagTable,
			Columns: []string{tagtranslation.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ttu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tagtranslation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TagTranslationUpdateOne is the builder for updating a single TagTranslation entity.
type TagTranslationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TagTranslationMutation
}

// SetContent sets the "content" field.
func (ttuo *TagTranslationUpdateOne) SetContent(s string) *TagTranslationUpdateOne {
	ttuo.mutation.SetContent(s)
	return ttuo
}

// SetTagID sets the "tag" edge to the Tag entity by ID.
func (ttuo *TagTranslationUpdateOne) SetTagID(id uuid.UUID) *TagTranslationUpdateOne {
	ttuo.mutation.SetTagID(id)
	return ttuo
}

// SetNillableTagID sets the "tag" edge to the Tag entity by ID if the given value is not nil.
func (ttuo *TagTranslationUpdateOne) SetNillableTagID(id *uuid.UUID) *TagTranslationUpdateOne {
	if id != nil {
		ttuo = ttuo.SetTagID(*id)
	}
	return ttuo
}

// SetTag sets the "tag" edge to the Tag entity.
func (ttuo *TagTranslationUpdateOne) SetTag(t *Tag) *TagTranslationUpdateOne {
	return ttuo.SetTagID(t.ID)
}

// Mutation returns the TagTranslationMutation object of the builder.
func (ttuo *TagTranslationUpdateOne) Mutation() *TagTranslationMutation {
	return ttuo.mutation
}

// ClearTag clears the "tag" edge to the Tag entity.
func (ttuo *TagTranslationUpdateOne) ClearTag() *TagTranslationUpdateOne {
	ttuo.mutation.ClearTag()
	return ttuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ttuo *TagTranslationUpdateOne) Select(field string, fields ...string) *TagTranslationUpdateOne {
	ttuo.fields = append([]string{field}, fields...)
	return ttuo
}

// Save executes the query and returns the updated TagTranslation entity.
func (ttuo *TagTranslationUpdateOne) Save(ctx context.Context) (*TagTranslation, error) {
	var (
		err  error
		node *TagTranslation
	)
	if len(ttuo.hooks) == 0 {
		if err = ttuo.check(); err != nil {
			return nil, err
		}
		node, err = ttuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TagTranslationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ttuo.check(); err != nil {
				return nil, err
			}
			ttuo.mutation = mutation
			node, err = ttuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ttuo.hooks) - 1; i >= 0; i-- {
			if ttuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ttuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ttuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TagTranslation)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TagTranslationMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ttuo *TagTranslationUpdateOne) SaveX(ctx context.Context) *TagTranslation {
	node, err := ttuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ttuo *TagTranslationUpdateOne) Exec(ctx context.Context) error {
	_, err := ttuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttuo *TagTranslationUpdateOne) ExecX(ctx context.Context) {
	if err := ttuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttuo *TagTranslationUpdateOne) check() error {
	if v, ok := ttuo.mutation.Content(); ok {
		if err := tagtranslation.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "TagTranslation.content": %w`, err)}
		}
	}
	return nil
}

func (ttuo *TagTranslationUpdateOne) sqlSave(ctx context.Context) (_node *TagTranslation, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tagtranslation.Table,
			Columns: tagtranslation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: tagtranslation.FieldID,
			},
		},
	}
	id, ok := ttuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TagTranslation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ttuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tagtranslation.FieldID)
		for _, f := range fields {
			if !tagtranslation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tagtranslation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ttuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ttuo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tagtranslation.FieldContent,
		})
	}
	if ttuo.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tagtranslation.TagTable,
			Columns: []string{tagtranslation.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttuo.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tagtranslation.TagTable,
			Columns: []string{tagtranslation.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TagTranslation{config: ttuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ttuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tagtranslation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
