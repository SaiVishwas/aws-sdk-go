// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestarconnections

import (
	"github.com/SaiVishwas/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// Exceeded the maximum limit for connections.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// Resource not found. Verify the connection resource ARN and try again.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"LimitExceededException":    newErrorLimitExceededException,
	"ResourceNotFoundException": newErrorResourceNotFoundException,
}
