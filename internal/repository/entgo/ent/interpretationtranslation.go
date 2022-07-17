// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/interpretation"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/interpretationtranslation"
)

// InterpretationTranslation is the model entity for the InterpretationTranslation schema.
type InterpretationTranslation struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Locale holds the value of the "locale" field.
	Locale interpretationtranslation.Locale `json:"locale,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InterpretationTranslationQuery when eager-loading is set.
	Edges                       InterpretationTranslationEdges `json:"edges"`
	interpretation_translations *uuid.UUID
}

// InterpretationTranslationEdges holds the relations/edges for other nodes in the graph.
type InterpretationTranslationEdges struct {
	// Interpretation holds the value of the interpretation edge.
	Interpretation *Interpretation `json:"interpretation,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// InterpretationOrErr returns the Interpretation value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InterpretationTranslationEdges) InterpretationOrErr() (*Interpretation, error) {
	if e.loadedTypes[0] {
		if e.Interpretation == nil {
			// The edge interpretation was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: interpretation.Label}
		}
		return e.Interpretation, nil
	}
	return nil, &NotLoadedError{edge: "interpretation"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*InterpretationTranslation) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case interpretationtranslation.FieldLocale, interpretationtranslation.FieldContent:
			values[i] = new(sql.NullString)
		case interpretationtranslation.FieldID:
			values[i] = new(uuid.UUID)
		case interpretationtranslation.ForeignKeys[0]: // interpretation_translations
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type InterpretationTranslation", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the InterpretationTranslation fields.
func (it *InterpretationTranslation) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case interpretationtranslation.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				it.ID = *value
			}
		case interpretationtranslation.FieldLocale:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field locale", values[i])
			} else if value.Valid {
				it.Locale = interpretationtranslation.Locale(value.String)
			}
		case interpretationtranslation.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				it.Content = value.String
			}
		case interpretationtranslation.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field interpretation_translations", values[i])
			} else if value.Valid {
				it.interpretation_translations = new(uuid.UUID)
				*it.interpretation_translations = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryInterpretation queries the "interpretation" edge of the InterpretationTranslation entity.
func (it *InterpretationTranslation) QueryInterpretation() *InterpretationQuery {
	return (&InterpretationTranslationClient{config: it.config}).QueryInterpretation(it)
}

// Update returns a builder for updating this InterpretationTranslation.
// Note that you need to call InterpretationTranslation.Unwrap() before calling this method if this InterpretationTranslation
// was returned from a transaction, and the transaction was committed or rolled back.
func (it *InterpretationTranslation) Update() *InterpretationTranslationUpdateOne {
	return (&InterpretationTranslationClient{config: it.config}).UpdateOne(it)
}

// Unwrap unwraps the InterpretationTranslation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (it *InterpretationTranslation) Unwrap() *InterpretationTranslation {
	_tx, ok := it.config.driver.(*txDriver)
	if !ok {
		panic("ent: InterpretationTranslation is not a transactional entity")
	}
	it.config.driver = _tx.drv
	return it
}

// String implements the fmt.Stringer.
func (it *InterpretationTranslation) String() string {
	var builder strings.Builder
	builder.WriteString("InterpretationTranslation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", it.ID))
	builder.WriteString("locale=")
	builder.WriteString(fmt.Sprintf("%v", it.Locale))
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(it.Content)
	builder.WriteByte(')')
	return builder.String()
}

// InterpretationTranslations is a parsable slice of InterpretationTranslation.
type InterpretationTranslations []*InterpretationTranslation

func (it InterpretationTranslations) config(cfg config) {
	for _i := range it {
		it[_i].config = cfg
	}
}
