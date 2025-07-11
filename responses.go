package openapi

import (
	"encoding/json"
	"regexp"
)

var ResponseCodePattern = regexp.MustCompile(`^[1-5](?:\d{2}|XX)$`)

// Responses is a container for the expected responses of an operation.
// The container maps a HTTP response code to the expected response.
// The documentation is not necessarily expected to cover all possible HTTP response codes because they may not be known in advance.
// However, documentation is expected to cover a successful operation response and any known errors.
// The default MAY be used as a default response object for all HTTP codes that are not covered individually by the Responses Object.
// The Responses Object MUST contain at least one response code, and if only one response code is provided
// it SHOULD be the response for a successful operation call.
//
// https://spec.openapis.org/oas/v3.1.1#responses-object
//
// Example:
//
//	'200':
//	  description: a pet to be returned
//	  content:
//	    application/json:
//	      schema:
//	        $ref: '#/components/schemas/Pet'
//	default:
//	  description: Unexpected error
//	  content:
//	    application/json:
//	      schema:
//	        $ref: '#/components/schemas/ErrorModel'
type Responses struct {
	// The documentation of responses other than the ones declared for specific HTTP response codes.
	// Use this field to cover undeclared responses.
	Default *RefOrSpec[Extendable[Response]] `json:"default,omitempty"`
	// Any HTTP status code can be used as the property name, but only one property per code,
	// to describe the expected response for that HTTP status code.
	// This field MUST be enclosed in quotation marks (for example, “200”) for compatibility between JSON and YAML.
	// To define a range of response codes, this field MAY contain the uppercase wildcard character X.
	// For example, 2XX represents all response codes between [200-299].
	// Only the following range definitions are allowed: 1XX, 2XX, 3XX, 4XX, and 5XX.
	// If a response is defined using an explicit code, the explicit code definition takes precedence over the range definition for that code.
	Response map[string]*RefOrSpec[Extendable[Response]] `json:"-"`
}

// MarshalJSON implements json.Marshaler interface.
func (o *Responses) MarshalJSON() ([]byte, error) {
	var raw map[string]json.RawMessage
	data, err := json.Marshal(&o.Response)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if o.Default != nil {
		data, err = json.Marshal(&o.Default)
		if err != nil {
			return nil, err
		}
		if raw == nil {
			raw = make(map[string]json.RawMessage, 1)
		}
		raw["default"] = data
	}
	return json.Marshal(&raw)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (o *Responses) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	if v, ok := raw["default"]; ok {
		if err := json.Unmarshal(v, &o.Default); err != nil {
			return err
		}
		delete(raw, "default")
	}
	data, err := json.Marshal(&raw)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &o.Response)
}

func (o *Responses) validateSpec(location string, validator *Validator) []*validationError {
	var errs []*validationError
	if o.Default != nil {
		errs = append(errs, o.Default.validateSpec(joinLoc(location, "default"), validator)...)
	}
	for k, v := range o.Response {
		if !ResponseCodePattern.MatchString(k) {
			errs = append(errs, newValidationError(joinLoc(location, k), "must match pattern '%s', but got '%s'", ResponseCodePattern, k))
		}
		errs = append(errs, v.validateSpec(joinLoc(location, k), validator)...)
	}
	return errs
}

type ResponsesBuilder struct {
	spec *RefOrSpec[Extendable[Responses]]
}

func NewResponsesBuilder() *ResponsesBuilder {
	return &ResponsesBuilder{
		spec: NewRefOrExtSpec[Responses](&Responses{}),
	}
}

func (b *ResponsesBuilder) Build() *RefOrSpec[Extendable[Responses]] {
	return b.spec
}

func (b *ResponsesBuilder) Extensions(v map[string]any) *ResponsesBuilder {
	b.spec.Spec.Extensions = v
	return b
}

func (b *ResponsesBuilder) AddExt(name string, value any) *ResponsesBuilder {
	b.spec.Spec.AddExt(name, value)
	return b
}

func (b *ResponsesBuilder) Default(v *RefOrSpec[Extendable[Response]]) *ResponsesBuilder {
	b.spec.Spec.Spec.Default = v
	return b
}

func (b *ResponsesBuilder) Response(v map[string]*RefOrSpec[Extendable[Response]]) *ResponsesBuilder {
	b.spec.Spec.Spec.Response = v
	return b
}

func (b *ResponsesBuilder) AddResponse(key string, value *RefOrSpec[Extendable[Response]]) *ResponsesBuilder {
	if b.spec.Spec.Spec.Response == nil {
		b.spec.Spec.Spec.Response = make(map[string]*RefOrSpec[Extendable[Response]], 1)
	}
	b.spec.Spec.Spec.Response[key] = value
	return b
}
