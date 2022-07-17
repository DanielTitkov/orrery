// Code generated by ent, DO NOT EDIT.

package usersession

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// Sid applies equality check predicate on the "sid" field. It's identical to SidEQ.
func Sid(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSid), v))
	})
}

// IP applies equality check predicate on the "ip" field. It's identical to IPEQ.
func IP(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIP), v))
	})
}

// UserAgent applies equality check predicate on the "user_agent" field. It's identical to UserAgentEQ.
func UserAgent(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserAgent), v))
	})
}

// LastActivity applies equality check predicate on the "last_activity" field. It's identical to LastActivityEQ.
func LastActivity(v time.Time) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastActivity), v))
	})
}

// Active applies equality check predicate on the "active" field. It's identical to ActiveEQ.
func Active(v bool) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActive), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.UserSession {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSession(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.UserSession {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSession(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.UserSession {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSession(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.UserSession {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSession(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// SidEQ applies the EQ predicate on the "sid" field.
func SidEQ(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSid), v))
	})
}

// SidNEQ applies the NEQ predicate on the "sid" field.
func SidNEQ(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSid), v))
	})
}

// SidIn applies the In predicate on the "sid" field.
func SidIn(vs ...string) predicate.UserSession {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSession(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSid), v...))
	})
}

// SidNotIn applies the NotIn predicate on the "sid" field.
func SidNotIn(vs ...string) predicate.UserSession {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSession(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSid), v...))
	})
}

// SidGT applies the GT predicate on the "sid" field.
func SidGT(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSid), v))
	})
}

// SidGTE applies the GTE predicate on the "sid" field.
func SidGTE(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSid), v))
	})
}

// SidLT applies the LT predicate on the "sid" field.
func SidLT(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSid), v))
	})
}

// SidLTE applies the LTE predicate on the "sid" field.
func SidLTE(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSid), v))
	})
}

// SidContains applies the Contains predicate on the "sid" field.
func SidContains(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSid), v))
	})
}

// SidHasPrefix applies the HasPrefix predicate on the "sid" field.
func SidHasPrefix(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSid), v))
	})
}

// SidHasSuffix applies the HasSuffix predicate on the "sid" field.
func SidHasSuffix(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSid), v))
	})
}

// SidEqualFold applies the EqualFold predicate on the "sid" field.
func SidEqualFold(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSid), v))
	})
}

// SidContainsFold applies the ContainsFold predicate on the "sid" field.
func SidContainsFold(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSid), v))
	})
}

// IPEQ applies the EQ predicate on the "ip" field.
func IPEQ(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIP), v))
	})
}

// IPNEQ applies the NEQ predicate on the "ip" field.
func IPNEQ(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIP), v))
	})
}

// IPIn applies the In predicate on the "ip" field.
func IPIn(vs ...string) predicate.UserSession {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSession(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIP), v...))
	})
}

// IPNotIn applies the NotIn predicate on the "ip" field.
func IPNotIn(vs ...string) predicate.UserSession {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSession(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIP), v...))
	})
}

// IPGT applies the GT predicate on the "ip" field.
func IPGT(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIP), v))
	})
}

// IPGTE applies the GTE predicate on the "ip" field.
func IPGTE(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIP), v))
	})
}

// IPLT applies the LT predicate on the "ip" field.
func IPLT(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIP), v))
	})
}

// IPLTE applies the LTE predicate on the "ip" field.
func IPLTE(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIP), v))
	})
}

// IPContains applies the Contains predicate on the "ip" field.
func IPContains(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIP), v))
	})
}

// IPHasPrefix applies the HasPrefix predicate on the "ip" field.
func IPHasPrefix(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIP), v))
	})
}

// IPHasSuffix applies the HasSuffix predicate on the "ip" field.
func IPHasSuffix(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIP), v))
	})
}

// IPEqualFold applies the EqualFold predicate on the "ip" field.
func IPEqualFold(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIP), v))
	})
}

