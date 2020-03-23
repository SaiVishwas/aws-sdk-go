// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// +build go1.10,integration

package cloudwatch_test

import (
	"context"
	"testing"
	"time"

	"github.com/SaiVishwas/aws-sdk-go/aws"
	"github.com/SaiVishwas/aws-sdk-go/aws/awserr"
	"github.com/SaiVishwas/aws-sdk-go/aws/request"
	"github.com/SaiVishwas/aws-sdk-go/awstesting/integration"
	"github.com/SaiVishwas/aws-sdk-go/service/cloudwatch"
)

var _ aws.Config
var _ awserr.Error
var _ request.Request

func TestInteg_00_ListMetrics(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-west-2")
	svc := cloudwatch.New(sess)
	params := &cloudwatch.ListMetricsInput{
		Namespace: aws.String("AWS/EC2"),
	}
	_, err := svc.ListMetricsWithContext(ctx, params, func(r *request.Request) {
		r.Handlers.Validate.RemoveByName("core.ValidateParametersHandler")
	})
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
func TestInteg_01_SetAlarmState(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-west-2")
	svc := cloudwatch.New(sess)
	params := &cloudwatch.SetAlarmStateInput{
		AlarmName:   aws.String("abc"),
		StateReason: aws.String("xyz"),
		StateValue:  aws.String("mno"),
	}
	_, err := svc.SetAlarmStateWithContext(ctx, params, func(r *request.Request) {
		r.Handlers.Validate.RemoveByName("core.ValidateParametersHandler")
	})
	if err == nil {
		t.Fatalf("expect request to fail")
	}
	aerr, ok := err.(awserr.RequestFailure)
	if !ok {
		t.Fatalf("expect awserr, was %T", err)
	}
	if len(aerr.Code()) == 0 {
		t.Errorf("expect non-empty error code")
	}
	if len(aerr.Message()) == 0 {
		t.Errorf("expect non-empty error message")
	}
	if v := aerr.Code(); v == request.ErrCodeSerialization {
		t.Errorf("expect API error code got serialization failure")
	}
}
