// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/tag"
)

// Tag is the model entity for the Tag schema.
type Tag struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Type holds the value of the "type" field.
	Type tag.Type `json:"type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TagQuery when eager-loading is set.
	Edges TagEdges `json:"edges"`
}

// TagEdges holds the relations/edges for other nodes in the graph.
type TagEdges struct {
	// Translations holds the value of the translations edge.
	Translations []*TagTranslation `json:"translations,omitempty"`
	// Test holds the value of the test edge.
	Test []*Test `json:"test,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TranslationsOrErr returns the Translations value or an error if the edge
// was not loaded in eager-loading.
func (e TagEdges) TranslationsOrErr() ([]*TagTranslation, error) {
	if e.loadedTypes[0] {
		return e.Translations, nil
	}
	return nil, &NotLoadedError{edge: "translations"}
}

// TestOrErr returns the Test value or an error if the edge
// was not loaded in eager-loading.
func (e TagEdges) TestOrErr() ([]*Test, error) {
	if e.loadedTypes[1] {
		return e.Test, nil
	}
	return nil, &NotLoadedError{edge: "test"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tag) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case tag.FieldCode, tag.FieldType:
			values[i] = new(sql.NullString)
		case tag.FieldCreateTime, tag.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case tag.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Tag", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tag fields.
func (t *Tag) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tag.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case tag.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				t.CreateTime = value.Time
			}
		case tag.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				t.UpdateTime = value.Time
			}
		case tag.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				t.Code = value.String
			}
		case tag.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				t.Type = tag.Type(value.String)
			}
		}
	}
	return nil
}

// QueryTranslations queries the "translations" edge of the Tag entity.
func (t *Tag) QueryTranslations() *TagTranslationQuery {
	return (&TagClient{config: t.config}).QueryTranslations(t)
}

// QueryTest queries the "test" edge of the Tag entity.
func (t *Tag) QueryTest() *TestQuery {
	return (&TagClient{config: t.config}).QueryTest(t)
}

// Update returns a builder for updating this Tag.
// Note that you need to call Tag.Unwrap() before calling this method if this Tag
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tag) Update() *TagUpdateOne {
	return (&TagClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Tag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tag) Unwrap() *Tag {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tag is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tag) String() string {
	var builder strings.Builder
	builder.WriteString("Tag(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("create_time=")
	builder.WriteString(t.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(t.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(t.Code)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", t.Type))
	builder.WriteByte(')')
	return builder.String()
}

// Tags is a parsable slice of Tag.
type Tags []*Tag

func (t Tags) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
