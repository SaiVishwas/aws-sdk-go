// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package pinpointsmsvoiceiface provides an interface to enable mocking the Amazon Pinpoint SMS and Voice Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package pinpointsmsvoiceiface

import (
	"github.com/SaiVishwas/aws-sdk-go/aws"
	"github.com/SaiVishwas/aws-sdk-go/aws/request"
	"github.com/SaiVishwas/aws-sdk-go/service/pinpointsmsvoice"
)

// PinpointSMSVoiceAPI provides an interface to enable mocking the
// pinpointsmsvoice.PinpointSMSVoice service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Pinpoint SMS and Voice Service.
//    func myFunc(svc pinpointsmsvoiceiface.PinpointSMSVoiceAPI) bool {
//        // Make svc.CreateConfigurationSet request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := pinpointsmsvoice.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockPinpointSMSVoiceClient struct {
//        pinpointsmsvoiceiface.PinpointSMSVoiceAPI
//    }
//    func (m *mockPinpointSMSVoiceClient) CreateConfigurationSet(input *pinpointsmsvoice.CreateConfigurationSetInput) (*pinpointsmsvoice.CreateConfigurationSetOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockPinpointSMSVoiceClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type PinpointSMSVoiceAPI interface {
	CreateConfigurationSet(*pinpointsmsvoice.CreateConfigurationSetInput) (*pinpointsmsvoice.CreateConfigurationSetOutput, error)
	CreateConfigurationSetWithContext(aws.Context, *pinpointsmsvoice.CreateConfigurationSetInput, ...request.Option) (*pinpointsmsvoice.CreateConfigurationSetOutput, error)
	CreateConfigurationSetRequest(*pinpointsmsvoice.CreateConfigurationSetInput) (*request.Request, *pinpointsmsvoice.CreateConfigurationSetOutput)

	CreateConfigurationSetEventDestination(*pinpointsmsvoice.CreateConfigurationSetEventDestinationInput) (*pinpointsmsvoice.CreateConfigurationSetEventDestinationOutput, error)
	CreateConfigurationSetEventDestinationWithContext(aws.Context, *pinpointsmsvoice.CreateConfigurationSetEventDestinationInput, ...request.Option) (*pinpointsmsvoice.CreateConfigurationSetEventDestinationOutput, error)
	CreateConfigurationSetEventDestinationRequest(*pinpointsmsvoice.CreateConfigurationSetEventDestinationInput) (*request.Request, *pinpointsmsvoice.CreateConfigurationSetEventDestinationOutput)

	DeleteConfigurationSet(*pinpointsmsvoice.DeleteConfigurationSetInput) (*pinpointsmsvoice.DeleteConfigurationSetOutput, error)
	DeleteConfigurationSetWithContext(aws.Context, *pinpointsmsvoice.DeleteConfigurationSetInput, ...request.Option) (*pinpointsmsvoice.DeleteConfigurationSetOutput, error)
	DeleteConfigurationSetRequest(*pinpointsmsvoice.DeleteConfigurationSetInput) (*request.Request, *pinpointsmsvoice.DeleteConfigurationSetOutput)

	DeleteConfigurationSetEventDestination(*pinpointsmsvoice.DeleteConfigurationSetEventDestinationInput) (*pinpointsmsvoice.DeleteConfigurationSetEventDestinationOutput, error)
	DeleteConfigurationSetEventDestinationWithContext(aws.Context, *pinpointsmsvoice.DeleteConfigurationSetEventDestinationInput, ...request.Option) (*pinpointsmsvoice.DeleteConfigurationSetEventDestinationOutput, error)
	DeleteConfigurationSetEventDestinationRequest(*pinpointsmsvoice.DeleteConfigurationSetEventDestinationInput) (*request.Request, *pinpointsmsvoice.DeleteConfigurationSetEventDestinationOutput)

	GetConfigurationSetEventDestinations(*pinpointsmsvoice.GetConfigurationSetEventDestinationsInput) (*pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput, error)
	GetConfigurationSetEventDestinationsWithContext(aws.Context, *pinpointsmsvoice.GetConfigurationSetEventDestinationsInput, ...request.Option) (*pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput, error)
	GetConfigurationSetEventDestinationsRequest(*pinpointsmsvoice.GetConfigurationSetEventDestinationsInput) (*request.Request, *pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput)

	ListConfigurationSets(*pinpointsmsvoice.ListConfigurationSetsInput) (*pinpointsmsvoice.ListConfigurationSetsOutput, error)
	ListConfigurationSetsWithContext(aws.Context, *pinpointsmsvoice.ListConfigurationSetsInput, ...request.Option) (*pinpointsmsvoice.ListConfigurationSetsOutput, error)
	ListConfigurationSetsRequest(*pinpointsmsvoice.ListConfigurationSetsInput) (*request.Request, *pinpointsmsvoice.ListConfigurationSetsOutput)

	SendVoiceMessage(*pinpointsmsvoice.SendVoiceMessageInput) (*pinpointsmsvoice.SendVoiceMessageOutput, error)
	SendVoiceMessageWithContext(aws.Context, *pinpointsmsvoice.SendVoiceMessageInput, ...request.Option) (*pinpointsmsvoice.SendVoiceMessageOutput, error)
	SendVoiceMessageRequest(*pinpointsmsvoice.SendVoiceMessageInput) (*request.Request, *pinpointsmsvoice.SendVoiceMessageOutput)

	UpdateConfigurationSetEventDestination(*pinpointsmsvoice.UpdateConfigurationSetEventDestinationInput) (*pinpointsmsvoice.UpdateConfigurationSetEventDestinationOutput, error)
	UpdateConfigurationSetEventDestinationWithContext(aws.Context, *pinpointsmsvoice.UpdateConfigurationSetEventDestinationInput, ...request.Option) (*pinpointsmsvoice.UpdateConfigurationSetEventDestinationOutput, error)
	UpdateConfigurationSetEventDestinationRequest(*pinpointsmsvoice.UpdateConfigurationSetEventDestinationInput) (*request.Request, *pinpointsmsvoice.UpdateConfigurationSetEventDestinationOutput)
}

var _ PinpointSMSVoiceAPI = (*pinpointsmsvoice.PinpointSMSVoice)(nil)
