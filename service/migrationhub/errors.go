// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhub

import (
	"github.com/SaiVishwas/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You do not have sufficient access to perform this action.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeDryRunOperation for service response error code
	// "DryRunOperation".
	//
	// Exception raised to indicate a successfully authorized action when the DryRun
	// flag is set to "true".
	ErrCodeDryRunOperation = "DryRunOperation"

	// ErrCodeHomeRegionNotSetException for service response error code
	// "HomeRegionNotSetException".
	//
	// The home region is not set. Set the home region to continue.
	ErrCodeHomeRegionNotSetException = "HomeRegionNotSetException"

	// ErrCodeInternalServerError for service response error code
	// "InternalServerError".
	//
	// Exception raised when an internal, configuration, or dependency error is
	// encountered.
	ErrCodeInternalServerError = "InternalServerError"

	// ErrCodeInvalidInputException for service response error code
	// "InvalidInputException".
	//
	// Exception raised when the provided input violates a policy constraint or
	// is entered in the wrong format or data type.
	ErrCodeInvalidInputException = "InvalidInputException"

	// ErrCodePolicyErrorException for service response error code
	// "PolicyErrorException".
	//
	// Exception raised when there are problems accessing Application Discovery
	// Service (Application Discovery Service); most likely due to a misconfigured
	// policy or the migrationhub-discovery role is missing or not configured correctly.
	ErrCodePolicyErrorException = "PolicyErrorException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// Exception raised when the request references a resource (Application Discovery
	// Service configuration, update stream, migration task, etc.) that does not
	// exist in Application Discovery Service (Application Discovery Service) or
	// in Migration Hub's repository.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// Exception raised when there is an internal, configuration, or dependency
	// error encountered.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"

	// ErrCodeUnauthorizedOperation for service response error code
	// "UnauthorizedOperation".
	//
	// Exception raised to indicate a request was not authorized when the DryRun
	// flag is set to "true".
	ErrCodeUnauthorizedOperation = "UnauthorizedOperation"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":       newErrorAccessDeniedException,
	"DryRunOperation":             newErrorDryRunOperation,
	"HomeRegionNotSetException":   newErrorHomeRegionNotSetException,
	"InternalServerError":         newErrorInternalServerError,
	"InvalidInputException":       newErrorInvalidInputException,
	"PolicyErrorException":        newErrorPolicyErrorException,
	"ResourceNotFoundException":   newErrorResourceNotFoundException,
	"ServiceUnavailableException": newErrorServiceUnavailableException,
	"UnauthorizedOperation":       newErrorUnauthorizedOperation,
}
