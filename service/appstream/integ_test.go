// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// +build go1.10,integration

package appstream_test

import (
	"context"
	"testing"
	"time"

	"github.com/SaiVishwas/aws-sdk-go/aws"
	"github.com/SaiVishwas/aws-sdk-go/aws/awserr"
	"github.com/SaiVishwas/aws-sdk-go/aws/request"
	"github.com/SaiVishwas/aws-sdk-go/awstesting/integration"
	"github.com/SaiVishwas/aws-sdk-go/service/appstream"
)

var _ aws.Config
var _ awserr.Error
var _ request.Request

func TestInteg_00_DescribeStacks(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-west-2")
	svc := appstream.New(sess)
	params := &appstream.DescribeStacksInput{}
	_, err := svc.DescribeStacksWithContext(ctx, params, func(r *request.Request) {
		r.Handlers.Validate.RemoveByName("core.ValidateParametersHandler")
	})
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
