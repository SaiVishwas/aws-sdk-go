// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mobile

import (
	"github.com/SaiVishwas/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccountActionRequiredException for service response error code
	// "AccountActionRequiredException".
	//
	// Account Action is required in order to continue the request.
	ErrCodeAccountActionRequiredException = "AccountActionRequiredException"

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	//
	// The request cannot be processed because some parameter is not valid or the
	// project state prevents the operation from being performed.
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeInternalFailureException for service response error code
	// "InternalFailureException".
	//
	// The service has encountered an unexpected error condition which prevents
	// it from servicing the request.
	ErrCodeInternalFailureException = "InternalFailureException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// There are too many AWS Mobile Hub projects in the account or the account
	// has exceeded the maximum number of resources in some AWS service. You should
	// create another sub-account using AWS Organizations or remove some resources
	// and retry your request.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	//
	// No entity can be found with the specified identifier.
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// The service is temporarily unavailable. The request should be retried after
	// some time delay.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"

	// ErrCodeTooManyRequestsException for service response error code
	// "TooManyRequestsException".
	//
	// Too many requests have been received for this AWS account in too short a
	// time. The request should be retried after some time delay.
	ErrCodeTooManyRequestsException = "TooManyRequestsException"

	// ErrCodeUnauthorizedException for service response error code
	// "UnauthorizedException".
	//
	// Credentials of the caller are insufficient to authorize the request.
	ErrCodeUnauthorizedException = "UnauthorizedException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccountActionRequiredException": newErrorAccountActionRequiredException,
	"BadRequestException":            newErrorBadRequestException,
	"InternalFailureException":       newErrorInternalFailureException,
	"LimitExceededException":         newErrorLimitExceededException,
	"NotFoundException":              newErrorNotFoundException,
	"ServiceUnavailableException":    newErrorServiceUnavailableException,
	"TooManyRequestsException":       newErrorTooManyRequestsException,
	"UnauthorizedException":          newErrorUnauthorizedException,
}
