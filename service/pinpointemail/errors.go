// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"github.com/SaiVishwas/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccountSuspendedException for service response error code
	// "AccountSuspendedException".
	//
	// The message can't be sent because the account's ability to send email has
	// been permanently restricted.
	ErrCodeAccountSuspendedException = "AccountSuspendedException"

	// ErrCodeAlreadyExistsException for service response error code
	// "AlreadyExistsException".
	//
	// The resource specified in your request already exists.
	ErrCodeAlreadyExistsException = "AlreadyExistsException"

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	//
	// The input you provided is invalid.
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeConcurrentModificationException for service response error code
	// "ConcurrentModificationException".
	//
	// The resource is being modified by another operation or thread.
	ErrCodeConcurrentModificationException = "ConcurrentModificationException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// There are too many instances of the specified resource type.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeMailFromDomainNotVerifiedException for service response error code
	// "MailFromDomainNotVerifiedException".
	//
	// The message can't be sent because the sending domain isn't verified.
	ErrCodeMailFromDomainNotVerifiedException = "MailFromDomainNotVerifiedException"

	// ErrCodeMessageRejected for service response error code
	// "MessageRejected".
	//
	// The message can't be sent because it contains invalid content.
	ErrCodeMessageRejected = "MessageRejected"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	//
	// The resource you attempted to access doesn't exist.
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeSendingPausedException for service response error code
	// "SendingPausedException".
	//
	// The message can't be sent because the account's ability to send email is
	// currently paused.
	ErrCodeSendingPausedException = "SendingPausedException"

	// ErrCodeTooManyRequestsException for service response error code
	// "TooManyRequestsException".
	//
	// Too many requests have been made to the operation.
	ErrCodeTooManyRequestsException = "TooManyRequestsException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccountSuspendedException":          newErrorAccountSuspendedException,
	"AlreadyExistsException":             newErrorAlreadyExistsException,
	"BadRequestException":                newErrorBadRequestException,
	"ConcurrentModificationException":    newErrorConcurrentModificationException,
	"LimitExceededException":             newErrorLimitExceededException,
	"MailFromDomainNotVerifiedException": newErrorMailFromDomainNotVerifiedException,
	"MessageRejected":                    newErrorMessageRejected,
	"NotFoundException":                  newErrorNotFoundException,
	"SendingPausedException":             newErrorSendingPausedException,
	"TooManyRequestsException":           newErrorTooManyRequestsException,
}
