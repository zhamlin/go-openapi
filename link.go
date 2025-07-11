package openapi

// Link represents a possible design-time link for a response.
// The presence of a link does not guarantee the caller’s ability to successfully invoke it,
// rather it provides a known relationship and traversal mechanism between responses and other operations.
// Unlike dynamic links (i.e. links provided in the response payload),
// the OAS linking mechanism does not require link information in the runtime response.
// For computing links, and providing instructions to execute them,
// a runtime expression is used for accessing values in an operation
// and using them as parameters while invoking the linked operation.
//
// https://spec.openapis.org/oas/v3.1.1#link-object
//
// Example:
//
//	paths:
//	  /users/{id}:
//	    parameters:
//	    - name: id
//	      in: path
//	      required: true
//	      description: the user identifier, as userId
//	      schema:
//	        type: string
//	    get:
//	      responses:
//	        '200':
//	          description: the user being returned
//	          content:
//	            application/json:
//	              schema:
//	                type: object
//	                properties:
//	                  uuid: # the unique user id
//	                    type: string
//	                    format: uuid
//	          links:
//	            address:
//	              # the target link operationId
//	              operationId: getUserAddress
//	              parameters:
//	                # get the `id` field from the request path parameter named `id`
//	                userId: $request.path.id
//	  # the path item of the linked operation
//	  /users/{userid}/address:
//	    parameters:
//	    - name: userid
//	      in: path
//	      required: true
//	      description: the user identifier, as userId
//	      schema:
//	        type: string
//	    # linked operation
//	    get:
//	      operationId: getUserAddress
//	      responses:
//	        '200':
//	          description: the user's address
type Link struct {
	// A literal value or {expression} to use as a request body when calling the target operation.
	RequestBody any `json:"requestBody,omitempty"`
	// A map representing parameters to pass to an operation as specified with operationId or identified via operationRef.
	// The key is the parameter name to be used, whereas the value can be a constant or an expression to be evaluated and
	// passed to the linked operation.
	// The parameter name can be qualified using the parameter location [{in}.]{name} for operations that use
	// the same parameter name in different locations (e.g. path.id).
	Parameters map[string]any `json:"parameters,omitempty"`
	// A server object to be used by the target operation.
	Server *Extendable[Server] `json:"server,omitempty"`
	// A relative or absolute URI reference to an OAS operation.
	// This field is mutually exclusive of the operationId field, and MUST point to an Operation Object.
	// Relative operationRef values MAY be used to locate an existing Operation Object in the OpenAPI definition.
	// See the rules for resolving Relative References.
	OperationRef string `json:"operationRef,omitempty"`
	// The name of an existing, resolvable OAS operation, as defined with a unique operationId.
	// This field is mutually exclusive of the operationRef field.
	OperationID string `json:"operationId,omitempty"`
	// A description of the link.
	// CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
}

func (o *Link) validateSpec(location string, validator *Validator) []*validationError {
	var errs []*validationError
	if o.OperationRef != "" && o.OperationID != "" {
		errs = append(errs, newValidationError(joinLoc(location, "operationRef&operationId"), ErrMutuallyExclusive))
	}
	if o.OperationID != "" {
		id := joinLoc("operations", o.OperationID)
		if !validator.visited[id] {
			validator.linkToOperationID[joinLoc(location, "operationId")] = o.OperationID
		}
	}
	// uncomment when JSONLookup is implemented
	// if o.OperationRef != "" {
	//	ref := NewRefOrExtSpec[Operation](o.OperationRef)
	//	errs = append(errs, ref.validateSpec(joinLoc(location, "operationRef"), validator)...)
	//}
	if o.Server != nil {
		errs = append(errs, o.Server.validateSpec(joinLoc(location, "server"), validator)...)
	}
	return errs
}

type LinkBuilder struct {
	spec *RefOrSpec[Extendable[Link]]
}

func NewLinkBuilder() *LinkBuilder {
	return &LinkBuilder{
		spec: NewRefOrExtSpec[Link](&Link{}),
	}
}

func (b *LinkBuilder) Build() *RefOrSpec[Extendable[Link]] {
	return b.spec
}

func (b *LinkBuilder) Extensions(v map[string]any) *LinkBuilder {
	b.spec.Spec.Extensions = v
	return b
}

func (b *LinkBuilder) AddExt(name string, value any) *LinkBuilder {
	b.spec.Spec.AddExt(name, value)
	return b
}

func (b *LinkBuilder) RequestBody(v any) *LinkBuilder {
	b.spec.Spec.Spec.RequestBody = v
	return b
}

func (b *LinkBuilder) Parameters(v map[string]any) *LinkBuilder {
	b.spec.Spec.Spec.Parameters = v
	return b
}

func (b *LinkBuilder) AddParameter(name string, value any) *LinkBuilder {
	if b.spec.Spec.Spec.Parameters == nil {
		b.spec.Spec.Spec.Parameters = make(map[string]any, 1)
	}
	b.spec.Spec.Spec.Parameters[name] = value
	return b
}

func (b *LinkBuilder) Server(v *Extendable[Server]) *LinkBuilder {
	b.spec.Spec.Spec.Server = v
	return b
}

func (b *LinkBuilder) OperationRef(v string) *LinkBuilder {
	b.spec.Spec.Spec.OperationRef = v
	return b
}

func (b *LinkBuilder) OperationID(v string) *LinkBuilder {
	b.spec.Spec.Spec.OperationID = v
	return b
}

func (b *LinkBuilder) Description(v string) *LinkBuilder {
	b.spec.Spec.Spec.Description = v
	return b
}
