// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dax

import (
	"github.com/SaiVishwas/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeClusterAlreadyExistsFault for service response error code
	// "ClusterAlreadyExistsFault".
	//
	// You already have a DAX cluster with the given identifier.
	ErrCodeClusterAlreadyExistsFault = "ClusterAlreadyExistsFault"

	// ErrCodeClusterNotFoundFault for service response error code
	// "ClusterNotFoundFault".
	//
	// The requested cluster ID does not refer to an existing DAX cluster.
	ErrCodeClusterNotFoundFault = "ClusterNotFoundFault"

	// ErrCodeClusterQuotaForCustomerExceededFault for service response error code
	// "ClusterQuotaForCustomerExceededFault".
	//
	// You have attempted to exceed the maximum number of DAX clusters for your
	// AWS account.
	ErrCodeClusterQuotaForCustomerExceededFault = "ClusterQuotaForCustomerExceededFault"

	// ErrCodeInsufficientClusterCapacityFault for service response error code
	// "InsufficientClusterCapacityFault".
	//
	// There are not enough system resources to create the cluster you requested
	// (or to resize an already-existing cluster).
	ErrCodeInsufficientClusterCapacityFault = "InsufficientClusterCapacityFault"

	// ErrCodeInvalidARNFault for service response error code
	// "InvalidARNFault".
	//
	// The Amazon Resource Name (ARN) supplied in the request is not valid.
	ErrCodeInvalidARNFault = "InvalidARNFault"

	// ErrCodeInvalidClusterStateFault for service response error code
	// "InvalidClusterStateFault".
	//
	// The requested DAX cluster is not in the available state.
	ErrCodeInvalidClusterStateFault = "InvalidClusterStateFault"

	// ErrCodeInvalidParameterCombinationException for service response error code
	// "InvalidParameterCombinationException".
	//
	// Two or more incompatible parameters were specified.
	ErrCodeInvalidParameterCombinationException = "InvalidParameterCombinationException"

	// ErrCodeInvalidParameterGroupStateFault for service response error code
	// "InvalidParameterGroupStateFault".
	//
	// One or more parameters in a parameter group are in an invalid state.
	ErrCodeInvalidParameterGroupStateFault = "InvalidParameterGroupStateFault"

	// ErrCodeInvalidParameterValueException for service response error code
	// "InvalidParameterValueException".
	//
	// The value for a parameter is invalid.
	ErrCodeInvalidParameterValueException = "InvalidParameterValueException"

	// ErrCodeInvalidSubnet for service response error code
	// "InvalidSubnet".
	//
	// An invalid subnet identifier was specified.
	ErrCodeInvalidSubnet = "InvalidSubnet"

	// ErrCodeInvalidVPCNetworkStateFault for service response error code
	// "InvalidVPCNetworkStateFault".
	//
	// The VPC network is in an invalid state.
	ErrCodeInvalidVPCNetworkStateFault = "InvalidVPCNetworkStateFault"

	// ErrCodeNodeNotFoundFault for service response error code
	// "NodeNotFoundFault".
	//
	// None of the nodes in the cluster have the given node ID.
	ErrCodeNodeNotFoundFault = "NodeNotFoundFault"

	// ErrCodeNodeQuotaForClusterExceededFault for service response error code
	// "NodeQuotaForClusterExceededFault".
	//
	// You have attempted to exceed the maximum number of nodes for a DAX cluster.
	ErrCodeNodeQuotaForClusterExceededFault = "NodeQuotaForClusterExceededFault"

	// ErrCodeNodeQuotaForCustomerExceededFault for service response error code
	// "NodeQuotaForCustomerExceededFault".
	//
	// You have attempted to exceed the maximum number of nodes for your AWS account.
	ErrCodeNodeQuotaForCustomerExceededFault = "NodeQuotaForCustomerExceededFault"

	// ErrCodeParameterGroupAlreadyExistsFault for service response error code
	// "ParameterGroupAlreadyExistsFault".
	//
	// The specified parameter group already exists.
	ErrCodeParameterGroupAlreadyExistsFault = "ParameterGroupAlreadyExistsFault"

	// ErrCodeParameterGroupNotFoundFault for service response error code
	// "ParameterGroupNotFoundFault".
	//
	// The specified parameter group does not exist.
	ErrCodeParameterGroupNotFoundFault = "ParameterGroupNotFoundFault"

	// ErrCodeParameterGroupQuotaExceededFault for service response error code
	// "ParameterGroupQuotaExceededFault".
	//
	// You have attempted to exceed the maximum number of parameter groups.
	ErrCodeParameterGroupQuotaExceededFault = "ParameterGroupQuotaExceededFault"

	// ErrCodeServiceLinkedRoleNotFoundFault for service response error code
	// "ServiceLinkedRoleNotFoundFault".
	//
	// The specified service linked role (SLR) was not found.
	ErrCodeServiceLinkedRoleNotFoundFault = "ServiceLinkedRoleNotFoundFault"

	// ErrCodeSubnetGroupAlreadyExistsFault for service response error code
	// "SubnetGroupAlreadyExistsFault".
	//
	// The specified subnet group already exists.
	ErrCodeSubnetGroupAlreadyExistsFault = "SubnetGroupAlreadyExistsFault"

	// ErrCodeSubnetGroupInUseFault for service response error code
	// "SubnetGroupInUseFault".
	//
	// The specified subnet group is currently in use.
	ErrCodeSubnetGroupInUseFault = "SubnetGroupInUseFault"

	// ErrCodeSubnetGroupNotFoundFault for service response error code
	// "SubnetGroupNotFoundFault".
	//
	// The requested subnet group name does not refer to an existing subnet group.
	ErrCodeSubnetGroupNotFoundFault = "SubnetGroupNotFoundFault"

	// ErrCodeSubnetGroupQuotaExceededFault for service response error code
	// "SubnetGroupQuotaExceededFault".
	//
	// The request cannot be processed because it would exceed the allowed number
	// of subnets in a subnet group.
	ErrCodeSubnetGroupQuotaExceededFault = "SubnetGroupQuotaExceededFault"

	// ErrCodeSubnetInUse for service response error code
	// "SubnetInUse".
	//
	// The requested subnet is being used by another subnet group.
	ErrCodeSubnetInUse = "SubnetInUse"

	// ErrCodeSubnetQuotaExceededFault for service response error code
	// "SubnetQuotaExceededFault".
	//
	// The request cannot be processed because it would exceed the allowed number
	// of subnets in a subnet group.
	ErrCodeSubnetQuotaExceededFault = "SubnetQuotaExceededFault"

	// ErrCodeTagNotFoundFault for service response error code
	// "TagNotFoundFault".
	//
	// The tag does not exist.
	ErrCodeTagNotFoundFault = "TagNotFoundFault"

	// ErrCodeTagQuotaPerResourceExceeded for service response error code
	// "TagQuotaPerResourceExceeded".
	//
	// You have exceeded the maximum number of tags for this DAX cluster.
	ErrCodeTagQuotaPerResourceExceeded = "TagQuotaPerResourceExceeded"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"ClusterAlreadyExistsFault":            newErrorClusterAlreadyExistsFault,
	"ClusterNotFoundFault":                 newErrorClusterNotFoundFault,
	"ClusterQuotaForCustomerExceededFault": newErrorClusterQuotaForCustomerExceededFault,
	"InsufficientClusterCapacityFault":     newErrorInsufficientClusterCapacityFault,
	"InvalidARNFault":                      newErrorInvalidARNFault,
	"InvalidClusterStateFault":             newErrorInvalidClusterStateFault,
	"InvalidParameterCombinationException": newErrorInvalidParameterCombinationException,
	"InvalidParameterGroupStateFault":      newErrorInvalidParameterGroupStateFault,
	"InvalidParameterValueException":       newErrorInvalidParameterValueException,
	"InvalidSubnet":                        newErrorInvalidSubnet,
	"InvalidVPCNetworkStateFault":          newErrorInvalidVPCNetworkStateFault,
	"NodeNotFoundFault":                    newErrorNodeNotFoundFault,
	"NodeQuotaForClusterExceededFault":     newErrorNodeQuotaForClusterExceededFault,
	"NodeQuotaForCustomerExceededFault":    newErrorNodeQuotaForCustomerExceededFault,
	"ParameterGroupAlreadyExistsFault":     newErrorParameterGroupAlreadyExistsFault,
	"ParameterGroupNotFoundFault":          newErrorParameterGroupNotFoundFault,
	"ParameterGroupQuotaExceededFault":     newErrorParameterGroupQuotaExceededFault,
	"ServiceLinkedRoleNotFoundFault":       newErrorServiceLinkedRoleNotFoundFault,
	"SubnetGroupAlreadyExistsFault":        newErrorSubnetGroupAlreadyExistsFault,
	"SubnetGroupInUseFault":                newErrorSubnetGroupInUseFault,
	"SubnetGroupNotFoundFault":             newErrorSubnetGroupNotFoundFault,
	"SubnetGroupQuotaExceededFault":        newErrorSubnetGroupQuotaExceededFault,
	"SubnetInUse":                          newErrorSubnetInUse,
	"SubnetQuotaExceededFault":             newErrorSubnetQuotaExceededFault,
	"TagNotFoundFault":                     newErrorTagNotFoundFault,
	"TagQuotaPerResourceExceeded":          newErrorTagQuotaPerResourceExceeded,
}
