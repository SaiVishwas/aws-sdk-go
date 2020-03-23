// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudsearchdomain

import (
	"github.com/SaiVishwas/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeDocumentServiceException for service response error code
	// "DocumentServiceException".
	//
	// Information about any problems encountered while processing an upload request.
	ErrCodeDocumentServiceException = "DocumentServiceException"

	// ErrCodeSearchException for service response error code
	// "SearchException".
	//
	// Information about any problems encountered while processing a search request.
	ErrCodeSearchException = "SearchException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"DocumentServiceException": newErrorDocumentServiceException,
	"SearchException":          newErrorSearchException,
}
