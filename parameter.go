package openapi

import (
	"regexp"
	"strings"
)

const (
	// InPath used together with Path Templating, where the parameter value is actually part of the operation’s URL.
	// This does not include the host or base path of the API.
	// For example, in /items/{itemId}, the path parameter is itemId.
	//
	// https://spec.openapis.org/oas/v3.1.1#parameter-locations
	InPath = "path"
	// InQuery used for parameters that are appended to the URL.
	// For example, in /items?id=###, the query parameter is id.
	//
	// https://spec.openapis.org/oas/v3.1.1#parameter-locations
	InQuery = "query"
	// InHeader used as custom headers that are expected as part of the request.
	// Note that [RFC7230] states header names are case insensitive.
	//
	// https://spec.openapis.org/oas/v3.1.1#parameter-locations
	InHeader = "header"
	// InCookie used to pass a specific cookie value to the API.
	//
	// https://spec.openapis.org/oas/v3.1.1#parameter-locations
	InCookie = "cookie"

	// StyleMatrix is the parameters defined by [RFC6570](https://www.rfc-editor.org/rfc/rfc6570#section-3.2.7)
	//
	// Supported types:
	//   - primitive
	//   - array
	//   - object
	//
	// Can be used with `in: path` location.
	StyleMatrix = "matrix"

	// StyleLabel is the parameters defined by [RFC6570](https://www.rfc-editor.org/rfc/rfc6570#section-3.2.5)
	//
	// Supported types:
	//   - primitive
	//   - array
	//   - object
	//
	// Can be used with `in: path` location.
	StyleLabel = "label"

	// StyleForm is the parameters defined by [RFC6570](https://www.rfc-editor.org/rfc/rfc6570#section-3.2.8)
	//
	// Supported types:
	//   - primitive
	//   - array
	//   - object
	//
	// Can be used with `in: query` and `in: cookie` locations.
	StyleForm = "form"

	// StyleSimple is the parameters defined by [RFC6570](https://www.rfc-editor.org/rfc/rfc6570#section-3.2.2)
	// This option replaces collectionFormat with a csv value from OpenAPI 2.0.
	//
	// Supported types:
	//   - array
	//
	// Can be used with `in: path` and `in: header` locations.
	StyleSimple = "simple"

	// StyleSpaceDelimited is space separated array or object values.
	// This option replaces collectionFormat with a ssv value from OpenAPI 2.0.
	//
	// Supported types:
	//   - array
	//   - object
	//
	// Can be used with `in: query` location.
	StyleSpaceDelimited = "spaceDelimited"

	// StylePipeDelimited is pipe separated array or object values.
	// This option replaces collectionFormat with a pipes value from OpenAPI 2.0.
	//
	// Supported types:
	//   - array
	//   - object
	//
	// Can be used with `in: query` location.
	StylePipeDelimited = "pipeDelimited"

	// StyleDeepObject provides a simple way of rendering nested objects using form parameters.
	//
	// Supported types:
	//   - object
	//
	// Can be used with `in: query` location.
	StyleDeepObject = "deepObject"

	ReservedCharacters = ":/?#[]@!$&'()*+,;="
)

var PathNamePattern = regexp.MustCompile(`[^/#?]+$`)

// Parameter describes a single operation parameter.
// A unique parameter is defined by a combination of a name and location.
//
// https://spec.openapis.org/oas/v3.1.1#parameter-object
//
// Example:
//
//	name: pet
//	description: Pets operations
type Parameter struct {
	// Example of the parameter’s potential value.
	// The example SHOULD match the specified schema and encoding properties if present.
	// The example field is mutually exclusive of the examples field.
	// Furthermore, if referencing a schema that contains an example, the example value SHALL override the example provided by the schema.
	// To represent examples of media types that cannot naturally be represented in JSON or YAML,
	// a string value can contain the example with escaping where necessary.
	Example any `json:"example,omitempty"`
	// A map containing the representations for the parameter.
	// The key is the media type and the value describes it.
	// The map MUST only contain one entry.
	Content map[string]*Extendable[MediaType] `json:"content,omitempty"`
	// Examples of the parameter’s potential value.
	// Each example SHOULD contain a value in the correct format as specified in the parameter encoding.
	// The examples field is mutually exclusive of the example field.
	// Furthermore, if referencing a schema that contains an example, the examples value SHALL override the example provided by the schema.
	Examples map[string]*RefOrSpec[Extendable[Example]] `json:"examples,omitempty"`
	// The schema defining the type used for the parameter.
	Schema *RefOrSpec[Schema] `json:"schema,omitempty"`
	// REQUIRED.
	// The location of the parameter.
	// Possible values are "query", "header", "path" or "cookie".
	In string `json:"in"`
	// A brief description of the parameter.
	// This could contain examples of use.
	// CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
	// Describes how the parameter value will be serialized depending on the type of the parameter value.
	// Default values (based on value of in):
	//   for query - form;
	//   for path - simple;
	//   for header - simple;
	//   for cookie - form.
	Style string `json:"style,omitempty"`
	// REQUIRED.
	// The name of the parameter.
	// Parameter names are case sensitive.
	// If in is "path", the name field MUST correspond to a template expression occurring within the path field in the Paths Object.
	// See Path Templating for further information.
	// If in is "header" and the name field is "Accept", "Content-Type" or "Authorization", the parameter definition SHALL be ignored.
	// For all other cases, the name corresponds to the parameter name used by the in property.
	Name string `json:"name"`
	// When this is true, parameter values of type array or object generate separate parameters
	// for each value of the array or key-value pair of the map.
	// For other types of parameters this property has no effect.
	// When style is form, the default value is true.
	// For all other styles, the default value is false.
	Explode bool `json:"explode"`
	// Determines whether the parameter value SHOULD allow reserved characters, as defined by [RFC3986]
	//   :/?#[]@!$&'()*+,;=
	// to be included without percent-encoding.
	// This property only applies to parameters with an in value of query.
	// The default value is false.
	AllowReserved bool `json:"allowReserved,omitempty"`
	// Sets the ability to pass empty-valued parameters.
	// This is valid only for query parameters and allows sending a parameter with an empty value.
	// Default value is false.
	// If style is used, and if behavior is n/a (cannot be serialized), the value of allowEmptyValue SHALL be ignored.
	// Use of this property is NOT RECOMMENDED, as it is likely to be removed in a later revision.
	AllowEmptyValue bool `json:"allowEmptyValue,omitempty"`
	// Specifies that a parameter is deprecated and SHOULD be transitioned out of usage.
	// Default value is false.
	Deprecated bool `json:"deprecated,omitempty"`
	// Determines whether this parameter is mandatory.
	// If the parameter location is "path", this property is REQUIRED and its value MUST be true.
	// Otherwise, the property MAY be included and its default value is false.
	Required bool `json:"required,omitempty"`
}

