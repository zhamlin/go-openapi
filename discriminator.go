package openapi

// Discriminator is used when request bodies or response payloads may be one of a number of different schemas,
// a discriminator object can be used to aid in serialization, deserialization, and validation.
// The discriminator is a specific object in a schema which is used to inform the consumer of the document of
// an alternative schema based on the value associated with it.
// When using the discriminator, inline schemas will not be considered.
//
// https://spec.openapis.org/oas/v3.1.1#discriminator-object
//
// Example:
//
//	MyResponseType:
//	  oneOf:
//	  - $ref: '#/components/schemas/Cat'
//	  - $ref: '#/components/schemas/Dog'
//	  - $ref: '#/components/schemas/Lizard'
//	  - $ref: 'https://gigantic-server.com/schemas/Monster/schema.json'
//	  discriminator:
//	    propertyName: petType
//	    mapping:
//	      dog: '#/components/schemas/Dog'
//	      monster: 'https://gigantic-server.com/schemas/Monster/schema.json'
type Discriminator struct {
	// An object to hold mappings between payload values and schema names or references.
	Mapping map[string]string `json:"mapping,omitempty"`
	// REQUIRED.
	// The name of the property in the payload that will hold the discriminator value.
	PropertyName string `json:"propertyName"`
}

func (o *Discriminator) validateSpec(location string, validator *Validator) []*validationError {
	var errs []*validationError
	if o.PropertyName == "" {
		errs = append(errs, newValidationError(joinLoc(location, "propertyName"), ErrRequired))
	}
	for k, v := range o.Mapping {
		ref := NewRefOrSpec[Schema](v)
		errs = append(errs, ref.validateSpec(joinLoc(location, "mapping", k), validator)...)
	}
	return errs
}

type DiscriminatorBuilder struct {
	spec *Discriminator
}

func NewDiscriminatorBuilder() *DiscriminatorBuilder {
	return &DiscriminatorBuilder{
		spec: &Discriminator{},
	}
}

func (b *DiscriminatorBuilder) Build() *Discriminator {
	return b.spec
}

func (b *DiscriminatorBuilder) Mapping(v map[string]string) *DiscriminatorBuilder {
	b.spec.Mapping = v
	return b
}

func (b *DiscriminatorBuilder) AddMapping(name, value string) *DiscriminatorBuilder {
	if b.spec.Mapping == nil {
		b.spec.Mapping = make(map[string]string, 1)
	}
	b.spec.Mapping[name] = value
	return b
}

func (b *DiscriminatorBuilder) PropertyName(v string) *DiscriminatorBuilder {
	b.spec.PropertyName = v
	return b
}
