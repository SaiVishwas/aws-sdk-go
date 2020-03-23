// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package opsworksiface provides an interface to enable mocking the AWS OpsWorks service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package opsworksiface

import (
	"github.com/SaiVishwas/aws-sdk-go/aws"
	"github.com/SaiVishwas/aws-sdk-go/aws/request"
	"github.com/SaiVishwas/aws-sdk-go/service/opsworks"
)

// OpsWorksAPI provides an interface to enable mocking the
// opsworks.OpsWorks service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS OpsWorks.
//    func myFunc(svc opsworksiface.OpsWorksAPI) bool {
//        // Make svc.AssignInstance request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := opsworks.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockOpsWorksClient struct {
//        opsworksiface.OpsWorksAPI
//    }
//    func (m *mockOpsWorksClient) AssignInstance(input *opsworks.AssignInstanceInput) (*opsworks.AssignInstanceOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockOpsWorksClient{}
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
type OpsWorksAPI interface {
	AssignInstance(*opsworks.AssignInstanceInput) (*opsworks.AssignInstanceOutput, error)
	AssignInstanceWithContext(aws.Context, *opsworks.AssignInstanceInput, ...request.Option) (*opsworks.AssignInstanceOutput, error)
	AssignInstanceRequest(*opsworks.AssignInstanceInput) (*request.Request, *opsworks.AssignInstanceOutput)

	AssignVolume(*opsworks.AssignVolumeInput) (*opsworks.AssignVolumeOutput, error)
	AssignVolumeWithContext(aws.Context, *opsworks.AssignVolumeInput, ...request.Option) (*opsworks.AssignVolumeOutput, error)
	AssignVolumeRequest(*opsworks.AssignVolumeInput) (*request.Request, *opsworks.AssignVolumeOutput)

	AssociateElasticIp(*opsworks.AssociateElasticIpInput) (*opsworks.AssociateElasticIpOutput, error)
	AssociateElasticIpWithContext(aws.Context, *opsworks.AssociateElasticIpInput, ...request.Option) (*opsworks.AssociateElasticIpOutput, error)
	AssociateElasticIpRequest(*opsworks.AssociateElasticIpInput) (*request.Request, *opsworks.AssociateElasticIpOutput)

	AttachElasticLoadBalancer(*opsworks.AttachElasticLoadBalancerInput) (*opsworks.AttachElasticLoadBalancerOutput, error)
	AttachElasticLoadBalancerWithContext(aws.Context, *opsworks.AttachElasticLoadBalancerInput, ...request.Option) (*opsworks.AttachElasticLoadBalancerOutput, error)
	AttachElasticLoadBalancerRequest(*opsworks.AttachElasticLoadBalancerInput) (*request.Request, *opsworks.AttachElasticLoadBalancerOutput)

	CloneStack(*opsworks.CloneStackInput) (*opsworks.CloneStackOutput, error)
	CloneStackWithContext(aws.Context, *opsworks.CloneStackInput, ...request.Option) (*opsworks.CloneStackOutput, error)
	CloneStackRequest(*opsworks.CloneStackInput) (*request.Request, *opsworks.CloneStackOutput)

	CreateApp(*opsworks.CreateAppInput) (*opsworks.CreateAppOutput, error)
	CreateAppWithContext(aws.Context, *opsworks.CreateAppInput, ...request.Option) (*opsworks.CreateAppOutput, error)
	CreateAppRequest(*opsworks.CreateAppInput) (*request.Request, *opsworks.CreateAppOutput)

	CreateDeployment(*opsworks.CreateDeploymentInput) (*opsworks.CreateDeploymentOutput, error)
	CreateDeploymentWithContext(aws.Context, *opsworks.CreateDeploymentInput, ...request.Option) (*opsworks.CreateDeploymentOutput, error)
	CreateDeploymentRequest(*opsworks.CreateDeploymentInput) (*request.Request, *opsworks.CreateDeploymentOutput)

	CreateInstance(*opsworks.CreateInstanceInput) (*opsworks.CreateInstanceOutput, error)
	CreateInstanceWithContext(aws.Context, *opsworks.CreateInstanceInput, ...request.Option) (*opsworks.CreateInstanceOutput, error)
	CreateInstanceRequest(*opsworks.CreateInstanceInput) (*request.Request, *opsworks.CreateInstanceOutput)

	CreateLayer(*opsworks.CreateLayerInput) (*opsworks.CreateLayerOutput, error)
	CreateLayerWithContext(aws.Context, *opsworks.CreateLayerInput, ...request.Option) (*opsworks.CreateLayerOutput, error)
	CreateLayerRequest(*opsworks.CreateLayerInput) (*request.Request, *opsworks.CreateLayerOutput)

	CreateStack(*opsworks.CreateStackInput) (*opsworks.CreateStackOutput, error)
	CreateStackWithContext(aws.Context, *opsworks.CreateStackInput, ...request.Option) (*opsworks.CreateStackOutput, error)
	CreateStackRequest(*opsworks.CreateStackInput) (*request.Request, *opsworks.CreateStackOutput)

	CreateUserProfile(*opsworks.CreateUserProfileInput) (*opsworks.CreateUserProfileOutput, error)
	CreateUserProfileWithContext(aws.Context, *opsworks.CreateUserProfileInput, ...request.Option) (*opsworks.CreateUserProfileOutput, error)
	CreateUserProfileRequest(*opsworks.CreateUserProfileInput) (*request.Request, *opsworks.CreateUserProfileOutput)

	DeleteApp(*opsworks.DeleteAppInput) (*opsworks.DeleteAppOutput, error)
	DeleteAppWithContext(aws.Context, *opsworks.DeleteAppInput, ...request.Option) (*opsworks.DeleteAppOutput, error)
	DeleteAppRequest(*opsworks.DeleteAppInput) (*request.Request, *opsworks.DeleteAppOutput)

	DeleteInstance(*opsworks.DeleteInstanceInput) (*opsworks.DeleteInstanceOutput, error)
	DeleteInstanceWithContext(aws.Context, *opsworks.DeleteInstanceInput, ...request.Option) (*opsworks.DeleteInstanceOutput, error)
	DeleteInstanceRequest(*opsworks.DeleteInstanceInput) (*request.Request, *opsworks.DeleteInstanceOutput)

	DeleteLayer(*opsworks.DeleteLayerInput) (*opsworks.DeleteLayerOutput, error)
	DeleteLayerWithContext(aws.Context, *opsworks.DeleteLayerInput, ...request.Option) (*opsworks.DeleteLayerOutput, error)
	DeleteLayerRequest(*opsworks.DeleteLayerInput) (*request.Request, *opsworks.DeleteLayerOutput)

	DeleteStack(*opsworks.DeleteStackInput) (*opsworks.DeleteStackOutput, error)
	DeleteStackWithContext(aws.Context, *opsworks.DeleteStackInput, ...request.Option) (*opsworks.DeleteStackOutput, error)
	DeleteStackRequest(*opsworks.DeleteStackInput) (*request.Request, *opsworks.DeleteStackOutput)

	DeleteUserProfile(*opsworks.DeleteUserProfileInput) (*opsworks.DeleteUserProfileOutput, error)
	DeleteUserProfileWithContext(aws.Context, *opsworks.DeleteUserProfileInput, ...request.Option) (*opsworks.DeleteUserProfileOutput, error)
	DeleteUserProfileRequest(*opsworks.DeleteUserProfileInput) (*request.Request, *opsworks.DeleteUserProfileOutput)

	DeregisterEcsCluster(*opsworks.DeregisterEcsClusterInput) (*opsworks.DeregisterEcsClusterOutput, error)
	DeregisterEcsClusterWithContext(aws.Context, *opsworks.DeregisterEcsClusterInput, ...request.Option) (*opsworks.DeregisterEcsClusterOutput, error)
	DeregisterEcsClusterRequest(*opsworks.DeregisterEcsClusterInput) (*request.Request, *opsworks.DeregisterEcsClusterOutput)

	DeregisterElasticIp(*opsworks.DeregisterElasticIpInput) (*opsworks.DeregisterElasticIpOutput, error)
	DeregisterElasticIpWithContext(aws.Context, *opsworks.DeregisterElasticIpInput, ...request.Option) (*opsworks.DeregisterElasticIpOutput, error)
	DeregisterElasticIpRequest(*opsworks.DeregisterElasticIpInput) (*request.Request, *opsworks.DeregisterElasticIpOutput)

	DeregisterInstance(*opsworks.DeregisterInstanceInput) (*opsworks.DeregisterInstanceOutput, error)
	DeregisterInstanceWithContext(aws.Context, *opsworks.DeregisterInstanceInput, ...request.Option) (*opsworks.DeregisterInstanceOutput, error)
	DeregisterInstanceRequest(*opsworks.DeregisterInstanceInput) (*request.Request, *opsworks.DeregisterInstanceOutput)

	DeregisterRdsDbInstance(*opsworks.DeregisterRdsDbInstanceInput) (*opsworks.DeregisterRdsDbInstanceOutput, error)
	DeregisterRdsDbInstanceWithContext(aws.Context, *opsworks.DeregisterRdsDbInstanceInput, ...request.Option) (*opsworks.DeregisterRdsDbInstanceOutput, error)
	DeregisterRdsDbInstanceRequest(*opsworks.DeregisterRdsDbInstanceInput) (*request.Request, *opsworks.DeregisterRdsDbInstanceOutput)

	DeregisterVolume(*opsworks.DeregisterVolumeInput) (*opsworks.DeregisterVolumeOutput, error)
	DeregisterVolumeWithContext(aws.Context, *opsworks.DeregisterVolumeInput, ...request.Option) (*opsworks.DeregisterVolumeOutput, error)
	DeregisterVolumeRequest(*opsworks.DeregisterVolumeInput) (*request.Request, *opsworks.DeregisterVolumeOutput)

	DescribeAgentVersions(*opsworks.DescribeAgentVersionsInput) (*opsworks.DescribeAgentVersionsOutput, error)
	DescribeAgentVersionsWithContext(aws.Context, *opsworks.DescribeAgentVersionsInput, ...request.Option) (*opsworks.DescribeAgentVersionsOutput, error)
	DescribeAgentVersionsRequest(*opsworks.DescribeAgentVersionsInput) (*request.Request, *opsworks.DescribeAgentVersionsOutput)

	DescribeApps(*opsworks.DescribeAppsInput) (*opsworks.DescribeAppsOutput, error)
	DescribeAppsWithContext(aws.Context, *opsworks.DescribeAppsInput, ...request.Option) (*opsworks.DescribeAppsOutput, error)
	DescribeAppsRequest(*opsworks.DescribeAppsInput) (*request.Request, *opsworks.DescribeAppsOutput)

	DescribeCommands(*opsworks.DescribeCommandsInput) (*opsworks.DescribeCommandsOutput, error)
	DescribeCommandsWithContext(aws.Context, *opsworks.DescribeCommandsInput, ...request.Option) (*opsworks.DescribeCommandsOutput, error)
	DescribeCommandsRequest(*opsworks.DescribeCommandsInput) (*request.Request, *opsworks.DescribeCommandsOutput)

	DescribeDeployments(*opsworks.DescribeDeploymentsInput) (*opsworks.DescribeDeploymentsOutput, error)
	DescribeDeploymentsWithContext(aws.Context, *opsworks.DescribeDeploymentsInput, ...request.Option) (*opsworks.DescribeDeploymentsOutput, error)
	DescribeDeploymentsRequest(*opsworks.DescribeDeploymentsInput) (*request.Request, *opsworks.DescribeDeploymentsOutput)

	DescribeEcsClusters(*opsworks.DescribeEcsClustersInput) (*opsworks.DescribeEcsClustersOutput, error)
	DescribeEcsClustersWithContext(aws.Context, *opsworks.DescribeEcsClustersInput, ...request.Option) (*opsworks.DescribeEcsClustersOutput, error)
	DescribeEcsClustersRequest(*opsworks.DescribeEcsClustersInput) (*request.Request, *opsworks.DescribeEcsClustersOutput)

	DescribeEcsClustersPages(*opsworks.DescribeEcsClustersInput, func(*opsworks.DescribeEcsClustersOutput, bool) bool) error
	DescribeEcsClustersPagesWithContext(aws.Context, *opsworks.DescribeEcsClustersInput, func(*opsworks.DescribeEcsClustersOutput, bool) bool, ...request.Option) error

	DescribeElasticIps(*opsworks.DescribeElasticIpsInput) (*opsworks.DescribeElasticIpsOutput, error)
	DescribeElasticIpsWithContext(aws.Context, *opsworks.DescribeElasticIpsInput, ...request.Option) (*opsworks.DescribeElasticIpsOutput, error)
	DescribeElasticIpsRequest(*opsworks.DescribeElasticIpsInput) (*request.Request, *opsworks.DescribeElasticIpsOutput)

	DescribeElasticLoadBalancers(*opsworks.DescribeElasticLoadBalancersInput) (*opsworks.DescribeElasticLoadBalancersOutput, error)
	DescribeElasticLoadBalancersWithContext(aws.Context, *opsworks.DescribeElasticLoadBalancersInput, ...request.Option) (*opsworks.DescribeElasticLoadBalancersOutput, error)
	DescribeElasticLoadBalancersRequest(*opsworks.DescribeElasticLoadBalancersInput) (*request.Request, *opsworks.DescribeElasticLoadBalancersOutput)

	DescribeInstances(*opsworks.DescribeInstancesInput) (*opsworks.DescribeInstancesOutput, error)
	DescribeInstancesWithContext(aws.Context, *opsworks.DescribeInstancesInput, ...request.Option) (*opsworks.DescribeInstancesOutput, error)
	DescribeInstancesRequest(*opsworks.DescribeInstancesInput) (*request.Request, *opsworks.DescribeInstancesOutput)

	DescribeLayers(*opsworks.DescribeLayersInput) (*opsworks.DescribeLayersOutput, error)
	DescribeLayersWithContext(aws.Context, *opsworks.DescribeLayersInput, ...request.Option) (*opsworks.DescribeLayersOutput, error)
	DescribeLayersRequest(*opsworks.DescribeLayersInput) (*request.Request, *opsworks.DescribeLayersOutput)

	DescribeLoadBasedAutoScaling(*opsworks.DescribeLoadBasedAutoScalingInput) (*opsworks.DescribeLoadBasedAutoScalingOutput, error)
	DescribeLoadBasedAutoScalingWithContext(aws.Context, *opsworks.DescribeLoadBasedAutoScalingInput, ...request.Option) (*opsworks.DescribeLoadBasedAutoScalingOutput, error)
	DescribeLoadBasedAutoScalingRequest(*opsworks.DescribeLoadBasedAutoScalingInput) (*request.Request, *opsworks.DescribeLoadBasedAutoScalingOutput)

	DescribeMyUserProfile(*opsworks.DescribeMyUserProfileInput) (*opsworks.DescribeMyUserProfileOutput, error)
	DescribeMyUserProfileWithContext(aws.Context, *opsworks.DescribeMyUserProfileInput, ...request.Option) (*opsworks.DescribeMyUserProfileOutput, error)
	DescribeMyUserProfileRequest(*opsworks.DescribeMyUserProfileInput) (*request.Request, *opsworks.DescribeMyUserProfileOutput)

	DescribeOperatingSystems(*opsworks.DescribeOperatingSystemsInput) (*opsworks.DescribeOperatingSystemsOutput, error)
	DescribeOperatingSystemsWithContext(aws.Context, *opsworks.DescribeOperatingSystemsInput, ...request.Option) (*opsworks.DescribeOperatingSystemsOutput, error)
	DescribeOperatingSystemsRequest(*opsworks.DescribeOperatingSystemsInput) (*request.Request, *opsworks.DescribeOperatingSystemsOutput)

	DescribePermissions(*opsworks.DescribePermissionsInput) (*opsworks.DescribePermissionsOutput, error)
	DescribePermissionsWithContext(aws.Context, *opsworks.DescribePermissionsInput, ...request.Option) (*opsworks.DescribePermissionsOutput, error)
	DescribePermissionsRequest(*opsworks.DescribePermissionsInput) (*request.Request, *opsworks.DescribePermissionsOutput)

	DescribeRaidArrays(*opsworks.DescribeRaidArraysInput) (*opsworks.DescribeRaidArraysOutput, error)
	DescribeRaidArraysWithContext(aws.Context, *opsworks.DescribeRaidArraysInput, ...request.Option) (*opsworks.DescribeRaidArraysOutput, error)
	DescribeRaidArraysRequest(*opsworks.DescribeRaidArraysInput) (*request.Request, *opsworks.DescribeRaidArraysOutput)

	DescribeRdsDbInstances(*opsworks.DescribeRdsDbInstancesInput) (*opsworks.DescribeRdsDbInstancesOutput, error)
	DescribeRdsDbInstancesWithContext(aws.Context, *opsworks.DescribeRdsDbInstancesInput, ...request.Option) (*opsworks.DescribeRdsDbInstancesOutput, error)
	DescribeRdsDbInstancesRequest(*opsworks.DescribeRdsDbInstancesInput) (*request.Request, *opsworks.DescribeRdsDbInstancesOutput)

	DescribeServiceErrors(*opsworks.DescribeServiceErrorsInput) (*opsworks.DescribeServiceErrorsOutput, error)
	DescribeServiceErrorsWithContext(aws.Context, *opsworks.DescribeServiceErrorsInput, ...request.Option) (*opsworks.DescribeServiceErrorsOutput, error)
	DescribeServiceErrorsRequest(*opsworks.DescribeServiceErrorsInput) (*request.Request, *opsworks.DescribeServiceErrorsOutput)

	DescribeStackProvisioningParameters(*opsworks.DescribeStackProvisioningParametersInput) (*opsworks.DescribeStackProvisioningParametersOutput, error)
	DescribeStackProvisioningParametersWithContext(aws.Context, *opsworks.DescribeStackProvisioningParametersInput, ...request.Option) (*opsworks.DescribeStackProvisioningParametersOutput, error)
	DescribeStackProvisioningParametersRequest(*opsworks.DescribeStackProvisioningParametersInput) (*request.Request, *opsworks.DescribeStackProvisioningParametersOutput)

	DescribeStackSummary(*opsworks.DescribeStackSummaryInput) (*opsworks.DescribeStackSummaryOutput, error)
	DescribeStackSummaryWithContext(aws.Context, *opsworks.DescribeStackSummaryInput, ...request.Option) (*opsworks.DescribeStackSummaryOutput, error)
	DescribeStackSummaryRequest(*opsworks.DescribeStackSummaryInput) (*request.Request, *opsworks.DescribeStackSummaryOutput)

	DescribeStacks(*opsworks.DescribeStacksInput) (*opsworks.DescribeStacksOutput, error)
	DescribeStacksWithContext(aws.Context, *opsworks.DescribeStacksInput, ...request.Option) (*opsworks.DescribeStacksOutput, error)
	DescribeStacksRequest(*opsworks.DescribeStacksInput) (*request.Request, *opsworks.DescribeStacksOutput)

	DescribeTimeBasedAutoScaling(*opsworks.DescribeTimeBasedAutoScalingInput) (*opsworks.DescribeTimeBasedAutoScalingOutput, error)
	DescribeTimeBasedAutoScalingWithContext(aws.Context, *opsworks.DescribeTimeBasedAutoScalingInput, ...request.Option) (*opsworks.DescribeTimeBasedAutoScalingOutput, error)
	DescribeTimeBasedAutoScalingRequest(*opsworks.DescribeTimeBasedAutoScalingInput) (*request.Request, *opsworks.DescribeTimeBasedAutoScalingOutput)

	DescribeUserProfiles(*opsworks.DescribeUserProfilesInput) (*opsworks.DescribeUserProfilesOutput, error)
	DescribeUserProfilesWithContext(aws.Context, *opsworks.DescribeUserProfilesInput, ...request.Option) (*opsworks.DescribeUserProfilesOutput, error)
	DescribeUserProfilesRequest(*opsworks.DescribeUserProfilesInput) (*request.Request, *opsworks.DescribeUserProfilesOutput)

	DescribeVolumes(*opsworks.DescribeVolumesInput) (*opsworks.DescribeVolumesOutput, error)
	DescribeVolumesWithContext(aws.Context, *opsworks.DescribeVolumesInput, ...request.Option) (*opsworks.DescribeVolumesOutput, error)
	DescribeVolumesRequest(*opsworks.DescribeVolumesInput) (*request.Request, *opsworks.DescribeVolumesOutput)

	DetachElasticLoadBalancer(*opsworks.DetachElasticLoadBalancerInput) (*opsworks.DetachElasticLoadBalancerOutput, error)
	DetachElasticLoadBalancerWithContext(aws.Context, *opsworks.DetachElasticLoadBalancerInput, ...request.Option) (*opsworks.DetachElasticLoadBalancerOutput, error)
	DetachElasticLoadBalancerRequest(*opsworks.DetachElasticLoadBalancerInput) (*request.Request, *opsworks.DetachElasticLoadBalancerOutput)

	DisassociateElasticIp(*opsworks.DisassociateElasticIpInput) (*opsworks.DisassociateElasticIpOutput, error)
	DisassociateElasticIpWithContext(aws.Context, *opsworks.DisassociateElasticIpInput, ...request.Option) (*opsworks.DisassociateElasticIpOutput, error)
	DisassociateElasticIpRequest(*opsworks.DisassociateElasticIpInput) (*request.Request, *opsworks.DisassociateElasticIpOutput)

	GetHostnameSuggestion(*opsworks.GetHostnameSuggestionInput) (*opsworks.GetHostnameSuggestionOutput, error)
	GetHostnameSuggestionWithContext(aws.Context, *opsworks.GetHostnameSuggestionInput, ...request.Option) (*opsworks.GetHostnameSuggestionOutput, error)
	GetHostnameSuggestionRequest(*opsworks.GetHostnameSuggestionInput) (*request.Request, *opsworks.GetHostnameSuggestionOutput)

	GrantAccess(*opsworks.GrantAccessInput) (*opsworks.GrantAccessOutput, error)
	GrantAccessWithContext(aws.Context, *opsworks.GrantAccessInput, ...request.Option) (*opsworks.GrantAccessOutput, error)
	GrantAccessRequest(*opsworks.GrantAccessInput) (*request.Request, *opsworks.GrantAccessOutput)

	ListTags(*opsworks.ListTagsInput) (*opsworks.ListTagsOutput, error)
	ListTagsWithContext(aws.Context, *opsworks.ListTagsInput, ...request.Option) (*opsworks.ListTagsOutput, error)
	ListTagsRequest(*opsworks.ListTagsInput) (*request.Request, *opsworks.ListTagsOutput)

	RebootInstance(*opsworks.RebootInstanceInput) (*opsworks.RebootInstanceOutput, error)
	RebootInstanceWithContext(aws.Context, *opsworks.RebootInstanceInput, ...request.Option) (*opsworks.RebootInstanceOutput, error)
	RebootInstanceRequest(*opsworks.RebootInstanceInput) (*request.Request, *opsworks.RebootInstanceOutput)

	RegisterEcsCluster(*opsworks.RegisterEcsClusterInput) (*opsworks.RegisterEcsClusterOutput, error)
	RegisterEcsClusterWithContext(aws.Context, *opsworks.RegisterEcsClusterInput, ...request.Option) (*opsworks.RegisterEcsClusterOutput, error)
	RegisterEcsClusterRequest(*opsworks.RegisterEcsClusterInput) (*request.Request, *opsworks.RegisterEcsClusterOutput)

	RegisterElasticIp(*opsworks.RegisterElasticIpInput) (*opsworks.RegisterElasticIpOutput, error)
	RegisterElasticIpWithContext(aws.Context, *opsworks.RegisterElasticIpInput, ...request.Option) (*opsworks.RegisterElasticIpOutput, error)
	RegisterElasticIpRequest(*opsworks.RegisterElasticIpInput) (*request.Request, *opsworks.RegisterElasticIpOutput)

	RegisterInstance(*opsworks.RegisterInstanceInput) (*opsworks.RegisterInstanceOutput, error)
	RegisterInstanceWithContext(aws.Context, *opsworks.RegisterInstanceInput, ...request.Option) (*opsworks.RegisterInstanceOutput, error)
	RegisterInstanceRequest(*opsworks.RegisterInstanceInput) (*request.Request, *opsworks.RegisterInstanceOutput)

	RegisterRdsDbInstance(*opsworks.RegisterRdsDbInstanceInput) (*opsworks.RegisterRdsDbInstanceOutput, error)
	RegisterRdsDbInstanceWithContext(aws.Context, *opsworks.RegisterRdsDbInstanceInput, ...request.Option) (*opsworks.RegisterRdsDbInstanceOutput, error)
	RegisterRdsDbInstanceRequest(*opsworks.RegisterRdsDbInstanceInput) (*request.Request, *opsworks.RegisterRdsDbInstanceOutput)

	RegisterVolume(*opsworks.RegisterVolumeInput) (*opsworks.RegisterVolumeOutput, error)
	RegisterVolumeWithContext(aws.Context, *opsworks.RegisterVolumeInput, ...request.Option) (*opsworks.RegisterVolumeOutput, error)
	RegisterVolumeRequest(*opsworks.RegisterVolumeInput) (*request.Request, *opsworks.RegisterVolumeOutput)

	SetLoadBasedAutoScaling(*opsworks.SetLoadBasedAutoScalingInput) (*opsworks.SetLoadBasedAutoScalingOutput, error)
	SetLoadBasedAutoScalingWithContext(aws.Context, *opsworks.SetLoadBasedAutoScalingInput, ...request.Option) (*opsworks.SetLoadBasedAutoScalingOutput, error)
	SetLoadBasedAutoScalingRequest(*opsworks.SetLoadBasedAutoScalingInput) (*request.Request, *opsworks.SetLoadBasedAutoScalingOutput)

	SetPermission(*opsworks.SetPermissionInput) (*opsworks.SetPermissionOutput, error)
	SetPermissionWithContext(aws.Context, *opsworks.SetPermissionInput, ...request.Option) (*opsworks.SetPermissionOutput, error)
	SetPermissionRequest(*opsworks.SetPermissionInput) (*request.Request, *opsworks.SetPermissionOutput)

	SetTimeBasedAutoScaling(*opsworks.SetTimeBasedAutoScalingInput) (*opsworks.SetTimeBasedAutoScalingOutput, error)
	SetTimeBasedAutoScalingWithContext(aws.Context, *opsworks.SetTimeBasedAutoScalingInput, ...request.Option) (*opsworks.SetTimeBasedAutoScalingOutput, error)
	SetTimeBasedAutoScalingRequest(*opsworks.SetTimeBasedAutoScalingInput) (*request.Request, *opsworks.SetTimeBasedAutoScalingOutput)

	StartInstance(*opsworks.StartInstanceInput) (*opsworks.StartInstanceOutput, error)
	StartInstanceWithContext(aws.Context, *opsworks.StartInstanceInput, ...request.Option) (*opsworks.StartInstanceOutput, error)
	StartInstanceRequest(*opsworks.StartInstanceInput) (*request.Request, *opsworks.StartInstanceOutput)

	StartStack(*opsworks.StartStackInput) (*opsworks.StartStackOutput, error)
	StartStackWithContext(aws.Context, *opsworks.StartStackInput, ...request.Option) (*opsworks.StartStackOutput, error)
	StartStackRequest(*opsworks.StartStackInput) (*request.Request, *opsworks.StartStackOutput)

	StopInstance(*opsworks.StopInstanceInput) (*opsworks.StopInstanceOutput, error)
	StopInstanceWithContext(aws.Context, *opsworks.StopInstanceInput, ...request.Option) (*opsworks.StopInstanceOutput, error)
	StopInstanceRequest(*opsworks.StopInstanceInput) (*request.Request, *opsworks.StopInstanceOutput)

	StopStack(*opsworks.StopStackInput) (*opsworks.StopStackOutput, error)
	StopStackWithContext(aws.Context, *opsworks.StopStackInput, ...request.Option) (*opsworks.StopStackOutput, error)
	StopStackRequest(*opsworks.StopStackInput) (*request.Request, *opsworks.StopStackOutput)

	TagResource(*opsworks.TagResourceInput) (*opsworks.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *opsworks.TagResourceInput, ...request.Option) (*opsworks.TagResourceOutput, error)
	TagResourceRequest(*opsworks.TagResourceInput) (*request.Request, *opsworks.TagResourceOutput)

	UnassignInstance(*opsworks.UnassignInstanceInput) (*opsworks.UnassignInstanceOutput, error)
	UnassignInstanceWithContext(aws.Context, *opsworks.UnassignInstanceInput, ...request.Option) (*opsworks.UnassignInstanceOutput, error)
	UnassignInstanceRequest(*opsworks.UnassignInstanceInput) (*request.Request, *opsworks.UnassignInstanceOutput)

	UnassignVolume(*opsworks.UnassignVolumeInput) (*opsworks.UnassignVolumeOutput, error)
	UnassignVolumeWithContext(aws.Context, *opsworks.UnassignVolumeInput, ...request.Option) (*opsworks.UnassignVolumeOutput, error)
	UnassignVolumeRequest(*opsworks.UnassignVolumeInput) (*request.Request, *opsworks.UnassignVolumeOutput)

	UntagResource(*opsworks.UntagResourceInput) (*opsworks.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *opsworks.UntagResourceInput, ...request.Option) (*opsworks.UntagResourceOutput, error)
	UntagResourceRequest(*opsworks.UntagResourceInput) (*request.Request, *opsworks.UntagResourceOutput)

	UpdateApp(*opsworks.UpdateAppInput) (*opsworks.UpdateAppOutput, error)
	UpdateAppWithContext(aws.Context, *opsworks.UpdateAppInput, ...request.Option) (*opsworks.UpdateAppOutput, error)
	UpdateAppRequest(*opsworks.UpdateAppInput) (*request.Request, *opsworks.UpdateAppOutput)

	UpdateElasticIp(*opsworks.UpdateElasticIpInput) (*opsworks.UpdateElasticIpOutput, error)
	UpdateElasticIpWithContext(aws.Context, *opsworks.UpdateElasticIpInput, ...request.Option) (*opsworks.UpdateElasticIpOutput, error)
	UpdateElasticIpRequest(*opsworks.UpdateElasticIpInput) (*request.Request, *opsworks.UpdateElasticIpOutput)

	UpdateInstance(*opsworks.UpdateInstanceInput) (*opsworks.UpdateInstanceOutput, error)
	UpdateInstanceWithContext(aws.Context, *opsworks.UpdateInstanceInput, ...request.Option) (*opsworks.UpdateInstanceOutput, error)
	UpdateInstanceRequest(*opsworks.UpdateInstanceInput) (*request.Request, *opsworks.UpdateInstanceOutput)

	UpdateLayer(*opsworks.UpdateLayerInput) (*opsworks.UpdateLayerOutput, error)
	UpdateLayerWithContext(aws.Context, *opsworks.UpdateLayerInput, ...request.Option) (*opsworks.UpdateLayerOutput, error)
	UpdateLayerRequest(*opsworks.UpdateLayerInput) (*request.Request, *opsworks.UpdateLayerOutput)

	UpdateMyUserProfile(*opsworks.UpdateMyUserProfileInput) (*opsworks.UpdateMyUserProfileOutput, error)
	UpdateMyUserProfileWithContext(aws.Context, *opsworks.UpdateMyUserProfileInput, ...request.Option) (*opsworks.UpdateMyUserProfileOutput, error)
	UpdateMyUserProfileRequest(*opsworks.UpdateMyUserProfileInput) (*request.Request, *opsworks.UpdateMyUserProfileOutput)

	UpdateRdsDbInstance(*opsworks.UpdateRdsDbInstanceInput) (*opsworks.UpdateRdsDbInstanceOutput, error)
	UpdateRdsDbInstanceWithContext(aws.Context, *opsworks.UpdateRdsDbInstanceInput, ...request.Option) (*opsworks.UpdateRdsDbInstanceOutput, error)
	UpdateRdsDbInstanceRequest(*opsworks.UpdateRdsDbInstanceInput) (*request.Request, *opsworks.UpdateRdsDbInstanceOutput)

	UpdateStack(*opsworks.UpdateStackInput) (*opsworks.UpdateStackOutput, error)
	UpdateStackWithContext(aws.Context, *opsworks.UpdateStackInput, ...request.Option) (*opsworks.UpdateStackOutput, error)
	UpdateStackRequest(*opsworks.UpdateStackInput) (*request.Request, *opsworks.UpdateStackOutput)

	UpdateUserProfile(*opsworks.UpdateUserProfileInput) (*opsworks.UpdateUserProfileOutput, error)
	UpdateUserProfileWithContext(aws.Context, *opsworks.UpdateUserProfileInput, ...request.Option) (*opsworks.UpdateUserProfileOutput, error)
	UpdateUserProfileRequest(*opsworks.UpdateUserProfileInput) (*request.Request, *opsworks.UpdateUserProfileOutput)

	UpdateVolume(*opsworks.UpdateVolumeInput) (*opsworks.UpdateVolumeOutput, error)
	UpdateVolumeWithContext(aws.Context, *opsworks.UpdateVolumeInput, ...request.Option) (*opsworks.UpdateVolumeOutput, error)
	UpdateVolumeRequest(*opsworks.UpdateVolumeInput) (*request.Request, *opsworks.UpdateVolumeOutput)

	WaitUntilAppExists(*opsworks.DescribeAppsInput) error
	WaitUntilAppExistsWithContext(aws.Context, *opsworks.DescribeAppsInput, ...request.WaiterOption) error

	WaitUntilDeploymentSuccessful(*opsworks.DescribeDeploymentsInput) error
	WaitUntilDeploymentSuccessfulWithContext(aws.Context, *opsworks.DescribeDeploymentsInput, ...request.WaiterOption) error

	WaitUntilInstanceOnline(*opsworks.DescribeInstancesInput) error
	WaitUntilInstanceOnlineWithContext(aws.Context, *opsworks.DescribeInstancesInput, ...request.WaiterOption) error

	WaitUntilInstanceRegistered(*opsworks.DescribeInstancesInput) error
	WaitUntilInstanceRegisteredWithContext(aws.Context, *opsworks.DescribeInstancesInput, ...request.WaiterOption) error

	WaitUntilInstanceStopped(*opsworks.DescribeInstancesInput) error
	WaitUntilInstanceStoppedWithContext(aws.Context, *opsworks.DescribeInstancesInput, ...request.WaiterOption) error

	WaitUntilInstanceTerminated(*opsworks.DescribeInstancesInput) error
	WaitUntilInstanceTerminatedWithContext(aws.Context, *opsworks.DescribeInstancesInput, ...request.WaiterOption) error
}

var _ OpsWorksAPI = (*opsworks.OpsWorks)(nil)
