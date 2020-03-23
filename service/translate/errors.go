// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package translate

import (
	"github.com/SaiVishwas/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeDetectedLanguageLowConfidenceException for service response error code
	// "DetectedLanguageLowConfidenceException".
	//
	// The confidence that Amazon Comprehend accurately detected the source language
	// is low. If a low confidence level is acceptable for your application, you
	// can use the language in the exception to call Amazon Translate again. For
	// more information, see the DetectDominantLanguage (https://docs.aws.amazon.com/comprehend/latest/dg/API_DetectDominantLanguage.html)
	// operation in the Amazon Comprehend Developer Guide.
	ErrCodeDetectedLanguageLowConfidenceException = "DetectedLanguageLowConfidenceException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// An internal server error occurred. Retry your request.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeInvalidFilterException for service response error code
	// "InvalidFilterException".
	//
	// The filter specified for the operation is invalid. Specify a different filter.
	ErrCodeInvalidFilterException = "InvalidFilterException"

	// ErrCodeInvalidParameterValueException for service response error code
	// "InvalidParameterValueException".
	//
	// The value of the parameter is invalid. Review the value of the parameter
	// you are using to correct it, and then retry your operation.
	ErrCodeInvalidParameterValueException = "InvalidParameterValueException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// The request that you made is invalid. Check your request to determine why
	// it's invalid and then retry the request.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The specified limit has been exceeded. Review your request and retry it with
	// a quantity below the stated limit.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The resource you are looking for has not been found. Review the resource
	// you're looking for and see if a different resource will accomplish your needs
	// before retrying the revised request.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// The Amazon Translate service is temporarily unavailable. Please wait a bit
	// and then retry your request.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"

	// ErrCodeTextSizeLimitExceededException for service response error code
	// "TextSizeLimitExceededException".
	//
	// The size of the text you submitted exceeds the size limit. Reduce the size
	// of the text or use a smaller document and then retry your request.
	ErrCodeTextSizeLimitExceededException = "TextSizeLimitExceededException"

	// ErrCodeTooManyRequestsException for service response error code
	// "TooManyRequestsException".
	//
	// You have made too many requests within a short period of time. Wait for a
	// short time and then try your request again.
	ErrCodeTooManyRequestsException = "TooManyRequestsException"

	// ErrCodeUnsupportedLanguagePairException for service response error code
	// "UnsupportedLanguagePairException".
	//
	// Amazon Translate does not support translation from the language of the source
	// text into the requested target language. For more information, see how-to-error-msg.
	ErrCodeUnsupportedLanguagePairException = "UnsupportedLanguagePairException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"DetectedLanguageLowConfidenceException": newErrorDetectedLanguageLowConfidenceException,
	"InternalServerException":                newErrorInternalServerException,
	"InvalidFilterException":                 newErrorInvalidFilterException,
	"InvalidParameterValueException":         newErrorInvalidParameterValueException,
	"InvalidRequestException":                newErrorInvalidRequestException,
	"LimitExceededException":                 newErrorLimitExceededException,
	"ResourceNotFoundException":              newErrorResourceNotFoundException,
	"ServiceUnavailableException":            newErrorServiceUnavailableException,
	"TextSizeLimitExceededException":         newErrorTextSizeLimitExceededException,
	"TooManyRequestsException":               newErrorTooManyRequestsException,
	"UnsupportedLanguagePairException":       newErrorUnsupportedLanguagePairException,
}
