// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package mediapackagevodiface provides an interface to enable mocking the AWS Elemental MediaPackage VOD service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package mediapackagevodiface

import (
	"github.com/SaiVishwas/aws-sdk-go/aws"
	"github.com/SaiVishwas/aws-sdk-go/aws/request"
	"github.com/SaiVishwas/aws-sdk-go/service/mediapackagevod"
)

// MediaPackageVodAPI provides an interface to enable mocking the
// mediapackagevod.MediaPackageVod service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Elemental MediaPackage VOD.
//    func myFunc(svc mediapackagevodiface.MediaPackageVodAPI) bool {
//        // Make svc.CreateAsset request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := mediapackagevod.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockMediaPackageVodClient struct {
//        mediapackagevodiface.MediaPackageVodAPI
//    }
//    func (m *mockMediaPackageVodClient) CreateAsset(input *mediapackagevod.CreateAssetInput) (*mediapackagevod.CreateAssetOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockMediaPackageVodClient{}
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
type MediaPackageVodAPI interface {
	CreateAsset(*mediapackagevod.CreateAssetInput) (*mediapackagevod.CreateAssetOutput, error)
	CreateAssetWithContext(aws.Context, *mediapackagevod.CreateAssetInput, ...request.Option) (*mediapackagevod.CreateAssetOutput, error)
	CreateAssetRequest(*mediapackagevod.CreateAssetInput) (*request.Request, *mediapackagevod.CreateAssetOutput)

	CreatePackagingConfiguration(*mediapackagevod.CreatePackagingConfigurationInput) (*mediapackagevod.CreatePackagingConfigurationOutput, error)
	CreatePackagingConfigurationWithContext(aws.Context, *mediapackagevod.CreatePackagingConfigurationInput, ...request.Option) (*mediapackagevod.CreatePackagingConfigurationOutput, error)
	CreatePackagingConfigurationRequest(*mediapackagevod.CreatePackagingConfigurationInput) (*request.Request, *mediapackagevod.CreatePackagingConfigurationOutput)

	CreatePackagingGroup(*mediapackagevod.CreatePackagingGroupInput) (*mediapackagevod.CreatePackagingGroupOutput, error)
	CreatePackagingGroupWithContext(aws.Context, *mediapackagevod.CreatePackagingGroupInput, ...request.Option) (*mediapackagevod.CreatePackagingGroupOutput, error)
	CreatePackagingGroupRequest(*mediapackagevod.CreatePackagingGroupInput) (*request.Request, *mediapackagevod.CreatePackagingGroupOutput)

	DeleteAsset(*mediapackagevod.DeleteAssetInput) (*mediapackagevod.DeleteAssetOutput, error)
	DeleteAssetWithContext(aws.Context, *mediapackagevod.DeleteAssetInput, ...request.Option) (*mediapackagevod.DeleteAssetOutput, error)
	DeleteAssetRequest(*mediapackagevod.DeleteAssetInput) (*request.Request, *mediapackagevod.DeleteAssetOutput)

	DeletePackagingConfiguration(*mediapackagevod.DeletePackagingConfigurationInput) (*mediapackagevod.DeletePackagingConfigurationOutput, error)
	DeletePackagingConfigurationWithContext(aws.Context, *mediapackagevod.DeletePackagingConfigurationInput, ...request.Option) (*mediapackagevod.DeletePackagingConfigurationOutput, error)
	DeletePackagingConfigurationRequest(*mediapackagevod.DeletePackagingConfigurationInput) (*request.Request, *mediapackagevod.DeletePackagingConfigurationOutput)

	DeletePackagingGroup(*mediapackagevod.DeletePackagingGroupInput) (*mediapackagevod.DeletePackagingGroupOutput, error)
	DeletePackagingGroupWithContext(aws.Context, *mediapackagevod.DeletePackagingGroupInput, ...request.Option) (*mediapackagevod.DeletePackagingGroupOutput, error)
	DeletePackagingGroupRequest(*mediapackagevod.DeletePackagingGroupInput) (*request.Request, *mediapackagevod.DeletePackagingGroupOutput)

	DescribeAsset(*mediapackagevod.DescribeAssetInput) (*mediapackagevod.DescribeAssetOutput, error)
	DescribeAssetWithContext(aws.Context, *mediapackagevod.DescribeAssetInput, ...request.Option) (*mediapackagevod.DescribeAssetOutput, error)
	DescribeAssetRequest(*mediapackagevod.DescribeAssetInput) (*request.Request, *mediapackagevod.DescribeAssetOutput)

	DescribePackagingConfiguration(*mediapackagevod.DescribePackagingConfigurationInput) (*mediapackagevod.DescribePackagingConfigurationOutput, error)
	DescribePackagingConfigurationWithContext(aws.Context, *mediapackagevod.DescribePackagingConfigurationInput, ...request.Option) (*mediapackagevod.DescribePackagingConfigurationOutput, error)
	DescribePackagingConfigurationRequest(*mediapackagevod.DescribePackagingConfigurationInput) (*request.Request, *mediapackagevod.DescribePackagingConfigurationOutput)

	DescribePackagingGroup(*mediapackagevod.DescribePackagingGroupInput) (*mediapackagevod.DescribePackagingGroupOutput, error)
	DescribePackagingGroupWithContext(aws.Context, *mediapackagevod.DescribePackagingGroupInput, ...request.Option) (*mediapackagevod.DescribePackagingGroupOutput, error)
	DescribePackagingGroupRequest(*mediapackagevod.DescribePackagingGroupInput) (*request.Request, *mediapackagevod.DescribePackagingGroupOutput)

	ListAssets(*mediapackagevod.ListAssetsInput) (*mediapackagevod.ListAssetsOutput, error)
	ListAssetsWithContext(aws.Context, *mediapackagevod.ListAssetsInput, ...request.Option) (*mediapackagevod.ListAssetsOutput, error)
	ListAssetsRequest(*mediapackagevod.ListAssetsInput) (*request.Request, *mediapackagevod.ListAssetsOutput)

	ListAssetsPages(*mediapackagevod.ListAssetsInput, func(*mediapackagevod.ListAssetsOutput, bool) bool) error
	ListAssetsPagesWithContext(aws.Context, *mediapackagevod.ListAssetsInput, func(*mediapackagevod.ListAssetsOutput, bool) bool, ...request.Option) error

	ListPackagingConfigurations(*mediapackagevod.ListPackagingConfigurationsInput) (*mediapackagevod.ListPackagingConfigurationsOutput, error)
	ListPackagingConfigurationsWithContext(aws.Context, *mediapackagevod.ListPackagingConfigurationsInput, ...request.Option) (*mediapackagevod.ListPackagingConfigurationsOutput, error)
	ListPackagingConfigurationsRequest(*mediapackagevod.ListPackagingConfigurationsInput) (*request.Request, *mediapackagevod.ListPackagingConfigurationsOutput)

	ListPackagingConfigurationsPages(*mediapackagevod.ListPackagingConfigurationsInput, func(*mediapackagevod.ListPackagingConfigurationsOutput, bool) bool) error
	ListPackagingConfigurationsPagesWithContext(aws.Context, *mediapackagevod.ListPackagingConfigurationsInput, func(*mediapackagevod.ListPackagingConfigurationsOutput, bool) bool, ...request.Option) error

	ListPackagingGroups(*mediapackagevod.ListPackagingGroupsInput) (*mediapackagevod.ListPackagingGroupsOutput, error)
	ListPackagingGroupsWithContext(aws.Context, *mediapackagevod.ListPackagingGroupsInput, ...request.Option) (*mediapackagevod.ListPackagingGroupsOutput, error)
	ListPackagingGroupsRequest(*mediapackagevod.ListPackagingGroupsInput) (*request.Request, *mediapackagevod.ListPackagingGroupsOutput)

	ListPackagingGroupsPages(*mediapackagevod.ListPackagingGroupsInput, func(*mediapackagevod.ListPackagingGroupsOutput, bool) bool) error
	ListPackagingGroupsPagesWithContext(aws.Context, *mediapackagevod.ListPackagingGroupsInput, func(*mediapackagevod.ListPackagingGroupsOutput, bool) bool, ...request.Option) error
}

var _ MediaPackageVodAPI = (*mediapackagevod.MediaPackageVod)(nil)
