// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/testdisplay"
)

// TestDisplayDelete is the builder for deleting a TestDisplay entity.
type TestDisplayDelete struct {
	config
	hooks    []Hook
	mutation *TestDisplayMutation
}

// Where appends a list predicates to the TestDisplayDelete builder.
func (tdd *TestDisplayDelete) Where(ps ...predicate.TestDisplay) *TestDisplayDelete {
	tdd.mutation.Where(ps...)
	return tdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tdd *TestDisplayDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tdd.hooks) == 0 {
		affected, err = tdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TestDisplayMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tdd.mutation = mutation
			affected, err = tdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tdd.hooks) - 1; i >= 0; i-- {
			if tdd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tdd *TestDisplayDelete) ExecX(ctx context.Context) int {
	n, err := tdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tdd *TestDisplayDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: testdisplay.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: testdisplay.FieldID,
			},
		},
	}
	if ps := tdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, tdd.driver, _spec)
}

// TestDisplayDeleteOne is the builder for deleting a single TestDisplay entity.
type TestDisplayDeleteOne struct {
	tdd *TestDisplayDelete
}

// Exec executes the deletion query.
func (tddo *TestDisplayDeleteOne) Exec(ctx context.Context) error {
	n, err := tddo.tdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{testdisplay.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tddo *TestDisplayDeleteOne) ExecX(ctx context.Context) {
	tddo.tdd.ExecX(ctx)
}