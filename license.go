package openapi

// License information for the exposed API.
//
// https://spec.openapis.org/oas/v3.1.1#license-object
//
// Example:
//
//	name: Apache 2.0
//	identifier: Apache-2.0
type License struct {
	// REQUIRED.
	// The license name used for the API.
	Name string `json:"name"`
	// An SPDX license expression for the API.
	// The identifier field is mutually exclusive of the url field.
	Identifier string `json:"identifier,omitempty"`
	// A URL to the license used for the API.
	// This MUST be in the form of a URL.
	// The url field is mutually exclusive of the identifier field.
	URL string `json:"url,omitempty"`
}

func (o *License) validateSpec(location string, _ *Validator) []*validationError {
	var errs []*validationError
	if o.Name == "" {
		errs = append(errs, newValidationError(joinLoc(location, "name"), ErrRequired))
	}
	if o.Identifier != "" && o.URL != "" {
		errs = append(errs, newValidationError(joinLoc(location, "identifier&url"), ErrMutuallyExclusive))
	}
	if err := checkURL(o.URL); err != nil {
		errs = append(errs, newValidationError(joinLoc(location, "url"), err))
	}
	return errs
}

type LicenseBuilder struct {
	spec *Extendable[License]
}

func NewLicenseBuilder() *LicenseBuilder {
	return &LicenseBuilder{
		spec: NewExtendable[License](&License{}),
	}
}

func (b *LicenseBuilder) Build() *Extendable[License] {
	return b.spec
}

func (b *LicenseBuilder) Extensions(v map[string]any) *LicenseBuilder {
	b.spec.Extensions = v
	return b
}

func (b *LicenseBuilder) AddExt(name string, value any) *LicenseBuilder {
	b.spec.AddExt(name, value)
	return b
}

func (b *LicenseBuilder) Name(v string) *LicenseBuilder {
	b.spec.Spec.Name = v
	return b
}

func (b *LicenseBuilder) Identifier(v string) *LicenseBuilder {
	b.spec.Spec.Identifier = v
	return b
}

func (b *LicenseBuilder) URL(v string) *LicenseBuilder {
	b.spec.Spec.URL = v
	return b
}
