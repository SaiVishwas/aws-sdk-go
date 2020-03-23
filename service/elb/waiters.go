// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elb

import (
	"time"

	"github.com/SaiVishwas/aws-sdk-go/aws"
	"github.com/SaiVishwas/aws-sdk-go/aws/request"
)

// WaitUntilAnyInstanceInService uses the Elastic Load Balancing API operation
// DescribeInstanceHealth to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *ELB) WaitUntilAnyInstanceInService(input *DescribeInstanceHealthInput) error {
	return c.WaitUntilAnyInstanceInServiceWithContext(aws.BackgroundContext(), input)
}

// WaitUntilAnyInstanceInServiceWithContext is an extended version of WaitUntilAnyInstanceInService.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ELB) WaitUntilAnyInstanceInServiceWithContext(ctx aws.Context, input *DescribeInstanceHealthInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilAnyInstanceInService",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "InstanceStates[].State",
				Expected: "InService",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeInstanceHealthInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeInstanceHealthRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilInstanceDeregistered uses the Elastic Load Balancing API operation
// DescribeInstanceHealth to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *ELB) WaitUntilInstanceDeregistered(input *DescribeInstanceHealthInput) error {
	return c.WaitUntilInstanceDeregisteredWithContext(aws.BackgroundContext(), input)
}

// WaitUntilInstanceDeregisteredWithContext is an extended version of WaitUntilInstanceDeregistered.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ELB) WaitUntilInstanceDeregisteredWithContext(ctx aws.Context, input *DescribeInstanceHealthInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilInstanceDeregistered",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "InstanceStates[].State",
				Expected: "OutOfService",
			},
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "InvalidInstance",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeInstanceHealthInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeInstanceHealthRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilInstanceInService uses the Elastic Load Balancing API operation
// DescribeInstanceHealth to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *ELB) WaitUntilInstanceInService(input *DescribeInstanceHealthInput) error {
	return c.WaitUntilInstanceInServiceWithContext(aws.BackgroundContext(), input)
}

// WaitUntilInstanceInServiceWithContext is an extended version of WaitUntilInstanceInService.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ELB) WaitUntilInstanceInServiceWithContext(ctx aws.Context, input *DescribeInstanceHealthInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilInstanceInService",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "InstanceStates[].State",
				Expected: "InService",
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "InvalidInstance",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeInstanceHealthInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeInstanceHealthRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
