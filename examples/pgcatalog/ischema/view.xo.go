// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/mattetti/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// View represents a row from 'information_schema.views'.
type View struct {
	TableCatalog            pgtypes.SQLIdentifier `json:"table_catalog"`              // table_catalog
	TableSchema             pgtypes.SQLIdentifier `json:"table_schema"`               // table_schema
	TableName               pgtypes.SQLIdentifier `json:"table_name"`                 // table_name
	ViewDefinition          pgtypes.CharacterData `json:"view_definition"`            // view_definition
	CheckOption             pgtypes.CharacterData `json:"check_option"`               // check_option
	IsUpdatable             pgtypes.YesOrNo       `json:"is_updatable"`               // is_updatable
	IsInsertableInto        pgtypes.YesOrNo       `json:"is_insertable_into"`         // is_insertable_into
	IsTriggerUpdatable      pgtypes.YesOrNo       `json:"is_trigger_updatable"`       // is_trigger_updatable
	IsTriggerDeletable      pgtypes.YesOrNo       `json:"is_trigger_deletable"`       // is_trigger_deletable
	IsTriggerInsertableInto pgtypes.YesOrNo       `json:"is_trigger_insertable_into"` // is_trigger_insertable_into
}
