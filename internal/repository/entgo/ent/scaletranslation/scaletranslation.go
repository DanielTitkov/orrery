// Code generated by ent, DO NOT EDIT.

package scaletranslation

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the scaletranslation type in the database.
	Label = "scale_translation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLocale holds the string denoting the locale field in the database.
	FieldLocale = "locale"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldAbbreviation holds the string denoting the abbreviation field in the database.
	FieldAbbreviation = "abbreviation"
	// EdgeScale holds the string denoting the scale edge name in mutations.
	EdgeScale = "scale"
	// Table holds the table name of the scaletranslation in the database.
	Table = "scale_translations"
	// ScaleTable is the table that holds the scale relation/edge.
	ScaleTable = "scale_translations"
	// ScaleInverseTable is the table name for the Scale entity.
	// It exists in this package in order to avoid circular dependency with the "scale" package.
	ScaleInverseTable = "scales"
	// ScaleColumn is the table column denoting the scale relation/edge.
	ScaleColumn = "scale_translations"
)

// Columns holds all SQL columns for scaletranslation fields.
var Columns = []string{
	FieldID,
	FieldLocale,
	FieldTitle,
	FieldDescription,
	FieldAbbreviation,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "scale_translations"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"scale_translations",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DefaultAbbreviation holds the default value on creation for the "abbreviation" field.
	DefaultAbbreviation string
	// AbbreviationValidator is a validator for the "abbreviation" field. It is called by the builders before save.
	AbbreviationValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Locale defines the type for the "locale" enum field.
type Locale string

// Locale values.
const (
	LocaleEn Locale = "en"
	LocaleRu Locale = "ru"
)

func (l Locale) String() string {
	return string(l)
}

// LocaleValidator is a validator for the "locale" field enum values. It is called by the builders before save.
func LocaleValidator(l Locale) error {
	switch l {
	case LocaleEn, LocaleRu:
		return nil
	default:
		return fmt.Errorf("scaletranslation: invalid enum value for locale field: %q", l)
	}
}
