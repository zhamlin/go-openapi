package openapi

// OAuthFlows allows configuration of the supported OAuth Flows.
//
// https://spec.openapis.org/oas/v3.1.1#oauth-flows-object
//
// Example:
//
//	type: oauth2
//	flows:
//	  implicit:
//	    authorizationUrl: https://example.com/api/oauth/dialog
//	    scopes:
//	      write:pets: modify pets in your account
//	      read:pets: read your pets
//	  authorizationCode:
//	    authorizationUrl: https://example.com/api/oauth/dialog
//	    tokenUrl: https://example.com/api/oauth/token
//	    scopes:
//	      write:pets: modify pets in your account
//	      read:pets: read your pets
type OAuthFlows struct {
	// Configuration for the OAuth Implicit flow.
	Implicit *Extendable[OAuthFlow] `json:"implicit,omitempty"`
	// Configuration for the OAuth Resource Owner Password flow.
	Password *Extendable[OAuthFlow] `json:"password,omitempty"`
	// Configuration for the OAuth Client Credentials flow.
	// Previously called application in OpenAPI 2.0.
	ClientCredentials *Extendable[OAuthFlow] `json:"clientCredentials,omitempty"`
	// Configuration for the OAuth Authorization Code flow.
	// Previously called accessCode in OpenAPI 2.0.
	AuthorizationCode *Extendable[OAuthFlow] `json:"authorizationCode,omitempty"`
}

func (o *OAuthFlows) validateSpec(location string, validator *Validator) []*validationError {
	var errs []*validationError
	if o.Implicit != nil {
		errs = append(errs, o.Implicit.validateSpec(joinLoc(location, "implicit"), validator)...)
		if o.Implicit.Spec.AuthorizationURL == "" {
			errs = append(errs, newValidationError(joinLoc(location, "implicit", "authorizationUrl"), ErrRequired))
		}
	}
	if o.Password != nil {
		errs = append(errs, o.Password.validateSpec(joinLoc(location, "password"), validator)...)
		if o.Password.Spec.TokenURL == "" {
			errs = append(errs, newValidationError(joinLoc(location, "password", "tokenUrl"), ErrRequired))
		}
	}
	if o.ClientCredentials != nil {
		errs = append(errs, o.ClientCredentials.validateSpec(joinLoc(location, "clientCredentials"), validator)...)
		if o.ClientCredentials.Spec.TokenURL == "" {
			errs = append(errs, newValidationError(joinLoc(location, "clientCredentials", "tokenUrl"), ErrRequired))
		}
	}
	if o.AuthorizationCode != nil {
		errs = append(errs, o.AuthorizationCode.validateSpec(joinLoc(location, "authorizationCode"), validator)...)
		if o.AuthorizationCode.Spec.AuthorizationURL == "" {
			errs = append(errs, newValidationError(joinLoc(location, "authorizationCode", "authorizationUrl"), ErrRequired))
		}
		if o.AuthorizationCode.Spec.TokenURL == "" {
			errs = append(errs, newValidationError(joinLoc(location, "authorizationCode", "tokenUrl"), ErrRequired))
		}
	}

	return errs
}

type OAuthFlowsBuilder struct {
	spec *Extendable[OAuthFlows]
}

func NewOAuthFlowsBuilder() *OAuthFlowsBuilder {
	return &OAuthFlowsBuilder{
		spec: NewExtendable[OAuthFlows](&OAuthFlows{}),
	}
}

func (b *OAuthFlowsBuilder) Build() *Extendable[OAuthFlows] {
	return b.spec
}

func (b *OAuthFlowsBuilder) Extensions(v map[string]any) *OAuthFlowsBuilder {
	b.spec.Extensions = v
	return b
}

func (b *OAuthFlowsBuilder) AddExt(name string, value any) *OAuthFlowsBuilder {
	b.spec.AddExt(name, value)
	return b
}

func (b *OAuthFlowsBuilder) Implicit(v *Extendable[OAuthFlow]) *OAuthFlowsBuilder {
	b.spec.Spec.Implicit = v
	return b
}

func (b *OAuthFlowsBuilder) Password(v *Extendable[OAuthFlow]) *OAuthFlowsBuilder {
	b.spec.Spec.Password = v
	return b
}

func (b *OAuthFlowsBuilder) ClientCredentials(v *Extendable[OAuthFlow]) *OAuthFlowsBuilder {
	b.spec.Spec.ClientCredentials = v
	return b
}

func (b *OAuthFlowsBuilder) AuthorizationCode(v *Extendable[OAuthFlow]) *OAuthFlowsBuilder {
	b.spec.Spec.AuthorizationCode = v
	return b
}
