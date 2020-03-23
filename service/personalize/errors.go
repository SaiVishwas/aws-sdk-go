// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"github.com/SaiVishwas/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeInvalidInputException for service response error code
	// "InvalidInputException".
	//
	// Provide a valid value for the field or parameter.
	ErrCodeInvalidInputException = "InvalidInputException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// The token is not valid.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The limit on the number of requests per second has been exceeded.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeResourceAlreadyExistsException for service response error code
	// "ResourceAlreadyExistsException".
	//
	// The specified resource already exists.
	ErrCodeResourceAlreadyExistsException = "ResourceAlreadyExistsException"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// The specified resource is in use.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// Could not find the specified resource.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"InvalidInputException":          newErrorInvalidInputException,
	"InvalidNextTokenException":      newErrorInvalidNextTokenException,
	"LimitExceededException":         newErrorLimitExceededException,
	"ResourceAlreadyExistsException": newErrorResourceAlreadyExistsException,
	"ResourceInUseException":         newErrorResourceInUseException,
	"ResourceNotFoundException":      newErrorResourceNotFoundException,
}
