// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/mattetti/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// SQLPart represents a row from 'information_schema.sql_parts'.
type SQLPart struct {
	Tableoid     pgtypes.Oid           `json:"tableoid"`       // tableoid
	Cmax         pgtypes.Cid           `json:"cmax"`           // cmax
	Xmax         pgtypes.Xid           `json:"xmax"`           // xmax
	Cmin         pgtypes.Cid           `json:"cmin"`           // cmin
	Xmin         pgtypes.Xid           `json:"xmin"`           // xmin
	Ctid         pgtypes.Tid           `json:"ctid"`           // ctid
	FeatureID    pgtypes.CharacterData `json:"feature_id"`     // feature_id
	FeatureName  pgtypes.CharacterData `json:"feature_name"`   // feature_name
	IsSupported  pgtypes.YesOrNo       `json:"is_supported"`   // is_supported
	IsVerifiedBy pgtypes.CharacterData `json:"is_verified_by"` // is_verified_by
	Comments     pgtypes.CharacterData `json:"comments"`       // comments
}
