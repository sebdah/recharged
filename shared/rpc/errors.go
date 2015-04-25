package rpc

type Errorer interface {
	Error() string
	GetCode() string
	GetDescription() string
	GetDetails() string
	SetCode(string)
	SetDescription(string)
	SetDetails(string)
}

type NotImplementedError struct {
	Code        string
	Description string
	Details     string
}

type NotSupportedError struct {
	Code        string
	Description string
	Details     string
}

type InternalError struct {
	Code        string
	Description string
	Details     string
}

type ProtocolError struct {
	Code        string
	Description string
	Details     string
}

type SecurityError struct {
	Code        string
	Description string
	Details     string
}

type FormationViolation struct {
	Code        string
	Description string
	Details     string
}

type PropertyConstraintViolation struct {
	Code        string
	Description string
	Details     string
}

type OccurenceConstraintViolation struct {
	Code        string
	Description string
	Details     string
}

type TypeConstraintViolation struct {
	Code        string
	Description string
	Details     string
}

type GenericError struct {
	Code        string
	Description string
	Details     string
}

func NewNotImplementedError() *NotImplementedError {
	err := new(NotImplementedError)
	err.Code = "NotImplemented"
	err.Description = "Requested Action is not known by receiver"
	err.Details = "{}"
	return err
}

func (this *NotImplementedError) Error() string {
	return this.Description
}

func (this *NotImplementedError) GetCode() string {
	return this.Code
}

func (this *NotImplementedError) GetDescription() string {
	return this.Description
}

func (this *NotImplementedError) GetDetails() string {
	return this.Details
}

func (this *NotImplementedError) SetCode(code string) {
	this.Code = code
}

func (this *NotImplementedError) SetDescription(description string) {
	this.Description = description
}

func (this *NotImplementedError) SetDetails(details string) {
	this.Details = details
}

func NewNotSupportedError() *NotSupportedError {
	err := new(NotSupportedError)
	err.Code = "NotSupported"
	err.Description = "Requested Action is recognized but not supported by the receiver"
	err.Details = "{}"
	return err
}

func (this *NotSupportedError) Error() string {
	return this.Description
}

func (this *NotSupportedError) GetCode() string {
	return this.Code
}

func (this *NotSupportedError) GetDescription() string {
	return this.Description
}

func (this *NotSupportedError) GetDetails() string {
	return this.Details
}

func (this *NotSupportedError) SetCode(code string) {
	this.Code = code
}

func (this *NotSupportedError) SetDescription(description string) {
	this.Description = description
}

func (this *NotSupportedError) SetDetails(details string) {
	this.Details = details
}

func NewProtocolError() *ProtocolError {
	err := new(ProtocolError)
	err.Code = "ProtocolError"
	err.Description = "Payload for Action is incomplete"
	err.Details = "{}"
	return err
}

func (this *ProtocolError) Error() string {
	return this.Description
}

func (this *ProtocolError) GetCode() string {
	return this.Code
}

func (this *ProtocolError) GetDescription() string {
	return this.Description
}

func (this *ProtocolError) GetDetails() string {
	return this.Details
}

func (this *ProtocolError) SetCode(code string) {
	this.Code = code
}

func (this *ProtocolError) SetDescription(description string) {
	this.Description = description
}

func (this *ProtocolError) SetDetails(details string) {
	this.Details = details
}

func NewFormationViolation() *FormationViolation {
	err := new(FormationViolation)
	err.Code = "FormationViolation"
	err.Description = "Payload for Action is syntactically incorrect or not conform the PDU structure for Action"
	err.Details = "{}"
	return err
}

func (this *FormationViolation) Error() string {
	return this.Description
}

func (this *FormationViolation) GetCode() string {
	return this.Code
}

func (this *FormationViolation) GetDescription() string {
	return this.Description
}

func (this *FormationViolation) GetDetails() string {
	return this.Details
}

func (this *FormationViolation) SetCode(code string) {
	this.Code = code
}

func (this *FormationViolation) SetDescription(description string) {
	this.Description = description
}

func (this *FormationViolation) SetDetails(details string) {
	this.Details = details
}

func NewOccurenceConstraintViolation() *OccurenceConstraintViolation {
	err := new(OccurenceConstraintViolation)
	err.Code = "OccurenceConstraintViolation"
	err.Description = "Payload for Action is syntactically correct but at least one of the fields violates occurence constraints"
	err.Details = "{}"
	return err
}

func (this *OccurenceConstraintViolation) Error() string {
	return this.Description
}

