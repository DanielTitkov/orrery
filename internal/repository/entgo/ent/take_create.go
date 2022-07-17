// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/response"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/result"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/take"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/test"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/user"
)

// TakeCreate is the builder for creating a Take entity.
type TakeCreate struct {
	config
	mutation *TakeMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (tc *TakeCreate) SetCreateTime(t time.Time) *TakeCreate {
	tc.mutation.SetCreateTime(t)
	return tc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (tc *TakeCreate) SetNillableCreateTime(t *time.Time) *TakeCreate {
	if t != nil {
		tc.SetCreateTime(*t)
	}
	return tc
}

// SetUpdateTime sets the "update_time" field.
func (tc *TakeCreate) SetUpdateTime(t time.Time) *TakeCreate {
	tc.mutation.SetUpdateTime(t)
	return tc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (tc *TakeCreate) SetNillableUpdateTime(t *time.Time) *TakeCreate {
	if t != nil {
		tc.SetUpdateTime(*t)
	}
	return tc
}

// SetSeed sets the "seed" field.
func (tc *TakeCreate) SetSeed(i int64) *TakeCreate {
	tc.mutation.SetSeed(i)
	return tc
}

// SetProgress sets the "progress" field.
func (tc *TakeCreate) SetProgress(i int) *TakeCreate {
	tc.mutation.SetProgress(i)
	return tc
}

// SetNillableProgress sets the "progress" field if the given value is not nil.
func (tc *TakeCreate) SetNillableProgress(i *int) *TakeCreate {
	if i != nil {
		tc.SetProgress(*i)
	}
	return tc
}

// SetPage sets the "page" field.
func (tc *TakeCreate) SetPage(i int) *TakeCreate {
	tc.mutation.SetPage(i)
	return tc
}

// SetNillablePage sets the "page" field if the given value is not nil.
func (tc *TakeCreate) SetNillablePage(i *int) *TakeCreate {
	if i != nil {
		tc.SetPage(*i)
	}
	return tc
}

// SetStartTime sets the "start_time" field.
func (tc *TakeCreate) SetStartTime(t time.Time) *TakeCreate {
	tc.mutation.SetStartTime(t)
	return tc
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (tc *TakeCreate) SetNillableStartTime(t *time.Time) *TakeCreate {
	if t != nil {
		tc.SetStartTime(*t)
	}
	return tc
}

// SetEndTime sets the "end_time" field.
func (tc *TakeCreate) SetEndTime(t time.Time) *TakeCreate {
	tc.mutation.SetEndTime(t)
	return tc
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (tc *TakeCreate) SetNillableEndTime(t *time.Time) *TakeCreate {
	if t != nil {
		tc.SetEndTime(*t)
	}
	return tc
}

// SetSuspicious sets the "suspicious" field.
func (tc *TakeCreate) SetSuspicious(b bool) *TakeCreate {
	tc.mutation.SetSuspicious(b)
	return tc
}

// SetNillableSuspicious sets the "suspicious" field if the given value is not nil.
func (tc *TakeCreate) SetNillableSuspicious(b *bool) *TakeCreate {
	if b != nil {
		tc.SetSuspicious(*b)
	}
	return tc
}

// SetStatus sets the "status" field.
func (tc *TakeCreate) SetStatus(t take.Status) *TakeCreate {
	tc.mutation.SetStatus(t)
	return tc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tc *TakeCreate) SetNillableStatus(t *take.Status) *TakeCreate {
	if t != nil {
		tc.SetStatus(*t)
	}
	return tc
}

// SetMark sets the "mark" field.
func (tc *TakeCreate) SetMark(i int) *TakeCreate {
	tc.mutation.SetMark(i)
	return tc
}

// SetNillableMark sets the "mark" field if the given value is not nil.
func (tc *TakeCreate) SetNillableMark(i *int) *TakeCreate {
	if i != nil {
		tc.SetMark(*i)
	}
	return tc
}

// SetInLocale sets the "in_locale" field.
func (tc *TakeCreate) SetInLocale(tl take.InLocale) *TakeCreate {
	tc.mutation.SetInLocale(tl)
	return tc
}

// SetNillableInLocale sets the "in_locale" field if the given value is not nil.
func (tc *TakeCreate) SetNillableInLocale(tl *take.InLocale) *TakeCreate {
	if tl != nil {
		tc.SetInLocale(*tl)
	}
	return tc
}

// SetMeta sets the "meta" field.
func (tc *TakeCreate) SetMeta(m map[string]interface{}) *TakeCreate {
	tc.mutation.SetMeta(m)
	return tc
}

// SetID sets the "id" field.
func (tc *TakeCreate) SetID(u uuid.UUID) *TakeCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TakeCreate) SetNillableID(u *uuid.UUID) *TakeCreate {
	if u != nil {
		tc.SetID(*u)
	}
	return tc
}

// AddResponseIDs adds the "responses" edge to the Response entity by IDs.
func (tc *TakeCreate) AddResponseIDs(ids ...uuid.UUID) *TakeCreate {
	tc.mutation.AddResponseIDs(ids...)
	return tc
}

// AddResponses adds the "responses" edges to the Response entity.
func (tc *TakeCreate) AddResponses(r ...*Response) *TakeCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tc.AddResponseIDs(ids...)
}

// AddResultIDs adds the "results" edge to the Result entity by IDs.
func (tc *TakeCreate) AddResultIDs(ids ...uuid.UUID) *TakeCreate {
	tc.mutation.AddResultIDs(ids...)
	return tc
}

// AddResults adds the "results" edges to the Result entity.
func (tc *TakeCreate) AddResults(r ...*Result) *TakeCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tc.AddResultIDs(ids...)
}

// SetTestID sets the "test" edge to the Test entity by ID.
func (tc *TakeCreate) SetTestID(id uuid.UUID) *TakeCreate {
	tc.mutation.SetTestID(id)
	return tc
}

// SetTest sets the "test" edge to the Test entity.
func (tc *TakeCreate) SetTest(t *Test) *TakeCreate {
	return tc.SetTestID(t.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (tc *TakeCreate) SetUserID(id uuid.UUID) *TakeCreate {
	tc.mutation.SetUserID(id)
	return tc
}

// SetUser sets the "user" edge to the User entity.
func (tc *TakeCreate) SetUser(u *User) *TakeCreate {
	return tc.SetUserID(u.ID)
}

// Mutation returns the TakeMutation object of the builder.
func (tc *TakeCreate) Mutation() *TakeMutation {
	return tc.mutation
}

// Save creates the Take in the database.
func (tc *TakeCreate) Save(ctx context.Context) (*Take, error) {
	var (
		err  error
		node *Take
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TakeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Take)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TakeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TakeCreate) SaveX(ctx context.Context) *Take {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TakeCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TakeCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TakeCreate) defaults() {
	if _, ok := tc.mutation.CreateTime(); !ok {
		v := take.DefaultCreateTime()
		tc.mutation.SetCreateTime(v)
	}
	if _, ok := tc.mutation.UpdateTime(); !ok {
		v := take.DefaultUpdateTime()
		tc.mutation.SetUpdateTime(v)
	}
	if _, ok := tc.mutation.Progress(); !ok {
		v := take.DefaultProgress
		tc.mutation.SetProgress(v)
	}
	if _, ok := tc.mutation.Page(); !ok {
		v := take.DefaultPage
		tc.mutation.SetPage(v)
	}
	if _, ok := tc.mutation.Suspicious(); !ok {
		v := take.DefaultSuspicious
		tc.mutation.SetSuspicious(v)
	}
	if _, ok := tc.mutation.Status(); !ok {
		v := take.DefaultStatus
		tc.mutation.SetStatus(v)
	}
	if _, ok := tc.mutation.InLocale(); !ok {
		v := take.DefaultInLocale
		tc.mutation.SetInLocale(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		v := take.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TakeCreate) check() error {
	if _, ok := tc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Take.create_time"`)}
	}
	if _, ok := tc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Take.update_time"`)}
	}
	if _, ok := tc.mutation.Seed(); !ok {
		return &ValidationError{Name: "seed", err: errors.New(`ent: missing required field "Take.seed"`)}
	}
	if _, ok := tc.mutation.Progress(); !ok {
		return &ValidationError{Name: "progress", err: errors.New(`ent: missing required field "Take.progress"`)}
	}
	if _, ok := tc.mutation.Page(); !ok {
		return &ValidationError{Name: "page", err: errors.New(`ent: missing required field "Take.page"`)}
	}
	if _, ok := tc.mutation.Suspicious(); !ok {
		return &ValidationError{Name: "suspicious", err: errors.New(`ent: missing required field "Take.suspicious"`)}
	}
	if _, ok := tc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Take.status"`)}
	}
	if v, ok := tc.mutation.Status(); ok {
		if err := take.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Take.status": %w`, err)}
		}
	}
	if v, ok := tc.mutation.Mark(); ok {
		if err := take.MarkValidator(v); err != nil {
			return &ValidationError{Name: "mark", err: fmt.Errorf(`ent: validator failed for field "Take.mark": %w`, err)}
		}
	}
	if _, ok := tc.mutation.InLocale(); !ok {
		return &ValidationError{Name: "in_locale", err: errors.New(`ent: missing required field "Take.in_locale"`)}
	}
	if v, ok := tc.mutation.InLocale(); ok {
		if err := take.InLocaleValidator(v); err != nil {
			return &ValidationError{Name: "in_locale", err: fmt.Errorf(`ent: validator failed for field "Take.in_locale": %w`, err)}
		}
	}
	if _, ok := tc.mutation.TestID(); !ok {
		return &ValidationError{Name: "test", err: errors.New(`ent: missing required edge "Take.test"`)}
	}
	if _, ok := tc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Take.user"`)}
	}
	return nil
}

func (tc *TakeCreate) sqlSave(ctx context.Context) (*Take, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
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

func (tc *TakeCreate) createSpec() (*Take, *sqlgraph.CreateSpec) {
	var (
		_node = &Take{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: take.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: take.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: take.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := tc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: take.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := tc.mutation.Seed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: take.FieldSeed,
		})
		_node.Seed = value
	}
	if value, ok := tc.mutation.Progress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: take.FieldProgress,
		})
		_node.Progress = value
	}
	if value, ok := tc.mutation.Page(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: take.FieldPage,
		})
		_node.Page = value
	}
	if value, ok := tc.mutation.StartTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: take.FieldStartTime,
		})
		_node.StartTime = &value
	}
	if value, ok := tc.mutation.EndTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: take.FieldEndTime,
		})
		_node.EndTime = &value
	}
	if value, ok := tc.mutation.Suspicious(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: take.FieldSuspicious,
		})
		_node.Suspicious = value
	}
	if value, ok := tc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: take.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := tc.mutation.Mark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: take.FieldMark,
		})
		_node.Mark = &value
	}
	if value, ok := tc.mutation.InLocale(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: take.FieldInLocale,
		})
		_node.InLocale = value
	}
	if value, ok := tc.mutation.Meta(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: take.FieldMeta,
		})
		_node.Meta = value
	}
	if nodes := tc.mutation.ResponsesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   take.ResponsesTable,
			Columns: []string{take.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: response.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ResultsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   take.ResultsTable,
			Columns: []string{take.ResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: result.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   take.TestTable,
			Columns: []string{take.TestColumn},
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
		_node.test_takes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   take.UserTable,
			Columns: []string{take.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_takes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TakeCreateBulk is the builder for creating many Take entities in bulk.
type TakeCreateBulk struct {
	config
	builders []*TakeCreate
}

// Save creates the Take entities in the database.
func (tcb *TakeCreateBulk) Save(ctx context.Context) ([]*Take, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Take, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TakeMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TakeCreateBulk) SaveX(ctx context.Context) []*Take {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TakeCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TakeCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
