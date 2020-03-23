// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package kinesisanalyticsv2iface provides an interface to enable mocking the Amazon Kinesis Analytics service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package kinesisanalyticsv2iface

import (
	"github.com/SaiVishwas/aws-sdk-go/aws"
	"github.com/SaiVishwas/aws-sdk-go/aws/request"
	"github.com/SaiVishwas/aws-sdk-go/service/kinesisanalyticsv2"
)

// KinesisAnalyticsV2API provides an interface to enable mocking the
// kinesisanalyticsv2.KinesisAnalyticsV2 service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Kinesis Analytics.
//    func myFunc(svc kinesisanalyticsv2iface.KinesisAnalyticsV2API) bool {
//        // Make svc.AddApplicationCloudWatchLoggingOption request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := kinesisanalyticsv2.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockKinesisAnalyticsV2Client struct {
//        kinesisanalyticsv2iface.KinesisAnalyticsV2API
//    }
//    func (m *mockKinesisAnalyticsV2Client) AddApplicationCloudWatchLoggingOption(input *kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionInput) (*kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockKinesisAnalyticsV2Client{}
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
type KinesisAnalyticsV2API interface {
	AddApplicationCloudWatchLoggingOption(*kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionInput) (*kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionOutput, error)
	AddApplicationCloudWatchLoggingOptionWithContext(aws.Context, *kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionInput, ...request.Option) (*kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionOutput, error)
	AddApplicationCloudWatchLoggingOptionRequest(*kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionInput) (*request.Request, *kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionOutput)

	AddApplicationInput(*kinesisanalyticsv2.AddApplicationInputInput) (*kinesisanalyticsv2.AddApplicationInputOutput, error)
	AddApplicationInputWithContext(aws.Context, *kinesisanalyticsv2.AddApplicationInputInput, ...request.Option) (*kinesisanalyticsv2.AddApplicationInputOutput, error)
	AddApplicationInputRequest(*kinesisanalyticsv2.AddApplicationInputInput) (*request.Request, *kinesisanalyticsv2.AddApplicationInputOutput)

	AddApplicationInputProcessingConfiguration(*kinesisanalyticsv2.AddApplicationInputProcessingConfigurationInput) (*kinesisanalyticsv2.AddApplicationInputProcessingConfigurationOutput, error)
	AddApplicationInputProcessingConfigurationWithContext(aws.Context, *kinesisanalyticsv2.AddApplicationInputProcessingConfigurationInput, ...request.Option) (*kinesisanalyticsv2.AddApplicationInputProcessingConfigurationOutput, error)
	AddApplicationInputProcessingConfigurationRequest(*kinesisanalyticsv2.AddApplicationInputProcessingConfigurationInput) (*request.Request, *kinesisanalyticsv2.AddApplicationInputProcessingConfigurationOutput)

	AddApplicationOutput(*kinesisanalyticsv2.AddApplicationOutputInput) (*kinesisanalyticsv2.AddApplicationOutputOutput, error)
	AddApplicationOutputWithContext(aws.Context, *kinesisanalyticsv2.AddApplicationOutputInput, ...request.Option) (*kinesisanalyticsv2.AddApplicationOutputOutput, error)
	AddApplicationOutputRequest(*kinesisanalyticsv2.AddApplicationOutputInput) (*request.Request, *kinesisanalyticsv2.AddApplicationOutputOutput)

	AddApplicationReferenceDataSource(*kinesisanalyticsv2.AddApplicationReferenceDataSourceInput) (*kinesisanalyticsv2.AddApplicationReferenceDataSourceOutput, error)
	AddApplicationReferenceDataSourceWithContext(aws.Context, *kinesisanalyticsv2.AddApplicationReferenceDataSourceInput, ...request.Option) (*kinesisanalyticsv2.AddApplicationReferenceDataSourceOutput, error)
	AddApplicationReferenceDataSourceRequest(*kinesisanalyticsv2.AddApplicationReferenceDataSourceInput) (*request.Request, *kinesisanalyticsv2.AddApplicationReferenceDataSourceOutput)

	AddApplicationVpcConfiguration(*kinesisanalyticsv2.AddApplicationVpcConfigurationInput) (*kinesisanalyticsv2.AddApplicationVpcConfigurationOutput, error)
	AddApplicationVpcConfigurationWithContext(aws.Context, *kinesisanalyticsv2.AddApplicationVpcConfigurationInput, ...request.Option) (*kinesisanalyticsv2.AddApplicationVpcConfigurationOutput, error)
	AddApplicationVpcConfigurationRequest(*kinesisanalyticsv2.AddApplicationVpcConfigurationInput) (*request.Request, *kinesisanalyticsv2.AddApplicationVpcConfigurationOutput)

	CreateApplication(*kinesisanalyticsv2.CreateApplicationInput) (*kinesisanalyticsv2.CreateApplicationOutput, error)
	CreateApplicationWithContext(aws.Context, *kinesisanalyticsv2.CreateApplicationInput, ...request.Option) (*kinesisanalyticsv2.CreateApplicationOutput, error)
	CreateApplicationRequest(*kinesisanalyticsv2.CreateApplicationInput) (*request.Request, *kinesisanalyticsv2.CreateApplicationOutput)

	CreateApplicationSnapshot(*kinesisanalyticsv2.CreateApplicationSnapshotInput) (*kinesisanalyticsv2.CreateApplicationSnapshotOutput, error)
	CreateApplicationSnapshotWithContext(aws.Context, *kinesisanalyticsv2.CreateApplicationSnapshotInput, ...request.Option) (*kinesisanalyticsv2.CreateApplicationSnapshotOutput, error)
	CreateApplicationSnapshotRequest(*kinesisanalyticsv2.CreateApplicationSnapshotInput) (*request.Request, *kinesisanalyticsv2.CreateApplicationSnapshotOutput)

	DeleteApplication(*kinesisanalyticsv2.DeleteApplicationInput) (*kinesisanalyticsv2.DeleteApplicationOutput, error)
	DeleteApplicationWithContext(aws.Context, *kinesisanalyticsv2.DeleteApplicationInput, ...request.Option) (*kinesisanalyticsv2.DeleteApplicationOutput, error)
	DeleteApplicationRequest(*kinesisanalyticsv2.DeleteApplicationInput) (*request.Request, *kinesisanalyticsv2.DeleteApplicationOutput)

	DeleteApplicationCloudWatchLoggingOption(*kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionInput) (*kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionOutput, error)
	DeleteApplicationCloudWatchLoggingOptionWithContext(aws.Context, *kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionInput, ...request.Option) (*kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionOutput, error)
	DeleteApplicationCloudWatchLoggingOptionRequest(*kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionInput) (*request.Request, *kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionOutput)

	DeleteApplicationInputProcessingConfiguration(*kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationInput) (*kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationOutput, error)
	DeleteApplicationInputProcessingConfigurationWithContext(aws.Context, *kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationInput, ...request.Option) (*kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationOutput, error)
	DeleteApplicationInputProcessingConfigurationRequest(*kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationInput) (*request.Request, *kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationOutput)

	DeleteApplicationOutput(*kinesisanalyticsv2.DeleteApplicationOutputInput) (*kinesisanalyticsv2.DeleteApplicationOutputOutput, error)
	DeleteApplicationOutputWithContext(aws.Context, *kinesisanalyticsv2.DeleteApplicationOutputInput, ...request.Option) (*kinesisanalyticsv2.DeleteApplicationOutputOutput, error)
	DeleteApplicationOutputRequest(*kinesisanalyticsv2.DeleteApplicationOutputInput) (*request.Request, *kinesisanalyticsv2.DeleteApplicationOutputOutput)

	DeleteApplicationReferenceDataSource(*kinesisanalyticsv2.DeleteApplicationReferenceDataSourceInput) (*kinesisanalyticsv2.DeleteApplicationReferenceDataSourceOutput, error)
	DeleteApplicationReferenceDataSourceWithContext(aws.Context, *kinesisanalyticsv2.DeleteApplicationReferenceDataSourceInput, ...request.Option) (*kinesisanalyticsv2.DeleteApplicationReferenceDataSourceOutput, error)
	DeleteApplicationReferenceDataSourceRequest(*kinesisanalyticsv2.DeleteApplicationReferenceDataSourceInput) (*request.Request, *kinesisanalyticsv2.DeleteApplicationReferenceDataSourceOutput)

	DeleteApplicationSnapshot(*kinesisanalyticsv2.DeleteApplicationSnapshotInput) (*kinesisanalyticsv2.DeleteApplicationSnapshotOutput, error)
	DeleteApplicationSnapshotWithContext(aws.Context, *kinesisanalyticsv2.DeleteApplicationSnapshotInput, ...request.Option) (*kinesisanalyticsv2.DeleteApplicationSnapshotOutput, error)
	DeleteApplicationSnapshotRequest(*kinesisanalyticsv2.DeleteApplicationSnapshotInput) (*request.Request, *kinesisanalyticsv2.DeleteApplicationSnapshotOutput)

	DeleteApplicationVpcConfiguration(*kinesisanalyticsv2.DeleteApplicationVpcConfigurationInput) (*kinesisanalyticsv2.DeleteApplicationVpcConfigurationOutput, error)
	DeleteApplicationVpcConfigurationWithContext(aws.Context, *kinesisanalyticsv2.DeleteApplicationVpcConfigurationInput, ...request.Option) (*kinesisanalyticsv2.DeleteApplicationVpcConfigurationOutput, error)
	DeleteApplicationVpcConfigurationRequest(*kinesisanalyticsv2.DeleteApplicationVpcConfigurationInput) (*request.Request, *kinesisanalyticsv2.DeleteApplicationVpcConfigurationOutput)

	DescribeApplication(*kinesisanalyticsv2.DescribeApplicationInput) (*kinesisanalyticsv2.DescribeApplicationOutput, error)
	DescribeApplicationWithContext(aws.Context, *kinesisanalyticsv2.DescribeApplicationInput, ...request.Option) (*kinesisanalyticsv2.DescribeApplicationOutput, error)
	DescribeApplicationRequest(*kinesisanalyticsv2.DescribeApplicationInput) (*request.Request, *kinesisanalyticsv2.DescribeApplicationOutput)

	DescribeApplicationSnapshot(*kinesisanalyticsv2.DescribeApplicationSnapshotInput) (*kinesisanalyticsv2.DescribeApplicationSnapshotOutput, error)
	DescribeApplicationSnapshotWithContext(aws.Context, *kinesisanalyticsv2.DescribeApplicationSnapshotInput, ...request.Option) (*kinesisanalyticsv2.DescribeApplicationSnapshotOutput, error)
	DescribeApplicationSnapshotRequest(*kinesisanalyticsv2.DescribeApplicationSnapshotInput) (*request.Request, *kinesisanalyticsv2.DescribeApplicationSnapshotOutput)

	DiscoverInputSchema(*kinesisanalyticsv2.DiscoverInputSchemaInput) (*kinesisanalyticsv2.DiscoverInputSchemaOutput, error)
	DiscoverInputSchemaWithContext(aws.Context, *kinesisanalyticsv2.DiscoverInputSchemaInput, ...request.Option) (*kinesisanalyticsv2.DiscoverInputSchemaOutput, error)
	DiscoverInputSchemaRequest(*kinesisanalyticsv2.DiscoverInputSchemaInput) (*request.Request, *kinesisanalyticsv2.DiscoverInputSchemaOutput)

	ListApplicationSnapshots(*kinesisanalyticsv2.ListApplicationSnapshotsInput) (*kinesisanalyticsv2.ListApplicationSnapshotsOutput, error)
	ListApplicationSnapshotsWithContext(aws.Context, *kinesisanalyticsv2.ListApplicationSnapshotsInput, ...request.Option) (*kinesisanalyticsv2.ListApplicationSnapshotsOutput, error)
	ListApplicationSnapshotsRequest(*kinesisanalyticsv2.ListApplicationSnapshotsInput) (*request.Request, *kinesisanalyticsv2.ListApplicationSnapshotsOutput)

	ListApplications(*kinesisanalyticsv2.ListApplicationsInput) (*kinesisanalyticsv2.ListApplicationsOutput, error)
	ListApplicationsWithContext(aws.Context, *kinesisanalyticsv2.ListApplicationsInput, ...request.Option) (*kinesisanalyticsv2.ListApplicationsOutput, error)
	ListApplicationsRequest(*kinesisanalyticsv2.ListApplicationsInput) (*request.Request, *kinesisanalyticsv2.ListApplicationsOutput)

	ListTagsForResource(*kinesisanalyticsv2.ListTagsForResourceInput) (*kinesisanalyticsv2.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *kinesisanalyticsv2.ListTagsForResourceInput, ...request.Option) (*kinesisanalyticsv2.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*kinesisanalyticsv2.ListTagsForResourceInput) (*request.Request, *kinesisanalyticsv2.ListTagsForResourceOutput)

	StartApplication(*kinesisanalyticsv2.StartApplicationInput) (*kinesisanalyticsv2.StartApplicationOutput, error)
	StartApplicationWithContext(aws.Context, *kinesisanalyticsv2.StartApplicationInput, ...request.Option) (*kinesisanalyticsv2.StartApplicationOutput, error)
	StartApplicationRequest(*kinesisanalyticsv2.StartApplicationInput) (*request.Request, *kinesisanalyticsv2.StartApplicationOutput)

	StopApplication(*kinesisanalyticsv2.StopApplicationInput) (*kinesisanalyticsv2.StopApplicationOutput, error)
	StopApplicationWithContext(aws.Context, *kinesisanalyticsv2.StopApplicationInput, ...request.Option) (*kinesisanalyticsv2.StopApplicationOutput, error)
	StopApplicationRequest(*kinesisanalyticsv2.StopApplicationInput) (*request.Request, *kinesisanalyticsv2.StopApplicationOutput)

	TagResource(*kinesisanalyticsv2.TagResourceInput) (*kinesisanalyticsv2.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *kinesisanalyticsv2.TagResourceInput, ...request.Option) (*kinesisanalyticsv2.TagResourceOutput, error)
	TagResourceRequest(*kinesisanalyticsv2.TagResourceInput) (*request.Request, *kinesisanalyticsv2.TagResourceOutput)

	UntagResource(*kinesisanalyticsv2.UntagResourceInput) (*kinesisanalyticsv2.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *kinesisanalyticsv2.UntagResourceInput, ...request.Option) (*kinesisanalyticsv2.UntagResourceOutput, error)
	UntagResourceRequest(*kinesisanalyticsv2.UntagResourceInput) (*request.Request, *kinesisanalyticsv2.UntagResourceOutput)

	UpdateApplication(*kinesisanalyticsv2.UpdateApplicationInput) (*kinesisanalyticsv2.UpdateApplicationOutput, error)
	UpdateApplicationWithContext(aws.Context, *kinesisanalyticsv2.UpdateApplicationInput, ...request.Option) (*kinesisanalyticsv2.UpdateApplicationOutput, error)
	UpdateApplicationRequest(*kinesisanalyticsv2.UpdateApplicationInput) (*request.Request, *kinesisanalyticsv2.UpdateApplicationOutput)
}

var _ KinesisAnalyticsV2API = (*kinesisanalyticsv2.KinesisAnalyticsV2)(nil)
