// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloudwatchiface provides an interface to enable mocking the Amazon CloudWatch service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloudwatchiface

import (
	"github.com/SaiVishwas/aws-sdk-go/aws"
	"github.com/SaiVishwas/aws-sdk-go/aws/request"
	"github.com/SaiVishwas/aws-sdk-go/service/cloudwatch"
)

// CloudWatchAPI provides an interface to enable mocking the
// cloudwatch.CloudWatch service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon CloudWatch.
//    func myFunc(svc cloudwatchiface.CloudWatchAPI) bool {
//        // Make svc.DeleteAlarms request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := cloudwatch.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCloudWatchClient struct {
//        cloudwatchiface.CloudWatchAPI
//    }
//    func (m *mockCloudWatchClient) DeleteAlarms(input *cloudwatch.DeleteAlarmsInput) (*cloudwatch.DeleteAlarmsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCloudWatchClient{}
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
type CloudWatchAPI interface {
	DeleteAlarms(*cloudwatch.DeleteAlarmsInput) (*cloudwatch.DeleteAlarmsOutput, error)
	DeleteAlarmsWithContext(aws.Context, *cloudwatch.DeleteAlarmsInput, ...request.Option) (*cloudwatch.DeleteAlarmsOutput, error)
	DeleteAlarmsRequest(*cloudwatch.DeleteAlarmsInput) (*request.Request, *cloudwatch.DeleteAlarmsOutput)

	DeleteAnomalyDetector(*cloudwatch.DeleteAnomalyDetectorInput) (*cloudwatch.DeleteAnomalyDetectorOutput, error)
	DeleteAnomalyDetectorWithContext(aws.Context, *cloudwatch.DeleteAnomalyDetectorInput, ...request.Option) (*cloudwatch.DeleteAnomalyDetectorOutput, error)
	DeleteAnomalyDetectorRequest(*cloudwatch.DeleteAnomalyDetectorInput) (*request.Request, *cloudwatch.DeleteAnomalyDetectorOutput)

	DeleteDashboards(*cloudwatch.DeleteDashboardsInput) (*cloudwatch.DeleteDashboardsOutput, error)
	DeleteDashboardsWithContext(aws.Context, *cloudwatch.DeleteDashboardsInput, ...request.Option) (*cloudwatch.DeleteDashboardsOutput, error)
	DeleteDashboardsRequest(*cloudwatch.DeleteDashboardsInput) (*request.Request, *cloudwatch.DeleteDashboardsOutput)

	DeleteInsightRules(*cloudwatch.DeleteInsightRulesInput) (*cloudwatch.DeleteInsightRulesOutput, error)
	DeleteInsightRulesWithContext(aws.Context, *cloudwatch.DeleteInsightRulesInput, ...request.Option) (*cloudwatch.DeleteInsightRulesOutput, error)
	DeleteInsightRulesRequest(*cloudwatch.DeleteInsightRulesInput) (*request.Request, *cloudwatch.DeleteInsightRulesOutput)

	DescribeAlarmHistory(*cloudwatch.DescribeAlarmHistoryInput) (*cloudwatch.DescribeAlarmHistoryOutput, error)
	DescribeAlarmHistoryWithContext(aws.Context, *cloudwatch.DescribeAlarmHistoryInput, ...request.Option) (*cloudwatch.DescribeAlarmHistoryOutput, error)
	DescribeAlarmHistoryRequest(*cloudwatch.DescribeAlarmHistoryInput) (*request.Request, *cloudwatch.DescribeAlarmHistoryOutput)

	DescribeAlarmHistoryPages(*cloudwatch.DescribeAlarmHistoryInput, func(*cloudwatch.DescribeAlarmHistoryOutput, bool) bool) error
	DescribeAlarmHistoryPagesWithContext(aws.Context, *cloudwatch.DescribeAlarmHistoryInput, func(*cloudwatch.DescribeAlarmHistoryOutput, bool) bool, ...request.Option) error

	DescribeAlarms(*cloudwatch.DescribeAlarmsInput) (*cloudwatch.DescribeAlarmsOutput, error)
	DescribeAlarmsWithContext(aws.Context, *cloudwatch.DescribeAlarmsInput, ...request.Option) (*cloudwatch.DescribeAlarmsOutput, error)
	DescribeAlarmsRequest(*cloudwatch.DescribeAlarmsInput) (*request.Request, *cloudwatch.DescribeAlarmsOutput)

	DescribeAlarmsPages(*cloudwatch.DescribeAlarmsInput, func(*cloudwatch.DescribeAlarmsOutput, bool) bool) error
	DescribeAlarmsPagesWithContext(aws.Context, *cloudwatch.DescribeAlarmsInput, func(*cloudwatch.DescribeAlarmsOutput, bool) bool, ...request.Option) error

	DescribeAlarmsForMetric(*cloudwatch.DescribeAlarmsForMetricInput) (*cloudwatch.DescribeAlarmsForMetricOutput, error)
	DescribeAlarmsForMetricWithContext(aws.Context, *cloudwatch.DescribeAlarmsForMetricInput, ...request.Option) (*cloudwatch.DescribeAlarmsForMetricOutput, error)
	DescribeAlarmsForMetricRequest(*cloudwatch.DescribeAlarmsForMetricInput) (*request.Request, *cloudwatch.DescribeAlarmsForMetricOutput)

	DescribeAnomalyDetectors(*cloudwatch.DescribeAnomalyDetectorsInput) (*cloudwatch.DescribeAnomalyDetectorsOutput, error)
	DescribeAnomalyDetectorsWithContext(aws.Context, *cloudwatch.DescribeAnomalyDetectorsInput, ...request.Option) (*cloudwatch.DescribeAnomalyDetectorsOutput, error)
	DescribeAnomalyDetectorsRequest(*cloudwatch.DescribeAnomalyDetectorsInput) (*request.Request, *cloudwatch.DescribeAnomalyDetectorsOutput)

	DescribeInsightRules(*cloudwatch.DescribeInsightRulesInput) (*cloudwatch.DescribeInsightRulesOutput, error)
	DescribeInsightRulesWithContext(aws.Context, *cloudwatch.DescribeInsightRulesInput, ...request.Option) (*cloudwatch.DescribeInsightRulesOutput, error)
	DescribeInsightRulesRequest(*cloudwatch.DescribeInsightRulesInput) (*request.Request, *cloudwatch.DescribeInsightRulesOutput)

	DescribeInsightRulesPages(*cloudwatch.DescribeInsightRulesInput, func(*cloudwatch.DescribeInsightRulesOutput, bool) bool) error
	DescribeInsightRulesPagesWithContext(aws.Context, *cloudwatch.DescribeInsightRulesInput, func(*cloudwatch.DescribeInsightRulesOutput, bool) bool, ...request.Option) error

	DisableAlarmActions(*cloudwatch.DisableAlarmActionsInput) (*cloudwatch.DisableAlarmActionsOutput, error)
	DisableAlarmActionsWithContext(aws.Context, *cloudwatch.DisableAlarmActionsInput, ...request.Option) (*cloudwatch.DisableAlarmActionsOutput, error)
	DisableAlarmActionsRequest(*cloudwatch.DisableAlarmActionsInput) (*request.Request, *cloudwatch.DisableAlarmActionsOutput)

	DisableInsightRules(*cloudwatch.DisableInsightRulesInput) (*cloudwatch.DisableInsightRulesOutput, error)
	DisableInsightRulesWithContext(aws.Context, *cloudwatch.DisableInsightRulesInput, ...request.Option) (*cloudwatch.DisableInsightRulesOutput, error)
	DisableInsightRulesRequest(*cloudwatch.DisableInsightRulesInput) (*request.Request, *cloudwatch.DisableInsightRulesOutput)

	EnableAlarmActions(*cloudwatch.EnableAlarmActionsInput) (*cloudwatch.EnableAlarmActionsOutput, error)
	EnableAlarmActionsWithContext(aws.Context, *cloudwatch.EnableAlarmActionsInput, ...request.Option) (*cloudwatch.EnableAlarmActionsOutput, error)
	EnableAlarmActionsRequest(*cloudwatch.EnableAlarmActionsInput) (*request.Request, *cloudwatch.EnableAlarmActionsOutput)

	EnableInsightRules(*cloudwatch.EnableInsightRulesInput) (*cloudwatch.EnableInsightRulesOutput, error)
	EnableInsightRulesWithContext(aws.Context, *cloudwatch.EnableInsightRulesInput, ...request.Option) (*cloudwatch.EnableInsightRulesOutput, error)
	EnableInsightRulesRequest(*cloudwatch.EnableInsightRulesInput) (*request.Request, *cloudwatch.EnableInsightRulesOutput)

	GetDashboard(*cloudwatch.GetDashboardInput) (*cloudwatch.GetDashboardOutput, error)
	GetDashboardWithContext(aws.Context, *cloudwatch.GetDashboardInput, ...request.Option) (*cloudwatch.GetDashboardOutput, error)
	GetDashboardRequest(*cloudwatch.GetDashboardInput) (*request.Request, *cloudwatch.GetDashboardOutput)

	GetInsightRuleReport(*cloudwatch.GetInsightRuleReportInput) (*cloudwatch.GetInsightRuleReportOutput, error)
	GetInsightRuleReportWithContext(aws.Context, *cloudwatch.GetInsightRuleReportInput, ...request.Option) (*cloudwatch.GetInsightRuleReportOutput, error)
	GetInsightRuleReportRequest(*cloudwatch.GetInsightRuleReportInput) (*request.Request, *cloudwatch.GetInsightRuleReportOutput)

	GetMetricData(*cloudwatch.GetMetricDataInput) (*cloudwatch.GetMetricDataOutput, error)
	GetMetricDataWithContext(aws.Context, *cloudwatch.GetMetricDataInput, ...request.Option) (*cloudwatch.GetMetricDataOutput, error)
	GetMetricDataRequest(*cloudwatch.GetMetricDataInput) (*request.Request, *cloudwatch.GetMetricDataOutput)

	GetMetricDataPages(*cloudwatch.GetMetricDataInput, func(*cloudwatch.GetMetricDataOutput, bool) bool) error
	GetMetricDataPagesWithContext(aws.Context, *cloudwatch.GetMetricDataInput, func(*cloudwatch.GetMetricDataOutput, bool) bool, ...request.Option) error

	GetMetricStatistics(*cloudwatch.GetMetricStatisticsInput) (*cloudwatch.GetMetricStatisticsOutput, error)
	GetMetricStatisticsWithContext(aws.Context, *cloudwatch.GetMetricStatisticsInput, ...request.Option) (*cloudwatch.GetMetricStatisticsOutput, error)
	GetMetricStatisticsRequest(*cloudwatch.GetMetricStatisticsInput) (*request.Request, *cloudwatch.GetMetricStatisticsOutput)

	GetMetricWidgetImage(*cloudwatch.GetMetricWidgetImageInput) (*cloudwatch.GetMetricWidgetImageOutput, error)
	GetMetricWidgetImageWithContext(aws.Context, *cloudwatch.GetMetricWidgetImageInput, ...request.Option) (*cloudwatch.GetMetricWidgetImageOutput, error)
	GetMetricWidgetImageRequest(*cloudwatch.GetMetricWidgetImageInput) (*request.Request, *cloudwatch.GetMetricWidgetImageOutput)

	ListDashboards(*cloudwatch.ListDashboardsInput) (*cloudwatch.ListDashboardsOutput, error)
	ListDashboardsWithContext(aws.Context, *cloudwatch.ListDashboardsInput, ...request.Option) (*cloudwatch.ListDashboardsOutput, error)
	ListDashboardsRequest(*cloudwatch.ListDashboardsInput) (*request.Request, *cloudwatch.ListDashboardsOutput)

	ListDashboardsPages(*cloudwatch.ListDashboardsInput, func(*cloudwatch.ListDashboardsOutput, bool) bool) error
	ListDashboardsPagesWithContext(aws.Context, *cloudwatch.ListDashboardsInput, func(*cloudwatch.ListDashboardsOutput, bool) bool, ...request.Option) error

	ListMetrics(*cloudwatch.ListMetricsInput) (*cloudwatch.ListMetricsOutput, error)
	ListMetricsWithContext(aws.Context, *cloudwatch.ListMetricsInput, ...request.Option) (*cloudwatch.ListMetricsOutput, error)
	ListMetricsRequest(*cloudwatch.ListMetricsInput) (*request.Request, *cloudwatch.ListMetricsOutput)

	ListMetricsPages(*cloudwatch.ListMetricsInput, func(*cloudwatch.ListMetricsOutput, bool) bool) error
	ListMetricsPagesWithContext(aws.Context, *cloudwatch.ListMetricsInput, func(*cloudwatch.ListMetricsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*cloudwatch.ListTagsForResourceInput) (*cloudwatch.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *cloudwatch.ListTagsForResourceInput, ...request.Option) (*cloudwatch.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*cloudwatch.ListTagsForResourceInput) (*request.Request, *cloudwatch.ListTagsForResourceOutput)

	PutAnomalyDetector(*cloudwatch.PutAnomalyDetectorInput) (*cloudwatch.PutAnomalyDetectorOutput, error)
	PutAnomalyDetectorWithContext(aws.Context, *cloudwatch.PutAnomalyDetectorInput, ...request.Option) (*cloudwatch.PutAnomalyDetectorOutput, error)
	PutAnomalyDetectorRequest(*cloudwatch.PutAnomalyDetectorInput) (*request.Request, *cloudwatch.PutAnomalyDetectorOutput)

	PutCompositeAlarm(*cloudwatch.PutCompositeAlarmInput) (*cloudwatch.PutCompositeAlarmOutput, error)
	PutCompositeAlarmWithContext(aws.Context, *cloudwatch.PutCompositeAlarmInput, ...request.Option) (*cloudwatch.PutCompositeAlarmOutput, error)
	PutCompositeAlarmRequest(*cloudwatch.PutCompositeAlarmInput) (*request.Request, *cloudwatch.PutCompositeAlarmOutput)

	PutDashboard(*cloudwatch.PutDashboardInput) (*cloudwatch.PutDashboardOutput, error)
	PutDashboardWithContext(aws.Context, *cloudwatch.PutDashboardInput, ...request.Option) (*cloudwatch.PutDashboardOutput, error)
	PutDashboardRequest(*cloudwatch.PutDashboardInput) (*request.Request, *cloudwatch.PutDashboardOutput)

	PutInsightRule(*cloudwatch.PutInsightRuleInput) (*cloudwatch.PutInsightRuleOutput, error)
	PutInsightRuleWithContext(aws.Context, *cloudwatch.PutInsightRuleInput, ...request.Option) (*cloudwatch.PutInsightRuleOutput, error)
	PutInsightRuleRequest(*cloudwatch.PutInsightRuleInput) (*request.Request, *cloudwatch.PutInsightRuleOutput)

	PutMetricAlarm(*cloudwatch.PutMetricAlarmInput) (*cloudwatch.PutMetricAlarmOutput, error)
	PutMetricAlarmWithContext(aws.Context, *cloudwatch.PutMetricAlarmInput, ...request.Option) (*cloudwatch.PutMetricAlarmOutput, error)
	PutMetricAlarmRequest(*cloudwatch.PutMetricAlarmInput) (*request.Request, *cloudwatch.PutMetricAlarmOutput)

	PutMetricData(*cloudwatch.PutMetricDataInput) (*cloudwatch.PutMetricDataOutput, error)
	PutMetricDataWithContext(aws.Context, *cloudwatch.PutMetricDataInput, ...request.Option) (*cloudwatch.PutMetricDataOutput, error)
	PutMetricDataRequest(*cloudwatch.PutMetricDataInput) (*request.Request, *cloudwatch.PutMetricDataOutput)

	SetAlarmState(*cloudwatch.SetAlarmStateInput) (*cloudwatch.SetAlarmStateOutput, error)
	SetAlarmStateWithContext(aws.Context, *cloudwatch.SetAlarmStateInput, ...request.Option) (*cloudwatch.SetAlarmStateOutput, error)
	SetAlarmStateRequest(*cloudwatch.SetAlarmStateInput) (*request.Request, *cloudwatch.SetAlarmStateOutput)

	TagResource(*cloudwatch.TagResourceInput) (*cloudwatch.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *cloudwatch.TagResourceInput, ...request.Option) (*cloudwatch.TagResourceOutput, error)
	TagResourceRequest(*cloudwatch.TagResourceInput) (*request.Request, *cloudwatch.TagResourceOutput)

	UntagResource(*cloudwatch.UntagResourceInput) (*cloudwatch.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *cloudwatch.UntagResourceInput, ...request.Option) (*cloudwatch.UntagResourceOutput, error)
	UntagResourceRequest(*cloudwatch.UntagResourceInput) (*request.Request, *cloudwatch.UntagResourceOutput)

	WaitUntilAlarmExists(*cloudwatch.DescribeAlarmsInput) error
	WaitUntilAlarmExistsWithContext(aws.Context, *cloudwatch.DescribeAlarmsInput, ...request.WaiterOption) error

	WaitUntilCompositeAlarmExists(*cloudwatch.DescribeAlarmsInput) error
	WaitUntilCompositeAlarmExistsWithContext(aws.Context, *cloudwatch.DescribeAlarmsInput, ...request.WaiterOption) error
}

var _ CloudWatchAPI = (*cloudwatch.CloudWatch)(nil)
