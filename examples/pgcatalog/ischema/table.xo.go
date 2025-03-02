// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/mattetti/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// Table represents a row from 'information_schema.tables'.
type Table struct {
	TableCatalog              pgtypes.SQLIdentifier `json:"table_catalog"`                // table_catalog
	TableSchema               pgtypes.SQLIdentifier `json:"table_schema"`                 // table_schema
	TableName                 pgtypes.SQLIdentifier `json:"table_name"`                   // table_name
	TableType                 pgtypes.CharacterData `json:"table_type"`                   // table_type
	SelfReferencingColumnName pgtypes.SQLIdentifier `json:"self_referencing_column_name"` // self_referencing_column_name
	ReferenceGeneration       pgtypes.CharacterData `json:"reference_generation"`         // reference_generation
	UserDefinedTypeCatalog    pgtypes.SQLIdentifier `json:"user_defined_type_catalog"`    // user_defined_type_catalog
	UserDefinedTypeSchema     pgtypes.SQLIdentifier `json:"user_defined_type_schema"`     // user_defined_type_schema
	UserDefinedTypeName       pgtypes.SQLIdentifier `json:"user_defined_type_name"`       // user_defined_type_name
	IsInsertableInto          pgtypes.YesOrNo       `json:"is_insertable_into"`           // is_insertable_into
	IsTyped                   pgtypes.YesOrNo       `json:"is_typed"`                     // is_typed
	CommitAction              pgtypes.CharacterData `json:"commit_action"`                // commit_action
}
