// Code generated by ent, DO NOT EDIT.

package tag

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the tag type in the database.
	Label = "tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeTranslations holds the string denoting the translations edge name in mutations.
	EdgeTranslations = "translations"
	// EdgeTest holds the string denoting the test edge name in mutations.
	EdgeTest = "test"
	// Table holds the table name of the tag in the database.
	Table = "tags"
	// TranslationsTable is the table that holds the translations relation/edge.
	TranslationsTable = "tag_translations"
	// TranslationsInverseTable is the table name for the TagTranslation entity.
	// It exists in this package in order to avoid circular dependency with the "tagtranslation" package.
	TranslationsInverseTable = "tag_translations"
	// TranslationsColumn is the table column denoting the translations relation/edge.
	TranslationsColumn = "tag_translations"
	// TestTable is the table that holds the test relation/edge. The primary key declared below.
	TestTable = "test_tags"
	// TestInverseTable is the table name for the Test entity.
	// It exists in this package in order to avoid circular dependency with the "test" package.
	TestInverseTable = "tests"
)

// Columns holds all SQL columns for tag fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldCode,
	FieldType,
}

var (
	// TestPrimaryKey and TestColumn2 are the table columns denoting the
	// primary key for the test relation (M2M).
	TestPrimaryKey = []string{"test_id", "tag_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Type defines the type for the "type" enum field.
type Type string

// TypeFeature is the default value of the Type enum.
const DefaultType = TypeFeature

// Type values.
const (
	TypeTheme   Type = "theme"
	TypeLen     Type = "len"
	TypeFeature Type = "feature"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeTheme, TypeLen, TypeFeature:
		return nil
	default:
		return fmt.Errorf("tag: invalid enum value for type field: %q", _type)
	}
}