func (this *OccurenceConstraintViolation) GetCode() string {
	return this.Code
}

func (this *OccurenceConstraintViolation) GetDescription() string {
	return this.Description
}

func (this *OccurenceConstraintViolation) GetDetails() string {
	return this.Details
}

func (this *OccurenceConstraintViolation) SetCode(code string) {
	this.Code = code
}

func (this *OccurenceConstraintViolation) SetDescription(description string) {
	this.Description = description
}

func (this *OccurenceConstraintViolation) SetDetails(details string) {
	this.Details = details
}

func NewInternalError() *InternalError {
	err := new(InternalError)
	err.Code = "InternalError"
	err.Description = "An internal error occurred and the receiver was not able to process the requested Action successfully"
	err.Details = "{}"
	return err
}

func (this *InternalError) Error() string {
	return this.Description
}

func (this *InternalError) GetCode() string {
	return this.Code
}

func (this *InternalError) GetDescription() string {
	return this.Description
}

func (this *InternalError) GetDetails() string {
	return this.Details
}

func (this *InternalError) SetCode(code string) {
	this.Code = code
}

func (this *InternalError) SetDescription(description string) {
	this.Description = description
}

func (this *InternalError) SetDetails(details string) {
	this.Details = details
}

func NewSecurityError() *SecurityError {
	err := new(SecurityError)
	err.Code = "SecurityError"
	err.Description = "During the processing of Action a security issue occurred preventing receiver from completing the Action successfully"
	err.Details = "{}"
	return err
}

func (this *SecurityError) Error() string {
	return this.Description
}

func (this *SecurityError) GetCode() string {
	return this.Code
}

func (this *SecurityError) GetDescription() string {
	return this.Description
}

func (this *SecurityError) GetDetails() string {
	return this.Details
}

func (this *SecurityError) SetCode(code string) {
	this.Code = code
}

func (this *SecurityError) SetDescription(description string) {
	this.Description = description
}

func (this *SecurityError) SetDetails(details string) {
	this.Details = details
}

func NewPropertyConstraintViolation() *PropertyConstraintViolation {
	err := new(PropertyConstraintViolation)
	err.Code = "PropertyConstraintViolation"
	err.Description = "Payload is syntactically correct but at least one field contains an invalid value"
	err.Details = "{}"
	return err
}

func (this *PropertyConstraintViolation) Error() string {
	return this.Description
}

func (this *PropertyConstraintViolation) GetCode() string {
	return this.Code
}

func (this *PropertyConstraintViolation) GetDescription() string {
	return this.Description
}

func (this *PropertyConstraintViolation) GetDetails() string {
	return this.Details
}

func (this *PropertyConstraintViolation) SetCode(code string) {
	this.Code = code
}

func (this *PropertyConstraintViolation) SetDescription(description string) {
	this.Description = description
}

func (this *PropertyConstraintViolation) SetDetails(details string) {
	this.Details = details
}

func NewTypeConstraintViolation() *TypeConstraintViolation {
	err := new(TypeConstraintViolation)
	err.Code = "TypeConstraintViolation"
	err.Description = "Payload for Action is syntactically correct but at least one of the fields violates data type constraints (e.g. “somestring”: 12)"
	err.Details = "{}"
	return err
}

func (this *TypeConstraintViolation) Error() string {
	return this.Description
}

func (this *TypeConstraintViolation) GetCode() string {
	return this.Code
}

func (this *TypeConstraintViolation) GetDescription() string {
	return this.Description
}

func (this *TypeConstraintViolation) GetDetails() string {
	return this.Details
}

func (this *TypeConstraintViolation) SetCode(code string) {
	this.Code = code
}

func (this *TypeConstraintViolation) SetDescription(description string) {
	this.Description = description
}

func (this *TypeConstraintViolation) SetDetails(details string) {
	this.Details = details
}

func NewGenericError() *GenericError {
	err := new(GenericError)
	err.Code = "GenericError"
	err.Description = "Any other error not covered by the previous ones"
	err.Details = "{}"
	return err
}

func (this *GenericError) Error() string {
	return this.Description
}

func (this *GenericError) GetCode() string {
	return this.Code
}

func (this *GenericError) GetDescription() string {
	return this.Description
}

func (this *GenericError) GetDetails() string {
	return this.Details
}

func (this *GenericError) SetCode(code string) {
	this.Code = code
}

func (this *GenericError) SetDescription(description string) {
	this.Description = description
}

func (this *GenericError) SetDetails(details string) {
	this.Details = details
}
