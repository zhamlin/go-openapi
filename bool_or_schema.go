package openapi

import (
	"encoding/json"
)

// BoolOrSchema handles Boolean or Schema type.
//
// It MUST be used as a pointer,
// otherwise the `false` can be omitted by json encoders in case of `omitempty` tag is set.
type BoolOrSchema struct {
	Schema  *RefOrSpec[Schema]
	Allowed bool
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (o *BoolOrSchema) UnmarshalJSON(data []byte) error {
	if json.Unmarshal(data, &o.Allowed) == nil {
		o.Schema = nil
		return nil
	}
	if err := json.Unmarshal(data, &o.Schema); err != nil {
		return err
	}
	o.Allowed = true
	return nil
}

// MarshalJSON implements json.Marshaler interface.
func (o *BoolOrSchema) MarshalJSON() ([]byte, error) {
	var v any
	if o.Schema != nil {
		v = o.Schema
	} else {
		v = o.Allowed
	}
	return json.Marshal(&v)
}

func (o *BoolOrSchema) validateSpec(path string, validator *Validator) []*validationError {
	var errs []*validationError
	if o.Schema != nil {
		errs = append(errs, o.Schema.validateSpec(path, validator)...)
	}
	return errs
}

func NewBoolOrSchema(v any) *BoolOrSchema {
	switch v := v.(type) {
	case bool:
		return &BoolOrSchema{Allowed: v}
	case *RefOrSpec[Schema]:
		return &BoolOrSchema{Schema: v}
	case *SchemaBuilder:
		return &BoolOrSchema{Schema: v.Build()}
	default:
		return nil
	}
}
