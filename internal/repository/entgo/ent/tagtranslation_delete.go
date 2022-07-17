// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/predicate"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/tagtranslation"
)

// TagTranslationDelete is the builder for deleting a TagTranslation entity.
type TagTranslationDelete struct {
	config
	hooks    []Hook
	mutation *TagTranslationMutation
}

// Where appends a list predicates to the TagTranslationDelete builder.
func (ttd *TagTranslationDelete) Where(ps ...predicate.TagTranslation) *TagTranslationDelete {
	ttd.mutation.Where(ps...)
	return ttd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ttd *TagTranslationDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ttd.hooks) == 0 {
		affected, err = ttd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TagTranslationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ttd.mutation = mutation
			affected, err = ttd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ttd.hooks) - 1; i >= 0; i-- {
			if ttd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ttd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttd *TagTranslationDelete) ExecX(ctx context.Context) int {
	n, err := ttd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ttd *TagTranslationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: tagtranslation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: tagtranslation.FieldID,
			},
		},
	}
	if ps := ttd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ttd.driver, _spec)
}

// TagTranslationDeleteOne is the builder for deleting a single TagTranslation entity.
type TagTranslationDeleteOne struct {
	ttd *TagTranslationDelete
}

// Exec executes the deletion query.
func (ttdo *TagTranslationDeleteOne) Exec(ctx context.Context) error {
	n, err := ttdo.ttd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tagtranslation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ttdo *TagTranslationDeleteOne) ExecX(ctx context.Context) {
	ttdo.ttd.ExecX(ctx)
}
