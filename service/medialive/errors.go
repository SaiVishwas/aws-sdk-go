// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"github.com/SaiVishwas/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeBadGatewayException for service response error code
	// "BadGatewayException".
	ErrCodeBadGatewayException = "BadGatewayException"

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	ErrCodeConflictException = "ConflictException"

	// ErrCodeForbiddenException for service response error code
	// "ForbiddenException".
	ErrCodeForbiddenException = "ForbiddenException"

	// ErrCodeGatewayTimeoutException for service response error code
	// "GatewayTimeoutException".
	ErrCodeGatewayTimeoutException = "GatewayTimeoutException"

	// ErrCodeInternalServerErrorException for service response error code
	// "InternalServerErrorException".
	ErrCodeInternalServerErrorException = "InternalServerErrorException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeTooManyRequestsException for service response error code
	// "TooManyRequestsException".
	ErrCodeTooManyRequestsException = "TooManyRequestsException"

	// ErrCodeUnprocessableEntityException for service response error code
	// "UnprocessableEntityException".
	ErrCodeUnprocessableEntityException = "UnprocessableEntityException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"BadGatewayException":          newErrorBadGatewayException,
	"BadRequestException":          newErrorBadRequestException,
	"ConflictException":            newErrorConflictException,
	"ForbiddenException":           newErrorForbiddenException,
	"GatewayTimeoutException":      newErrorGatewayTimeoutException,
	"InternalServerErrorException": newErrorInternalServerErrorException,
	"NotFoundException":            newErrorNotFoundException,
	"TooManyRequestsException":     newErrorTooManyRequestsException,
	"UnprocessableEntityException": newErrorUnprocessableEntityException,
}