func (o *Parameter) validateSpec(location string, validator *Validator) []*validationError {
	var errs []*validationError
	if o.Schema != nil && o.Content != nil {
		errs = append(errs, newValidationError(joinLoc(location, "schema&content"), ErrMutuallyExclusive))
	}
	if o.Example != nil && len(o.Examples) > 0 {
		errs = append(errs, newValidationError(joinLoc(location, "example&examples"), ErrMutuallyExclusive))
	}

	if l := len(o.Content); l > 0 {
		if l != 1 {
			errs = append(errs, newValidationError(joinLoc(location, "content"), "invalid number of items, expected only one, but got '%d'", l))
		}
		for k, v := range o.Content {
			errs = append(errs, v.validateSpec(joinLoc(location, "content", k), validator)...)
		}
	}
	if o.Schema != nil {
		errs = append(errs, o.Schema.validateSpec(joinLoc(location, "schema"), validator)...)
	}

	switch o.In {
	case InQuery, InHeader, InPath, InCookie:
	case "":
		errs = append(errs, newValidationError(joinLoc(location, "in"), ErrRequired))
	default:
		errs = append(errs, newValidationError(joinLoc(location, "in"), "invalid value, expected one of [%s, %s, %s, %s], but got '%s'", InQuery, InHeader, InPath, InCookie, o.In))
	}

	switch o.Style {
	case "":
	case StyleMatrix, StyleLabel:
		if o.In != InPath {
			errs = append(errs, newValidationError(joinLoc(location, "style"), "only allowed when `in` is '%s'", InPath))
		}
	case StyleForm:
		if o.In != InQuery && o.In != InCookie {
			errs = append(errs, newValidationError(joinLoc(location, "style"), "only allowed when `in` is '%s' or '%s' ", InQuery, InCookie))
		}
	case StyleSimple:
		if o.In != InPath && o.In != InHeader {
			errs = append(errs, newValidationError(joinLoc(location, "style"), "only allowed when `in` is '%s' or '%s' ", InPath, InHeader))
		}
	case StyleSpaceDelimited, StylePipeDelimited, StyleDeepObject:
		if o.In != InQuery {
			errs = append(errs, newValidationError(joinLoc(location, "style"), "only allowed when `in` is '%s'", InQuery))
		}
	default:
		errs = append(errs, newValidationError(joinLoc(location, "style"), "invalid value, expected one of [%s, %s, %s, %s, %s, %s, %s], but got '%s'", StyleMatrix, StyleLabel, StyleForm, StyleSimple, StyleSpaceDelimited, StylePipeDelimited, StyleDeepObject, o.Style))
	}

	switch {
	case o.Name == "":
		errs = append(errs, newValidationError(joinLoc(location, "name"), ErrRequired))
	case o.In == InPath && !PathNamePattern.MatchString(o.Name):
		errs = append(errs, newValidationError(joinLoc(location, "name"), "must match pattern '%s', but got '%s'", PathNamePattern, o.Name))
	case !o.AllowReserved && o.In == InQuery && strings.ContainsAny(o.Name, ReservedCharacters):
		errs = append(errs, newValidationError(joinLoc(location, "name"), "'%s' contains reserved characters: '%s'", o.Name, ReservedCharacters))
	}

	if o.AllowReserved && o.In != InQuery {
		errs = append(errs, newValidationError(joinLoc(location, "allowReserved"), "only allowed when `in` is '%s'", InQuery))
	}

	if o.AllowEmptyValue && o.In != InQuery {
		errs = append(errs, newValidationError(joinLoc(location, "allowEmptyValue"), "only allowed when `in` is '%s'", InQuery))
	}

	if !o.Required && o.In == InPath {
		errs = append(errs, newValidationError(joinLoc(location, "required"), "must be `true` when `in` is '%s'", InPath))
	}

	if validator.opts.doNotValidateExamples {
		return errs
	}
	var schemaRef string
	if o.Schema != nil {
		schemaRef = o.Schema.getLocationOrRef(joinLoc(location, "schema"))
	} else if len(o.Content) > 0 {
		for k, v := range o.Content {
			schemaRef = v.Spec.Schema.getLocationOrRef(joinLoc(location, "content", k, "schema"))
			break
		}
	}

	if schemaRef == "'" {
		errs = append(errs, newValidationError(location, "unable to validate examples without schema or content"))
		return errs
	}

	if o.Example != nil {
		if e := validator.ValidateData(joinLoc(location, "schema"), o.Example); e != nil {
			errs = append(errs, newValidationError(joinLoc(location, "example"), e))
		}
	}
	if len(o.Examples) > 0 {
		for k, v := range o.Examples {
			example, err := v.GetSpec(validator.spec.Spec.Components)
			if err != nil {
				// do not add the error, because it is already validated earlier
				continue
			}
			if value := example.Spec.Value; value != nil {
				if e := validator.ValidateData(joinLoc(location, "schema"), value); e != nil {
					errs = append(errs, newValidationError(joinLoc(location, "examples", k), e))
				}
			}
		}
	}
	return errs
}

