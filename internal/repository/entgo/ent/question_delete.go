// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/predicate"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/question"
)

// QuestionDelete is the builder for deleting a Question entity.
type QuestionDelete struct {
	config
	hooks    []Hook
	mutation *QuestionMutation
}

// Where appends a list predicates to the QuestionDelete builder.
func (qd *QuestionDelete) Where(ps ...predicate.Question) *QuestionDelete {
	qd.mutation.Where(ps...)
	return qd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (qd *QuestionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(qd.hooks) == 0 {
		affected, err = qd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QuestionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			qd.mutation = mutation
			affected, err = qd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(qd.hooks) - 1; i >= 0; i-- {
			if qd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = qd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (qd *QuestionDelete) ExecX(ctx context.Context) int {
	n, err := qd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (qd *QuestionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: question.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: question.FieldID,
			},
		},
	}
	if ps := qd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, qd.driver, _spec)
}

// QuestionDeleteOne is the builder for deleting a single Question entity.
type QuestionDeleteOne struct {
	qd *QuestionDelete
}

// Exec executes the deletion query.
func (qdo *QuestionDeleteOne) Exec(ctx context.Context) error {
	n, err := qdo.qd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{question.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (qdo *QuestionDeleteOne) ExecX(ctx context.Context) {
	qdo.qd.ExecX(ctx)
}
