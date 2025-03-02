// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/mattetti/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// TableConstraint represents a row from 'information_schema.table_constraints'.
type TableConstraint struct {
	ConstraintCatalog pgtypes.SQLIdentifier `json:"constraint_catalog"` // constraint_catalog
	ConstraintSchema  pgtypes.SQLIdentifier `json:"constraint_schema"`  // constraint_schema
	ConstraintName    pgtypes.SQLIdentifier `json:"constraint_name"`    // constraint_name
	TableCatalog      pgtypes.SQLIdentifier `json:"table_catalog"`      // table_catalog
	TableSchema       pgtypes.SQLIdentifier `json:"table_schema"`       // table_schema
	TableName         pgtypes.SQLIdentifier `json:"table_name"`         // table_name
	ConstraintType    pgtypes.CharacterData `json:"constraint_type"`    // constraint_type
	IsDeferrable      pgtypes.YesOrNo       `json:"is_deferrable"`      // is_deferrable
	InitiallyDeferred pgtypes.YesOrNo       `json:"initially_deferred"` // initially_deferred
}
