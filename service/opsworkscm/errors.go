// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworkscm

import (
	"github.com/SaiVishwas/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// This occurs when the provided nextToken is not valid.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeInvalidStateException for service response error code
	// "InvalidStateException".
	//
	// The resource is in a state that does not allow you to perform a specified
	// action.
	ErrCodeInvalidStateException = "InvalidStateException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The limit of servers or backups has been reached.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeResourceAlreadyExistsException for service response error code
	// "ResourceAlreadyExistsException".
	//
	// The requested resource cannot be created because it already exists.
	ErrCodeResourceAlreadyExistsException = "ResourceAlreadyExistsException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The requested resource does not exist, or access was denied.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// One or more of the provided request parameters are not valid.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"InvalidNextTokenException":      newErrorInvalidNextTokenException,
	"InvalidStateException":          newErrorInvalidStateException,
	"LimitExceededException":         newErrorLimitExceededException,
	"ResourceAlreadyExistsException": newErrorResourceAlreadyExistsException,
	"ResourceNotFoundException":      newErrorResourceNotFoundException,
	"ValidationException":            newErrorValidationException,
}
