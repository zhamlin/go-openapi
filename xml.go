package openapi

// XML is a metadata object that allows for more fine-tuned XML model definitions.
// When using arrays, XML element names are not inferred (for singular/plural forms) and the name property SHOULD
// be used to add that information.
// See examples for expected behavior.
//
// https://spec.openapis.org/oas/v3.1.1#xml-object
//
// Example:
//
//	Person:
//	  type: object
//	  properties:
//	    id:
//	      type: integer
//	      format: int32
//	      xml:
//	        attribute: true
//	    name:
//	      type: string
//	      xml:
//	        namespace: https://example.com/schema/sample
//	        prefix: sample
//
//	<Person id="123">
//	    <sample:name xmlns:sample="https://example.com/schema/sample">example</sample:name>
//	</Person>
type XML struct {
	// Replaces the name of the element/attribute used for the described schema property.
	// When defined within items, it will affect the name of the individual XML elements within the list.
	// When defined alongside type being array (outside the items), it will affect the wrapping element and only if wrapped is true.
	// If wrapped is false, it will be ignored.
	Name string `json:"name,omitempty"`
	// The URI of the namespace definition.
	// This MUST be in the form of an absolute URI.
	Namespace string `json:"namespace,omitempty"`
	// The prefix to be used for the name.
	Prefix string `json:"prefix,omitempty"`
	// Declares whether the property definition translates to an attribute instead of an element. Default value is false.
	Attribute bool `json:"attribute,omitempty"`
	// MAY be used only for an array definition.
	// Signifies whether the array is wrapped (for example, <books><book/><book/></books>) or unwrapped (<book/><book/>).
	// Default value is false.
	// The definition takes effect only when defined alongside type being array (outside the items).
	Wrapped bool `json:"wrapped,omitempty"`
}

func (o *XML) validateSpec(path string, validator *Validator) []*validationError {
	return nil // nothing to validate
}

type XMLBuilder struct {
	spec *Extendable[XML]
}

func NewXMLBuilder() *XMLBuilder {
	return &XMLBuilder{
		spec: NewExtendable[XML](&XML{}),
	}
}

func (b *XMLBuilder) Build() *Extendable[XML] {
	return b.spec
}

func (b *XMLBuilder) Extensions(v map[string]any) *XMLBuilder {
	b.spec.Extensions = v
	return b
}

func (b *XMLBuilder) AddExt(name string, value any) *XMLBuilder {
	b.spec.AddExt(name, value)
	return b
}

func (b *XMLBuilder) Name(v string) *XMLBuilder {
	b.spec.Spec.Name = v
	return b
}

func (b *XMLBuilder) Namespace(v string) *XMLBuilder {
	b.spec.Spec.Namespace = v
	return b
}

func (b *XMLBuilder) Prefix(v string) *XMLBuilder {
	b.spec.Spec.Prefix = v
	return b
}

func (b *XMLBuilder) Attribute(v bool) *XMLBuilder {
	b.spec.Spec.Attribute = v
	return b
}

func (b *XMLBuilder) Wrapped(v bool) *XMLBuilder {
	b.spec.Spec.Wrapped = v
	return b
}
