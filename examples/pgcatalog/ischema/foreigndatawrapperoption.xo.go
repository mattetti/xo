// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/mattetti/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// ForeignDataWrapperOption represents a row from 'information_schema.foreign_data_wrapper_options'.
type ForeignDataWrapperOption struct {
	ForeignDataWrapperCatalog pgtypes.SQLIdentifier `json:"foreign_data_wrapper_catalog"` // foreign_data_wrapper_catalog
	ForeignDataWrapperName    pgtypes.SQLIdentifier `json:"foreign_data_wrapper_name"`    // foreign_data_wrapper_name
	OptionName                pgtypes.SQLIdentifier `json:"option_name"`                  // option_name
	OptionValue               pgtypes.CharacterData `json:"option_value"`                 // option_value
}