// IPContainsFold applies the ContainsFold predicate on the "ip" field.
func IPContainsFold(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIP), v))
	})
}

// UserAgentEQ applies the EQ predicate on the "user_agent" field.
func UserAgentEQ(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserAgent), v))
	})
}

// UserAgentNEQ applies the NEQ predicate on the "user_agent" field.
func UserAgentNEQ(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserAgent), v))
	})
}

// UserAgentIn applies the In predicate on the "user_agent" field.
func UserAgentIn(vs ...string) predicate.UserSession {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSession(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserAgent), v...))
	})
}

// UserAgentNotIn applies the NotIn predicate on the "user_agent" field.
func UserAgentNotIn(vs ...string) predicate.UserSession {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSession(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserAgent), v...))
	})
}

// UserAgentGT applies the GT predicate on the "user_agent" field.
func UserAgentGT(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserAgent), v))
	})
}

// UserAgentGTE applies the GTE predicate on the "user_agent" field.
func UserAgentGTE(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserAgent), v))
	})
}

// UserAgentLT applies the LT predicate on the "user_agent" field.
func UserAgentLT(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserAgent), v))
	})
}

// UserAgentLTE applies the LTE predicate on the "user_agent" field.
func UserAgentLTE(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserAgent), v))
	})
}

// UserAgentContains applies the Contains predicate on the "user_agent" field.
func UserAgentContains(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserAgent), v))
	})
}

// UserAgentHasPrefix applies the HasPrefix predicate on the "user_agent" field.
func UserAgentHasPrefix(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserAgent), v))
	})
}

// UserAgentHasSuffix applies the HasSuffix predicate on the "user_agent" field.
func UserAgentHasSuffix(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserAgent), v))
	})
}

// UserAgentEqualFold applies the EqualFold predicate on the "user_agent" field.
func UserAgentEqualFold(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserAgent), v))
	})
}

// UserAgentContainsFold applies the ContainsFold predicate on the "user_agent" field.
func UserAgentContainsFold(v string) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserAgent), v))
	})
}

// LastActivityEQ applies the EQ predicate on the "last_activity" field.
func LastActivityEQ(v time.Time) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastActivity), v))
	})
}

// LastActivityNEQ applies the NEQ predicate on the "last_activity" field.
func LastActivityNEQ(v time.Time) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastActivity), v))
	})
}

// LastActivityIn applies the In predicate on the "last_activity" field.
func LastActivityIn(vs ...time.Time) predicate.UserSession {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSession(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastActivity), v...))
	})
}

// LastActivityNotIn applies the NotIn predicate on the "last_activity" field.
func LastActivityNotIn(vs ...time.Time) predicate.UserSession {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSession(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastActivity), v...))
	})
}

// LastActivityGT applies the GT predicate on the "last_activity" field.
func LastActivityGT(v time.Time) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastActivity), v))
	})
}

// LastActivityGTE applies the GTE predicate on the "last_activity" field.
func LastActivityGTE(v time.Time) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastActivity), v))
	})
}

// LastActivityLT applies the LT predicate on the "last_activity" field.
func LastActivityLT(v time.Time) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastActivity), v))
	})
}

// LastActivityLTE applies the LTE predicate on the "last_activity" field.
func LastActivityLTE(v time.Time) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastActivity), v))
	})
}

// ActiveEQ applies the EQ predicate on the "active" field.
func ActiveEQ(v bool) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActive), v))
	})
}

// ActiveNEQ applies the NEQ predicate on the "active" field.
func ActiveNEQ(v bool) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldActive), v))
	})
}

// MetaIsNil applies the IsNil predicate on the "meta" field.
func MetaIsNil() predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMeta)))
	})
}

// MetaNotNil applies the NotNil predicate on the "meta" field.
func MetaNotNil() predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMeta)))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserSession) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserSession) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
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
func Not(p predicate.UserSession) predicate.UserSession {
	return predicate.UserSession(func(s *sql.Selector) {
		p(s.Not())
	})
}
