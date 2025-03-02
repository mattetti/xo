// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/mattetti/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// DomainConstraint represents a row from 'information_schema.domain_constraints'.
type DomainConstraint struct {
	ConstraintCatalog pgtypes.SQLIdentifier `json:"constraint_catalog"` // constraint_catalog
	ConstraintSchema  pgtypes.SQLIdentifier `json:"constraint_schema"`  // constraint_schema
	ConstraintName    pgtypes.SQLIdentifier `json:"constraint_name"`    // constraint_name
	DomainCatalog     pgtypes.SQLIdentifier `json:"domain_catalog"`     // domain_catalog
	DomainSchema      pgtypes.SQLIdentifier `json:"domain_schema"`      // domain_schema
	DomainName        pgtypes.SQLIdentifier `json:"domain_name"`        // domain_name
	IsDeferrable      pgtypes.YesOrNo       `json:"is_deferrable"`      // is_deferrable
	InitiallyDeferred pgtypes.YesOrNo       `json:"initially_deferred"` // initially_deferred
}