type ParameterBuilder struct {
	spec *RefOrSpec[Extendable[Parameter]]
}

func NewParameterBuilder() *ParameterBuilder {
	return &ParameterBuilder{
		spec: NewRefOrExtSpec[Parameter](&Parameter{}),
	}
}

func (b *ParameterBuilder) Build() *RefOrSpec[Extendable[Parameter]] {
	return b.spec
}

func (b *ParameterBuilder) Extensions(v map[string]any) *ParameterBuilder {
	b.spec.Spec.Extensions = v
	return b
}

func (b *ParameterBuilder) AddExt(name string, value any) *ParameterBuilder {
	b.spec.Spec.AddExt(name, value)
	return b
}

func (b *ParameterBuilder) Example(v any) *ParameterBuilder {
	b.spec.Spec.Spec.Example = v
	return b
}

func (b *ParameterBuilder) Content(v map[string]*Extendable[MediaType]) *ParameterBuilder {
	b.spec.Spec.Spec.Content = v
	return b
}

func (b *ParameterBuilder) AddContent(name string, value *Extendable[MediaType]) *ParameterBuilder {
	if b.spec.Spec.Spec.Content == nil {
		b.spec.Spec.Spec.Content = make(map[string]*Extendable[MediaType], 1)
	}
	b.spec.Spec.Spec.Content[name] = value
	return b
}

func (b *ParameterBuilder) Examples(v map[string]*RefOrSpec[Extendable[Example]]) *ParameterBuilder {
	b.spec.Spec.Spec.Examples = v
	return b
}

func (b *ParameterBuilder) AddExample(name string, value *RefOrSpec[Extendable[Example]]) *ParameterBuilder {
	if b.spec.Spec.Spec.Examples == nil {
		b.spec.Spec.Spec.Examples = make(map[string]*RefOrSpec[Extendable[Example]], 1)
	}
	b.spec.Spec.Spec.Examples[name] = value
	return b
}

func (b *ParameterBuilder) Schema(v *RefOrSpec[Schema]) *ParameterBuilder {
	b.spec.Spec.Spec.Schema = v
	return b
}

func (b *ParameterBuilder) In(v string) *ParameterBuilder {
	b.spec.Spec.Spec.In = v
	return b
}

func (b *ParameterBuilder) Description(v string) *ParameterBuilder {
	b.spec.Spec.Spec.Description = v
	return b
}

func (b *ParameterBuilder) Style(v string) *ParameterBuilder {
	b.spec.Spec.Spec.Style = v
	return b
}

func (b *ParameterBuilder) Name(v string) *ParameterBuilder {
	b.spec.Spec.Spec.Name = v
	return b
}

func (b *ParameterBuilder) Explode(v bool) *ParameterBuilder {
	b.spec.Spec.Spec.Explode = v
	return b
}

func (b *ParameterBuilder) AllowReserved(v bool) *ParameterBuilder {
	b.spec.Spec.Spec.AllowReserved = v
	return b
}

func (b *ParameterBuilder) AllowEmptyValue(v bool) *ParameterBuilder {
	b.spec.Spec.Spec.AllowEmptyValue = v
	return b
}

func (b *ParameterBuilder) Deprecated(v bool) *ParameterBuilder {
	b.spec.Spec.Spec.Deprecated = v
	return b
}

func (b *ParameterBuilder) Required(v bool) *ParameterBuilder {
	b.spec.Spec.Spec.Required = v
	return b
}
