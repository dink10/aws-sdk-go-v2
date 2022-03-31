// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The resource with the name requested already exists.
type AlreadyExistsException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *AlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AlreadyExistsException) ErrorCode() string             { return "AlreadyExistsException" }
func (e *AlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified client token has already been used in another resource request.
// It's best practice for client tokens to be unique for each resource operation
// request. However, client token expire after 36 hours.
type ClientTokenConflictException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ClientTokenConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ClientTokenConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ClientTokenConflictException) ErrorCode() string             { return "ClientTokenConflictException" }
func (e *ClientTokenConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource is currently being modified by another operation.
type ConcurrentModificationException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ConcurrentModificationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentModificationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentModificationException) ErrorCode() string {
	return "ConcurrentModificationException"
}
func (e *ConcurrentModificationException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Another resource operation is currently being performed on this resource.
type ConcurrentOperationException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ConcurrentOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentOperationException) ErrorCode() string             { return "ConcurrentOperationException" }
func (e *ConcurrentOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource handler has returned that the downstream service generated an error
// that doesn't map to any other handler error code.
type GeneralServiceException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *GeneralServiceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *GeneralServiceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *GeneralServiceException) ErrorCode() string             { return "GeneralServiceException" }
func (e *GeneralServiceException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource handler has failed without a returning a more specific error code.
// This can include timeouts.
type HandlerFailureException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *HandlerFailureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *HandlerFailureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *HandlerFailureException) ErrorCode() string             { return "HandlerFailureException" }
func (e *HandlerFailureException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The resource handler has returned that an unexpected error occurred within the
// resource handler.
type HandlerInternalFailureException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *HandlerInternalFailureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *HandlerInternalFailureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *HandlerInternalFailureException) ErrorCode() string {
	return "HandlerInternalFailureException"
}
func (e *HandlerInternalFailureException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The resource handler has returned that the credentials provided by the user are
// invalid.
type InvalidCredentialsException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidCredentialsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidCredentialsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidCredentialsException) ErrorCode() string             { return "InvalidCredentialsException" }
func (e *InvalidCredentialsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource handler has returned that invalid input from the user has generated
// a generic exception.
type InvalidRequestException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string             { return "InvalidRequestException" }
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource handler has returned that the request couldn't be completed due to
// networking issues, such as a failure to receive a response from the server.
type NetworkFailureException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *NetworkFailureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NetworkFailureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NetworkFailureException) ErrorCode() string             { return "NetworkFailureException" }
func (e *NetworkFailureException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The resource handler has returned that the downstream resource failed to
// complete all of its ready-state checks.
type NotStabilizedException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *NotStabilizedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotStabilizedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotStabilizedException) ErrorCode() string             { return "NotStabilizedException" }
func (e *NotStabilizedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more properties included in this resource operation are defined as
// create-only, and therefore can't be updated.
type NotUpdatableException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *NotUpdatableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotUpdatableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotUpdatableException) ErrorCode() string             { return "NotUpdatableException" }
func (e *NotUpdatableException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Cloud Control API hasn't received a valid response from the resource handler,
// due to a configuration error. This includes issues such as the resource handler
// returning an invalid response, or timing out.
type PrivateTypeException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *PrivateTypeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PrivateTypeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PrivateTypeException) ErrorCode() string             { return "PrivateTypeException" }
func (e *PrivateTypeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A resource operation with the specified request token can't be found.
type RequestTokenNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *RequestTokenNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RequestTokenNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RequestTokenNotFoundException) ErrorCode() string             { return "RequestTokenNotFoundException" }
func (e *RequestTokenNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource is temporarily unavailable to be acted upon. For example, if the
// resource is currently undergoing an operation and can't be acted upon until that
// operation is finished.
type ResourceConflictException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ResourceConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceConflictException) ErrorCode() string             { return "ResourceConflictException" }
func (e *ResourceConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A resource with the specified identifier can't be found.
type ResourceNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource handler has returned that the downstream service returned an
// internal error, typically with a 5XX HTTP status code.
type ServiceInternalErrorException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ServiceInternalErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceInternalErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceInternalErrorException) ErrorCode() string             { return "ServiceInternalErrorException" }
func (e *ServiceInternalErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The resource handler has returned that a non-transient resource limit was
// reached on the service side.
type ServiceLimitExceededException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ServiceLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceLimitExceededException) ErrorCode() string             { return "ServiceLimitExceededException" }
func (e *ServiceLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was denied due to request throttling.
type ThrottlingException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottlingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottlingException) ErrorCode() string             { return "ThrottlingException" }
func (e *ThrottlingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified extension doesn't exist in the CloudFormation registry.
type TypeNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *TypeNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TypeNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TypeNotFoundException) ErrorCode() string             { return "TypeNotFoundException" }
func (e *TypeNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource doesn't support this resource operation.
type UnsupportedActionException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *UnsupportedActionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedActionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedActionException) ErrorCode() string             { return "UnsupportedActionException" }
func (e *UnsupportedActionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
