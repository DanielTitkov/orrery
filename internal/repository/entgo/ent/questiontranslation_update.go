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
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/question"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/questiontranslation"
)

// QuestionTranslationUpdate is the builder for updating QuestionTranslation entities.
type QuestionTranslationUpdate struct {
	config
	hooks    []Hook
	mutation *QuestionTranslationMutation
}

// Where appends a list predicates to the QuestionTranslationUpdate builder.
func (qtu *QuestionTranslationUpdate) Where(ps ...predicate.QuestionTranslation) *QuestionTranslationUpdate {
	qtu.mutation.Where(ps...)
	return qtu
}

// SetContent sets the "content" field.
func (qtu *QuestionTranslationUpdate) SetContent(s string) *QuestionTranslationUpdate {
	qtu.mutation.SetContent(s)
	return qtu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (qtu *QuestionTranslationUpdate) SetNillableContent(s *string) *QuestionTranslationUpdate {
	if s != nil {
		qtu.SetContent(*s)
	}
	return qtu
}

// ClearContent clears the value of the "content" field.
func (qtu *QuestionTranslationUpdate) ClearContent() *QuestionTranslationUpdate {
	qtu.mutation.ClearContent()
	return qtu
}

// SetHeaderContent sets the "header_content" field.
func (qtu *QuestionTranslationUpdate) SetHeaderContent(s string) *QuestionTranslationUpdate {
	qtu.mutation.SetHeaderContent(s)
	return qtu
}

// SetNillableHeaderContent sets the "header_content" field if the given value is not nil.
func (qtu *QuestionTranslationUpdate) SetNillableHeaderContent(s *string) *QuestionTranslationUpdate {
	if s != nil {
		qtu.SetHeaderContent(*s)
	}
	return qtu
}

// ClearHeaderContent clears the value of the "header_content" field.
func (qtu *QuestionTranslationUpdate) ClearHeaderContent() *QuestionTranslationUpdate {
	qtu.mutation.ClearHeaderContent()
	return qtu
}

// SetFooterContent sets the "footer_content" field.
func (qtu *QuestionTranslationUpdate) SetFooterContent(s string) *QuestionTranslationUpdate {
	qtu.mutation.SetFooterContent(s)
	return qtu
}

// SetNillableFooterContent sets the "footer_content" field if the given value is not nil.
func (qtu *QuestionTranslationUpdate) SetNillableFooterContent(s *string) *QuestionTranslationUpdate {
	if s != nil {
		qtu.SetFooterContent(*s)
	}
	return qtu
}

// ClearFooterContent clears the value of the "footer_content" field.
func (qtu *QuestionTranslationUpdate) ClearFooterContent() *QuestionTranslationUpdate {
	qtu.mutation.ClearFooterContent()
	return qtu
}

// SetQuestionID sets the "question" edge to the Question entity by ID.
func (qtu *QuestionTranslationUpdate) SetQuestionID(id uuid.UUID) *QuestionTranslationUpdate {
	qtu.mutation.SetQuestionID(id)
	return qtu
}

// SetNillableQuestionID sets the "question" edge to the Question entity by ID if the given value is not nil.
func (qtu *QuestionTranslationUpdate) SetNillableQuestionID(id *uuid.UUID) *QuestionTranslationUpdate {
	if id != nil {
		qtu = qtu.SetQuestionID(*id)
	}
	return qtu
}

// SetQuestion sets the "question" edge to the Question entity.
func (qtu *QuestionTranslationUpdate) SetQuestion(q *Question) *QuestionTranslationUpdate {
	return qtu.SetQuestionID(q.ID)
}

// Mutation returns the QuestionTranslationMutation object of the builder.
func (qtu *QuestionTranslationUpdate) Mutation() *QuestionTranslationMutation {
	return qtu.mutation
}

// ClearQuestion clears the "question" edge to the Question entity.
func (qtu *QuestionTranslationUpdate) ClearQuestion() *QuestionTranslationUpdate {
	qtu.mutation.ClearQuestion()
	return qtu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (qtu *QuestionTranslationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(qtu.hooks) == 0 {
		affected, err = qtu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QuestionTranslationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			qtu.mutation = mutation
			affected, err = qtu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(qtu.hooks) - 1; i >= 0; i-- {
			if qtu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = qtu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qtu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (qtu *QuestionTranslationUpdate) SaveX(ctx context.Context) int {
	affected, err := qtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (qtu *QuestionTranslationUpdate) Exec(ctx context.Context) error {
	_, err := qtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qtu *QuestionTranslationUpdate) ExecX(ctx context.Context) {
	if err := qtu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (qtu *QuestionTranslationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   questiontranslation.Table,
			Columns: questiontranslation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: questiontranslation.FieldID,
			},
		},
	}
	if ps := qtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := qtu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: questiontranslation.FieldContent,
		})
	}
	if qtu.mutation.ContentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: questiontranslation.FieldContent,
		})
	}
	if value, ok := qtu.mutation.HeaderContent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: questiontranslation.FieldHeaderContent,
		})
	}
	if qtu.mutation.HeaderContentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: questiontranslation.FieldHeaderContent,
		})
	}
	if value, ok := qtu.mutation.FooterContent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: questiontranslation.FieldFooterContent,
		})
	}
	if qtu.mutation.FooterContentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: questiontranslation.FieldFooterContent,
		})
	}
	if qtu.mutation.QuestionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   questiontranslation.QuestionTable,
			Columns: []string{questiontranslation.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: question.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qtu.mutation.QuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   questiontranslation.QuestionTable,
			Columns: []string{questiontranslation.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: question.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, qtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{questiontranslation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// QuestionTranslationUpdateOne is the builder for updating a single QuestionTranslation entity.
type QuestionTranslationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *QuestionTranslationMutation
}

// SetContent sets the "content" field.
func (qtuo *QuestionTranslationUpdateOne) SetContent(s string) *QuestionTranslationUpdateOne {
	qtuo.mutation.SetContent(s)
	return qtuo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (qtuo *QuestionTranslationUpdateOne) SetNillableContent(s *string) *QuestionTranslationUpdateOne {
	if s != nil {
		qtuo.SetContent(*s)
	}
	return qtuo
}

// ClearContent clears the value of the "content" field.
func (qtuo *QuestionTranslationUpdateOne) ClearContent() *QuestionTranslationUpdateOne {
	qtuo.mutation.ClearContent()
	return qtuo
}

// SetHeaderContent sets the "header_content" field.
func (qtuo *QuestionTranslationUpdateOne) SetHeaderContent(s string) *QuestionTranslationUpdateOne {
	qtuo.mutation.SetHeaderContent(s)
	return qtuo
}

// SetNillableHeaderContent sets the "header_content" field if the given value is not nil.
func (qtuo *QuestionTranslationUpdateOne) SetNillableHeaderContent(s *string) *QuestionTranslationUpdateOne {
	if s != nil {
		qtuo.SetHeaderContent(*s)
	}
	return qtuo
}

// ClearHeaderContent clears the value of the "header_content" field.
func (qtuo *QuestionTranslationUpdateOne) ClearHeaderContent() *QuestionTranslationUpdateOne {
	qtuo.mutation.ClearHeaderContent()
	return qtuo
}

// SetFooterContent sets the "footer_content" field.
func (qtuo *QuestionTranslationUpdateOne) SetFooterContent(s string) *QuestionTranslationUpdateOne {
	qtuo.mutation.SetFooterContent(s)
	return qtuo
}

// SetNillableFooterContent sets the "footer_content" field if the given value is not nil.
func (qtuo *QuestionTranslationUpdateOne) SetNillableFooterContent(s *string) *QuestionTranslationUpdateOne {
	if s != nil {
		qtuo.SetFooterContent(*s)
	}
	return qtuo
}

// ClearFooterContent clears the value of the "footer_content" field.
func (qtuo *QuestionTranslationUpdateOne) ClearFooterContent() *QuestionTranslationUpdateOne {
	qtuo.mutation.ClearFooterContent()
	return qtuo
}

// SetQuestionID sets the "question" edge to the Question entity by ID.
func (qtuo *QuestionTranslationUpdateOne) SetQuestionID(id uuid.UUID) *QuestionTranslationUpdateOne {
	qtuo.mutation.SetQuestionID(id)
	return qtuo
}

// SetNillableQuestionID sets the "question" edge to the Question entity by ID if the given value is not nil.
func (qtuo *QuestionTranslationUpdateOne) SetNillableQuestionID(id *uuid.UUID) *QuestionTranslationUpdateOne {
	if id != nil {
		qtuo = qtuo.SetQuestionID(*id)
	}
	return qtuo
}

// SetQuestion sets the "question" edge to the Question entity.
func (qtuo *QuestionTranslationUpdateOne) SetQuestion(q *Question) *QuestionTranslationUpdateOne {
	return qtuo.SetQuestionID(q.ID)
}

// Mutation returns the QuestionTranslationMutation object of the builder.
func (qtuo *QuestionTranslationUpdateOne) Mutation() *QuestionTranslationMutation {
	return qtuo.mutation
}

// ClearQuestion clears the "question" edge to the Question entity.
func (qtuo *QuestionTranslationUpdateOne) ClearQuestion() *QuestionTranslationUpdateOne {
	qtuo.mutation.ClearQuestion()
	return qtuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (qtuo *QuestionTranslationUpdateOne) Select(field string, fields ...string) *QuestionTranslationUpdateOne {
	qtuo.fields = append([]string{field}, fields...)
	return qtuo
}

// Save executes the query and returns the updated QuestionTranslation entity.
func (qtuo *QuestionTranslationUpdateOne) Save(ctx context.Context) (*QuestionTranslation, error) {
	var (
		err  error
		node *QuestionTranslation
	)
	if len(qtuo.hooks) == 0 {
		node, err = qtuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QuestionTranslationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			qtuo.mutation = mutation
			node, err = qtuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(qtuo.hooks) - 1; i >= 0; i-- {
			if qtuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = qtuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, qtuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*QuestionTranslation)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from QuestionTranslationMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (qtuo *QuestionTranslationUpdateOne) SaveX(ctx context.Context) *QuestionTranslation {
	node, err := qtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (qtuo *QuestionTranslationUpdateOne) Exec(ctx context.Context) error {
	_, err := qtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qtuo *QuestionTranslationUpdateOne) ExecX(ctx context.Context) {
	if err := qtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (qtuo *QuestionTranslationUpdateOne) sqlSave(ctx context.Context) (_node *QuestionTranslation, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   questiontranslation.Table,
			Columns: questiontranslation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: questiontranslation.FieldID,
			},
		},
	}
	id, ok := qtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "QuestionTranslation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := qtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, questiontranslation.FieldID)
		for _, f := range fields {
			if !questiontranslation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != questiontranslation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := qtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := qtuo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: questiontranslation.FieldContent,
		})
	}
	if qtuo.mutation.ContentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: questiontranslation.FieldContent,
		})
	}
	if value, ok := qtuo.mutation.HeaderContent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: questiontranslation.FieldHeaderContent,
		})
	}
	if qtuo.mutation.HeaderContentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: questiontranslation.FieldHeaderContent,
		})
	}
	if value, ok := qtuo.mutation.FooterContent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: questiontranslation.FieldFooterContent,
		})
	}
	if qtuo.mutation.FooterContentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: questiontranslation.FieldFooterContent,
		})
	}
	if qtuo.mutation.QuestionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   questiontranslation.QuestionTable,
			Columns: []string{questiontranslation.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: question.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qtuo.mutation.QuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   questiontranslation.QuestionTable,
			Columns: []string{questiontranslation.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: question.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &QuestionTranslation{config: qtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, qtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{questiontranslation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
