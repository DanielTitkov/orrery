// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/tag"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/tagtranslation"
)

// TagTranslationCreate is the builder for creating a TagTranslation entity.
type TagTranslationCreate struct {
	config
	mutation *TagTranslationMutation
	hooks    []Hook
}

// SetLocale sets the "locale" field.
func (ttc *TagTranslationCreate) SetLocale(t tagtranslation.Locale) *TagTranslationCreate {
	ttc.mutation.SetLocale(t)
	return ttc
}

// SetContent sets the "content" field.
func (ttc *TagTranslationCreate) SetContent(s string) *TagTranslationCreate {
	ttc.mutation.SetContent(s)
	return ttc
}

// SetID sets the "id" field.
func (ttc *TagTranslationCreate) SetID(u uuid.UUID) *TagTranslationCreate {
	ttc.mutation.SetID(u)
	return ttc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ttc *TagTranslationCreate) SetNillableID(u *uuid.UUID) *TagTranslationCreate {
	if u != nil {
		ttc.SetID(*u)
	}
	return ttc
}

// SetTagID sets the "tag" edge to the Tag entity by ID.
func (ttc *TagTranslationCreate) SetTagID(id uuid.UUID) *TagTranslationCreate {
	ttc.mutation.SetTagID(id)
	return ttc
}

// SetNillableTagID sets the "tag" edge to the Tag entity by ID if the given value is not nil.
func (ttc *TagTranslationCreate) SetNillableTagID(id *uuid.UUID) *TagTranslationCreate {
	if id != nil {
		ttc = ttc.SetTagID(*id)
	}
	return ttc
}

// SetTag sets the "tag" edge to the Tag entity.
func (ttc *TagTranslationCreate) SetTag(t *Tag) *TagTranslationCreate {
	return ttc.SetTagID(t.ID)
}

// Mutation returns the TagTranslationMutation object of the builder.
func (ttc *TagTranslationCreate) Mutation() *TagTranslationMutation {
	return ttc.mutation
}

// Save creates the TagTranslation in the database.
func (ttc *TagTranslationCreate) Save(ctx context.Context) (*TagTranslation, error) {
	var (
		err  error
		node *TagTranslation
	)
	ttc.defaults()
	if len(ttc.hooks) == 0 {
		if err = ttc.check(); err != nil {
			return nil, err
		}
		node, err = ttc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TagTranslationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ttc.check(); err != nil {
				return nil, err
			}
			ttc.mutation = mutation
			if node, err = ttc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ttc.hooks) - 1; i >= 0; i-- {
			if ttc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ttc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ttc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (ttc *TagTranslationCreate) SaveX(ctx context.Context) *TagTranslation {
	v, err := ttc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ttc *TagTranslationCreate) Exec(ctx context.Context) error {
	_, err := ttc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttc *TagTranslationCreate) ExecX(ctx context.Context) {
	if err := ttc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ttc *TagTranslationCreate) defaults() {
	if _, ok := ttc.mutation.ID(); !ok {
		v := tagtranslation.DefaultID()
		ttc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttc *TagTranslationCreate) check() error {
	if _, ok := ttc.mutation.Locale(); !ok {
		return &ValidationError{Name: "locale", err: errors.New(`ent: missing required field "TagTranslation.locale"`)}
	}
	if v, ok := ttc.mutation.Locale(); ok {
		if err := tagtranslation.LocaleValidator(v); err != nil {
			return &ValidationError{Name: "locale", err: fmt.Errorf(`ent: validator failed for field "TagTranslation.locale": %w`, err)}
		}
	}
	if _, ok := ttc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "TagTranslation.content"`)}
	}
	if v, ok := ttc.mutation.Content(); ok {
		if err := tagtranslation.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "TagTranslation.content": %w`, err)}
		}
	}
	return nil
}

func (ttc *TagTranslationCreate) sqlSave(ctx context.Context) (*TagTranslation, error) {
	_node, _spec := ttc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ttc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (ttc *TagTranslationCreate) createSpec() (*TagTranslation, *sqlgraph.CreateSpec) {
	var (
		_node = &TagTranslation{config: ttc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tagtranslation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: tagtranslation.FieldID,
			},
		}
	)
	if id, ok := ttc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ttc.mutation.Locale(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: tagtranslation.FieldLocale,
		})
		_node.Locale = value
	}
	if value, ok := ttc.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tagtranslation.FieldContent,
		})
		_node.Content = value
	}
	if nodes := ttc.mutation.TagIDs(); len(nodes) > 0 {
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
		_node.tag_translations = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TagTranslationCreateBulk is the builder for creating many TagTranslation entities in bulk.
type TagTranslationCreateBulk struct {
	config
	builders []*TagTranslationCreate
}

// Save creates the TagTranslation entities in the database.
func (ttcb *TagTranslationCreateBulk) Save(ctx context.Context) ([]*TagTranslation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ttcb.builders))
	nodes := make([]*TagTranslation, len(ttcb.builders))
	mutators := make([]Mutator, len(ttcb.builders))
	for i := range ttcb.builders {
		func(i int, root context.Context) {
			builder := ttcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TagTranslationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ttcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ttcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ttcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ttcb *TagTranslationCreateBulk) SaveX(ctx context.Context) []*TagTranslation {
	v, err := ttcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ttcb *TagTranslationCreateBulk) Exec(ctx context.Context) error {
	_, err := ttcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttcb *TagTranslationCreateBulk) ExecX(ctx context.Context) {
	if err := ttcb.Exec(ctx); err != nil {
		panic(err)
	}
}
