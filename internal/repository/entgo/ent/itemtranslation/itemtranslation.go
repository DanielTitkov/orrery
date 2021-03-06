// Code generated by ent, DO NOT EDIT.

package itemtranslation

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the itemtranslation type in the database.
	Label = "item_translation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLocale holds the string denoting the locale field in the database.
	FieldLocale = "locale"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// EdgeItem holds the string denoting the item edge name in mutations.
	EdgeItem = "item"
	// Table holds the table name of the itemtranslation in the database.
	Table = "item_translations"
	// ItemTable is the table that holds the item relation/edge.
	ItemTable = "item_translations"
	// ItemInverseTable is the table name for the Item entity.
	// It exists in this package in order to avoid circular dependency with the "item" package.
	ItemInverseTable = "items"
	// ItemColumn is the table column denoting the item relation/edge.
	ItemColumn = "item_translations"
)

// Columns holds all SQL columns for itemtranslation fields.
var Columns = []string{
	FieldID,
	FieldLocale,
	FieldContent,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "item_translations"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"item_translations",
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
	// ContentValidator is a validator for the "content" field. It is called by the builders before save.
	ContentValidator func(string) error
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
		return fmt.Errorf("itemtranslation: invalid enum value for locale field: %q", l)
	}
}
