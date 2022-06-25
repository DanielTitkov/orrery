// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/result"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/scale"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/take"
	"github.com/google/uuid"
)

// Result is the model entity for the Result schema.
type Result struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// RawScore holds the value of the "raw_score" field.
	RawScore float64 `json:"raw_score,omitempty"`
	// FinalScore holds the value of the "final_score" field.
	FinalScore float64 `json:"final_score,omitempty"`
	// Meta holds the value of the "meta" field.
	Meta map[string]interface{} `json:"meta,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ResultQuery when eager-loading is set.
	Edges         ResultEdges `json:"edges"`
	scale_results *uuid.UUID
	take_results  *uuid.UUID
}

// ResultEdges holds the relations/edges for other nodes in the graph.
type ResultEdges struct {
	// Scale holds the value of the scale edge.
	Scale *Scale `json:"scale,omitempty"`
	// Take holds the value of the take edge.
	Take *Take `json:"take,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ScaleOrErr returns the Scale value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResultEdges) ScaleOrErr() (*Scale, error) {
	if e.loadedTypes[0] {
		if e.Scale == nil {
			// The edge scale was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: scale.Label}
		}
		return e.Scale, nil
	}
	return nil, &NotLoadedError{edge: "scale"}
}

// TakeOrErr returns the Take value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResultEdges) TakeOrErr() (*Take, error) {
	if e.loadedTypes[1] {
		if e.Take == nil {
			// The edge take was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: take.Label}
		}
		return e.Take, nil
	}
	return nil, &NotLoadedError{edge: "take"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Result) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case result.FieldMeta:
			values[i] = new([]byte)
		case result.FieldRawScore, result.FieldFinalScore:
			values[i] = new(sql.NullFloat64)
		case result.FieldCreateTime, result.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case result.FieldID:
			values[i] = new(uuid.UUID)
		case result.ForeignKeys[0]: // scale_results
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case result.ForeignKeys[1]: // take_results
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Result", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Result fields.
func (r *Result) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case result.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				r.ID = *value
			}
		case result.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				r.CreateTime = value.Time
			}
		case result.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				r.UpdateTime = value.Time
			}
		case result.FieldRawScore:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field raw_score", values[i])
			} else if value.Valid {
				r.RawScore = value.Float64
			}
		case result.FieldFinalScore:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field final_score", values[i])
			} else if value.Valid {
				r.FinalScore = value.Float64
			}
		case result.FieldMeta:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field meta", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.Meta); err != nil {
					return fmt.Errorf("unmarshal field meta: %w", err)
				}
			}
		case result.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field scale_results", values[i])
			} else if value.Valid {
				r.scale_results = new(uuid.UUID)
				*r.scale_results = *value.S.(*uuid.UUID)
			}
		case result.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field take_results", values[i])
			} else if value.Valid {
				r.take_results = new(uuid.UUID)
				*r.take_results = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryScale queries the "scale" edge of the Result entity.
func (r *Result) QueryScale() *ScaleQuery {
	return (&ResultClient{config: r.config}).QueryScale(r)
}

// QueryTake queries the "take" edge of the Result entity.
func (r *Result) QueryTake() *TakeQuery {
	return (&ResultClient{config: r.config}).QueryTake(r)
}

// Update returns a builder for updating this Result.
// Note that you need to call Result.Unwrap() before calling this method if this Result
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Result) Update() *ResultUpdateOne {
	return (&ResultClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Result entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Result) Unwrap() *Result {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Result is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Result) String() string {
	var builder strings.Builder
	builder.WriteString("Result(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("create_time=")
	builder.WriteString(r.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(r.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("raw_score=")
	builder.WriteString(fmt.Sprintf("%v", r.RawScore))
	builder.WriteString(", ")
	builder.WriteString("final_score=")
	builder.WriteString(fmt.Sprintf("%v", r.FinalScore))
	builder.WriteString(", ")
	builder.WriteString("meta=")
	builder.WriteString(fmt.Sprintf("%v", r.Meta))
	builder.WriteByte(')')
	return builder.String()
}

// Results is a parsable slice of Result.
type Results []*Result

func (r Results) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}