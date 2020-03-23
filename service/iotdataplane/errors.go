// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotdataplane

import (
	"github.com/SaiVishwas/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// The specified version does not match the version of the document.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalFailureException for service response error code
	// "InternalFailureException".
	//
	// An unexpected error has occurred.
	ErrCodeInternalFailureException = "InternalFailureException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// The request is not valid.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeMethodNotAllowedException for service response error code
	// "MethodNotAllowedException".
	//
	// The specified combination of HTTP verb and URI is not supported.
	ErrCodeMethodNotAllowedException = "MethodNotAllowedException"

	// ErrCodeRequestEntityTooLargeException for service response error code
	// "RequestEntityTooLargeException".
	//
	// The payload exceeds the maximum size allowed.
	ErrCodeRequestEntityTooLargeException = "RequestEntityTooLargeException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified resource does not exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// The service is temporarily unavailable.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The rate exceeds the limit.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeUnauthorizedException for service response error code
	// "UnauthorizedException".
	//
	// You are not authorized to perform this operation.
	ErrCodeUnauthorizedException = "UnauthorizedException"

	// ErrCodeUnsupportedDocumentEncodingException for service response error code
	// "UnsupportedDocumentEncodingException".
	//
	// The document encoding is not supported.
	ErrCodeUnsupportedDocumentEncodingException = "UnsupportedDocumentEncodingException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"ConflictException":                    newErrorConflictException,
	"InternalFailureException":             newErrorInternalFailureException,
	"InvalidRequestException":              newErrorInvalidRequestException,
	"MethodNotAllowedException":            newErrorMethodNotAllowedException,
	"RequestEntityTooLargeException":       newErrorRequestEntityTooLargeException,
	"ResourceNotFoundException":            newErrorResourceNotFoundException,
	"ServiceUnavailableException":          newErrorServiceUnavailableException,
	"ThrottlingException":                  newErrorThrottlingException,
	"UnauthorizedException":                newErrorUnauthorizedException,
	"UnsupportedDocumentEncodingException": newErrorUnsupportedDocumentEncodingException,
}
