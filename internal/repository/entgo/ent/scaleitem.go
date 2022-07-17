// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/item"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/scale"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/scaleitem"
)

// ScaleItem is the model entity for the ScaleItem schema.
type ScaleItem struct {
	config `json:"-"`
	// Reverse holds the value of the "reverse" field.
	Reverse bool `json:"reverse,omitempty"`
	// ScaleID holds the value of the "scale_id" field.
	ScaleID uuid.UUID `json:"scale_id,omitempty"`
	// ItemID holds the value of the "item_id" field.
	ItemID uuid.UUID `json:"item_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ScaleItemQuery when eager-loading is set.
	Edges ScaleItemEdges `json:"edges"`
}

// ScaleItemEdges holds the relations/edges for other nodes in the graph.
type ScaleItemEdges struct {
	// Item holds the value of the item edge.
	Item *Item `json:"item,omitempty"`
	// Scale holds the value of the scale edge.
	Scale *Scale `json:"scale,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ItemOrErr returns the Item value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ScaleItemEdges) ItemOrErr() (*Item, error) {
	if e.loadedTypes[0] {
		if e.Item == nil {
			// The edge item was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: item.Label}
		}
		return e.Item, nil
	}
	return nil, &NotLoadedError{edge: "item"}
}

// ScaleOrErr returns the Scale value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ScaleItemEdges) ScaleOrErr() (*Scale, error) {
	if e.loadedTypes[1] {
		if e.Scale == nil {
			// The edge scale was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: scale.Label}
		}
		return e.Scale, nil
	}
	return nil, &NotLoadedError{edge: "scale"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ScaleItem) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case scaleitem.FieldReverse:
			values[i] = new(sql.NullBool)
		case scaleitem.FieldScaleID, scaleitem.FieldItemID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ScaleItem", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ScaleItem fields.
func (si *ScaleItem) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case scaleitem.FieldReverse:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field reverse", values[i])
			} else if value.Valid {
				si.Reverse = value.Bool
			}
		case scaleitem.FieldScaleID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field scale_id", values[i])
			} else if value != nil {
				si.ScaleID = *value
			}
		case scaleitem.FieldItemID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field item_id", values[i])
			} else if value != nil {
				si.ItemID = *value
			}
		}
	}
	return nil
}

// QueryItem queries the "item" edge of the ScaleItem entity.
func (si *ScaleItem) QueryItem() *ItemQuery {
	return (&ScaleItemClient{config: si.config}).QueryItem(si)
}

// QueryScale queries the "scale" edge of the ScaleItem entity.
func (si *ScaleItem) QueryScale() *ScaleQuery {
	return (&ScaleItemClient{config: si.config}).QueryScale(si)
}

// Update returns a builder for updating this ScaleItem.
// Note that you need to call ScaleItem.Unwrap() before calling this method if this ScaleItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (si *ScaleItem) Update() *ScaleItemUpdateOne {
	return (&ScaleItemClient{config: si.config}).UpdateOne(si)
}

// Unwrap unwraps the ScaleItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (si *ScaleItem) Unwrap() *ScaleItem {
	_tx, ok := si.config.driver.(*txDriver)
	if !ok {
		panic("ent: ScaleItem is not a transactional entity")
	}
	si.config.driver = _tx.drv
	return si
}

// String implements the fmt.Stringer.
func (si *ScaleItem) String() string {
	var builder strings.Builder
	builder.WriteString("ScaleItem(")
	builder.WriteString("reverse=")
	builder.WriteString(fmt.Sprintf("%v", si.Reverse))
	builder.WriteString(", ")
	builder.WriteString("scale_id=")
	builder.WriteString(fmt.Sprintf("%v", si.ScaleID))
	builder.WriteString(", ")
	builder.WriteString("item_id=")
	builder.WriteString(fmt.Sprintf("%v", si.ItemID))
	builder.WriteByte(')')
	return builder.String()
}

// ScaleItems is a parsable slice of ScaleItem.
type ScaleItems []*ScaleItem

func (si ScaleItems) config(cfg config) {
	for _i := range si {
		si[_i].config = cfg
	}
}
