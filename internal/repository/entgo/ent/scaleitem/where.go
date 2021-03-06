// Code generated by ent, DO NOT EDIT.

package scaleitem

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/predicate"
)

// Reverse applies equality check predicate on the "reverse" field. It's identical to ReverseEQ.
func Reverse(v bool) predicate.ScaleItem {
	return predicate.ScaleItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReverse), v))
	})
}

// ScaleID applies equality check predicate on the "scale_id" field. It's identical to ScaleIDEQ.
func ScaleID(v uuid.UUID) predicate.ScaleItem {
	return predicate.ScaleItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScaleID), v))
	})
}

// ItemID applies equality check predicate on the "item_id" field. It's identical to ItemIDEQ.
func ItemID(v uuid.UUID) predicate.ScaleItem {
	return predicate.ScaleItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldItemID), v))
	})
}

// ReverseEQ applies the EQ predicate on the "reverse" field.
func ReverseEQ(v bool) predicate.ScaleItem {
	return predicate.ScaleItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReverse), v))
	})
}

// ReverseNEQ applies the NEQ predicate on the "reverse" field.
func ReverseNEQ(v bool) predicate.ScaleItem {
	return predicate.ScaleItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReverse), v))
	})
}

// ScaleIDEQ applies the EQ predicate on the "scale_id" field.
func ScaleIDEQ(v uuid.UUID) predicate.ScaleItem {
	return predicate.ScaleItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScaleID), v))
	})
}

// ScaleIDNEQ applies the NEQ predicate on the "scale_id" field.
func ScaleIDNEQ(v uuid.UUID) predicate.ScaleItem {
	return predicate.ScaleItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldScaleID), v))
	})
}

// ScaleIDIn applies the In predicate on the "scale_id" field.
func ScaleIDIn(vs ...uuid.UUID) predicate.ScaleItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ScaleItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldScaleID), v...))
	})
}

// ScaleIDNotIn applies the NotIn predicate on the "scale_id" field.
func ScaleIDNotIn(vs ...uuid.UUID) predicate.ScaleItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ScaleItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldScaleID), v...))
	})
}

// ItemIDEQ applies the EQ predicate on the "item_id" field.
func ItemIDEQ(v uuid.UUID) predicate.ScaleItem {
	return predicate.ScaleItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldItemID), v))
	})
}

// ItemIDNEQ applies the NEQ predicate on the "item_id" field.
func ItemIDNEQ(v uuid.UUID) predicate.ScaleItem {
	return predicate.ScaleItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldItemID), v))
	})
}

// ItemIDIn applies the In predicate on the "item_id" field.
func ItemIDIn(vs ...uuid.UUID) predicate.ScaleItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ScaleItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldItemID), v...))
	})
}

// ItemIDNotIn applies the NotIn predicate on the "item_id" field.
func ItemIDNotIn(vs ...uuid.UUID) predicate.ScaleItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ScaleItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldItemID), v...))
	})
}

// HasItem applies the HasEdge predicate on the "item" edge.
func HasItem() predicate.ScaleItem {
	return predicate.ScaleItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, ItemColumn),
			sqlgraph.To(ItemInverseTable, ItemFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ItemTable, ItemColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasItemWith applies the HasEdge predicate on the "item" edge with a given conditions (other predicates).
func HasItemWith(preds ...predicate.Item) predicate.ScaleItem {
	return predicate.ScaleItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, ItemColumn),
			sqlgraph.To(ItemInverseTable, ItemFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ItemTable, ItemColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasScale applies the HasEdge predicate on the "scale" edge.
func HasScale() predicate.ScaleItem {
	return predicate.ScaleItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, ScaleColumn),
			sqlgraph.To(ScaleInverseTable, ScaleFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ScaleTable, ScaleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasScaleWith applies the HasEdge predicate on the "scale" edge with a given conditions (other predicates).
func HasScaleWith(preds ...predicate.Scale) predicate.ScaleItem {
	return predicate.ScaleItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, ScaleColumn),
			sqlgraph.To(ScaleInverseTable, ScaleFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ScaleTable, ScaleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ScaleItem) predicate.ScaleItem {
	return predicate.ScaleItem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ScaleItem) predicate.ScaleItem {
	return predicate.ScaleItem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ScaleItem) predicate.ScaleItem {
	return predicate.ScaleItem(func(s *sql.Selector) {
		p(s.Not())
	})
}
