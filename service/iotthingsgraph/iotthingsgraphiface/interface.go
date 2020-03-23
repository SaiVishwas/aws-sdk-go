// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package iotthingsgraphiface provides an interface to enable mocking the AWS IoT Things Graph service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package iotthingsgraphiface

import (
	"github.com/SaiVishwas/aws-sdk-go/aws"
	"github.com/SaiVishwas/aws-sdk-go/aws/request"
	"github.com/SaiVishwas/aws-sdk-go/service/iotthingsgraph"
)

// IoTThingsGraphAPI provides an interface to enable mocking the
// iotthingsgraph.IoTThingsGraph service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS IoT Things Graph.
//    func myFunc(svc iotthingsgraphiface.IoTThingsGraphAPI) bool {
//        // Make svc.AssociateEntityToThing request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := iotthingsgraph.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockIoTThingsGraphClient struct {
//        iotthingsgraphiface.IoTThingsGraphAPI
//    }
//    func (m *mockIoTThingsGraphClient) AssociateEntityToThing(input *iotthingsgraph.AssociateEntityToThingInput) (*iotthingsgraph.AssociateEntityToThingOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockIoTThingsGraphClient{}
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
type IoTThingsGraphAPI interface {
	AssociateEntityToThing(*iotthingsgraph.AssociateEntityToThingInput) (*iotthingsgraph.AssociateEntityToThingOutput, error)
	AssociateEntityToThingWithContext(aws.Context, *iotthingsgraph.AssociateEntityToThingInput, ...request.Option) (*iotthingsgraph.AssociateEntityToThingOutput, error)
	AssociateEntityToThingRequest(*iotthingsgraph.AssociateEntityToThingInput) (*request.Request, *iotthingsgraph.AssociateEntityToThingOutput)

	CreateFlowTemplate(*iotthingsgraph.CreateFlowTemplateInput) (*iotthingsgraph.CreateFlowTemplateOutput, error)
	CreateFlowTemplateWithContext(aws.Context, *iotthingsgraph.CreateFlowTemplateInput, ...request.Option) (*iotthingsgraph.CreateFlowTemplateOutput, error)
	CreateFlowTemplateRequest(*iotthingsgraph.CreateFlowTemplateInput) (*request.Request, *iotthingsgraph.CreateFlowTemplateOutput)

	CreateSystemInstance(*iotthingsgraph.CreateSystemInstanceInput) (*iotthingsgraph.CreateSystemInstanceOutput, error)
	CreateSystemInstanceWithContext(aws.Context, *iotthingsgraph.CreateSystemInstanceInput, ...request.Option) (*iotthingsgraph.CreateSystemInstanceOutput, error)
	CreateSystemInstanceRequest(*iotthingsgraph.CreateSystemInstanceInput) (*request.Request, *iotthingsgraph.CreateSystemInstanceOutput)

	CreateSystemTemplate(*iotthingsgraph.CreateSystemTemplateInput) (*iotthingsgraph.CreateSystemTemplateOutput, error)
	CreateSystemTemplateWithContext(aws.Context, *iotthingsgraph.CreateSystemTemplateInput, ...request.Option) (*iotthingsgraph.CreateSystemTemplateOutput, error)
	CreateSystemTemplateRequest(*iotthingsgraph.CreateSystemTemplateInput) (*request.Request, *iotthingsgraph.CreateSystemTemplateOutput)

	DeleteFlowTemplate(*iotthingsgraph.DeleteFlowTemplateInput) (*iotthingsgraph.DeleteFlowTemplateOutput, error)
	DeleteFlowTemplateWithContext(aws.Context, *iotthingsgraph.DeleteFlowTemplateInput, ...request.Option) (*iotthingsgraph.DeleteFlowTemplateOutput, error)
	DeleteFlowTemplateRequest(*iotthingsgraph.DeleteFlowTemplateInput) (*request.Request, *iotthingsgraph.DeleteFlowTemplateOutput)

	DeleteNamespace(*iotthingsgraph.DeleteNamespaceInput) (*iotthingsgraph.DeleteNamespaceOutput, error)
	DeleteNamespaceWithContext(aws.Context, *iotthingsgraph.DeleteNamespaceInput, ...request.Option) (*iotthingsgraph.DeleteNamespaceOutput, error)
	DeleteNamespaceRequest(*iotthingsgraph.DeleteNamespaceInput) (*request.Request, *iotthingsgraph.DeleteNamespaceOutput)

	DeleteSystemInstance(*iotthingsgraph.DeleteSystemInstanceInput) (*iotthingsgraph.DeleteSystemInstanceOutput, error)
	DeleteSystemInstanceWithContext(aws.Context, *iotthingsgraph.DeleteSystemInstanceInput, ...request.Option) (*iotthingsgraph.DeleteSystemInstanceOutput, error)
	DeleteSystemInstanceRequest(*iotthingsgraph.DeleteSystemInstanceInput) (*request.Request, *iotthingsgraph.DeleteSystemInstanceOutput)

	DeleteSystemTemplate(*iotthingsgraph.DeleteSystemTemplateInput) (*iotthingsgraph.DeleteSystemTemplateOutput, error)
	DeleteSystemTemplateWithContext(aws.Context, *iotthingsgraph.DeleteSystemTemplateInput, ...request.Option) (*iotthingsgraph.DeleteSystemTemplateOutput, error)
	DeleteSystemTemplateRequest(*iotthingsgraph.DeleteSystemTemplateInput) (*request.Request, *iotthingsgraph.DeleteSystemTemplateOutput)

	DeploySystemInstance(*iotthingsgraph.DeploySystemInstanceInput) (*iotthingsgraph.DeploySystemInstanceOutput, error)
	DeploySystemInstanceWithContext(aws.Context, *iotthingsgraph.DeploySystemInstanceInput, ...request.Option) (*iotthingsgraph.DeploySystemInstanceOutput, error)
	DeploySystemInstanceRequest(*iotthingsgraph.DeploySystemInstanceInput) (*request.Request, *iotthingsgraph.DeploySystemInstanceOutput)

	DeprecateFlowTemplate(*iotthingsgraph.DeprecateFlowTemplateInput) (*iotthingsgraph.DeprecateFlowTemplateOutput, error)
	DeprecateFlowTemplateWithContext(aws.Context, *iotthingsgraph.DeprecateFlowTemplateInput, ...request.Option) (*iotthingsgraph.DeprecateFlowTemplateOutput, error)
	DeprecateFlowTemplateRequest(*iotthingsgraph.DeprecateFlowTemplateInput) (*request.Request, *iotthingsgraph.DeprecateFlowTemplateOutput)

	DeprecateSystemTemplate(*iotthingsgraph.DeprecateSystemTemplateInput) (*iotthingsgraph.DeprecateSystemTemplateOutput, error)
	DeprecateSystemTemplateWithContext(aws.Context, *iotthingsgraph.DeprecateSystemTemplateInput, ...request.Option) (*iotthingsgraph.DeprecateSystemTemplateOutput, error)
	DeprecateSystemTemplateRequest(*iotthingsgraph.DeprecateSystemTemplateInput) (*request.Request, *iotthingsgraph.DeprecateSystemTemplateOutput)

	DescribeNamespace(*iotthingsgraph.DescribeNamespaceInput) (*iotthingsgraph.DescribeNamespaceOutput, error)
	DescribeNamespaceWithContext(aws.Context, *iotthingsgraph.DescribeNamespaceInput, ...request.Option) (*iotthingsgraph.DescribeNamespaceOutput, error)
	DescribeNamespaceRequest(*iotthingsgraph.DescribeNamespaceInput) (*request.Request, *iotthingsgraph.DescribeNamespaceOutput)

	DissociateEntityFromThing(*iotthingsgraph.DissociateEntityFromThingInput) (*iotthingsgraph.DissociateEntityFromThingOutput, error)
	DissociateEntityFromThingWithContext(aws.Context, *iotthingsgraph.DissociateEntityFromThingInput, ...request.Option) (*iotthingsgraph.DissociateEntityFromThingOutput, error)
	DissociateEntityFromThingRequest(*iotthingsgraph.DissociateEntityFromThingInput) (*request.Request, *iotthingsgraph.DissociateEntityFromThingOutput)

	GetEntities(*iotthingsgraph.GetEntitiesInput) (*iotthingsgraph.GetEntitiesOutput, error)
	GetEntitiesWithContext(aws.Context, *iotthingsgraph.GetEntitiesInput, ...request.Option) (*iotthingsgraph.GetEntitiesOutput, error)
	GetEntitiesRequest(*iotthingsgraph.GetEntitiesInput) (*request.Request, *iotthingsgraph.GetEntitiesOutput)

	GetFlowTemplate(*iotthingsgraph.GetFlowTemplateInput) (*iotthingsgraph.GetFlowTemplateOutput, error)
	GetFlowTemplateWithContext(aws.Context, *iotthingsgraph.GetFlowTemplateInput, ...request.Option) (*iotthingsgraph.GetFlowTemplateOutput, error)
	GetFlowTemplateRequest(*iotthingsgraph.GetFlowTemplateInput) (*request.Request, *iotthingsgraph.GetFlowTemplateOutput)

	GetFlowTemplateRevisions(*iotthingsgraph.GetFlowTemplateRevisionsInput) (*iotthingsgraph.GetFlowTemplateRevisionsOutput, error)
	GetFlowTemplateRevisionsWithContext(aws.Context, *iotthingsgraph.GetFlowTemplateRevisionsInput, ...request.Option) (*iotthingsgraph.GetFlowTemplateRevisionsOutput, error)
	GetFlowTemplateRevisionsRequest(*iotthingsgraph.GetFlowTemplateRevisionsInput) (*request.Request, *iotthingsgraph.GetFlowTemplateRevisionsOutput)

	GetFlowTemplateRevisionsPages(*iotthingsgraph.GetFlowTemplateRevisionsInput, func(*iotthingsgraph.GetFlowTemplateRevisionsOutput, bool) bool) error
	GetFlowTemplateRevisionsPagesWithContext(aws.Context, *iotthingsgraph.GetFlowTemplateRevisionsInput, func(*iotthingsgraph.GetFlowTemplateRevisionsOutput, bool) bool, ...request.Option) error

	GetNamespaceDeletionStatus(*iotthingsgraph.GetNamespaceDeletionStatusInput) (*iotthingsgraph.GetNamespaceDeletionStatusOutput, error)
	GetNamespaceDeletionStatusWithContext(aws.Context, *iotthingsgraph.GetNamespaceDeletionStatusInput, ...request.Option) (*iotthingsgraph.GetNamespaceDeletionStatusOutput, error)
	GetNamespaceDeletionStatusRequest(*iotthingsgraph.GetNamespaceDeletionStatusInput) (*request.Request, *iotthingsgraph.GetNamespaceDeletionStatusOutput)

	GetSystemInstance(*iotthingsgraph.GetSystemInstanceInput) (*iotthingsgraph.GetSystemInstanceOutput, error)
	GetSystemInstanceWithContext(aws.Context, *iotthingsgraph.GetSystemInstanceInput, ...request.Option) (*iotthingsgraph.GetSystemInstanceOutput, error)
	GetSystemInstanceRequest(*iotthingsgraph.GetSystemInstanceInput) (*request.Request, *iotthingsgraph.GetSystemInstanceOutput)

	GetSystemTemplate(*iotthingsgraph.GetSystemTemplateInput) (*iotthingsgraph.GetSystemTemplateOutput, error)
	GetSystemTemplateWithContext(aws.Context, *iotthingsgraph.GetSystemTemplateInput, ...request.Option) (*iotthingsgraph.GetSystemTemplateOutput, error)
	GetSystemTemplateRequest(*iotthingsgraph.GetSystemTemplateInput) (*request.Request, *iotthingsgraph.GetSystemTemplateOutput)

	GetSystemTemplateRevisions(*iotthingsgraph.GetSystemTemplateRevisionsInput) (*iotthingsgraph.GetSystemTemplateRevisionsOutput, error)
	GetSystemTemplateRevisionsWithContext(aws.Context, *iotthingsgraph.GetSystemTemplateRevisionsInput, ...request.Option) (*iotthingsgraph.GetSystemTemplateRevisionsOutput, error)
	GetSystemTemplateRevisionsRequest(*iotthingsgraph.GetSystemTemplateRevisionsInput) (*request.Request, *iotthingsgraph.GetSystemTemplateRevisionsOutput)

	GetSystemTemplateRevisionsPages(*iotthingsgraph.GetSystemTemplateRevisionsInput, func(*iotthingsgraph.GetSystemTemplateRevisionsOutput, bool) bool) error
	GetSystemTemplateRevisionsPagesWithContext(aws.Context, *iotthingsgraph.GetSystemTemplateRevisionsInput, func(*iotthingsgraph.GetSystemTemplateRevisionsOutput, bool) bool, ...request.Option) error

	GetUploadStatus(*iotthingsgraph.GetUploadStatusInput) (*iotthingsgraph.GetUploadStatusOutput, error)
	GetUploadStatusWithContext(aws.Context, *iotthingsgraph.GetUploadStatusInput, ...request.Option) (*iotthingsgraph.GetUploadStatusOutput, error)
	GetUploadStatusRequest(*iotthingsgraph.GetUploadStatusInput) (*request.Request, *iotthingsgraph.GetUploadStatusOutput)

	ListFlowExecutionMessages(*iotthingsgraph.ListFlowExecutionMessagesInput) (*iotthingsgraph.ListFlowExecutionMessagesOutput, error)
	ListFlowExecutionMessagesWithContext(aws.Context, *iotthingsgraph.ListFlowExecutionMessagesInput, ...request.Option) (*iotthingsgraph.ListFlowExecutionMessagesOutput, error)
	ListFlowExecutionMessagesRequest(*iotthingsgraph.ListFlowExecutionMessagesInput) (*request.Request, *iotthingsgraph.ListFlowExecutionMessagesOutput)

	ListFlowExecutionMessagesPages(*iotthingsgraph.ListFlowExecutionMessagesInput, func(*iotthingsgraph.ListFlowExecutionMessagesOutput, bool) bool) error
	ListFlowExecutionMessagesPagesWithContext(aws.Context, *iotthingsgraph.ListFlowExecutionMessagesInput, func(*iotthingsgraph.ListFlowExecutionMessagesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*iotthingsgraph.ListTagsForResourceInput) (*iotthingsgraph.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *iotthingsgraph.ListTagsForResourceInput, ...request.Option) (*iotthingsgraph.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*iotthingsgraph.ListTagsForResourceInput) (*request.Request, *iotthingsgraph.ListTagsForResourceOutput)

	ListTagsForResourcePages(*iotthingsgraph.ListTagsForResourceInput, func(*iotthingsgraph.ListTagsForResourceOutput, bool) bool) error
	ListTagsForResourcePagesWithContext(aws.Context, *iotthingsgraph.ListTagsForResourceInput, func(*iotthingsgraph.ListTagsForResourceOutput, bool) bool, ...request.Option) error

	SearchEntities(*iotthingsgraph.SearchEntitiesInput) (*iotthingsgraph.SearchEntitiesOutput, error)
	SearchEntitiesWithContext(aws.Context, *iotthingsgraph.SearchEntitiesInput, ...request.Option) (*iotthingsgraph.SearchEntitiesOutput, error)
	SearchEntitiesRequest(*iotthingsgraph.SearchEntitiesInput) (*request.Request, *iotthingsgraph.SearchEntitiesOutput)

	SearchEntitiesPages(*iotthingsgraph.SearchEntitiesInput, func(*iotthingsgraph.SearchEntitiesOutput, bool) bool) error
	SearchEntitiesPagesWithContext(aws.Context, *iotthingsgraph.SearchEntitiesInput, func(*iotthingsgraph.SearchEntitiesOutput, bool) bool, ...request.Option) error

	SearchFlowExecutions(*iotthingsgraph.SearchFlowExecutionsInput) (*iotthingsgraph.SearchFlowExecutionsOutput, error)
	SearchFlowExecutionsWithContext(aws.Context, *iotthingsgraph.SearchFlowExecutionsInput, ...request.Option) (*iotthingsgraph.SearchFlowExecutionsOutput, error)
	SearchFlowExecutionsRequest(*iotthingsgraph.SearchFlowExecutionsInput) (*request.Request, *iotthingsgraph.SearchFlowExecutionsOutput)

	SearchFlowExecutionsPages(*iotthingsgraph.SearchFlowExecutionsInput, func(*iotthingsgraph.SearchFlowExecutionsOutput, bool) bool) error
	SearchFlowExecutionsPagesWithContext(aws.Context, *iotthingsgraph.SearchFlowExecutionsInput, func(*iotthingsgraph.SearchFlowExecutionsOutput, bool) bool, ...request.Option) error

	SearchFlowTemplates(*iotthingsgraph.SearchFlowTemplatesInput) (*iotthingsgraph.SearchFlowTemplatesOutput, error)
	SearchFlowTemplatesWithContext(aws.Context, *iotthingsgraph.SearchFlowTemplatesInput, ...request.Option) (*iotthingsgraph.SearchFlowTemplatesOutput, error)
	SearchFlowTemplatesRequest(*iotthingsgraph.SearchFlowTemplatesInput) (*request.Request, *iotthingsgraph.SearchFlowTemplatesOutput)

	SearchFlowTemplatesPages(*iotthingsgraph.SearchFlowTemplatesInput, func(*iotthingsgraph.SearchFlowTemplatesOutput, bool) bool) error
	SearchFlowTemplatesPagesWithContext(aws.Context, *iotthingsgraph.SearchFlowTemplatesInput, func(*iotthingsgraph.SearchFlowTemplatesOutput, bool) bool, ...request.Option) error

	SearchSystemInstances(*iotthingsgraph.SearchSystemInstancesInput) (*iotthingsgraph.SearchSystemInstancesOutput, error)
	SearchSystemInstancesWithContext(aws.Context, *iotthingsgraph.SearchSystemInstancesInput, ...request.Option) (*iotthingsgraph.SearchSystemInstancesOutput, error)
	SearchSystemInstancesRequest(*iotthingsgraph.SearchSystemInstancesInput) (*request.Request, *iotthingsgraph.SearchSystemInstancesOutput)

	SearchSystemInstancesPages(*iotthingsgraph.SearchSystemInstancesInput, func(*iotthingsgraph.SearchSystemInstancesOutput, bool) bool) error
	SearchSystemInstancesPagesWithContext(aws.Context, *iotthingsgraph.SearchSystemInstancesInput, func(*iotthingsgraph.SearchSystemInstancesOutput, bool) bool, ...request.Option) error

	SearchSystemTemplates(*iotthingsgraph.SearchSystemTemplatesInput) (*iotthingsgraph.SearchSystemTemplatesOutput, error)
	SearchSystemTemplatesWithContext(aws.Context, *iotthingsgraph.SearchSystemTemplatesInput, ...request.Option) (*iotthingsgraph.SearchSystemTemplatesOutput, error)
	SearchSystemTemplatesRequest(*iotthingsgraph.SearchSystemTemplatesInput) (*request.Request, *iotthingsgraph.SearchSystemTemplatesOutput)

	SearchSystemTemplatesPages(*iotthingsgraph.SearchSystemTemplatesInput, func(*iotthingsgraph.SearchSystemTemplatesOutput, bool) bool) error
	SearchSystemTemplatesPagesWithContext(aws.Context, *iotthingsgraph.SearchSystemTemplatesInput, func(*iotthingsgraph.SearchSystemTemplatesOutput, bool) bool, ...request.Option) error

	SearchThings(*iotthingsgraph.SearchThingsInput) (*iotthingsgraph.SearchThingsOutput, error)
	SearchThingsWithContext(aws.Context, *iotthingsgraph.SearchThingsInput, ...request.Option) (*iotthingsgraph.SearchThingsOutput, error)
	SearchThingsRequest(*iotthingsgraph.SearchThingsInput) (*request.Request, *iotthingsgraph.SearchThingsOutput)

	SearchThingsPages(*iotthingsgraph.SearchThingsInput, func(*iotthingsgraph.SearchThingsOutput, bool) bool) error
	SearchThingsPagesWithContext(aws.Context, *iotthingsgraph.SearchThingsInput, func(*iotthingsgraph.SearchThingsOutput, bool) bool, ...request.Option) error

	TagResource(*iotthingsgraph.TagResourceInput) (*iotthingsgraph.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *iotthingsgraph.TagResourceInput, ...request.Option) (*iotthingsgraph.TagResourceOutput, error)
	TagResourceRequest(*iotthingsgraph.TagResourceInput) (*request.Request, *iotthingsgraph.TagResourceOutput)

	UndeploySystemInstance(*iotthingsgraph.UndeploySystemInstanceInput) (*iotthingsgraph.UndeploySystemInstanceOutput, error)
	UndeploySystemInstanceWithContext(aws.Context, *iotthingsgraph.UndeploySystemInstanceInput, ...request.Option) (*iotthingsgraph.UndeploySystemInstanceOutput, error)
	UndeploySystemInstanceRequest(*iotthingsgraph.UndeploySystemInstanceInput) (*request.Request, *iotthingsgraph.UndeploySystemInstanceOutput)

	UntagResource(*iotthingsgraph.UntagResourceInput) (*iotthingsgraph.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *iotthingsgraph.UntagResourceInput, ...request.Option) (*iotthingsgraph.UntagResourceOutput, error)
	UntagResourceRequest(*iotthingsgraph.UntagResourceInput) (*request.Request, *iotthingsgraph.UntagResourceOutput)

	UpdateFlowTemplate(*iotthingsgraph.UpdateFlowTemplateInput) (*iotthingsgraph.UpdateFlowTemplateOutput, error)
	UpdateFlowTemplateWithContext(aws.Context, *iotthingsgraph.UpdateFlowTemplateInput, ...request.Option) (*iotthingsgraph.UpdateFlowTemplateOutput, error)
	UpdateFlowTemplateRequest(*iotthingsgraph.UpdateFlowTemplateInput) (*request.Request, *iotthingsgraph.UpdateFlowTemplateOutput)

	UpdateSystemTemplate(*iotthingsgraph.UpdateSystemTemplateInput) (*iotthingsgraph.UpdateSystemTemplateOutput, error)
	UpdateSystemTemplateWithContext(aws.Context, *iotthingsgraph.UpdateSystemTemplateInput, ...request.Option) (*iotthingsgraph.UpdateSystemTemplateOutput, error)
	UpdateSystemTemplateRequest(*iotthingsgraph.UpdateSystemTemplateInput) (*request.Request, *iotthingsgraph.UpdateSystemTemplateOutput)

	UploadEntityDefinitions(*iotthingsgraph.UploadEntityDefinitionsInput) (*iotthingsgraph.UploadEntityDefinitionsOutput, error)
	UploadEntityDefinitionsWithContext(aws.Context, *iotthingsgraph.UploadEntityDefinitionsInput, ...request.Option) (*iotthingsgraph.UploadEntityDefinitionsOutput, error)
	UploadEntityDefinitionsRequest(*iotthingsgraph.UploadEntityDefinitionsInput) (*request.Request, *iotthingsgraph.UploadEntityDefinitionsOutput)
}

var _ IoTThingsGraphAPI = (*iotthingsgraph.IoTThingsGraph)(nil)
