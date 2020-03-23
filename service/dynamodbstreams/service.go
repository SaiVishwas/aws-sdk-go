// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodbstreams

import (
	"github.com/SaiVishwas/aws-sdk-go/aws"
	"github.com/SaiVishwas/aws-sdk-go/aws/client"
	"github.com/SaiVishwas/aws-sdk-go/aws/client/metadata"
	"github.com/SaiVishwas/aws-sdk-go/aws/request"
	"github.com/SaiVishwas/aws-sdk-go/aws/signer/v4"
	"github.com/SaiVishwas/aws-sdk-go/private/protocol"
	"github.com/SaiVishwas/aws-sdk-go/private/protocol/jsonrpc"
)

// DynamoDBStreams provides the API operation methods for making requests to
// Amazon DynamoDB Streams. See this package's package overview docs
// for details on the service.
//
// DynamoDBStreams methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type DynamoDBStreams struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "streams.dynamodb" // Name of service.
	EndpointsID = ServiceName        // ID to lookup a service endpoint with.
	ServiceID   = "DynamoDB Streams" // ServiceID is a unique identifier of a specific service.
)

// New creates a new instance of the DynamoDBStreams client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     mySession := session.Must(session.NewSession())
//
//     // Create a DynamoDBStreams client from just a session.
//     svc := dynamodbstreams.New(mySession)
//
//     // Create a DynamoDBStreams client with additional configuration
//     svc := dynamodbstreams.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *DynamoDBStreams {
	c := p.ClientConfig(EndpointsID, cfgs...)
	if c.SigningNameDerived || len(c.SigningName) == 0 {
		c.SigningName = "dynamodb"
	}
	return newClient(*c.Config, c.Handlers, c.PartitionID, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, partitionID, endpoint, signingRegion, signingName string) *DynamoDBStreams {
	svc := &DynamoDBStreams{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				ServiceID:     ServiceID,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				PartitionID:   partitionID,
				Endpoint:      endpoint,
				APIVersion:    "2012-08-10",
				JSONVersion:   "1.0",
				TargetPrefix:  "DynamoDBStreams_20120810",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(
		protocol.NewUnmarshalErrorHandler(jsonrpc.NewUnmarshalTypedError(exceptionFromCode)).NamedHandler(),
	)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a DynamoDBStreams operation and runs any
// custom request initialization.
func (c *DynamoDBStreams) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
