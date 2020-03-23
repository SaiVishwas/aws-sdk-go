// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscalingplans

import (
	"github.com/SaiVishwas/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeConcurrentUpdateException for service response error code
	// "ConcurrentUpdateException".
	//
	// Concurrent updates caused an exception, for example, if you request an update
	// to a scaling plan that already has a pending update.
	ErrCodeConcurrentUpdateException = "ConcurrentUpdateException"

	// ErrCodeInternalServiceException for service response error code
	// "InternalServiceException".
	//
	// The service encountered an internal error.
	ErrCodeInternalServiceException = "InternalServiceException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// The token provided is not valid.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// Your account exceeded a limit. This exception is thrown when a per-account
	// resource limit is exceeded.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeObjectNotFoundException for service response error code
	// "ObjectNotFoundException".
	//
	// The specified object could not be found.
	ErrCodeObjectNotFoundException = "ObjectNotFoundException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// An exception was thrown for a validation issue. Review the parameters provided.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"ConcurrentUpdateException": newErrorConcurrentUpdateException,
	"InternalServiceException":  newErrorInternalServiceException,
	"InvalidNextTokenException": newErrorInvalidNextTokenException,
	"LimitExceededException":    newErrorLimitExceededException,
	"ObjectNotFoundException":   newErrorObjectNotFoundException,
	"ValidationException":       newErrorValidationException,
}
