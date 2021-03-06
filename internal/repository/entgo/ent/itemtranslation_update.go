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
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/item"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/itemtranslation"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/predicate"
)

// ItemTranslationUpdate is the builder for updating ItemTranslation entities.
type ItemTranslationUpdate struct {
	config
	hooks    []Hook
	mutation *ItemTranslationMutation
}

// Where appends a list predicates to the ItemTranslationUpdate builder.
func (itu *ItemTranslationUpdate) Where(ps ...predicate.ItemTranslation) *ItemTranslationUpdate {
	itu.mutation.Where(ps...)
	return itu
}

// SetContent sets the "content" field.
func (itu *ItemTranslationUpdate) SetContent(s string) *ItemTranslationUpdate {
	itu.mutation.SetContent(s)
	return itu
}

// SetItemID sets the "item" edge to the Item entity by ID.
func (itu *ItemTranslationUpdate) SetItemID(id uuid.UUID) *ItemTranslationUpdate {
	itu.mutation.SetItemID(id)
	return itu
}

// SetNillableItemID sets the "item" edge to the Item entity by ID if the given value is not nil.
func (itu *ItemTranslationUpdate) SetNillableItemID(id *uuid.UUID) *ItemTranslationUpdate {
	if id != nil {
		itu = itu.SetItemID(*id)
	}
	return itu
}

// SetItem sets the "item" edge to the Item entity.
func (itu *ItemTranslationUpdate) SetItem(i *Item) *ItemTranslationUpdate {
	return itu.SetItemID(i.ID)
}

// Mutation returns the ItemTranslationMutation object of the builder.
func (itu *ItemTranslationUpdate) Mutation() *ItemTranslationMutation {
	return itu.mutation
}

// ClearItem clears the "item" edge to the Item entity.
func (itu *ItemTranslationUpdate) ClearItem() *ItemTranslationUpdate {
	itu.mutation.ClearItem()
	return itu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (itu *ItemTranslationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(itu.hooks) == 0 {
		if err = itu.check(); err != nil {
			return 0, err
		}
		affected, err = itu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemTranslationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = itu.check(); err != nil {
				return 0, err
			}
			itu.mutation = mutation
			affected, err = itu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(itu.hooks) - 1; i >= 0; i-- {
			if itu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = itu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, itu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (itu *ItemTranslationUpdate) SaveX(ctx context.Context) int {
	affected, err := itu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (itu *ItemTranslationUpdate) Exec(ctx context.Context) error {
	_, err := itu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (itu *ItemTranslationUpdate) ExecX(ctx context.Context) {
	if err := itu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (itu *ItemTranslationUpdate) check() error {
	if v, ok := itu.mutation.Content(); ok {
		if err := itemtranslation.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "ItemTranslation.content": %w`, err)}
		}
	}
	return nil
}

func (itu *ItemTranslationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   itemtranslation.Table,
			Columns: itemtranslation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: itemtranslation.FieldID,
			},
		},
	}
	if ps := itu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := itu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: itemtranslation.FieldContent,
		})
	}
	if itu.mutation.ItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   itemtranslation.ItemTable,
			Columns: []string{itemtranslation.ItemColumn},
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
	if nodes := itu.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   itemtranslation.ItemTable,
			Columns: []string{itemtranslation.ItemColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, itu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{itemtranslation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ItemTranslationUpdateOne is the builder for updating a single ItemTranslation entity.
type ItemTranslationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ItemTranslationMutation
}

// SetContent sets the "content" field.
func (ituo *ItemTranslationUpdateOne) SetContent(s string) *ItemTranslationUpdateOne {
	ituo.mutation.SetContent(s)
	return ituo
}

// SetItemID sets the "item" edge to the Item entity by ID.
func (ituo *ItemTranslationUpdateOne) SetItemID(id uuid.UUID) *ItemTranslationUpdateOne {
	ituo.mutation.SetItemID(id)
	return ituo
}

// SetNillableItemID sets the "item" edge to the Item entity by ID if the given value is not nil.
func (ituo *ItemTranslationUpdateOne) SetNillableItemID(id *uuid.UUID) *ItemTranslationUpdateOne {
	if id != nil {
		ituo = ituo.SetItemID(*id)
	}
	return ituo
}

// SetItem sets the "item" edge to the Item entity.
func (ituo *ItemTranslationUpdateOne) SetItem(i *Item) *ItemTranslationUpdateOne {
	return ituo.SetItemID(i.ID)
}

// Mutation returns the ItemTranslationMutation object of the builder.
func (ituo *ItemTranslationUpdateOne) Mutation() *ItemTranslationMutation {
	return ituo.mutation
}

// ClearItem clears the "item" edge to the Item entity.
func (ituo *ItemTranslationUpdateOne) ClearItem() *ItemTranslationUpdateOne {
	ituo.mutation.ClearItem()
	return ituo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ituo *ItemTranslationUpdateOne) Select(field string, fields ...string) *ItemTranslationUpdateOne {
	ituo.fields = append([]string{field}, fields...)
	return ituo
}

// Save executes the query and returns the updated ItemTranslation entity.
func (ituo *ItemTranslationUpdateOne) Save(ctx context.Context) (*ItemTranslation, error) {
	var (
		err  error
		node *ItemTranslation
	)
	if len(ituo.hooks) == 0 {
		if err = ituo.check(); err != nil {
			return nil, err
		}
		node, err = ituo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemTranslationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ituo.check(); err != nil {
				return nil, err
			}
			ituo.mutation = mutation
			node, err = ituo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ituo.hooks) - 1; i >= 0; i-- {
			if ituo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ituo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ituo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ItemTranslation)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ItemTranslationMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ituo *ItemTranslationUpdateOne) SaveX(ctx context.Context) *ItemTranslation {
	node, err := ituo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ituo *ItemTranslationUpdateOne) Exec(ctx context.Context) error {
	_, err := ituo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ituo *ItemTranslationUpdateOne) ExecX(ctx context.Context) {
	if err := ituo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ituo *ItemTranslationUpdateOne) check() error {
	if v, ok := ituo.mutation.Content(); ok {
		if err := itemtranslation.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "ItemTranslation.content": %w`, err)}
		}
	}
	return nil
}

func (ituo *ItemTranslationUpdateOne) sqlSave(ctx context.Context) (_node *ItemTranslation, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   itemtranslation.Table,
			Columns: itemtranslation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: itemtranslation.FieldID,
			},
		},
	}
	id, ok := ituo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ItemTranslation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ituo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, itemtranslation.FieldID)
		for _, f := range fields {
			if !itemtranslation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != itemtranslation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ituo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ituo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: itemtranslation.FieldContent,
		})
	}
	if ituo.mutation.ItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   itemtranslation.ItemTable,
			Columns: []string{itemtranslation.ItemColumn},
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
	if nodes := ituo.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   itemtranslation.ItemTable,
			Columns: []string{itemtranslation.ItemColumn},
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
	_node = &ItemTranslation{config: ituo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ituo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{itemtranslation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
